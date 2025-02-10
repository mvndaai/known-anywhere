//go:generate go run ./scripts/error_codes/error_codes.go -format uuid -location . -fix true
//go:generate sh ./docker/build_frontend.sh

package main

import (
	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/router"
)

func main() {
	err := router.StartServer()
	if err != nil {
		ctxerr.Handle(err)
		return
	}
}
