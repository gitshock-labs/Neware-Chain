package keeper

import (
	"context"

    "scavenge/x/scavenge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) SubmitScavenge(goCtx context.Context,  msg *types.MsgSubmitScavenge) (*types.MsgSubmitScavengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgSubmitScavengeResponse{}, nil
}
