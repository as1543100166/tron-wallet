package grpcClient

import (
	"bytes"
	"fmt"

	"github.com/ranjbar-dev/tron-wallet/grpcClient/proto/api"
	"github.com/ranjbar-dev/tron-wallet/grpcClient/proto/core"
	"github.com/ranjbar-dev/tron-wallet/util"
	"google.golang.org/protobuf/proto"
)

func (g *GrpcClient) GetAccount(addr string) (*core.Account, error) {
	account := new(core.Account)
	var err error

	account.Address, err = util.DecodeCheck(addr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := g.getContext()
	defer cancel()

	acc, err := g.Client.GetAccount(ctx, account)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(acc.Address, account.Address) {
		return nil, fmt.Errorf("account not found")
	}
	return acc, nil
}

func (g *GrpcClient) GetAccountResource(addr string) (*api.AccountResourceMessage, error) {
	account := new(core.Account)
	var err error

	account.Address, err = util.DecodeCheck(addr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := g.getContext()
	defer cancel()

	acc, err := g.Client.GetAccountResource(ctx, account)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

// CreateAccount activate tron account
func (g *GrpcClient) CreateAccount(from, addr string) (*api.TransactionExtention, error) {
	var err error

	contract := &core.AccountCreateContract{}
	if contract.OwnerAddress, err = util.DecodeCheck(from); err != nil {
		return nil, err
	}
	if contract.AccountAddress, err = util.DecodeCheck(addr); err != nil {
		return nil, err
	}
	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.CreateAccount2(ctx, contract)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("bad transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return tx, nil
}
