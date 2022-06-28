package queue

// LinkNode 链表节点
type LinkNode struct {
	Pre   *LinkNode
	Next  *LinkNode
	Value interface{}
	Key   string
}

// LinkQueue 链表队列
type LinkQueue struct {
	head *LinkNode
	tail *LinkNode
	Cap  int
	Size int
}

// InitLinkQueue 初始化一个链表队列
//  - cap: 队列的容积
func InitLinkQueue(cap int) (linkQueue *LinkQueue) {
	linkQueue = &LinkQueue{Cap: cap}
	linkQueue.head = linkQueue.tail
	return
}

// Push 向队列推送一个值
//  - value: 需要推送到队列里面的值
func (l *LinkQueue) Push(value interface{}) {
	node := LinkNode{Pre: l.tail, Value: value}
	l.Size++
	if l.tail == nil {
		l.tail = &node
		l.head = l.tail
		return
	}
	l.tail.Next = &node
	l.tail = l.tail.Next
	if l.Size > l.Cap {
		l.Pop()
	}
}

// Pop 从队列出一个值
//  - isCnt: 是否需要忽略计数，默认不需要忽略
func (l *LinkQueue) Pop() (result interface{}) {
	if l.head == nil || l.Size == 0 {
		return
	}
	l.Size--
	node := l.head
	l.head = l.head.Next
	node.Pre = nil
	node.Next = nil
	result = node.Value
	if l.Size == 0 {
		l.tail = l.head
	}
	return
}

// PushNode 向队列推送一个值
//  - value: 需要推送到队列里面的值
func (l *LinkQueue) PushNode(node *LinkNode) (isFlow bool, flowKey string) {
	node.Pre = l.tail
	l.Size++
	if l.tail == nil {
		l.tail = node
		l.head = l.tail
		return
	}
	l.tail.Next = node
	l.tail = l.tail.Next
	if l.Size > l.Cap {
		isFlow = true
		flowKey = l.PopNode().Key
	}
	return
}

// PopNode 从队列出一个值
//  - isCnt: 是否需要忽略计数，默认不需要忽略
func (l *LinkQueue) PopNode() (node *LinkNode) {
	if l.head == nil || l.Size == 0 {
		return
	}
	l.Size--
	node = l.head
	l.head = l.head.Next
	l.head.Pre = nil
	node.Pre = nil
	node.Next = nil
	if l.Size == 0 {
		l.tail = l.head
	}
	return
}
