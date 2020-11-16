package types

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
)

type Message struct {
	Version    uint64          `json:"Version"`
	To         address.Address `json:"To"`
	From       address.Address `json:"From"`
	Nonce      uint64          `json:"Nonce"`
	Value      abi.TokenAmount `json:"Value"`
	GasLimit   int64           `json:"GasLimit"`
	GasFeeCap  abi.TokenAmount `json:"GasFeeCap"`
	GasPremium abi.TokenAmount `json:"GasPremium"`
	Method     uint64          `json:"Method"`
	Params     []byte          `json:"Params"`
}
