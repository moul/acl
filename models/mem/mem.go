package memrepo

import (
	"fmt"

	"github.com/moul/acl/gen/pb"
)

type Repo struct{}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) GetToken(token string) (*aclpb.Token, error) {
	return nil, fmt.Errorf("not implemented")
}
