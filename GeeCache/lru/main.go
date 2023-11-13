package lru

import (
	"container/list"
)

type Cache struct {
	maxBytes  int64
	nbytes    int64
	ll        *list.List                    //双向链表
	cache     map[string]*list.Element      //值是对应节点指针
	OnEvicted func(key string, value Value) //？？？OnEvicted 是某条记录被移除时的回调函数，
}
type entry struct {
	key   string
	value Value
}
type Value interface {
	len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 查找：
// 1.从字典中找到对应的双向链表节点，
// 2.将节点移动到队尾
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 删除
// 从链表删除对头元素
// 从字典中删除对应元素
// 更新nbytes
func (c Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(kv.value.len() + len(kv.key))
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 新增/修改
// 从字典里寻得到就更新寻不到就新增
// 更新已使用内存，更新位置
func (c Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.len()) - int64(kv.value.len())
		kv.value = value
	} else {
		ele := entry{
			key:   key,
			value: value,
		}
		c.ll.PushFront(ele)
		c.nbytes += int64(value.len()) + int64(len(key))
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}
func (c *Cache) Len() int {
	return c.ll.Len()
}
