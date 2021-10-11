package funciones

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func Funcion10(num int, root *TreeNode) (bool, int) {
	node := root
	counter := 0

	for node != nil {
		counter++
		if node.val > num {
			node = node.left
		} else if node.val < num {
			node = node.right
		} else {
			return true, counter
		}
	}
	return false, counter
}
