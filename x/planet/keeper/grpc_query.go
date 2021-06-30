package keeper

import (
	"github.com/fernando-romero/planet/x/planet/types"
)

var _ types.QueryServer = Keeper{}
