package router

import (
	"net/http"

	"github.com/mvndaai/ctxerr"
)

func loginsHandler(r *http.Request) (data, meta any, status int, _ error) {
	return nil, nil, http.StatusBadGateway, ctxerr.New(r.Context(), "c691a249-1413-407d-a309-b44d52eb51ac", "test error")
}

// https://work.meta.com/help/624731485413747?helpref=faq_content
// https://docs.x.com/resources/fundamentals/authentication/guides/log-in-with-x
// https://developer.apple.com/help/account/configure-app-capabilities/configure-sign-in-with-apple-for-the-web/
// https://developers.google.com/identity/sign-in/web/sign-in
// https://github.com/lucia-auth/example-sveltekit-github-oauth
