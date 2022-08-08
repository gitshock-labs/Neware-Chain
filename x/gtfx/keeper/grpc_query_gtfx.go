package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gtfx/x/gtfx/types"
)

func (k Keeper) Gtfx(goCtx context.Context, req *types.QueryGtfxRequest) (*types.QueryGtfxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryGtfxResponse{}, nil
}
