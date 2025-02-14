package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/jamescrystal/mychain/testutil/keeper"
	"github.com/jamescrystal/mychain/x/stablecoin/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.StablecoinKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
