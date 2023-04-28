package util

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	SUCCESS                = 1
	FAILED                 = 0
	PENDING                = 2
	TXTYPE_REWARD_SHIPPING = 3
	TXTYPE_SELLER_WITHDRAW = 4
	TIMEOUT                = 5
)

type Channel struct {
	NewTxBroadcast chan TxInfo
	Seller         chan *ReceiptSt
	Receipt        chan *ReceiptSt
}

type TxInfo struct {
	TxHash  string
	Type    int
	Holders []common.Address
}

type ReceiptSt struct {
	Code    int
	Receipt *types.Receipt
	Hodlers []common.Address
}

func NewChannel() *Channel {

	newTxBroadcast_chan := make(chan TxInfo)
	seller_chan := make(chan *ReceiptSt)
	receipt_chan := make(chan *ReceiptSt)
	return &Channel{
		newTxBroadcast_chan,
		seller_chan,
		receipt_chan,
	}
}
