// +build unit

package ethereum

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

func TestIsValidAddress(t *testing.T) {
	t.Parallel()
	validAddress := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	invalidAddress := "0xabc"
	invalidAddress2 := "323b5d4c32345ced77393b3530b1eed0f346429d"
	{
		got := IsValidAddress(validAddress)
		expected := true

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}

	{
		got := IsValidAddress(invalidAddress)
		expected := false

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}

	{
		got := IsValidAddress(invalidAddress2)
		expected := false

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}
}

func TestIsZeroAddress(t *testing.T) {
	t.Parallel()
	validAddress := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	zeroAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")

	// Test common.Address input

	{
		isZeroAddress := IsZeroAddress(validAddress)

		if isZeroAddress {
			t.Error("Expected to be false")
		}
	}

	{
		isZeroAddress := IsZeroAddress(zeroAddress)

		if !isZeroAddress {
			t.Error("Expected to be true")
		}
	}

	// Test string input

	{
		isZeroAddress := IsZeroAddress(validAddress.Hex())

		if isZeroAddress {
			t.Error("Expected to be false")
		}
	}

	{
		isZeroAddress := IsZeroAddress(zeroAddress.Hex())

		if !isZeroAddress {
			t.Error("Expected to be true")
		}
	}
}

func TestGetSigRSV(t *testing.T) {
	t.Parallel()
	sig, err := hex.DecodeString("1f4ab7e26711f235331edc67bd697fd0c7628dd5ffcab333870640dee329914b2bce958fb3ee54817b1d5102e364a9164f46f732f4a02a9d5cd9569b085f211200")
	if err != nil {
		t.Errorf("Got error %s", err)
	}

	expectedR := "0x1f4ab7e26711f235331edc67bd697fd0c7628dd5ffcab333870640dee329914b"
	expectedS := "0x2bce958fb3ee54817b1d5102e364a9164f46f732f4a02a9d5cd9569b085f2112"
	expectedV := 27
	rsv := GetSigRSV(sig)
	if expectedR != rsv.R {
		t.Errorf("Expected %s, got %s", expectedR, rsv.R)
	}

	if expectedS != rsv.S {
		t.Errorf("Expected %s, got %s", expectedS, rsv.S)
	}

	if expectedV != rsv.V {
		t.Errorf("Expected %v, got %v", expectedV, rsv.V)
	}
}

func TestGetSigRSVBytes(t *testing.T) {
	t.Parallel()
	// TODO
}

func TestDecimalsToWei(t *testing.T) {
	t.Parallel()
	amount := decimal.NewFromFloat(0.02)
	got := DecimalsToWei(amount, 18)
	expected := new(big.Int)
	expected.SetString("20000000000000000", 10)
	if got.Cmp(expected) != 0 {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestEthToWei(t *testing.T) {
	t.Parallel()
	amountInEth := decimal.NewFromFloat(0.02)
	got := EthToWei(amountInEth)
	expected := new(big.Int)
	expected.SetString("20000000000000000", 10)
	if got.Cmp(expected) != 0 {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestWeiToDecimals(t *testing.T) {
	t.Parallel()
	weiAmount := big.NewInt(0)
	weiAmount.SetString("20000000000000000", 10)
	ethAmount := WeiToDecimals(weiAmount, int(18))
	expected := decimal.NewFromFloat(0.02)
	if !ethAmount.Equals(expected) {
		t.Error("%s does not equal expected %s", ethAmount, expected)
	}
}

func TestCalcGasLimit(t *testing.T) {
	t.Parallel()
	gasPrice := big.NewInt(0)
	gasPrice.SetString("2000000000", 10)
	gasLimit := uint64(21000)
	expected := big.NewInt(0)
	expected.SetString("42000000000000", 10)
	gasCost := CalcGasCost(gasLimit, gasPrice)
	if gasCost.Cmp(expected) != 0 {
		t.Error("expected %s, got %s", gasCost, expected)
	}
}
