package jwt

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const subjectContextKey = contextKey("subject")

func ContextWithSubject(ctx context.Context, subject string) context.Context {
	return context.WithValue(ctx, subjectContextKey, subject)
}

func SubjectFromContext(ctx context.Context) uuid.UUID {
	if v, ok := ctx.Value(subjectContextKey).(string); ok {
		if id, err := uuid.Parse(v); err == nil {
			return id
		}
	}
	return uuid.UUID{}
}
