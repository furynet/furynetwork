package keeper

import (
	"github.com/furynet/furynetwork/x/strategicreserve/types"
)

var _ types.QueryServer = Keeper{}
