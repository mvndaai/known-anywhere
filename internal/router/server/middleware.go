package server

import (
	"net/http"
	"strings"

	"github.com/mvndaai/ctxerr"
)

func TrimSpaces(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Trim values of query params
		q := r.URL.Query()
		for k, vs := range q {
			for i, v := range vs {
				q[k][i] = strings.TrimSpace(v)
			}
		}
		r.URL.RawQuery = q.Encode()

		// Trim values of headers
		for k, vs := range r.Header {
			for i, v := range vs {
				r.Header[k][i] = strings.TrimSpace(v)
			}
		}
		next.ServeHTTP(w, r)
	}
}

var ignorableHeaderPrefixes = []string{
	"Accept",
	"Authorization",
	"Cache-",
	"Connection",
	"Content-",
	"Proto",
	"Sec-",
	"X-Forwarded-",
	"X-Real-IP",
}

func LogHeadersAndParams(l http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		headers := http.Header{}
	outer:
		for k, vs := range req.Header {
			k := http.CanonicalHeaderKey(k)
			for _, prefix := range ignorableHeaderPrefixes {
				if strings.HasPrefix(k, prefix) {
					continue outer
				}
			}
			for _, v := range vs {
				headers.Add(k, v)
			}
		}

		ctx := req.Context()
		if len(headers) > 0 {
			clearDuplicates(headers)
			ctx = ctxerr.SetField(ctx, "headers", headers)
		}
		if len(req.URL.Query()) > 0 {
			q := req.URL.Query()
			clearDuplicates(q)
			ctx = ctxerr.SetField(ctx, "parameters", q)
		}

		l(res, req.WithContext(ctx))
	}
}

func clearDuplicates(m map[string][]string) {
	for k, vs := range m {
		if len(vs) < 2 {
			continue
		}
		r := make([]string, 0, len(vs))
		existing := map[string]struct{}{}
		for _, v := range vs {
			if _, ok := existing[v]; ok {
				continue
			}
			r = append(r, v)
			existing[v] = struct{}{}
		}
		m[k] = r
	}
}
