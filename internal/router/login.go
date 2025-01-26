package router

import (
	"net/http"

	"github.com/mvndaai/ctxerr"
)

func loginsHandler(r *http.Request) (data, meta any, status int, _ error) {
	return nil, nil, http.StatusBadGateway, ctxerr.New(r.Context(), "82bd850e-23aa-4fe0-8b9b-801aab34de0a", "test error")
}

// https://work.meta.com/help/624731485413747?helpref=faq_content
// https://docs.x.com/resources/fundamentals/authentication/guides/log-in-with-x
// https://developer.apple.com/help/account/configure-app-capabilities/configure-sign-in-with-apple-for-the-web/
// https://developers.google.com/identity/sign-in/web/sign-in
// https://github.com/lucia-auth/example-sveltekit-github-oauth
