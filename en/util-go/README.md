# Collection of Utility Functions

The utility functions' implementation are found below in the [full code](#full-code) section.

Here are examples of usage:

```go
valid := util.IsValidAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
fmt.Println(valid) // true
```

```go
zeroed := util.IsZeroAddress("0x0")
fmt.Println(zeroed) // true
```

```go
wei := util.ToWei(0.02, 18)
fmt.Println(wei) // 20000000000000000
```

```go
wei := new(big.Int)
wei.SetString("20000000000000000", 10)
eth := util.ToDecimal(wei, 18)
fmt.Println(eth) // 0.02
```

```go
gasLimit := uint64(21000)
gasPrice := new(big.Int)
gasPrice.SetString("2000000000", 10)
gasCost := util.CalcGasCost(gasLimit, gasPrice)
fmt.Println(gasCost) // 42000000000000
```

```go
sig := "0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301"
r, s, v := util.SigRSV(sig)
fmt.Println(hexutil.Encode(r[:])[2:]) // 789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c6
fmt.Println(hexutil.Encode(s[:])[2:]) // 2621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde023
fmt.Println(v)                        // 28
```

---

### Full code

[util.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/util/util.go)

```go
package util

import (
	"math/big"
	"reflect"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
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

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV signatures R S V returned as arrays
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}
```

test file: [util_test.go](https://github.com/miguelmota/ethereum-development-with-go-book/blob/master/code/util/util_test.go)
