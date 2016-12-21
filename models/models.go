package models

import "github.com/moul/acl/gen/pb"

type Repository interface {
	GetToken(string) (*aclpb.Token, error)
}
