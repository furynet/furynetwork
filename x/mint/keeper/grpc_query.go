package keeper

import (
	"github.com/furynet/furynetwork/x/mint/types"
)

var _ types.QueryServer = Keeper{}
