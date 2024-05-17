package keeper

import (
    "context"
    "strconv"

    errorsmod "cosmossdk.io/errors"
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

    "loan/x/loan/types"
)

func (k msgServer) LiquidateLoan(goCtx context.Context, msg *types.MsgLiquidateLoan) (*types.MsgLiquidateLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.Lender != msg.Creator{
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot liquidate: not the lender")
	}
	if loan.State != "approved"{
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	deadline, err := strconv.ParseInt(loan.Deadline, 10, 64)
	if err != nil{
		panic(err)
	}
	if ctx.BlockHeight() < deadline {
		return nil, errorsmod.Wrap(types.ErrDeadline, "Cannot liquidate before deadline")
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, collateral)
	if err != nil{
		return nil, err
	}
	loan.State = "liquidated"
	k.SetLoan(ctx, loan)
	

	return &types.MsgLiquidateLoanResponse{}, nil
}
