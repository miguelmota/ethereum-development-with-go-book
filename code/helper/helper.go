package ethereum

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"crypto/ecdsa"

	"github.com/coincircle/exchange/ethereum/token"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
 * Helper functions interact with the blockchain.
 */

// Service service
type Service struct {
	Client *ethclient.Client
}

// Options service options
type Options struct {
	ProviderURI string
}

// New returns new service
func New(opts *Options) (*Service, error) {
	if opts.ProviderURI == "" {
		return nil, errors.New("ethereum provider uri is required")
	}
	client, err := ethclient.Dial(opts.ProviderURI)
	if err != nil {
		return nil, err
	}
	return &Service{
		Client: client,
	}, nil
}

// GetAccountBalance get account balance
func (s *Service) GetAccountBalance(address string) (*big.Int, error) {
	accountAddress := common.HexToAddress(address)
	rawBalance, err := s.Client.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		return nil, err
	}

	return rawBalance, nil
}

// GetTokenBalance get token balance
func (s *Service) GetTokenBalance(_tokenAddress string, _accountAddress string) (*big.Int, error) {
	var bal *big.Int
	tokenAddress := common.HexToAddress(_tokenAddress)
	accountAddress := common.HexToAddress(_accountAddress)
	instance, err := token.NewTokenCaller(tokenAddress, s.Client)
	if err != nil {
		return bal, err
	}

	bal, err = instance.BalanceOf(&bind.CallOpts{Pending: false}, accountAddress)
	if err != nil {
		return bal, err
	}

	return bal, nil
}

// GetTokenDecimals get token decimals
func (s *Service) GetTokenDecimals(tokenAddress string) (*big.Int, error) {
	var decimals *big.Int
	instance, err := token.NewTokenCaller(common.HexToAddress(tokenAddress), s.Client)
	if err != nil {
		return decimals, err
	}

	decimalsInt8, err := instance.Decimals(&bind.CallOpts{})
	decimals = big.NewInt(int64(decimalsInt8))
	if err != nil {
		return decimals, err
	}

	return decimals, nil
}

// TransferEth transfer ETH to an address
func (s *Service) TransferEth(privateKey string, _toAddress string, amount *big.Int) (*types.Transaction, error) {
	toAddress := common.HexToAddress(_toAddress)

	chainID := big.NewInt(1)
	signer := types.NewEIP155Signer(chainID)
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return &types.Transaction{}, err
	}
	fromAddress, err := s.GetPublicAddressFromPrivateKey(key)
	if err != nil {
		return &types.Transaction{}, err
	}
	nonce, err := s.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return &types.Transaction{}, err
	}
	gasLimit := uint64(121000) // standard limit for sending
	gasPrice, err := s.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return &types.Transaction{}, err
	}

	tx, err := types.SignTx(types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil), signer, key)
	if err != nil {
		return tx, err
	}

	err = s.SendTx(tx)
	if err != nil {
		return tx, err
	}

	return tx, nil
}

// TransferTokens transfer tokens to an address
func (s *Service) TransferTokens(auth bind.TransactOpts, _tokenAddress string, _toAddress string, amount *big.Int) (*types.Transaction, error) {
	tokenAddress := common.HexToAddress(_tokenAddress)
	toAddress := common.HexToAddress(_toAddress)
	instance, err := token.NewToken(tokenAddress, s.Client)
	if err != nil {
		return &types.Transaction{}, err
	}

	tx, err := instance.Transfer(&auth, toAddress, amount)
	if err != nil {
		return tx, err
	}

	return tx, nil
}

// SendTx send a transaction to the network
func (s *Service) SendTx(tx *types.Transaction) error {
	err := s.Client.SendTransaction(context.Background(), tx)
	if err != nil {
		return err
	}

	return nil
}

// SignTx sign a transaction with a private key
func (s *Service) SignTx(nonce uint64, _toAddress string, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, privateKey string) (*types.Transaction, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return &types.Transaction{}, err
	}

	toAddress := common.HexToAddress(_toAddress)
	chainID := big.NewInt(1)
	signer := types.NewEIP155Signer(chainID)

	if gasPrice == nil {
		gasPrice, err = s.Client.SuggestGasPrice(context.Background())

		if err != nil {
			return &types.Transaction{}, err
		}
	}

	tx, err := types.SignTx(types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, data), signer, key)
	if err != nil {
		return &types.Transaction{}, err
	}

	return tx, nil
}

// TransferTokensTxData generate transaction data for transfer token call
func (s *Service) TransferTokensTxData(_toAddress string, amount *big.Int) ([]byte, error) {
	toAddress := common.HexToAddress(_toAddress)
	parsed, err := abi.JSON(strings.NewReader(token.TokenABI))
	if err != nil {
		return nil, err
	}

	data, err := parsed.Pack("transfer", toAddress, amount)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetLatestBlockNumber get the latest block number
func (s *Service) GetLatestBlockNumber() (*big.Int, error) {
	block, err := s.Client.HeaderByNumber(context.Background(), nil)
	return block.Number, err
}

// GetPublicAddressFromPrivateKey returns public address from private key
func (s *Service) GetPublicAddressFromPrivateKey(priv *ecdsa.PrivateKey) (common.Address, error) {
	var address common.Address
	pub := priv.Public()
	pubECDSA, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return address, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address = crypto.PubkeyToAddress(*pubECDSA)
	return address, nil
}

// GetGasPrice gets clamped gas price
func (s *Service) GetGasPrice() *big.Int {
	maxGasPrice := big.NewInt(90000000000)     // 90 gwei
	defaultGasPrice := big.NewInt(20000000000) // 20 gwei
	gasPrice, err := s.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return defaultGasPrice
	}

	// cap gas price in case SuggestGasPrice goes off the rails
	if gasPrice.Cmp(maxGasPrice) == 1 {
		return maxGasPrice
	}

	return gasPrice
}
