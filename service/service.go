package aclsvc

import (
	"fmt"

	"golang.org/x/net/context"
)

type Service struct{}

func New() aclpb.AclServiceServer {
	return &Service{}
}

func (svc Service) HasPerm(ctx context.Context, input *aclpb.HasPermRequest) (*aclpb.HasPermResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
