package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestLoan{}

func NewMsgRequestLoan(creator string, amount string, fee string, collateral string, deadline string) *MsgRequestLoan {
	return &MsgRequestLoan{
		Creator:    creator,
		Amount:     amount,
		Fee:        fee,
		Collateral: collateral,
		Deadline:   deadline,
	}
}

func (msg *MsgRequestLoan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
