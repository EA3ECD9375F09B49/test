package libs

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	GoCacheLocal = cache.New(2*time.Minute, 5*time.Minute)
	GoCacheTimes = map[int]time.Duration{
		0: 50 * time.Second,
		1: 2 * time.Hour,
		2: 5 * time.Second,
	}
)

func GSet(keyName string, value interface{}, expire time.Duration) {
	GoCacheLocal.Set(keyName, value, expire)
}

func GGet(keyName string) (result interface{}, ok bool) {
	result, ok = GoCacheLocal.Get(keyName)
	return
}

func GDel(keyName string) {
	GoCacheLocal.Delete(keyName)
}
