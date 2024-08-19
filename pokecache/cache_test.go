package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Second)

	if cache.CacheMap == nil {
		t.Error("cache is empty")
	}
}

func TestAddToCache(t *testing.T) {
	cache := NewCache(time.Second)

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "key1",
			val: []byte("val1"),
		},
		{
			key: "key2",
			val: []byte("val2"),
		},
		{
			key: "",
			val: []byte("val3"),
		},
	}

	for _, case_ := range cases {
		cache.Add(case_.key, case_.val)

		actual, ok := cache.Get(case_.key)
		if !ok {
			t.Errorf("%s is not found", case_.key)
			continue
		}

		if string(actual) != string(case_.val) {
			t.Errorf("%s does not match %s", string(actual), string(case_.val))
			continue
		}
	}

}

func TestReapCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	case_ := struct {
		key string
		val []byte
	}{
		key: "key1",
		val: []byte("val1"),
	}

	cache.Add(case_.key, case_.val)

	time.Sleep(interval * 2)

	_, ok := cache.Get(case_.key)
	if ok {
		t.Errorf("value was not reaped %v", ok)
	}

}

func TestReapCacheFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	case_ := struct {
		key string
		val []byte
	}{
		key: "key1",
		val: []byte("val1"),
	}

	cache.Add(case_.key, case_.val)

	time.Sleep(interval / 2)

	v, ok := cache.Get(case_.key)
	if !ok {
		t.Errorf("value was not found when it should be present %v %v", ok, v)
	}
}
