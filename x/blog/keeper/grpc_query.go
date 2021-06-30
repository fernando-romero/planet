package keeper

import (
	"github.com/fernando-romero/planet/x/blog/types"
)

var _ types.QueryServer = Keeper{}
