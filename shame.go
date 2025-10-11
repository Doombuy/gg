package gg

import (
	"sync"
	"time"
)

// "fmt"
type item struct {
	value     int
	expiresAt time.Time
}

type Cache struct {
	data map[string]item
	mu   sync.RWMutex
}

func New() *Cache {
	c := &Cache{
		data: make(map[string]item),
	}
	go c.cleanup()
	return c

}

func (c *Cache) Set(key string, value int, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = item{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
}

func (c *Cache) Get(key string) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	it, ok := c.data[key]
	if !ok {
		return 0, false
	}

	// Проверяем, не истёк ли срок
	if time.Now().After(it.expiresAt) {
		delete(c.data, key)
		return 0, false
	}

	return it.value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache) All() map[string]item {
	return c.data
}

func (c *Cache) cleanup() {
	for {
		time.Sleep(time.Second * 10)

		for k, v := range c.data {
			if time.Now().After(v.expiresAt) {
				delete(c.data, k)
			}
		}
	}
}
