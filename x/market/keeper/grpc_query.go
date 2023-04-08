package keeper

import (
	"github.com/furynet/furynetwork/x/market/types"
)

var _ types.QueryServer = Keeper{}
