package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "loan/testutil/keeper"
	"loan/x/loan/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LoanKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
