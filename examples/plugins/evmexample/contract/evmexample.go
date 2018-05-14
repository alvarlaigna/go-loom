// +build evm

package main

import (
	"github.com/loomnetwork/go-loom/examples/plugins/evmexample/types"
	"github.com/loomnetwork/go-loom/plugin"
	"github.com/loomnetwork/go-loom/plugin/contractpb"
)

type EvmExample struct {
}

func (c *EvmExample) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "EvmExample",
		Version: "1.0.0",
	}, nil
}

func (c *EvmExample) SetValue(ctx contractpb.Context, tx *types.EthTransaction) error {
	simpleStoreAddr, err := ctx.Resolve("SimpleStore")
	if err != nil {
		return err
	}
	evmOut := []byte{}
	err = contractpb.CallEVM(ctx, simpleStoreAddr, tx.Data, &evmOut)
	return err
}

func (c *EvmExample) GetValue(ctx contractpb.Context, tx *types.EthCall) (*types.EthCallResult, error) {
	simpleStoreAddr, err := ctx.Resolve("SimpleStore")
	if err != nil {
		return nil, err
	}
	evmOut := []byte{}
	err = contractpb.CallEVM(ctx, simpleStoreAddr, tx.Data, &evmOut)
	if err != nil {
		return nil, err
	}
	return &types.EthCallResult{
		Data: evmOut,
	}, err
}

var Contract = contractpb.MakePluginContract(&EvmExample{})

func main() {
	plugin.Serve(Contract)
}
