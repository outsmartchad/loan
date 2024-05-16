package keeper

import (
	"context"
	"encoding/binary"

	"loan/x/loan/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetLoanCount get the total number of loan
func (k Keeper) GetLoanCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LoanCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLoanCount set the total number of loan
func (k Keeper) SetLoanCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LoanCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLoan appends a loan in the store with a new id and update the count
func (k Keeper) AppendLoan(
	ctx context.Context,
	loan types.Loan,
) uint64 {
	// Create the loan
	count := k.GetLoanCount(ctx)

	// Set the ID of the appended value
	loan.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LoanKey))
	appendedValue := k.cdc.MustMarshal(&loan)
	store.Set(GetLoanIDBytes(loan.Id), appendedValue)

	// Update loan count
	k.SetLoanCount(ctx, count+1)

	return count
}

// SetLoan set a specific loan in the store
func (k Keeper) SetLoan(ctx context.Context, loan types.Loan) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LoanKey))
	b := k.cdc.MustMarshal(&loan)
	store.Set(GetLoanIDBytes(loan.Id), b)
}

// GetLoan returns a loan from its id
func (k Keeper) GetLoan(ctx context.Context, id uint64) (val types.Loan, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LoanKey))
	b := store.Get(GetLoanIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLoan removes a loan from the store
func (k Keeper) RemoveLoan(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LoanKey))
	store.Delete(GetLoanIDBytes(id))
}

// GetAllLoan returns all loan
func (k Keeper) GetAllLoan(ctx context.Context) (list []types.Loan) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LoanKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Loan
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLoanIDBytes returns the byte representation of the ID
func GetLoanIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.LoanKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
