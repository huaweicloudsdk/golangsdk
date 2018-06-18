package signer

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

var maxCacheItemCount = 30

func TestInMultipleThread(test *testing.T) {
	cache := NewCache(maxCacheItemCount)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go multipleThreadTest(&wg, cache, test, i)
	}

	wg.Wait()
}

func multipleThreadTest(wg *sync.WaitGroup, cache *MemoryCache, test *testing.T, tId int) {
	defer wg.Done()
	fmt.Println("Enter into multipleThreadTest:", tId)
	time.Sleep(3 * time.Second)
	fmt.Println("Begin cache test", tId)

	for i := 0; i < 20; i++ {
		timeNano := time.Now().UnixNano()
		key := fmt.Sprint(timeNano, "_", i)
		cache.Add(key, strconv.Itoa(i))

		if len(cache.cacheHolder) > maxCacheItemCount || len(cache.cacheKeys) > maxCacheItemCount {
			test.Error("Too much cache items")
		}
	}
}

func TestCacheInSingleThread(test *testing.T) {
	cache := NewCache(2)

	cache.Add("a1", "a1V")

	if len(cache.cacheKeys) != 1 && len(cache.cacheHolder) != 1 {
		test.Fail()
	}

	if cache.Get("a1") != "a1V" {
		test.Fail()
	}

	cache.Add("a2", "a2V")

	if len(cache.cacheKeys) != 2 && len(cache.cacheHolder) != 2 {
		test.Fail()
	}

	if cache.Get("a1") != "a1V" {
		test.Fail()
	}

	if cache.Get("a2") != "a2V" {
		test.Fail()
	}

	cache.Add("a3", "a3V")

	if len(cache.cacheKeys) != 2 && len(cache.cacheHolder) != 2 {
		test.Fail()
	}

	if cache.Get("a1") != "" {
		test.Fail()
	}

	if cache.Get("a2") != "a2V" {
		test.Fail()
	}

	if cache.Get("a3") != "a3V" {
		test.Fail()
	}
	cache.Add("a4", "a4V")

	if len(cache.cacheKeys) != 2 && len(cache.cacheHolder) != 2 {
		test.Fail()
	}

	if cache.Get("a1") != "" {
		test.Fail()
	}

	if cache.Get("a2") != "" {
		test.Fail()
	}

	if cache.Get("a3") != "a3V" {
		test.Fail()
	}

	if cache.Get("a4") != "a4V" {
		test.Fail()
	}
}
