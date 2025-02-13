package keeper

import (
	"github.com/jamescrystal/mychain/x/mychain/types"
)

var _ types.QueryServer = Keeper{}
