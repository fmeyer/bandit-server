package repository

import (
	"fmt"
	"github.com/dropbox/godropbox/memcache"
	"github.com/dropbox/godropbox/net2"
	"strconv"
	"time"
)

const HIT = "h"
const REWARD = "r"

type Memcached struct {
	mc memcache.Client
}

func NewMemcached(address string) Memcached {
	var maxIdleTime time.Duration = 30 * time.Second

	options := net2.ConnectionOptions{MaxActiveConnections: 2, MaxIdleConnections: 1, MaxIdleTime: &maxIdleTime}

	manager := memcache.NewStaticShardManager([]string{address}, func(key string, numShard int) int { return 0 }, options)
	return Memcached{mc: memcache.NewShardedClient(manager, false)}
}

func (v Memcached) mcGet(kind string, experiment string, arm string) int64 {
	key := fmt.Sprintf("%s:%s:%s", kind, experiment, arm)
	val := v.mc.Get(key).Value()
	if val == nil {
		return 0
	}
	count, _ := strconv.ParseInt(string(val), 10, 64)
	return count
}

func (v Memcached) mcIncr(kind string, experiment string, arm string) {
	key := fmt.Sprintf("%s:%s:%s", kind, experiment, arm)
	v.mc.Increment(key, 1, 1, 0)
}

func (v Memcached) Get(experiment string, arms []string) ExperimentData {
	var expData = ExperimentData{0, make(map[string]ArmData)}

	for _, arm := range arms {
		armData := ArmData{Hits: v.mcGet(HIT, experiment, arm), Rewards: v.mcGet(REWARD, experiment, arm)}

		expData.Arms[arm] = armData
		expData.TotalHits += armData.Hits
	}

	return expData
}

func (v Memcached) Hit(experiment string, arm string) {
	v.mcIncr(HIT, experiment, arm)
}

func (v Memcached) Reward(experiment string, arm string) {
	v.mcIncr(REWARD, experiment, arm)
}
