package db

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mvndaai/known-socially/internal/types"
	"github.com/patrickmn/go-cache"
)

const cacheFmt = "%s:%s:%s" //env:type:columnName:columnValue
const cacheExpire = 5 * time.Minute

type CacheImpl struct {
	cache *cache.Cache
}

func newCache() *CacheImpl {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &CacheImpl{cache: c}
}

func logoutKey(userID uuid.UUID) string {
	return fmt.Sprintf(cacheFmt, "logout", "user_id", userID.String())
}
func (c *CacheImpl) SetJWTLogout(userID uuid.UUID, items []types.Logout) {
	c.cache.Set(logoutKey(userID), items, cacheExpire)
}

func (c *CacheImpl) GetJWTLogout(userID uuid.UUID) ([]types.Logout, bool) {
	items, ok := c.cache.Get(logoutKey(userID))
	if !ok {
		return nil, false
	}
	return items.([]types.Logout), true
}

func (c *CacheImpl) DeleteJWTLogout(userID uuid.UUID) {
	c.cache.Delete(logoutKey(userID))
}
