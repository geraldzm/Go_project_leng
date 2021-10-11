package funciones

func Funcion10(num int, tree Tree) (bool, int) {
	node := tree.Root
	counter := 1

	for node != nil {

		if node.Value > num {
			node = node.LeftChild
		} else if node.Value < num {
			node = node.RightChild
		} else {
			counter += 2
			return true, counter
		}

		counter += 3
	}

	return false, counter
}
