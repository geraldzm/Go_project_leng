package funciones

func InsertNode(val int, n *Nodo) int {

	comparations := 0

	if n.value < val {

		comparations += 2
		if n.rightChild == nil {
			n.rightChild = &Nodo{value: val, leftChild: nil, rightChild: nil}
			return comparations
		}

		return comparations + InsertNode(val, n.rightChild)
	} else if n.value > val {

		comparations += 3
		if n.leftChild == nil {
			n.leftChild = &Nodo{value: val, leftChild: nil, rightChild: nil}
			return comparations
		}

		return comparations + InsertNode(val, n.leftChild)
	}

	return 2
}

func (t *Tree) InsertInTree(val int) int {

	if t.root == nil {
		t.root = &Nodo{value: val, leftChild: nil, rightChild: nil}
		return 1
	}

	return 1 + InsertNode(val, t.root)
}

// func main() {
// 	t := Tree{root: nil}

// 	x := t.InsertInTree(4)
// 	fmt.Println(*t.root, x)

// 	x = t.InsertInTree(2)

// 	fmt.Println(*t.root.leftChild, x)

// 	x = t.InsertInTree(2)

// 	fmt.Println(*t.root.leftChild, x)

// 	x = t.InsertInTree(5)

// 	fmt.Println(*t.root.rightChild, x)

// }
