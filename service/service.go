package aclsvc

import (
	"fmt"
	"strings"

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
	return &aclpb.HasPermResponse{
		HasPerm: hasPerm(token, input.Service, input.Name, input.Resource),
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

func hasPerm(token *aclpb.Token, service, name, resource string) bool {
	return len(getResources(token, service, name, resource)) > 0
}

func getResources(token *aclpb.Token, service, name, resource string) []string {
	ret := []string{}

	for _, perm := range token.Permissions {
		if perm.Service != service {
			continue
		}

		if !permMatches(name, perm.Name) {
			continue
		}

		for _, effectiveResource := range perm.Resources {
			if permMatches(resource, effectiveResource) {
				ret = append(ret, effectiveResource)
			}
		}
	}

	return ret
}

func permMatches(request, effective string) bool {
	if request == "" {
		return true
	}

	requestParts := strings.Split(request, ":")
	requestLen := len(requestParts)
	effectiveParts := strings.Split(effective, ":")
	effectiveLen := len(effectiveParts)

	longest := len(requestParts)
	if effectiveLen > longest {
		longest = effectiveLen
	}

	for i := 0; i < longest; i++ {
		if i >= effectiveLen {
			break
		}
		if effectiveParts[i] == "*" {
			continue
		}
		if i < requestLen && requestParts[i] == effectiveParts[i] {
			continue
		}
		return false
	}
	return true
}
