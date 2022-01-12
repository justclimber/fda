package commonapi

import (
	"github.com/justclimber/fda/common/api/fdagrpc"
)

var AuthMethods = map[string]bool{
	fdagrpc.UrlPrefix + "Register": true,
	fdagrpc.UrlPrefix + "Login":    true,
}
