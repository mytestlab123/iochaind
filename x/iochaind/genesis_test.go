package iochaind_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iochaind/testutil/keeper"
	"github.com/mytestlab123/iochaind/testutil/nullify"
	"github.com/mytestlab123/iochaind/x/iochaind"
	"github.com/mytestlab123/iochaind/x/iochaind/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IochaindKeeper(t)
	iochaind.InitGenesis(ctx, *k, genesisState)
	got := iochaind.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
