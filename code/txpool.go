package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
)

type testBlockChain struct {
	statedb       *state.StateDB
	gasLimit      uint64
	chainHeadFeed *event.Feed
}

func (bc *testBlockChain) CurrentBlock() *types.Block {
	return types.NewBlock(&types.Header{
		GasLimit: bc.gasLimit,
	}, nil, nil, nil)
}

func (bc *testBlockChain) GetBlock(hash common.Hash, number uint64) *types.Block {
	return bc.CurrentBlock()
}

func (bc *testBlockChain) StateAt(common.Hash) (*state.StateDB, error) {
	return bc.statedb, nil
}

func (bc *testBlockChain) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return bc.chainHeadFeed.Subscribe(ch)
}

func main() {
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(ethdb.NewMemDatabase()))
	blockchain := &testBlockChain{statedb, 1000000, new(event.Feed)}

	txpool := core.NewTxPool(core.TxPoolConfig{}, params.RinkebyChainConfig, blockchain)

	pendingTxs, queuedTxs := txpool.Content()

	spew.Dump(pendingTxs)
	spew.Dump(queuedTxs)

	pendingTxs, err = txpool.Pending()
	if err != nil {
		panic(err)
	}

	spew.Dump(pendingTxs)

	pendingTxs, err = txpool.Pending()
	if err != nil {
		panic(err)
	}

	spew.Dump(pendingTxs)

	tx := types.Transaction{} // types.NewTransaction(...)
	if err := txpool.AddRemote(tx); err != nil {
		panic(err)
	}
}
