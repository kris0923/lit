package lite

import (
	"net/http"
)

type Core struct{}

func NewCore() *Core {
	return &Core{}
}

func (*Core) ServeHTTP(req *http.Request, resp http.ResponseWriter) {
	//
}
