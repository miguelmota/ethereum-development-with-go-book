// +build integration

package ethereum

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var s *Service

func init() {
	service, err := New(&Options{
		ProviderURI: "wss://mainnet.infura.io/ws",
	})
	if err != nil {
		panic(err)
	}

	s = service
}

func TestGetAccountBalance(t *testing.T) {
	t.Parallel()
	accountAddress := "0x85E4B84D784eE9eEB7489F0B0c66B343AF2a0BE5"
	balance, err := s.GetAccountBalance(accountAddress)
	if err != nil {
		t.FailNow()
	}

	if balance.Cmp(big.NewInt(0)) != 1 {
		t.FailNow()
	}
}

func TestGetTokenBalance(t *testing.T) {
	t.Parallel()
	accountAddress := "0x85E4B84D784eE9eEB7489F0B0c66B343AF2a0BE5"
	// BAT
	tokenAddress := "0x0d8775f648430679a709e98d2b0cb6250d2887ef"
	balance, err := s.GetTokenBalance(tokenAddress, accountAddress)
	if err != nil {
		t.FailNow()
	}

	if balance.Cmp(big.NewInt(0)) != 1 {
		t.FailNow()
	}
}

func TestGetTokenAllowance(t *testing.T) {
	t.Parallel()
	accountAddress := "0x85E4B84D784eE9eEB7489F0B0c66B343AF2a0BE5"
	// BAT
	tokenAddress := "0x0d8775f648430679a709e98d2b0cb6250d2887ef"
	balance, err := s.GetTokenBalance(tokenAddress, accountAddress)
	if err != nil {
		t.FailNow()
	}

	if balance.Cmp(big.NewInt(0)) != 1 {
		t.FailNow()
	}
}

func TestGetLastestBlockNumber(t *testing.T) {
	t.Parallel()
	blockNumber, err := s.GetLatestBlockNumber()
	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	if blockNumber.Cmp(big.NewInt(4000000)) != 1 {
		t.Errorf("Expected latest block number to be larger than 4,000,000, instead got %s", blockNumber)
	}
}

func TestTransferEth(t *testing.T) {
	t.Parallel()
	t.Skip("Skipping TransferEth")
	amount := big.NewInt(0)
	// testrpc
	privateKey := "7e96f2011b1ec4f470bc6df5c195cd24869546d41b9fd7dc2a6fe9867d4b9f8e"
	toAddress := "0xe5b072d5320dcf3ee3aae76f29704a09dd6fce5e"
	tx, err := s.TransferEth(privateKey, toAddress, amount)
	if err != nil {
		t.Errorf("Error transfering eth, got error %s:", err)
	}

	t.Log(tx)
}

func TestTransferTokens(t *testing.T) {
	t.Parallel()
	// TODO
}

func TestTransferTokensTxData(t *testing.T) {
	t.Parallel()
	//t.Skip("Skipping TransferTokensTxData")
	amount := big.NewInt(0)
	// testrpc
	toAddress := "0xe5b072d5320dcf3ee3aae76f29704a09dd6fce5e"
	tx, err := s.TransferTokensTxData(toAddress, amount)
	if err != nil {
		t.Errorf("Got error %s:", err)
	}

	t.Log(tx)
}

func TestSignTx(t *testing.T) {
	t.Parallel()
	t.Skip("Skipping SignTx")
	// testrpc
	privateKey := "950ef72991706f0f651819372508d56fd5a71b43a5124de77c1dfe37f0b0bb3c"
	fromAddress := "0x785fcea75ae5a82153ff4f2250748d2925e31c6e"
	toAddress := "0xe5b072d5320dcf3ee3aae76f29704a09dd6fce5e"
	// BAT token
	tokenAddress := "0x0d8775f648430679a709e98d2b0cb6250d2887ef"
	tokenAmount := big.NewInt(0)
	data, err := s.TransferTokensTxData(toAddress, tokenAmount)
	if err != nil {
		t.Errorf("Got error %s:", err)
	}
	amount := big.NewInt(0)
	nonce, err := s.Client.PendingNonceAt(context.Background(), common.HexToAddress(fromAddress))
	if err != nil {
		t.Errorf("Got error %s:", err)
	}
	gasLimit := uint64(121000)
	tx, err := s.SignTx(nonce, tokenAddress, amount, gasLimit, nil, data, privateKey)
	if err != nil {
		t.Errorf("Got error %s:", err)
	}

	t.Log(tx)
}

func TestSendTx(t *testing.T) {
	t.Parallel()
	t.Skip("Skipping SendTx")
	// testrpc
	privateKey := "950ef72991706f0f651819372508d56fd5a71b43a5124de77c1dfe37f0b0bb3c"
	fromAddress := "0x785fcea75ae5a82153ff4f2250748d2925e31c6e"
	toAddress := "0xe5b072d5320dcf3ee3aae76f29704a09dd6fce5e"
	// BAT token
	tokenAddress := "0x0d8775f648430679a709e98d2b0cb6250d2887ef"
	tokenAmount := big.NewInt(0)
	tokenAmount.SetString("1000000000000000000", 10)
	data, err := s.TransferTokensTxData(toAddress, tokenAmount)
	if err != nil {
		t.Errorf("Got error %s:", err)
	}
	amount := big.NewInt(0)
	nonce, err := s.Client.PendingNonceAt(context.Background(), common.HexToAddress(fromAddress))
	if err != nil {
		t.Errorf("Got error %s:", err)
	}
	gasLimit := uint64(121000)
	tx, err := s.SignTx(nonce, tokenAddress, amount, gasLimit, nil, data, privateKey)
	if err != nil {
		t.Errorf("Got error %s:", err)
	}

	err = s.SendTx(tx)
	if err != nil {
		t.Errorf("Got error %s:", err)
	}

	t.Log(tx)
}

func TestGetTokenDecimals(t *testing.T) {
	t.Parallel()
	// BAT token
	tokenAddress := "0x0d8775f648430679a709e98d2b0cb6250d2887ef"
	decimals, err := s.GetTokenDecimals(tokenAddress)
	if err != nil {
		t.Errorf("Got error: %s", err)
	}

	expected := big.NewInt(18)
	if decimals.Cmp(expected) != 0 {
		t.Errorf("Expected %s, got %s", expected, decimals)
	}
}

func TestGetPublicAddressFromPrivateKey(t *testing.T) {
	t.Parallel()
	privStr := "4da37b3d296c29f4d56b8b46f8189292adfb7e14d24958bfded5dae50ab32039"
	priv, err := crypto.HexToECDSA(privStr)
	if err != nil {
		t.FailNow()
	}

	expectedAddress := common.HexToAddress("0x107e5320964430027a442fc5cc85a972adaa6f52")
	address, err := s.GetPublicAddressFromPrivateKey(priv)
	if err != nil {
		t.FailNow()
	}
	if address.Hex() != expectedAddress.Hex() {
		t.FailNow()
	}
}

func TestGetGasPrice(t *testing.T) {
	t.Parallel()
	min := big.NewInt(0)
	max := big.NewInt(0)
	max.SetString("90000000000", 10)
	gasPrice := s.GetGasPrice()
	if gasPrice.Cmp(min) != 1 {
		t.FailNow()
	}

	if gasPrice.Cmp(max) != -1 {
		t.FailNow()
	}
}
