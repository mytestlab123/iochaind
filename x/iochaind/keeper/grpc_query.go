package keeper

import (
	"github.com/mytestlab123/iochaind/x/iochaind/types"
)

var _ types.QueryServer = Keeper{}
