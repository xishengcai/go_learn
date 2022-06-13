package main

type LRUCache struct {
	capacity int
	cache    map[int]*node
	head     *node
	end      *node
	count    int
}

type node struct {
	pre   *node
	next  *node
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node, 0),
		head:     &node{nil, nil, -1, -1},
		end:      &node{nil, nil, -1, -1},
	}
}

func (this *LRUCache) get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.add(node)
	return node.value
}

func (this *LRUCache) put(key int, value int) {
	node := &node{
		key:   key,
		value: value,
	}

	x, ok := this.cache[key]

	if ok {
		this.remove(x)
	} else if this.capacity > this.count {
		this.count++
	} else if this.capacity == this.count {
		delete(this.cache, this.head.next.key)
		this.remove(this.head.next)
	}
	this.cache[key] = node
	this.add(node)

}

func (this *LRUCache) remove(node *node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) add(node *node) {
	// set head
	if this.head.next == nil {
		this.head.next = node
		this.end.pre = node
		node.pre = this.head
		node.next = this.end
		return
	}

	// set end
	node.pre = this.end.pre
	node.next = this.end
	this.end.pre = node
	node.pre.next = node

}

func main() {
	lRUCache := Constructor(2)
	lRUCache.put(1, 1)
	lRUCache.put(2, 2)
	lRUCache.get(1)    // 返回 1
	lRUCache.put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.get(2)    // 返回 -1 (未找到)
	lRUCache.put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.get(1)    // 返回 -1 (未找到)
	lRUCache.get(3)    // 返回 3
	lRUCache.get(4)

}
