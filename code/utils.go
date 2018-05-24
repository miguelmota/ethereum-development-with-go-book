package ethereum

import (
	"math/big"
	"reflect"
	"regexp"
	"strconv"

	types "github.com/coincircle/exchange/ethereum/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// IsValidAddress validate hex address
func IsValidAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(addressIfc interface{}) bool {
	var address common.Address
	switch v := addressIfc.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// WeiToDecimals wei to decimals
func WeiToDecimals(n *big.Int, decimals int) decimal.Decimal {
	mul := decimal.NewFromFloat(float64(0.1)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(n.String())
	result := num.Mul(mul)

	return result
}

// DecimalsToWei decimals to wei
func DecimalsToWei(amount decimal.Decimal, decimals int) *big.Int {
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// EthToWei eth to wei
func EthToWei(amount decimal.Decimal) *big.Int {
	wei := DecimalsToWei(amount, 18)
	return wei
}

// GetSigRSV signatures R S V  returned as strings
func GetSigRSV(sig []byte) *types.RSV {
	R := "0x" + common.Bytes2Hex(sig)[0:64]
	S := "0x" + common.Bytes2Hex(sig)[64:128]
	vStr := common.Bytes2Hex(sig)[128:130]
	V, _ := strconv.Atoi(vStr)
	V = V + 27
	rsv := &types.RSV{R: R, S: S, V: V}
	return rsv
}

// GetSigRSVBytes signatures R S V returned as arrays
func GetSigRSVBytes(sig []byte) ([32]byte, [32]byte, uint8) {
	sigstr := common.Bytes2Hex(sig)
	_R := sigstr[0:64]
	_S := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(_R))
	copy(S[:], common.FromHex(_S))
	vStr := sigstr[128:130]
	_V, _ := strconv.Atoi(vStr)
	V := uint8(_V + 27)

	return R, S, V
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}
