package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iochaind/testutil/keeper"
	"github.com/mytestlab123/iochaind/x/iochaind/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IochaindKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
