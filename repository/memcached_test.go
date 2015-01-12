package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
  "github.com/dropbox/godropbox/memcache"
)

func newMemcachedMocked() Memcached {
	return Memcached{mc: memcache.NewMockClient()}
}

func TestMemcachedBlank(t *testing.T) {
	repo := newMemcachedMocked()

	data := repo.Get("experiment", []string{"arm1"})
	assert.Equal(t, 0, data.TotalHits)
	assert.Equal(t, 0, data.Arms["arm1"].Hits)
	assert.Equal(t, 0, data.Arms["arm1"].Rewards)
}

func TestMemcachedHitCounter(t *testing.T) {
	repo := newMemcachedMocked()
	repo.Hit("experiment", "arm1")
	repo.Hit("experiment", "arm2")
	repo.Hit("experiment", "arm2")

	data := repo.Get("experiment", []string{"arm1", "arm2"})
	assert.Equal(t, 3, data.TotalHits)
	assert.Equal(t, 1, data.Arms["arm1"].Hits)
	assert.Equal(t, 0, data.Arms["arm1"].Rewards)
	assert.Equal(t, 2, data.Arms["arm2"].Hits)
	assert.Equal(t, 0, data.Arms["arm2"].Rewards)
}

func TestMemcachedRewardCounter(t *testing.T) {
	repo := newMemcachedMocked()
	repo.Reward("experiment", "arm1")
	repo.Reward("experiment", "arm2")
	repo.Reward("experiment", "arm2")

	data := repo.Get("experiment", []string{"arm1", "arm2"})
	assert.Equal(t, 0, data.TotalHits)
	assert.Equal(t, 0, data.Arms["arm1"].Hits)
	assert.Equal(t, 1, data.Arms["arm1"].Rewards)
	assert.Equal(t, 0, data.Arms["arm2"].Hits)
	assert.Equal(t, 2, data.Arms["arm2"].Rewards)
}

func TestMemcachedTotalHitsWhenSubSet(t *testing.T) {
	repo := newMemcachedMocked()
	repo.Hit("experiment", "arm1")
	repo.Hit("experiment", "arm2")
	repo.Hit("experiment", "arm2")
	repo.Hit("experiment", "arm3")
	repo.Hit("experiment", "arm3")
	repo.Hit("experiment", "arm3")

	data := repo.Get("experiment", []string{"arm1", "arm2"})
	assert.Equal(t, 3, data.TotalHits)
}
