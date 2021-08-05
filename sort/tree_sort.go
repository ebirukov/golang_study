package sort

type tree struct {
	value       int
	left, right *tree
}

func BuildTree(values []int) *tree {
	var root *tree
	for _, value := range values {
		root = add(value, root)
	}
	return root
}

func add(value int, t *tree) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(value, t.left)
	} else {
		t.right = add(value, t.right)
	}
	return t
}
