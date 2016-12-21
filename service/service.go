package aclsvc

import (
	"fmt"

	"github.com/moul/acl/gen/pb"
	"github.com/moul/acl/models"
	"golang.org/x/net/context"
)

type Service struct {
	repo models.Repository
}

func New(repo models.Repository) aclpb.AclServiceServer {
	return &Service{
		repo: repo,
	}
}

func (svc Service) Hasperm(ctx context.Context, input *aclpb.HasPermRequest) (*aclpb.HasPermResponse, error) {
	token, err := svc.repo.GetToken(input.Token)
	if err != nil {
		return nil, err
	}
	fmt.Println(token)
	return nil, fmt.Errorf("not implemented")
}
