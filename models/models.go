package models

import "github.com/moul/acl/gen/pb"

type Repository interface {
	AddToken(aclpb.Token) (string, error)
	GetToken(string) (*aclpb.Token, error)
}
