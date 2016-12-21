package aclsvc

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/moul/acl/gen/pb"
)

func Test_hasPerm(t *testing.T) {
	Convey("Testing hasPerm()", t, func() {
		token := &aclpb.Token{
			Permissions: []*aclpb.Permission{
				{
					Service:  "compute",
					Name:     "servers:read",
					Resource: "abcdef",
				},
			},
		}

		result, err := hasPerm(token, "compute", "servers:read", "abcdef")
		So(err, ShouldBeNil)
		So(result, ShouldBeTrue)
	})
}

func Test_permMatches(t *testing.T) {
	Convey("Testing permMatches()", t, func() {
	})
}
