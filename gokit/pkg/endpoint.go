package pkg

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeHomeEndpoint(svc HomeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HomeRequest)
		resp, err := svc.Get(ctx, req)
		if err != nil {
			return HomeResponse{}, nil
		}
		return resp, nil
	}
}
