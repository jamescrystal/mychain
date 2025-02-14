package keeper

import (
	"github.com/jamescrystal/mychain/x/stablecoin/types"
)

var _ types.QueryServer = Keeper{}
