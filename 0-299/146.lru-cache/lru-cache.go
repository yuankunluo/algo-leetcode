type LRUCache struct {
	capacity int
	cache map[int]*Node 
	head, end *Node
}

type Node struct {
	key, value int
	pre, next *Node
}


func Constructor(capacity int) LRUCache {
	cache := make(map[int]*Node, capacity)
	return LRUCache{
		capacity: capacity,
		cache: cache,
		head: nil, 
		end: nil,
	}
}


func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.setHead(node)
	return node.value
}


func (this *LRUCache) remove(node *Node) {
	if node.pre == nil {
		this.head = node.next
	} else {
		node.pre.next = node.next
	}

	if node.next == nil {
		this.end = node.pre 
	} else {
		node.next.pre = node.pre
	}
}

func (this *LRUCache) setHead(node *Node){
	node.pre = nil
	if this.head == nil {
		this.head = node
	} else {
		this.head.pre, node.next = node, this.head
		this.head = node
	}

	if this.end == nil {
		this.end = node
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.remove(node)
		this.setHead(node)
	} else {
		if len(this.cache) >= this.capacity {
			delete(this.cache, this.end.key)
			this.remove(this.end)
		}
		node := &Node{key:key, value:value}
		this.cache[key] = node 
		this.setHead(node)
	}
}
