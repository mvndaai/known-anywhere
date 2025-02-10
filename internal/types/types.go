package types

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mvndaai/ctxerr"
)

func JSONTag(a any, name string) string {
	typeOf := reflect.TypeOf(a)
	switch typeOf.Kind() {
	case reflect.Struct:
		for i := range typeOf.NumField() {
			if typeOf.Field(i).Name == name {
				return typeOf.Field(i).Tag.Get("json")
			}
		}
	}
	panic(fmt.Sprintf("field %s not found", name))
}

type (
	Pagination struct {
		Limit       int    `json:"limit"`
		Cursor      string `json:"cursor"`
		ShowDeleted bool   `json:"show_deleted,omitempty"`
		ShowPending bool   `json:"show_pending,omitempty"`
	}

	PaginationResponse struct {
		Total  int    `json:"total"`
		Cursor string `json:"cursor"`
	}
)

func (p *Pagination) Normalize() {
	if p.Limit == 0 {
		p.Limit = 10
	}
	p.Cursor = strings.TrimSpace(p.Cursor)
}

func (p *Pagination) Fill(ctx context.Context, q url.Values) error {
	var err error
	key := JSONTag(*p, "Limit")
	if v := strings.TrimSpace(q.Get(key)); v != "" {
		p.Limit, err = strconv.Atoi(v)
		if err != nil {
			ctx = ctxerr.SetField(ctx, key, v)
			return ctxerr.NewHTTP(ctx, "da799ef8-f059-4794-90e2-d5a1cff2886e", "invalid limit", http.StatusBadRequest, "invalid limit")
		}
	}
	p.Cursor = q.Get(JSONTag(*p, "Cursor"))
	p.Normalize()
	return nil
}

type (
	DomainCreate struct {
		DisplayName string `json:"display_name"`
		Description string `json:"description"`
		Notes       string `json:"notes"`
	}

	Domain struct {
		ID uuid.UUID `json:"id"`
		DomainCreate
	}

	DomainLink struct {
		DomainID    uuid.UUID `json:"domain_id"`
		Link        string    `json:"link"`
		CountryCode string    `json:"country_code"`
	}

	DomainList struct {
		Pagination Pagination   `json:"pagination"`
		Filters    DomainCreate `json:"filters"`
	}
)

type (
	UserCreate struct {
		Username    string `json:"username"`
		DisplayName string `json:"display_name"`
	}

	User struct {
		ID uuid.UUID `json:"id"`
		UserCreate
	}

	UserList struct {
		Pagination Pagination `json:"pagination"`
		Filters    UserCreate `json:"filters"`
	}
)

type (
	Logout struct {
		JWTID      uuid.UUID `json:"jwt_id"`
		Expiration time.Time `json:"expiration"`
		UserID     uuid.UUID `json:"user_id"`
	}
)

func (v DomainCreate) Validate(ctx context.Context) error {
	var err error
	if v.DisplayName == "" {
		err = errors.Join(err, ctxerr.NewHTTP(ctx, "56f79cb4-b081-4447-b0fe-a0317d57f809", "missing display_name", http.StatusBadRequest, "missing display_name"))
	}
	return ctxerr.QuickWrap(ctx, err)
}

func (v DomainLink) Validate(ctx context.Context) error {
	var err error
	if v.Link == "" {
		err = errors.Join(err, ctxerr.NewHTTP(ctx, "c4173215-7a81-467f-ac8b-0dc525074bd0", "missing link", http.StatusBadRequest, "missing link"))
	}
	return ctxerr.QuickWrap(ctx, err)
}

func (v *DomainList) Normalize() {
	v.Filters.DisplayName = strings.TrimSpace(v.Filters.DisplayName)
	v.Filters.Description = strings.TrimSpace(v.Filters.Description)
	v.Filters.Notes = strings.TrimSpace(v.Filters.Notes)
	v.Pagination.Normalize()
}

func (v *DomainList) Fill(ctx context.Context, q url.Values) error {
	v.Filters.DisplayName = q.Get(JSONTag(v.Filters, "DisplayName"))
	v.Filters.Description = q.Get(JSONTag(v.Filters, "Description"))
	v.Filters.Notes = q.Get(JSONTag(v.Filters, "Notes"))
	err := v.Pagination.Fill(ctx, q)
	if err != nil {
		return ctxerr.QuickWrap(ctx, err)
	}
	v.Normalize()
	return nil
}

func (v *UserList) Normalize() {
	v.Filters.Username = strings.TrimSpace(v.Filters.Username)
	v.Filters.DisplayName = strings.TrimSpace(v.Filters.DisplayName)
	v.Pagination.Normalize()
}

func (v *UserList) Fill(ctx context.Context, q url.Values) error {
	v.Filters.Username = q.Get(JSONTag(v.Filters, "Username"))
	v.Filters.DisplayName = q.Get(JSONTag(v.Filters, "DisplayName"))
	err := v.Pagination.Fill(ctx, q)
	if err != nil {
		return ctxerr.QuickWrap(ctx, err)
	}
	v.Normalize()
	return nil
}

func (d Domain) GetID() uuid.UUID { return d.ID }
func (u User) GetID() uuid.UUID   { return u.ID }
