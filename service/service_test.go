package aclsvc

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/moul/acl/gen/pb"
)

func Test_hasPerm(t *testing.T) {
	Convey("Testing hasPerm()", t, FailureContinues, func() {
		token := aclpb.Token{}
		tokenJson := `{
  "permissions": [
    { "service": "compute", "name": "can_boot",    "resources": ["server1", "server2"] },
    { "service": "compute", "name": "can_delete",  "resources": ["server1"] },
    { "service": "account", "name": "token:*",     "resources": ["token1", "token2"] },
    { "service": "account", "name": "token:read",  "resources": ["token2", "token3"] },
    { "service": "account", "name": "token:write", "resources": ["token4"] }
  ]
}`
		So(json.Unmarshal([]byte(tokenJson), &token), ShouldBeNil)
		So(hasPerm(&token, "compute", "can_boot", "server1"), ShouldBeTrue)
		So(hasPerm(&token, "compute", "can_boot", "server2"), ShouldBeTrue)
		So(hasPerm(&token, "compute", "can_boot", "server3"), ShouldBeFalse)
		So(hasPerm(&token, "account", "token:read", "token1"), ShouldBeTrue)
		So(hasPerm(&token, "account", "token:write", "token1"), ShouldBeTrue)
		So(hasPerm(&token, "account", "token:write", "token1"), ShouldBeTrue)
		So(hasPerm(&token, "account", "token:write", "token4"), ShouldBeTrue)
		So(hasPerm(&token, "account", "token:write", "token3"), ShouldBeFalse)
	})
}

func Test_permMatches(t *testing.T) {
	Convey("Testing permMatches()", t, FailureContinues, func() {
		// simple
		So(permMatches("read", "read"), ShouldBeTrue)
		So(permMatches("", "read"), ShouldBeTrue)
		So(permMatches("write", "read"), ShouldBeFalse)

		// wildcard
		So(permMatches("read", "*"), ShouldBeTrue)

		// nested
		So(permMatches("object:read", "object:read"), ShouldBeTrue)
		So(permMatches("object:read", "object:*"), ShouldBeTrue)
		So(permMatches("", "object:*"), ShouldBeTrue)
		So(permMatches("object:write", "object:read"), ShouldBeFalse)

		// nested with different sizes
		So(permMatches("object:read:subperm", "*"), ShouldBeTrue)
		So(permMatches("object:read:subperm", "object"), ShouldBeTrue)
		So(permMatches("object", "object:read:hello"), ShouldBeFalse)

	})
}
