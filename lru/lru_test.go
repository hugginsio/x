package lru_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/hugginsio/x/lru"
)

func TestNew(t *testing.T) {
	t.Run("valid capacity", func(t *testing.T) {
		if cache, err := lru.New[string, int](3); err != nil || cache == nil {
			t.FailNow()
		}
	})

	t.Run("invalid capacity", func(t *testing.T) {
		if cache, err := lru.New[string, int](0); err == nil || cache != nil {
			t.FailNow()
		}

		if cache, err := lru.New[string, int](-1); err == nil || cache != nil {
			t.FailNow()
		}
	})
}

func TestGetPut(t *testing.T) {
	t.Run("put and get single item", func(t *testing.T) {
		cache, err := lru.New[string, int](3)
		if err != nil {
			t.Error(err)
		}

		if err := cache.Put("item", 1); err != nil {
			t.Error(err)
		}

		if res, err := cache.Get("item"); err != nil {
			t.Error(err)
		} else if res != 1 {
			t.Error("expected item to equal one")
		}
	})

	t.Run("get non-existent key", func(t *testing.T) {
		cache, err := lru.New[string, int](3)
		if err != nil {
			t.Error(err)
		}

		if _, err := cache.Get("non-existent"); err == nil {
			t.FailNow()
		}
	})

	t.Run("update existing key", func(t *testing.T) {
		cache, err := lru.New[string, int](1)
		if err != nil {
			t.Error(err)
		}

		if err := cache.Put("item", 1); err != nil {
			t.Error(err)
		}

		if err := cache.Put("item", 1); err != nil {
			t.Error(err)
		}
	})
}

func TestEviction(t *testing.T) {
	t.Run("evict LRU when at capacity", func(t *testing.T) {
		cache, err := lru.New[string, int](2)
		if err != nil {
			t.Error(err)
		}

		cache.Put("first", 1)
		cache.Put("second", 2)
		cache.Put("third", 3)

		if res, err := cache.Get("first"); res != 0 || err != lru.ErrNotFound {
			t.Error("expected first key to be evicted")
		}
	})

	t.Run("access order affects eviction", func(t *testing.T) {
		// TODO: Test that accessing an item prevents it from being evicted
		// Example:
		// cache := New(2)
		// Put(1, "a")
		// Put(2, "b")
		// Get(1) // Access key 1, making it recently used
		// Put(3, "c") // Should evict key 2, not key 1
		cache, err := lru.New[string, int](2)
		if err != nil {
			t.Error(err)
		}

		cache.Put("first", 1)
		cache.Put("second", 2)

		if _, err := cache.Get("first"); err != nil {
			t.Error(err)
		}

		cache.Put("third", 3) // should evict second, not first

		if res, err := cache.Get("first"); res != 1 || err != nil {
			t.Error("expected second key to remain in cache")
		}

		if res, err := cache.Get("second"); res != 0 || err != lru.ErrNotFound {
			t.Error("expected second key to be evicted")
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("remove existing key", func(t *testing.T) {
		cache, err := lru.New[string, int](2)
		if err != nil {
			t.Error(err)
		}

		if err := cache.Put("first", 1); err != nil {
			t.Error(err)
		}

		if cache.Len() != 1 {
			t.Error("expected cache length to increment after put")
		}

		if !cache.Remove("first") {
			t.Error("expected remove to return true")
		}

		if cache.Len() != 0 {
			t.Error("expected cache length to be zero after last item removed")
		}
	})

	t.Run("remove non-existent key", func(t *testing.T) {
		cache, err := lru.New[string, int](1)
		if err != nil {
			t.Error(err)
		}

		if cache.Remove("nah") {
			t.Error("expected removal of nonexistent key to return false")
		}
	})
}

func TestLen(t *testing.T) {
	cache, err := lru.New[string, int](2)
	if err != nil {
		t.Error(err)
	}

	if cache.Len() != 0 {
		t.Error("expected cache length to be zero")
	}

	cache.Put("first", 1)
	cache.Put("second", 2)

	if cache.Len() != 2 {
		t.Error("expected cache length to be two")
	}

	cache.Put("third", 2)

	if cache.Len() != 2 {
		t.Error("expected cache length to be two")
	}

	cache.Remove("second")

	if cache.Len() != 1 {
		t.Error("expected cache length to be one")
	}
}

// Benchmark tests
func BenchmarkGet(b *testing.B) {
	cache, err := lru.New[string, int](100)
	if err != nil {
		b.Fatal(err)
	}

	for i := range 100 {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		if res, err := cache.Get("key50"); err != nil {
			b.Fatalf("error getting key: %v", err)
		} else if res != 50 {
			b.Fatalf("incorrect result: %d, expected 50", res)
		}
	}
}

func BenchmarkPut(b *testing.B) {
	cache, err := lru.New[string, int](math.MaxInt)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; b.Loop(); i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}
}

func BenchmarkEviction(b *testing.B) {
	cache, err := lru.New[string, int](1)
	if err != nil {
		b.Fatal(err)
	}

	if err := cache.Put("preload", 1); err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; b.Loop(); i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}
}
