package tree

type BSTree struct {
	head      *Node
	nodeCount int
	//rwm sync.RWMutex
}

func NewBSTree() *BSTree {
	return &BSTree{}
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (b *BSTree) Insert(value int) {
	if b.head == nil {
		b.head = &Node{value, nil, nil}
		return
	}

	node := b.head
	for node != nil {
		if value >= node.Value {
			if node.Right == nil {
				node.Right = &Node{value, nil, nil}
				return
			}
			node = node.Right
		} else {
			if node.Left == nil {
				node.Left = &Node{value, nil, nil}
				return
			}
			node = node.Left
		}
	}

	return
}

func (b *BSTree) Delete(value int) bool {
	return true
}

func (b BSTree) ModPrint() []int {
	result := make([]int, 0, b.nodeCount)
	node := b.head

	var modPrintCore func(node *Node)
	modPrintCore = func(node *Node) {
		if node.Left != nil {
			modPrintCore(node.Left)
		}
		result = append(result, node.Value)
		if node.Right != nil {
			modPrintCore(node.Right)
		}
	}

	modPrintCore(node)
	return result
}

func (b BSTree) Max() *Node {
	return &Node{}
}

func (b BSTree) Min() *Node {
	return &Node{}
}

func (b BSTree) Front(value int) *Node {
	return &Node{}
}

func (b BSTree) Backend(value int) *Node {
	return &Node{}
}

func (b BSTree) Find(value int) *Node {
	node := b.head
	for node != nil {
		if value > node.Value {
			node = node.Right
			continue
		}
		if value < node.Value {
			node = node.Left
			continue
		}
		return node
	}
	return nil
}
