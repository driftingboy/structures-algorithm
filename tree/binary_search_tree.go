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
	// 0.find it and parent node
	node := b.head
	var pNode *Node
	for node != nil {
		if node.Value == value {
			break
		} else if value > node.Value {
			pNode = node
			node = node.Right
		} else {
			pNode = node
			node = node.Left
		}
	}
	if node == nil {
		return true
	}

	// 1.如果存在两个子结点
	if node.Left != nil && node.Right != nil {
		// 查找右子树最小结点(或左子树最大节点)
		minRightSubPNode := node
		minRightSubNode := node.Right
		for minRightSubNode.Left != nil {
			minRightSubPNode = minRightSubNode
			minRightSubNode = minRightSubNode.Left
		}
		node.Value = minRightSubNode.Value
		// 待删除结点变为一个、零个子节点的情况
		pNode = minRightSubPNode
		node = minRightSubNode
	}

	// 2.如果只有1、0个子节点
	onlyHasRight := node.Left == nil && node.Right != nil
	onlyHasLeft := node.Left != nil && node.Right == nil
	hasNoChild := node.Left == nil && node.Right == nil
	if onlyHasLeft {
		if pNode == nil {
			b.head = node.Left
		} else {
			pNode.Left = node.Left
		}
	} else if onlyHasRight {
		if pNode == nil {
			b.head = node.Right
		} else {
			pNode.Right = node.Right
		}
	} else if hasNoChild {
		if pNode == nil {
			b.head = nil
		} else if pNode.Left == node {
			pNode.Left = nil
		} else {
			pNode.Right = nil
		}
	}

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
