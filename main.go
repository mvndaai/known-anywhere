//go:generate go run ./scripts/ctxerrcodes/ctxerrcodes.go -format uuid -location . -fix true
//go:generate sh ./docker/build_frontend.sh

package main

import (
	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-anywhere/internal/router"
)

func main() {
	err := router.StartServer()
	if err != nil {
		ctxerr.Handle(err)
		return
	}
}
