package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLiquidateLoan{}

func NewMsgLiquidateLoan(creator string, id uint64) *MsgLiquidateLoan {
	return &MsgLiquidateLoan{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgLiquidateLoan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
