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

func (t *Tree) Size() int {
	return t.Root.Size()
}

// Retorna el tamano del arbol
func (n *Nodo) Size() int {
	if n == nil {
		return 0
	}

	counter := 1

	if n.RightChild != nil {
		counter += n.RightChild.Size()
	}

	if n.LeftChild != nil {
		counter += n.RightChild.Value
	}

	return counter
}
