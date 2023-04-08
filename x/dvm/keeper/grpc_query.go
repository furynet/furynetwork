package keeper

import (
	"github.com/furynet/furynetwork/x/dvm/types"
)

var _ types.QueryServer = Keeper{}
