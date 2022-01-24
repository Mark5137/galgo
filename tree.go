package galgo

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (tree *Tree) Insert(data int) *Tree {
	if tree.Root != nil {
		tree.Root.Insert(data)
	} else {
		tree.Root = &Node{Data: data, Left: nil, Right: nil}
	}

	return tree
}

func (node *Node) Insert(data int) {
	if data > node.Data {
		if node.Right != nil {
			node.Right.Insert(data)
		} else {
			node.Right = &Node{Data: data, Left: nil, Right: nil}
		}
	}

	if data < node.Data {
		if node.Left != nil {
			node.Left.Insert(data)
		} else {
			node.Left = &Node{Data: data, Left: nil, Right: nil}
		}
	}
}

func (tree *Tree) Search(data int) *Node {
	if tree.Root != nil {
		return tree.Root.Search(data)
	}

	return nil
}

func (node *Node) Search(data int) *Node {
	if node.Data == data {
		return node
	}

	if data > node.Data {
		if node.Right != nil {
			return node.Right.Search(data)
		}
	}

	if data < node.Data {
		if node.Left != nil {
			return node.Left.Search(data)
		}
	}

	return nil
}

func (tree *Tree) Remove(data int)  {
	if tree.Root == nil {
		return
	}

	if tree.Root.Data == data {
		if tree.Root.Left == nil && tree.Root.Right == nil {
			tree.Root = nil
		}

		if tree.Root.Left == nil {
			tree.Root = tree.Root.Right
		}

		if tree.Root.Right == nil {
			tree.Root = tree.Root.Left
		}

		if tree.Root.Left != nil && tree.Root.Right != nil {
			if tree.Root.Right.Left == nil {
				tree.Root.Right.Left = tree.Root.Left
				tree.Root = tree.Root.Right
			}
		}
	}
}

func (tree *Tree) Traverse() {
}
