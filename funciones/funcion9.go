package funciones

func InsertNode(val int, n *Nodo) int {

	comparations := 0

	if n.Value < val {

		comparations += 2
		if n.RightChild == nil {
			n.RightChild = &Nodo{Value: val, LeftChild: nil, RightChild: nil}
			return comparations
		}

		return comparations + InsertNode(val, n.RightChild)
	} else if n.Value > val {

		comparations += 3
		if n.LeftChild == nil {
			n.LeftChild = &Nodo{Value: val, LeftChild: nil, RightChild: nil}
			return comparations
		}

		return comparations + InsertNode(val, n.LeftChild)
	}

	return 2
}

func (t *Tree) Funcion9(val int) int {

	if t.Root == nil {
		t.Root = &Nodo{Value: val, LeftChild: nil, RightChild: nil}
		return 1
	}

	return 1 + InsertNode(val, t.Root)
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func HeightOf(root *Nodo) int {
	if root == nil {
		return -1
	}
	return 1 + Max(HeightOf(root.LeftChild), HeightOf(root.RightChild))
}

// func main() {
// 	t := Tree{Root: nil}

// 	x := t.InsertInTree(4)
// 	fmt.Println(*t.Root, x)

// 	x = t.InsertInTree(2)

// 	fmt.Println(*t.Root.LeftChild, x)

// 	x = t.InsertInTree(2)

// 	fmt.Println(*t.Root.LeftChild, x)

// 	x = t.InsertInTree(5)

// 	fmt.Println(*t.Root.RightChild, x)

// }
