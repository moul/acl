package aclsvc

import (
	"fmt"

	"github.com/moul/acl/gen/pb"
	"golang.org/x/net/context"
)

type Service struct{}

func New() aclpb.AclServiceServer {
	return &Service{}
}

func (svc Service) Hasperm(ctx context.Context, input *aclpb.HasPermRequest) (*aclpb.HasPermResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
