package queue

type LinkedStackOfString struct {
	first *Node
}

type Node struct {
	item string
	next *Node
}

func (ls *LinkedStackOfString) Pop() string {
	i := ls.first.item
	ls.first = ls.first.next
	return i
}

func (ls *LinkedStackOfString) Push(i string) {
	oldFirst := ls.first
	ls.first = &Node{
		item: i,
		next: oldFirst,
	}
}

func (ls *LinkedStackOfString) IsEmpty() bool {
	return ls.first == nil
}