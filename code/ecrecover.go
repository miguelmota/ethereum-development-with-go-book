// TODO
ec recover golang

func recoverSig(r, s, v *big.Int) ([]byte) {
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
}



func (c Constructor) signTx(tx *types.Transaction, s types.Signer, prv *ecdsa.PrivateKey) ([]byte, error) {
	h := s.Hash(tx)
	sig, err := ethCrypto.Sign(h[:], prv)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func recoverPlain(R, S, Vb *big.Int, homestead bool) ([]byte, error) {
	if Vb.BitLen() > 8 {
		return nil, ErrInvalidSig
	}
	V := byte(Vb.Uint64() - 27)
	if !crypto.ValidateSignatureValues(V, R, S, homestead) {
		return nil, ErrInvalidSig
	}
	// encode the snature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, 65)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V
    return sig, nil
}







package evmcompat

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	ssha "github.com/miguelmota/go-solidity-sha3"
)

type SignatureType uint8

const (
	SignatureType_EIP712 SignatureType = 0
	SignatureType_GETH   SignatureType = 1
	SignatureType_TREZOR SignatureType = 2
)

// SoliditySign signs the given data with the specified private key and returns the 65-byte signature.
// The signature is in a format that's compatible with the ecverify() Solidity function.
func SoliditySign(data []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(data, privKey)
	if err != nil {
		return nil, err
	}

	v := sig[len(sig)-1]
	sig[len(sig)-1] = v + 27
	return sig, nil
}

// SolidityRecover recovers the Ethereum address from the signed hash and the 65-byte signature.
func SolidityRecover(hash []byte, sig []byte) (common.Address, error) {
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes, got %d bytes", len(sig))
	}
	stdSig := make([]byte, 65)
	copy(stdSig[:], sig[:])
	stdSig[len(sig)-1] -= 27

	var signer common.Address
	pubKey, err := crypto.Ecrecover(hash, stdSig)
	if err != nil {
		return signer, err
	}

	copy(signer[:], crypto.Keccak256(pubKey[1:])[12:])
	return signer, nil
}

// GenerateTypedSig signs the given data with the specified private key and returns the 66-byte signature
// (the first byte of which is used to denote the SignatureType).
func GenerateTypedSig(data []byte, privKey *ecdsa.PrivateKey, sigType SignatureType) ([]byte, error) {
	if sigType != SignatureType_EIP712 {
		return nil, errors.New("signing failed, sig type not implemented")
	}

	sig, err := SoliditySign(data, privKey)
	if err != nil {
		return nil, err
	}
	// Prefix the sig with a single byte indicating the sig type, in this case EIP712
	typedSig := append(make([]byte, 0, 66), byte(SignatureType_EIP712))
	return append(typedSig, sig...), nil
}

// RecoverAddressFromTypedSig recovers the Ethereum address from a signed hash and a 66-byte signature
// (the first byte of which is expected to denote the SignatureType).
func RecoverAddressFromTypedSig(hash []byte, sig []byte) (common.Address, error) {
	var signer common.Address

	if len(sig) != 66 {
		return signer, fmt.Errorf("signature must be 66 bytes, not %d bytes", len(sig))
	}

	switch SignatureType(sig[0]) {
	case SignatureType_EIP712:
	case SignatureType_GETH:
		hash = ssha.SoliditySHA3(
			ssha.String("\x19Ethereum Signed Message:\n32"),
			ssha.Bytes32(hash),
		)
	case SignatureType_TREZOR:
		hash = ssha.SoliditySHA3(
			ssha.String("\x19Ethereum Signed Message:\n\x20"),
			ssha.Bytes32(hash),
		)
	default:
		return signer, fmt.Errorf("invalid signature type: %d", sig[0])
	}

	signer, err := SolidityRecover(hash, sig[1:])
	return signer, err
}

//TODO in future all interfaces and not do conversions from strings
type Pair struct {
	Type  string
	Value string
}

// NOTE: This function is deprecated, use the one in github.com/miguelmota/go-solidity-sha3 instead!
func SoliditySHA3(pairs []*Pair) ([]byte, error) {
	//convert to packed bytes like solidity
	data, err := SolidityPackedBytes(pairs)
	if err != nil {
		return nil, err
	}

	d := sha3.NewLegacyKeccak256()
	d.Write(data)
	return d.Sum(nil), nil
}

func SolidityPackedBytes(pairs []*Pair) ([]byte, error) {
	var b bytes.Buffer

	for _, pair := range pairs {
		fmt.Printf("%v\n", pair)
		switch strings.ToLower(pair.Type) {
		case "address":
			decoded, err := hex.DecodeString(pair.Value)
			if err != nil {
				return nil, err
			}
			if len(decoded) != 20 {
				return nil, fmt.Errorf("we don't support partial addresses, the len was %d we wanted 20", len(decoded))
			}
			b.Write(decoded)
		case "uint16": //"uint", "uint16", "uint64":
			//pack integers
			u, err := strconv.ParseUint(pair.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			var bTest []byte = make([]byte, 2)
			//			binary.LittleEndian.PutUint32(bTest, uint32(u))
			//			fmt.Printf("little-%v\n", bTest)
			binary.BigEndian.PutUint16(bTest, uint16(u))
			b.Write(bTest)
		case "uint32": //"uint", "uint16", "uint64":
			//pack integers
			u, err := strconv.ParseUint(pair.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			var bTest []byte = make([]byte, 4)
			//			binary.LittleEndian.PutUint32(bTest, uint32(u))
			//			fmt.Printf("little-%v\n", bTest)
			binary.BigEndian.PutUint32(bTest, uint32(u))
			b.Write(bTest)
		case "uint64": //"uint", "uint16", "uint64":
			//pack integers
			u, err := strconv.ParseUint(pair.Value, 10, 64)
			if err != nil {
				return nil, err
			}
			var bTest []byte = make([]byte, 8)
			//			binary.LittleEndian.PutUint32(bTest, uint32(u))
			//			fmt.Printf("little-%v\n", bTest)
			binary.BigEndian.PutUint64(bTest, u)
			b.Write(bTest)
		case "uint256":
			n := new(big.Int)
			_, valid := n.SetString(pair.Value, 10)
			if !valid {
				return nil, errors.New("invalid big int")
			}

			bytes := n.Bytes()
			padlen := 32 - len(bytes)
			if padlen < 0 {
				return nil, errors.New("big int byte length too large")
			}
			pad := make([]byte, padlen, padlen)
			b.Write(pad)
			b.Write(bytes)
		}
	}

	return b.Bytes(), nil
}




// +build evm

package evmcompat

import (
	"encoding/hex"
	"math/big"
	"testing"

	"golang.org/x/crypto/sha3"
	ssha "github.com/miguelmota/go-solidity-sha3"
)

/*
describe('solidity tight packing multiple arguments', function () {
  it('should equal', function () {
    var a = abi.solidityPack(
      [ 'bytes32', 'uint32', 'uint32', 'uint32', 'uint32' ],
      [ new Buffer('123456', 'hex'), 6, 7, 8, 9 ]
    )
    var b = '123456000000000000000000000000000000000000000000000000000000000000000006000000070000000800000009'
    assert.equal(a.toString('hex'), b.toString('hex'))
  })
})
*/
func TestSolidityPackedBytes(t *testing.T) {
	want := "0000000843989fb883ba8111221e8912389753847589383700000007"

	pairs := []*Pair{
		&Pair{"uint32", "8"},
		&Pair{"Address", "43989fb883ba8111221e89123897538475893837"},
		&Pair{"uint32", "7"},
	}

	g, err := SolidityPackedBytes(pairs)
	if err != nil {
		t.Errorf("TestSolidityPackedBytes failed got error %q", err)
	}
	got := hex.EncodeToString(g)

	if got != want {
		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
	}

	g = ssha.ConcatByteSlices(
		ssha.Uint32(8),
		ssha.Address("43989fb883ba8111221e89123897538475893837"),
		ssha.Uint32(7),
	)
	got = hex.EncodeToString(g)

	if got != want {
		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
	}

	wantsha3 := "5611aae8648e01a2e4721917fd1706014b8f4d387928e3cad536be41e5af4f77"

	g2, err := SoliditySHA3(pairs)
	if err != nil {
		t.Errorf("TestSolidityPackedBytes failed got error %q", err)
	}
	gotsha3 := hex.EncodeToString(g2)

	if gotsha3 != wantsha3 {
		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", gotsha3, wantsha3)
	}

	g2 = ssha.SoliditySHA3(g)
	gotsha3 = hex.EncodeToString(g2)
	if gotsha3 != wantsha3 {
		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", gotsha3, wantsha3)
	}
}

func TestSoliditySha3(t *testing.T) {
	want := "43989fb883ba8111221e8912389753847589383700000000000000000000000000000000000000002710564fe203"

	pairs := []*Pair{
		&Pair{"Address", "43989fb883ba8111221e89123897538475893837"},
		&Pair{"Address", "0000000000000000000000000000000000000000"},
		&Pair{"uint16", "10000"},
		&Pair{"uint32", "1448075779"},
	}

	g, err := SolidityPackedBytes(pairs)
	if err != nil {
		t.Errorf("TestSoliditySha3 failed got error %q", err)
	}
	got := hex.EncodeToString(g)

	if got != want {
		t.Errorf("TestSoliditySha3 failed got %q, want %q", got, want)
	}

	slices := [][]byte{
		ssha.Address("43989fb883ba8111221e89123897538475893837"),
		ssha.Address("0000000000000000000000000000000000000000"),
		ssha.Uint16(uint16(10000)),
		ssha.Uint32(uint32(1448075779)),
	}
	g = ssha.ConcatByteSlices(slices...)
	got = hex.EncodeToString(g)

	if got != want {
		t.Errorf("TestSoliditySha3 failed got %q, want %q", got, want)
	}

	wantsha3 := "7221df1d75e4baccbccd8a1fb33dbc5fca5f3c543e4acbb37c1b9edf990d3e1e"

	g2, err := SoliditySHA3(pairs)
	if err != nil {
		t.Errorf("TestSoliditySha3 failed got error %q", err)
	}
	gotsha3 := hex.EncodeToString(g2)

	if gotsha3 != wantsha3 {
		t.Errorf("TestSoliditySha3 failed got %q, want %q", gotsha3, wantsha3)
	}

	g2 = ssha.SoliditySHA3(slices...)
	gotsha3 = hex.EncodeToString(g2)

	if gotsha3 != wantsha3 {
		t.Errorf("TestSoliditySha3 failed got %q, want %q", gotsha3, wantsha3)
	}
}

func TestSolidityPackedBytesTypeAddress(t *testing.T) {
	want := "43989fb883ba8111221e89123897538475893837"
	pairs := []*Pair{
		&Pair{"Address", "43989fb883ba8111221e89123897538475893837"},
	}

	g, err := SolidityPackedBytes(pairs)
	if err != nil {
		t.Errorf("TestSolidityPackedBytesTypeAddress failed got error %q", err)
	}
	got := hex.EncodeToString([]byte(g))

	if got != want {
		t.Errorf("TestSolidityPackedBytesTypeAddress failed got %q, want %q", got, want)
	}

	g = ssha.Address("43989fb883ba8111221e89123897538475893837")
	got = hex.EncodeToString(g)

	if got != want {
		t.Errorf("TestSolidityPackedBytesTypeAddress failed got %q, want %q", got, want)
	}
}

func TestSolidityPackedUint16(t *testing.T) {
	want := "002a"

	pairs := []*Pair{&Pair{"uint16", "42"}}

	g, err := SolidityPackedBytes(pairs)
	if err != nil {
		t.Errorf("TestSolidityPackedBytes failed got error %q", err)
	}
	got := hex.EncodeToString([]byte(g))

	if got != want {
		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
	}

	g = ssha.Uint16(uint16(42))
	got = hex.EncodeToString(g)

	if got != want {
		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
	}
}

func TestSolidityPackedUint256(t *testing.T) {
	want := "000000000000000000000000000000000000000000000000000000000000002a"

	pairs := []*Pair{&Pair{"uint256", "42"}}

	g, err := SolidityPackedBytes(pairs)
	if err != nil {
		t.Errorf("TestSolidityPackedBytes failed got error %q", err)
	}
	got := hex.EncodeToString([]byte(g))

	if got != want {
		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
	}

	g = ssha.Uint256(big.NewInt(42))
	got = hex.EncodeToString(g)

	if got != want {
		t.Errorf("TestSolidityPackedBytes failed got %q, want %q", got, want)
	}
}

func TestSoliditySha3With256(t *testing.T) {
	want := "9f022fbbf24efa13621bbc6c2fc2f3b1f742d3320123acde9a25a9b5e25d81a9"

	pairs := []*Pair{
		&Pair{"uint256", "42"},
		&Pair{"Address", "32be343b94f860124dc4fee278fdcbd38c102d88"},
		&Pair{"Address", "74ff65739a88fdaf9675ff33405f760b53832ad7"},
		&Pair{"uint256", "52"},
	}

	g, err := SolidityPackedBytes(pairs)
	if err != nil {
		t.Errorf("TestSoliditySha3With256 failed got error %q", err)
	}
	if len(g) != 104 {
		t.Errorf("length unexpected")
	}

	d := sha3.NewLegacyKeccak256()
	d.Write(g)
	hash := d.Sum(nil)
	if hex.EncodeToString(hash) != want {
		t.Errorf("hashes don't match")
	}

	g = ssha.ConcatByteSlices(
		ssha.Uint256(big.NewInt(42)),
		ssha.Address("32be343b94f860124dc4fee278fdcbd38c102d88"),
		ssha.Address("74ff65739a88fdaf9675ff33405f760b53832ad7"),
		ssha.Uint256(big.NewInt(52)),
	)
	if len(g) != 104 {
		t.Errorf("length unexpected")
	}

	d = sha3.NewLegacyKeccak256()
	d.Write(g)
	hash = d.Sum(nil)
	if hex.EncodeToString(hash) != want {
		t.Errorf("hashes don't match")
	}
}

func TestAnotherSoliditySha3With256(t *testing.T) {
	want := "00000000000000000000000000000000000000000000000000000000564fe20343989fb883ba8111221e8912389753847589383766989fb883ba8111221e8912389753847589386700000000000000000000000000000000000000000000000000000000564fe203"

	pairs := []*Pair{
		&Pair{"uint256", "1448075779"},
		&Pair{"Address", "43989fb883ba8111221e89123897538475893837"},
		&Pair{"Address", "66989fb883ba8111221e89123897538475893867"},
		&Pair{"uint256", "1448075779"},
	}

	g, err := SolidityPackedBytes(pairs)
	if err != nil {
		t.Errorf("TestSoliditySha3With256 failed got error %q", err)
	}
	if len(g) != 104 {
		t.Errorf("length unexpected")
	}
	got := hex.EncodeToString([]byte(g))
	if want != got {
		t.Errorf("hashes don't match -\n%s\n%s", want, g)
	}

	g = ssha.ConcatByteSlices(
		ssha.Uint256(big.NewInt(1448075779)),
		ssha.Address("43989fb883ba8111221e89123897538475893837"),
		ssha.Address("66989fb883ba8111221e89123897538475893867"),
		ssha.Uint256(big.NewInt(1448075779)),
	)
	if len(g) != 104 {
		t.Errorf("length unexpected")
	}
	got = hex.EncodeToString(g)
	if want != got {
		t.Errorf("hashes don't match -\n%s\n%s", want, g)
	}
}

func TestAnotherSoliditySha3WithUnit64(t *testing.T) {
	want := "920ae4155769cd69c30626f054134b5f003772473f57f84837402df6d166e663"

	pairs := []*Pair{
		&Pair{"uint32", "5"},
	}

	g, err := SoliditySHA3(pairs)
	if err != nil {
		t.Errorf("TestSoliditySha3With256 failed got error %q", err)
	}
	got := hex.EncodeToString([]byte(g))
	if want != got {
		t.Errorf("hashes don't match -\n%s\n%s", want, got)
	}

	g = ssha.SoliditySHA3(ssha.Uint32(uint32(5)))
	got = hex.EncodeToString(g)
	if want != got {
		t.Errorf("hashes don't match -\n%s\n%s", want, got)
	}

	want2 := "fe07a98784cd1850eae35ede546d7028e6bf9569108995fc410868db775e5e6a"

	pairs2 := []*Pair{
		&Pair{"uint64", "5"},
	}

	g2, err := SoliditySHA3(pairs2)
	if err != nil {
		t.Errorf("TestSoliditySha3With256 failed got error %q", err)
	}
	got2 := hex.EncodeToString([]byte(g2))
	if want2 != got2 {
		t.Errorf("hashes don't match -\n%s\n%s", want2, got2)
	}

	g2 = ssha.SoliditySHA3(ssha.Uint64(uint64(5)))
	got2 = hex.EncodeToString(g2)
	if want2 != got2 {
		t.Errorf("hashes don't match -\n%s\n%s", want2, got2)
	}
}
