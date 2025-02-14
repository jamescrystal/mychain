package stablecoin_test

import (
	"testing"

	keepertest "github.com/jamescrystal/mychain/testutil/keeper"
	"github.com/jamescrystal/mychain/testutil/nullify"
	stablecoin "github.com/jamescrystal/mychain/x/stablecoin/module"
	"github.com/jamescrystal/mychain/x/stablecoin/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StablecoinKeeper(t)
	stablecoin.InitGenesis(ctx, k, genesisState)
	got := stablecoin.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
