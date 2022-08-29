package grpc

// THIS FILE IS GENERATED CODE, DO NOT EDIT
// SOURCE AT `proto/osmosis/twap/v1beta1/query.yml`

import (
	context "context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Anmol1696/osmosis/v11/x/twap/client"
	"github.com/Anmol1696/osmosis/v11/x/twap/client/queryproto"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Querier struct {
	Q client.Querier
}

var _ queryproto.QueryServer = Querier{}

func (q Querier) GetArithmeticTwap(grpcCtx context.Context,
	req *queryproto.GetArithmeticTwapRequest,
) (*queryproto.GetArithmeticTwapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(grpcCtx)
	return q.Q.GetArithmeticTwap(ctx, *req)
}
