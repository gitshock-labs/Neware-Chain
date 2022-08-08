package gtfx_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "gtfx/testutil/keeper"
	"gtfx/testutil/nullify"
	"gtfx/x/gtfx"
	"gtfx/x/gtfx/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GtfxKeeper(t)
	gtfx.InitGenesis(ctx, *k, genesisState)
	got := gtfx.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
