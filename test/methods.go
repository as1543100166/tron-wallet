package test

import (
	tronWallet "github.com/ranjbar-dev/tron-wallet"
	"github.com/ranjbar-dev/tron-wallet/enums"
)

var node = enums.MAIN_NODE
var validPrivateKey = "ee01b1139022f6cecf76c57745227f2cd55583f2b6d18286d60b7b32e0a6111e"
var invalidPrivateKey = "invalid"
var validOwnerAddress = "TK729QYqVuRmPxWB12SBdnZhNYPfoE44Z2"
var invalidOwnerAddress = "T2w5FSuWhxcaJmBUVFY93UY4ihwx55668b6"
var validToAddress = "TJnsY5bGiwuPCQFismQDwyVTPAn7M88888"
var invalidToAddress = "TJnsY5bGiwuPCQQDwyVTPAnM88888"
var trxAmount int64 = 10000
var trc20Amount int64 = 10000

func wallet() *tronWallet.TronWallet {
	w, _ := tronWallet.CreateTronWallet(node, validPrivateKey)
	return w
}

func token() *tronWallet.Token {
	return &tronWallet.Token{
		ContractAddress: enums.MAIN_Tether_USDT,
	}
}

func crawler() *tronWallet.Crawler {
	return &tronWallet.Crawler{
		Node:      node,
		Addresses: []string{validOwnerAddress},
	}
}
