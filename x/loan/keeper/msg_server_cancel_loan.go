package keeper

import (
	"context"

	"loan/x/loan/types"
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CancelLoan(goCtx context.Context, msg *types.MsgCancelLoan) (*types.MsgCancelLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, found := k.GetLoan(ctx, msg.Id)
	if !found{
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.Borrower != msg.Creator{
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot cancel: not the borrower")
	}
	if loan.State != "requested"{
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	borrower, _ := sdk.AccAddressFromBech32(msg.Creator)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil{
		return nil, err
	}
	loan.State = "cancelled"
	k.SetLoan(ctx, loan)

	return &types.MsgCancelLoanResponse{}, nil
}
