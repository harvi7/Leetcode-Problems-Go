package design

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	length   int
	head     *Node // Dummy head node
	tail     *Node // Dummy tail node
	data     map[int]*Node
}

// addNode adds a new node to the head of the doubly linked list.
func (this *LRUCache) addNode(node *Node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

// removeNode removes a node from the doubly linked list.
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// moveToHead moves an existing node to the head of the linked list.
// This signifies that it has been "recently used."
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addNode(node)
}

func ConstructorLRUCache(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		length:   0,
		head:     head,
		tail:     tail,
		data:     make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.data[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; ok {
		// Key exists, update the value and move to head.
		node.value = value
		this.moveToHead(node)
		return
	}

	// Key is new
	node := &Node{key: key, value: value}
	this.data[key] = node
	this.addNode(node)
	this.length++

	// If we've exceeded capacity, remove the tail node.
	if this.length > this.capacity {
		nodeToRemove := this.tail.prev
		this.removeNode(nodeToRemove)
		delete(this.data, nodeToRemove.key)
		this.length--
	}
}
