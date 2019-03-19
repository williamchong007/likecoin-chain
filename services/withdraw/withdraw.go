package withdraw

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/tendermint/tendermint/crypto"
	tmRPC "github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/types"

	"github.com/likecoin/likechain/services/abi/relay"
	"github.com/likecoin/likechain/services/eth"
	logger "github.com/likecoin/likechain/services/log"
	"github.com/likecoin/likechain/services/tendermint"
	"github.com/likecoin/likechain/services/utils"
)

var log = logger.L

// type AppHashContractProof struct {
// 	Height     uint64
// 	Round      uint64
// 	Payload struct {
// 		SuffixLen  uint8
// 		Suffix     []byte
// 		VotesCount uint8
// 		Votes      []struct {
// 			TimeLen uint8
// 			Time    []byte
// 			Sig     [65]byte
// 		}
// 		AppHashLen   uint8
// 		AppHash      []byte
// 		AppHashProof [4][32]byte
// 	}
// }

func encodeTimestamp(vote *types.CanonicalVote) []byte {
	cdc := types.GetCodec()
	buf := new(bytes.Buffer)
	// Field 4, typ3 = variable length struct (2), so field index = 00100|010 = 0x22
	// See comments in the Relay smart contract  for more details
	buf.WriteByte(0x22)
	buf.Write(cdc.MustMarshalBinaryLengthPrefixed(vote.Timestamp))
	return buf.Bytes()
}

func encodeSuffix(vote *types.CanonicalVote) []byte {
	cdc := types.GetCodec()
	buf := new(bytes.Buffer)
	// Field 5, typ3 = variable length struct (2), so field index = 00101|010 = 0x2A
	// See comments in the Relay smart contract  for more details
	buf.WriteByte(0x2A)
	buf.Write(cdc.MustMarshalBinaryLengthPrefixed(vote.BlockID))
	// Field 6, typ3 = variable length struct (2), so field index = 00110|010 = 0x32
	buf.WriteByte(0x32)
	buf.Write(cdc.MustMarshalBinaryBare(vote.ChainID))
	return buf.Bytes()
}

func genContractProofPayload(signedHeader *types.SignedHeader, validators []types.Validator) []byte {
	header := signedHeader.Header
	rawVotes := signedHeader.Commit.Precommits
	votes := []*types.Vote{}
	power := int64(0)

	for _, vote := range rawVotes {
		if vote != nil && len(vote.BlockID.Hash) > 0 {
			votes = append(votes, vote)
			power += validators[vote.ValidatorIndex].VotingPower
		}
	}

	totalPower := int64(0)
	for _, validator := range validators {
		totalPower += validator.VotingPower
	}
	if power*3 <= totalPower*2 {
		return nil
	}

	cv := types.CanonicalizeVote(header.ChainID, votes[0])
	suffix := encodeSuffix(&cv)

	buf := new(bytes.Buffer)
	buf.WriteByte(uint8(len(suffix)))
	buf.Write(suffix)
	buf.WriteByte(uint8(len(votes)))

	tmToEthAddr := tendermint.MapValidatorIndexToEthAddr(validators)
	for _, vote := range votes {
		cv := types.CanonicalizeVote(header.ChainID, vote)
		time := encodeTimestamp(&cv)
		buf.WriteByte(uint8(len(time)))
		buf.Write(time)

		signBytes := vote.SignBytes(header.ChainID)
		ethAddr := tmToEthAddr[vote.ValidatorIndex]
		ethSig := tendermint.SignatureToEthereumSig(vote.Signature, crypto.Sha256(signBytes), ethAddr)
		buf.Write(ethSig[64:])
		buf.Write(ethSig[:64])
	}

	buf.WriteByte(uint8(len(header.AppHash)))
	buf.Write(header.AppHash)
	_, proof := Proof(header)
	for _, pf := range proof {
		buf.Write(pf)
	}
	return buf.Bytes()
}

func doWithdraw(tmClient *tmRPC.HTTP, lb *eth.LoadBalancer, refAuth *bind.TransactOpts, contractAddr common.Address, packedTx []byte, nonce int64) {
	lb.Do(func(ethClient *ethclient.Client) error {
		contract, err := relay.NewRelay(contractAddr, ethClient)
		if err != nil {
			log.
				WithField("relay_addr", contractAddr.Hex()).
				WithError(err).
				Error("Cannot initialize relay contract in withdraw")
			return err
		}
		withdrawIDBytes := crypto.Sha256(packedTx)
		withdrawIDBytes32 := [32]byte{}
		copy(withdrawIDBytes32[:], withdrawIDBytes)
		consumed, err := contract.ConsumedIds(nil, withdrawIDBytes32)
		if err != nil {
			log.
				WithField("packed_tx", common.Bytes2Hex(packedTx)).
				WithError(err).
				Error("Cannot check if the withdraw is already consumed from Ethereum contract")
			return err
		}
		if consumed {
			log.
				WithField("packed_tx", common.Bytes2Hex(packedTx)).
				Info("The withdraw is already processed on Ethereum contract, skipping")
			return nil
		}
		contractHeight := getContractHeight(lb, contractAddr)
		queryResult, err := tmClient.ABCIQueryWithOptions("withdraw_proof", packedTx, tmRPC.ABCIQueryOptions{Height: contractHeight})
		if err != nil {
			log.
				WithField("packed_tx", common.Bytes2Hex(packedTx)).
				WithError(err).
				Error("Cannot get withdraw_proof from LikeChain")
			return err
		}
		proof := ParseRangeProof(queryResult.Response.Value)
		if proof == nil {
			log.
				WithField("packed_tx", common.Bytes2Hex(packedTx)).
				WithField("range_proof_json", string(queryResult.Response.Value)).
				Error("Cannot parse RangeProof")
			return err
		}
		log.
			WithField("packed_tx", common.Bytes2Hex(packedTx)).
			WithField("root_hash", common.Bytes2Hex(proof.ComputeRootHash())).
			Debug("Computed RangeProof root hash")
		contractProof := proof.ContractProof()

		log.
			WithField("packed_tx", common.Bytes2Hex(packedTx)).
			WithField("contract_proof", common.Bytes2Hex(contractProof)).
			Info("Calling withdraw on Ethereum")

		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		auth := &bind.TransactOpts{
			From:    refAuth.From,
			Nonce:   big.NewInt(nonce),
			Signer:  refAuth.Signer,
			Context: c,
		}

		tx, err := contract.Withdraw(auth, packedTx, contractProof)
		if err != nil {
			log.
				WithField("packed_tx", common.Bytes2Hex(packedTx)).
				WithField("contract_proof", common.Bytes2Hex(contractProof)).
				WithError(err).
				Error("Cannot call withdraw on relay contract")
			return err
		}
		receipt := eth.WaitForReceipt(lb, tx.Hash())
		log.
			WithField("packed_tx", common.Bytes2Hex(packedTx)).
			WithField("eth_tx_hash", tx.Hash().Hex()).
			WithField("gas_used", receipt.GasUsed).
			WithField("status", receipt.Status).
			Info("withdraw call executed on Ethereum")
		return nil
	})
}

func getContractHeight(lb *eth.LoadBalancer, contractAddr common.Address) int64 {
	height := int64(0)
	lb.Do(func(ethClient *ethclient.Client) error {
		contract, err := relay.NewRelay(contractAddr, ethClient)
		if err != nil {
			log.
				WithField("relay_addr", contractAddr.Hex()).
				WithError(err).
				Error("Cannot initialize relay contract when getting contract height")
			return err
		}
		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		opts := bind.CallOpts{
			Context: c,
		}
		h, err := contract.LatestBlockHeight(&opts)
		if err != nil {
			log.
				WithField("relay_addr", contractAddr.Hex()).
				WithError(err).
				Error("Cannot get contract height")
			return err
		}
		height = h.Int64()
		return nil
	})
	return height
}

func commitWithdrawHash(tmClient *tmRPC.HTTP, lb *eth.LoadBalancer, refAuth *bind.TransactOpts, contractAddr common.Address, height int64) bool {
	result := false
	lb.Do(func(ethClient *ethclient.Client) error {
		validators := tendermint.GetValidators(tmClient)
		signedHeader := tendermint.GetSignedHeader(tmClient, height)

		log.
			WithField("header_block_hash", signedHeader.Commit.BlockID.Hash).
			Debug("Got SignedHeader")
		contractPayload := genContractProofPayload(&signedHeader, validators)
		if len(contractPayload) == 0 {
			result = false
			return nil
		}
		contract, err := relay.NewRelay(contractAddr, ethClient)
		if err != nil {
			log.
				WithField("relay_addr", contractAddr.Hex()).
				WithError(err).
				Error("Cannot initialize relay contract in commit withdraw hash")
			return err
		}

		round := uint64(signedHeader.Commit.Round())
		log.
			WithField("height", height).
			WithField("round", round).
			WithField("contract_payload", common.Bytes2Hex(contractPayload)).
			Info("Calling commitWithdrawHash on Ethereum")

		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		auth := &bind.TransactOpts{
			From:    refAuth.From,
			Signer:  refAuth.Signer,
			Context: c,
		}
		tx, err := contract.CommitWithdrawHash(auth, uint64(height), round, contractPayload)
		if err != nil {
			log.
				WithField("height", height).
				WithField("round", round).
				WithField("contract_payload", common.Bytes2Hex(contractPayload)).
				WithError(err).
				Error("Cannot call withdraw on relay contract")
			return err
		}

		receipt := eth.WaitForReceipt(lb, tx.Hash())
		log.
			WithField("eth_tx_hash", tx.Hash().Hex()).
			WithField("gas_used", receipt.GasUsed).
			WithField("status", receipt.Status).
			Info("commitWithdrawHash call executed on Ethereum")
		result = true
		return nil
	})
	return result
}

func getWithdrawPackedTxs(tmClient *tmRPC.HTTP, lastHeight, newHeight int64) [][]byte {
	log.
		WithField("last_height", lastHeight).
		WithField("new_height", newHeight).
		Info("Searching withdraws on LikeChain")
	resultTxs := tendermint.TxSearch(tmClient, "withdraw.height", lastHeight+1, newHeight)
	if len(resultTxs) == 0 {
		return nil
	}
	packedTxs := make([][]byte, 0, len(resultTxs))
	for _, resultTx := range resultTxs {
		packedTx := resultTx.TxResult.Data
		log.
			WithField("tx_hash", resultTx.Hash).
			WithField("packed_tx", common.Bytes2Hex(packedTx)).
			Debug("Withdraw search result")
		packedTxs = append(packedTxs, packedTx)
	}
	return packedTxs
}

type runState struct {
	LastHeight int64
}

func loadState(path string) (*runState, error) {
	jsonBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	state := runState{}
	err = json.Unmarshal(jsonBytes, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (state *runState) save(path string) error {
	jsonBytes, err := json.Marshal(&state)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, jsonBytes, 0644)
	return err
}

// Run starts the subscription to the withdraws on LikeChain and commits proofs onto Ethereum
func Run(tmClient *tmRPC.HTTP, lb *eth.LoadBalancer, auth *bind.TransactOpts, contractAddr common.Address, statePath string) {
	state, err := loadState(statePath)
	if err != nil {
		log.
			WithField("state_path", statePath).
			WithError(err).
			Info("Failed to load state, creating empty state")
		state = &runState{}
		state.save(statePath)
	}
	for ; ; time.Sleep(time.Minute) {
		var newHeight int64
		utils.RetryIfPanic(5, func() {
			newHeight = tendermint.GetHeight(tmClient)
		})
		if newHeight == state.LastHeight {
			log.
				WithField("last_height", state.LastHeight).
				Info("No new LikeChain block since last height")
			continue
		}
		log.
			WithField("last_height", state.LastHeight).
			WithField("new_height", newHeight).
			Info("New LikeChain height")
		var packedTxs [][]byte
		utils.RetryIfPanic(5, func() {
			packedTxs = getWithdrawPackedTxs(tmClient, state.LastHeight, newHeight)
		})
		if len(packedTxs) <= 0 {
			log.
				WithField("last_height", state.LastHeight).
				WithField("new_height", newHeight).
				Info("No withdraw transaction within range")
			state.LastHeight = newHeight
			state.save(statePath)
			continue
		}
		contractHeight := getContractHeight(lb, contractAddr)
		if contractHeight < newHeight {
			commitOk := commitWithdrawHash(tmClient, lb, auth, contractAddr, newHeight)
			if !commitOk {
				log.
					WithField("new_height", newHeight).
					Error("Commit withdraw hash failed for this height, skipping")
				continue
			}
		} else if contractHeight > newHeight {
			log.
				WithField("contract_height", contractHeight).
				WithField("new_height", newHeight).
				Error("New height is less than contract height")
		}
		wg := sync.WaitGroup{}
		wg.Add(len(packedTxs))
		nonce := eth.GetNonce(lb, auth.From)
		for _, packedTx := range packedTxs {
			go func(packedTx []byte, nonce int64) {
				doWithdraw(tmClient, lb, auth, contractAddr, packedTx, nonce)
				wg.Done()
			}(packedTx, nonce)
			nonce++
		}
		wg.Wait()
		state.LastHeight = newHeight
		state.save(statePath)
	}
}