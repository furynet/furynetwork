package keeper

import (
	"github.com/furynet/furynetwork/x/house/types"
)

var _ types.QueryServer = Keeper{}
