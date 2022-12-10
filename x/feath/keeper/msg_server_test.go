package keeper_test

import (
	"context"
	"testing"

	keepertest "feath/testutil/keeper"
	"feath/x/feath/keeper"
	"feath/x/feath/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FeathKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
