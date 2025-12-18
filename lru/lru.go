// Package lru implements a simple LRU cache.
package lru

import (
	"errors"
)

var (
	ErrNotFound        = errors.New("key not found")
	ErrInvalidCapacity = errors.New("capacity must be greater than 0")
)

// node represents a doubly linked list node
type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

// Cache is a generic LRU cache implementation
type Cache[K comparable, V any] struct {
	capacity int
	size     int
	cache    map[K]*node[K, V]
	head     *node[K, V] // most recently used
	tail     *node[K, V] // least recently used
}

// New creates a new LRU cache with the given capacity
func New[K comparable, V any](capacity int) (*Cache[K, V], error) {
	if capacity <= 0 {
		return nil, ErrInvalidCapacity
	}

	cache := &Cache[K, V]{
		capacity: capacity,
		size:     0,
		cache:    map[K]*node[K, V]{},
		head:     &node[K, V]{},
		tail:     &node[K, V]{},
	}

	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	return cache, nil
}

// Get retrieves a value from the cache and marks it as recently used
func (c *Cache[K, V]) Get(key K) (V, error) {
	var zero V

	res := c.cache[key]
	if res == nil {
		return zero, ErrNotFound
	}

	c.moveToFront(res)
	return res.value, nil
}

// Put adds or updates a key-value pair in the cache
func (c *Cache[K, V]) Put(key K, value V) error {
	if current := c.cache[key]; current != nil {
		current.value = value
		c.moveToFront(current)
		return nil
	}

	if c.size >= c.capacity {
		tail := c.removeTail()
		delete(c.cache, tail.key)
	} else {
		c.size++
	}

	node := &node[K, V]{
		key:   key,
		value: value,
	}

	c.cache[node.key] = node
	c.addToFront(node)

	return nil
}

// Remove removes a key from the cache
func (c *Cache[K, V]) Remove(key K) bool {
	node := c.cache[key]
	if node == nil {
		return false
	}

	delete(c.cache, node.key)
	c.removeNode(node)
	c.size--

	return true
}

// Len returns the current number of items in the cache
func (c *Cache[K, V]) Len() int {
	return c.size
}

// Cap returns the capacity of the cache
func (c *Cache[K, V]) Cap() int {
	return c.capacity
}

// moveToFront moves a node to the front of the list (most recently used position)
func (c *Cache[K, V]) moveToFront(n *node[K, V]) {
	c.removeNode(n)
	c.addToFront(n)
}

// removeNode removes a node from the doubly linked list
func (c *Cache[K, V]) removeNode(n *node[K, V]) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

// addToFront adds a new node to the front of the list
func (c *Cache[K, V]) addToFront(n *node[K, V]) {
	n.next = c.head.next
	n.prev = c.head
	c.head.next.prev = n
	c.head.next = n
}

// removeTail removes and returns the tail node (LRU item)
func (c *Cache[K, V]) removeTail() *node[K, V] {
	lru := c.tail.prev
	c.removeNode(lru)
	return lru
}
