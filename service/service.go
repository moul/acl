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

func (svc Service) HasPerm(ctx context.Context, input *aclpb.HasPermRequest) (*aclpb.HasPermResponse, error) {
	token, err := svc.repo.GetToken(input.TokenId)
	if err != nil {
		return nil, err
	}
	hasPerm, err := hasPerm(token, input.Service, input.Name, input.Resource)
	if err != nil {
		return nil, err
	}

	return &aclpb.HasPermResponse{
		HasPerm: hasPerm,
	}, nil
}

func (svc Service) GetToken(ctx context.Context, input *aclpb.GetTokenRequest) (*aclpb.GetTokenResponse, error) {
	token, err := svc.repo.GetToken(input.Id)
	if err != nil {
		return nil, err
	}

	return &aclpb.GetTokenResponse{
		Token: token,
	}, nil
}

func (svc Service) AddToken(ctx context.Context, input *aclpb.AddTokenRequest) (*aclpb.AddTokenResponse, error) {
	if input.Token == nil {
		return nil, fmt.Errorf("invalid token")
	}
	id, err := svc.repo.AddToken(*input.Token)
	if err != nil {
		return nil, err
	}

	return &aclpb.AddTokenResponse{
		Id: id,
	}, nil
}

func hasPerm(token *aclpb.Token, service, name, resource string) (bool, error) {
	return false, fmt.Errorf("not implemented")
}
