package design_hashmap

type Node struct {
	key   int
	value int
	next  *Node
}

type MyHashMap struct {
	size int
	data []*Node
}

func Constructor() MyHashMap {
	return MyHashMap{
		size: 1000,
		data: make([]*Node, 1000),
	}
}
func (m *MyHashMap) Put(key int, value int) {
	position := key % m.size

	newNode := &Node{key: key, value: value, next: nil}
	if m.data[position] == nil {
		m.data[position] = newNode
		return
	}

	head := m.data[position]
	for head != nil {
		if head.key == key {
			head.value = value
			return
		}

		head = head.next
	}

	newNode.next = m.data[position]
	m.data[position] = newNode
}
func (m *MyHashMap) Get(key int) int {
	position := key % m.size

	head := m.data[position]

	for head != nil {
		if head.key == key {
			return head.value
		}

		head = head.next
	}

	return -1
}
func (m *MyHashMap) Remove(key int) {
	position := key % m.size
	head := m.data[position]

	if head == nil {
		return
	}

	if head.key == key {
		m.data[position] = head.next
		return
	}

	prev := head
	curr := head.next
	for curr != nil {
		if curr.key == key {
			prev.next = curr.next
			return
		}
		prev = curr
		curr = curr.next
	}
}
