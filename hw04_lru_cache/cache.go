package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mx       sync.Mutex
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	defer c.mx.Unlock()
	c.mx.Lock()

	wasInCache := false

	if _, ok := c.items[key]; ok {
		wasInCache = ok
	} else if c.queue.Len() == c.capacity {
		delete(c.items, c.queue.Back().GetKey())
		c.queue.Remove(c.queue.Back())
	}

	item := c.queue.PushFront(value)
	item.SetKey(key)

	c.items[key] = item

	return wasInCache
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	defer c.mx.Unlock()

	c.mx.Lock()
	if v, ok := c.items[key]; ok {
		c.queue.MoveToFront(v)
		return v.GetValue(), true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
