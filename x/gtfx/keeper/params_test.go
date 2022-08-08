package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "gtfx/testutil/keeper"
	"gtfx/x/gtfx/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GtfxKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
