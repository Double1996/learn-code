package alg

// Q: 设计一个 LRU 缓存机制
/**
设计思路: 双向链表 + hashmap
如何写:
1. 两个结构体, LruCache,  DLinkNode //
2. 两个工厂方法: 构造 Lru, DLinkNode //
3. 2个 Lru 操作, Get. Put		  //
4. 4个双向链表操作	 添加到头，移动到头， 删除节点， 删除尾巴节点				  //
*/

type LruCache struct {
	size       int                // 元素个数
	capacity   int                // 容量
	cache      map[int]*DLinkNode // map 缓存
	head, tail *DLinkNode         // 头尾两个双向的链表, 用于放置元素
}

type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

func initDLinkNode(key, value int) *DLinkNode { // 记住入参
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LruCache { // 记住构造方法
	l := LruCache{capacity: capacity,
		cache: map[int]*DLinkNode{},
		head:  initDLinkNode(0, 0),
		tail:  initDLinkNode(0, 0),
	}
	l.head.next = l.tail //   定义头和尾部
	l.tail.prev = l.head //
	return l
}

func (this *LruCache) Get(key int) int { // 从key 获取
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node) // 获取到的话 放入 头部
	return node.value
}

func (this *LruCache) Put(key, value int) {
	if _, ok := this.cache[key]; !ok { // Get 和 Put 开头都要判断有没有 node
		node := initDLinkNode(key, value) // 没有的话新建 节点
		this.cache[key] = node
		this.addToHead(node) // 把节点 放入到 头之中
		this.size++          // size++
	}

	if this.size > this.capacity { // 超过了，容量
		removed := this.removeTail()
		delete(this.cache, removed.key)
		this.size--
	} else {
		node := this.cache[key] // 进行兜底操作，把这个 node
		node.value = value
		this.moveToHead(node)
	}
}

// 将此节点添加到头
func (this *LruCache) addToHead(node *DLinkNode) {
	node.prev = this.head      // 构建这个 节点的 链表指向
	node.next = this.head.next //
	this.head.next.prev = node // 头的前一个节点是 node ||  this.head   node   this.head.next
	this.head.next = node      //  构建指向
}

// 移动到头节点
func (this *LruCache) moveToHead(node *DLinkNode) {
	this.removeNode(node) // 删除原来的位置
	this.addToHead(node)  // 移动到头部节点
}

// 删除节点, 把中间的节点删除
func (this *LruCache) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 容量不够的时候
func (this *LruCache) removeTail() *DLinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
