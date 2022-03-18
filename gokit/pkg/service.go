package pkg

import "context"

type HomeService interface {
	Get(context.Context, HomeRequest) (HomeResponse, error)
}

type service struct {
}

func NewHomeService() HomeService {
	return service{}
}

func (s service) Get(_ context.Context, _ HomeRequest) (HomeResponse, error) {
	return HomeResponse{}, nil
}
