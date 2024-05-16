package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRepayLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLiquidateLoan{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelLoan{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
