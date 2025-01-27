package types_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/mvndaai/known-socially/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDomainListFill(t *testing.T) {
	tests := []struct {
		name          string
		q             url.Values
		expected      types.DomainList
		errorContains string
	}{
		{
			name: "empty",
			q:    url.Values{},
			expected: types.DomainList{
				Filters: types.DomainCreate{},
				Pagination: types.Pagination{
					Limit: 10,
				},
			},
			errorContains: "",
		},
		{
			name: "everything",
			q: url.Values{
				"limit":        {" 20 "},
				"cursor":       {" c "},
				"display_name": {" dn "},
				"description":  {" d "},
				"notes":        {" n "},
			},
			expected: types.DomainList{
				Filters: types.DomainCreate{
					DisplayName: "dn",
					Description: "d",
					Notes:       "n",
				},
				Pagination: types.Pagination{
					Limit:  20,
					Cursor: "c",
				},
			},
			errorContains: "",
		},
		{
			name:          "non number limit",
			q:             url.Values{"limit": {"i"}},
			expected:      types.DomainList{},
			errorContains: "invalid limit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := types.DomainList{}
			err := l.Fill(context.Background(), tt.q)
			require.Equal(t, tt.errorContains == "", err == nil)
			assert.Equal(t, tt.expected, l)
			if err != nil {
				assert.Contains(t, tt.errorContains, err.Error())
			}
		})
	}
}
