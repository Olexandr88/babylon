package finality_test

import (
	"testing"

	keepertest "github.com/babylonlabs-io/babylon/testutil/keeper"
	"github.com/babylonlabs-io/babylon/testutil/nullify"
	"github.com/babylonlabs-io/babylon/x/finality"
	"github.com/babylonlabs-io/babylon/x/finality/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.FinalityKeeper(t, nil, nil)
	finality.InitGenesis(ctx, *k, genesisState)
	got := finality.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

}
