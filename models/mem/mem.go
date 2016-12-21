package memrepo

import (
	"fmt"

	"github.com/moul/acl/gen/pb"
	uuid "github.com/satori/go.uuid"
)

type Repo struct {
	tokens []*aclpb.Token
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) AddToken(token aclpb.Token) (string, error) {
	if token.Id == "" {
		token.Id = uuid.NewV4().String()
	}
	r.tokens = append(r.tokens, &token)
	return token.Id, nil
}

func (r *Repo) GetToken(tokenID string) (*aclpb.Token, error) {
	for _, token := range r.tokens {
		if token.Id == tokenID {
			return token, nil
		}
	}
	return nil, fmt.Errorf("token not found")
}
