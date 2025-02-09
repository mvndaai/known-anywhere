package jwt

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/config"
)

const (
	HeaderAuthorization       = "Authorization"
	HeaderAuthorizationPrefix = "Bearer "
)

type JWTClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func issuer() string {
	env := config.Get().Env
	return "ks-" + env
}

func (c JWTClaims) Valid() error {
	return jwt.NewValidator().Validate(c.RegisteredClaims)
}

func (c JWTClaims) Validate(ctx context.Context) error {
	if c.ExpiresAt != nil {
		maxExpiresAt := time.Now().Add(24 * 7 * time.Hour)
		if !c.ExpiresAt.Before(maxExpiresAt) {
			return ctxerr.NewHTTP(ctx, "921a5bf2-e9e8-46d0-b9d0-e4950b176d98", "expiration more than 7 days in the future",
				http.StatusBadRequest, "invalid expiration")
		}
	}

	if err := c.Valid(); err != nil {
		return ctxerr.WrapHTTP(ctx, err, "4009836a-30db-427e-985b-ba8be12b23fc", err.Error(), http.StatusBadGateway)
	}

	return nil
}

func (c *JWTClaims) Normalize(ctx context.Context) {
	now := time.Now()

	c.RegisteredClaims.ID = uuid.Must(uuid.NewV7()).String()
	c.RegisteredClaims.Issuer = issuer()
	c.RegisteredClaims.NotBefore = jwt.NewNumericDate(now)
	c.RegisteredClaims.IssuedAt = jwt.NewNumericDate(now)

	if c.RegisteredClaims.ExpiresAt == nil {
		c.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(now.Add(5 * time.Minute))
	}

	for i := range c.Audience {
		c.Audience[i] = strings.ToLower(c.Audience[i])
	}
}

func (c *JWTClaims) EnsureClaims(ctx context.Context, method, path string, params url.Values, aud string) error {
	ctx = ctxerr.SetField(ctx, "method", method)
	ctx = ctxerr.SetField(ctx, "path", path)
	ctx = ctxerr.SetField(ctx, "params", params)
	_ = aud // TODO: use aud

	if c == nil {
		return ctxerr.New(ctx, "1a071b29-4c11-41ca-b765-7e50bd7227ad", "nil claims")
	}

	if c.ExpiresAt == nil {
		return ctxerr.NewHTTP(ctx, "5c424190-9b00-43cb-8dcc-1fe06d061140", "'exp' required in token", http.StatusUnauthorized, "misisng jwt exp")
	}

	if is := issuer(); c.Issuer != is {
		ctx = ctxerr.SetField(ctx, "issuerToken", c.Issuer)
		ctx = ctxerr.SetField(ctx, "issuerEnv", is)
		return ctxerr.NewHTTP(ctx, "a525f59f-f5e9-40d9-9956-aa089e84ea02", "invalid issuer", http.StatusUnauthorized, "invalid issuer")
	}

	return nil
}

func GetJWTClaims(r *http.Request) (*JWTClaims, error) {
	ctx := r.Context()
	jwtToken := strings.TrimPrefix(r.Header.Get(HeaderAuthorization), HeaderAuthorizationPrefix)
	if jwtToken == "" {
		return nil, ctxerr.NewHTTP(ctx, "670a3fb3-1539-4a20-acdd-0bb0f425bf95", "missing auth token", http.StatusUnauthorized, "missing token")
	}
	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", ctxerr.NewHTTP(ctx, "e4cd2a10-e084-4186-85a5-1788fcfad945", "", http.StatusUnauthorized, "bad token signing method")
		}
		return config.Get().JWTSecret, nil
	})
	if err != nil {
		return nil, ctxerr.WrapHTTP(ctx, err, "a6c3217d-0f6b-4153-8789-128e5cefc43d", err.Error(), http.StatusUnauthorized)
	}

	if err := claims.Valid(); err != nil {
		return nil, ctxerr.WrapHTTP(ctx, err, "012ed653-319b-47c5-9e71-d8282c375b62", err.Error(), http.StatusUnauthorized, "invalid jwt claims")
	}
	return claims, nil
}

// GenerateJWT creates a JWT with our claims
func GenerateJWT(ctx context.Context, claims JWTClaims) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(config.Get().JWTSecret)
	if err != nil {
		return "", ctxerr.Wrap(ctx, err, "757b2749-1bea-4d79-91e0-826766773357", "could not sign jwt")
	}
	return token, nil
}

func ExtractJWTSubject(headers http.Header) string {
	jwtToken := strings.TrimPrefix(headers.Get(HeaderAuthorization), HeaderAuthorizationPrefix)
	claims := &JWTClaims{}
	_, _ = jwt.ParseWithClaims(jwtToken, claims, nil)
	return claims.Subject
}
