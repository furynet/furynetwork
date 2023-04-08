package mint_test

import (
	"testing"

	"github.com/furynet/furynetwork/testutil/nullify"
	simappUtil "github.com/furynet/furynetwork/testutil/simapp"
	"github.com/furynet/furynetwork/x/mint"
	"github.com/furynet/furynetwork/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	mint.InitGenesis(ctx, tApp.MintKeeper, genesisState)
	got := mint.ExportGenesis(ctx, tApp.MintKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
