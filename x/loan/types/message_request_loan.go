package types

import (
	"strconv"

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
	amount, _ := sdk.ParseCoinsNormalized(msg.Amount)
	if !amount.IsValid(){
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coins object")
	}
	if amount.Empty(){
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount is empty")
	}
	fee, _ := sdk.ParseCoinsNormalized(msg.Fee)
	if !fee.IsValid(){
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "fee is not a valid Coins object")
	}
	deadline, err := strconv.ParseInt(msg.Deadline, 10, 64)
	if err != nil{
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "deadline is not an integer")
	}
	if deadline <= 0{
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "deadline should be positive integer")
	}
	collateral, _ := sdk.ParseCoinsNormalized(msg.Collateral)
	if !collateral.IsValid(){
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collateral is not a valid Coins object")
	}
	if collateral.Empty(){
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collateral is empty")
	}
	return nil
}
