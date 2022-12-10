package feath_test

import (
	"testing"

	keepertest "feath/testutil/keeper"
	"feath/testutil/nullify"
	"feath/x/feath"
	"feath/x/feath/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FeathKeeper(t)
	feath.InitGenesis(ctx, *k, genesisState)
	got := feath.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
