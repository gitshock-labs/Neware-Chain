package keeper

import (
	"context"

	"gtfx/x/gtfx/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Gtfx(goCtx context.Context, req *types.QueryGtfxRequest) (*types.QueryGtfxResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	return &types.QueryGtfxResponse{Text: "Hello, Gitshock Finance On Cosmos Ecosystems!"}, nil // <--
}
