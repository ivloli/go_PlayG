package main

import (
	"container/list"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("vim-go")
}

type LruCache struct {
	size     int
	cache    *list.List
	cacheMap map[interface{}]*list.Element
	m        sync.Mutex
}

type entry struct {
	key   interface{}
	value interface{}
}

func NewCache(size int) *LruCache {
	l := &LruCache{}
	l.size = size
	l.cache = list.New()
	l.cacheMap = make(map[interface{}]*list.Element)
	return l
}

func (c *LruCache) Set(key, value interface{}) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.cache == nil {
		c.cache = list.New()
		c.cacheMap = make(map[interface{}]*list.Element)
	}
	if v, ok := c.cacheMap[key]; ok {
		v.Value.(*entry).value = value
		c.cache.MoveToFront(v)
	} else {
		en := &entry{
			key:   key,
			value: value,
		}
		ele := c.cache.PushFront(en)
		c.cacheMap[key] = ele
	}
	if c.size != 0 && c.cache.Len() > c.size {
		c.RemoveOld()
	}
}

func (c *LruCache) RemoveOld() {
	if c.cache.Len() == 0 {
		return
	}
	ele := c.cache.Back()
	c.cache.Remove(ele)
	key := ele.Value.(*entry).key
	delete(c.cacheMap, key)
}

func (c *LruCache) Get(key interface{}) interface{} {
	c.m.Lock()
	defer c.m.Unlock()
	if v, ok := c.cacheMap[key]; ok {
		return v.Value.(*entry).value
	}
	return nil
}

func (c *LruCache) Delete(key interface{}) {
	c.m.Lock()
	defer c.m.Unlock()
	if v, ok := c.cacheMap[key]; ok {
		c.cache.Remove(v)
		delete(c.cacheMap, key)
	}
}
