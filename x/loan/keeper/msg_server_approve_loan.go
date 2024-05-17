package keeper

import (
	"context"

	"loan/x/loan/types"
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found{
		// cannot find the loan in the blockchain
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.Borrower == msg.Creator{
		// should not be the same one to lend and borrow
		return nil, errorsmod.Wrap(sdkerrors.ErrConflict, "the borrower and the lender should not be the same one")
	}
	if loan.State != "requested"{
		// check if the state is requested first
		return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	amt, err := sdk.ParseCoinsNormalized(loan.Amount)	
	if err != nil{
		return nil, errorsmod.Wrap(types.ErrWrongLoanState, "Cannot parse coins in loan amount")
	}
	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amt)
	if err != nil{
		return nil, err
	}
	loan.Lender = msg.Creator
	loan.State = "approved"
	k.SetLoan(ctx, loan) 
	return &types.MsgApproveLoanResponse{}, nil
}
