package heap

type Node struct {
	Key int
}

type Heap struct {
	heapArray []*Node
}

func (heap *Heap) PrintHeap() {
	for _, node := range heap.heapArray {
		println(node.Key)
	}
}

func (heap *Heap) Insert(key int) int {
	if heap.heapArray == nil {
		heap.heapArray = make([]*Node, 0)
	}
	heap.heapArray = append(heap.heapArray, &Node{Key: key})
	return heap.trickleUp(len(heap.heapArray) - 1)
}

func (heap *Heap) trickleUp(index int) int {
	parent := (index - 1) / 2
	bottom := heap.heapArray[index]
	for index > 0 && heap.heapArray[parent].Key < bottom.Key {
		heap.heapArray[index] = heap.heapArray[parent]
		index = parent
		parent = (index - 1) / 2
	}
	heap.heapArray[index] = bottom
	return index
}

func (heap *Heap) Remove() *Node {
	if heap.heapArray == nil || len(heap.heapArray) == 0 {
		return nil
	}
	root := heap.heapArray[0]
	heap.heapArray[0] = heap.heapArray[len(heap.heapArray)-1]
	heap.heapArray = heap.heapArray[:len(heap.heapArray)-1] // cut last element of slice
	heap.trickleDown(0)
	heap.PrintHeap()
	return root
}

func (heap *Heap) trickleDown(index int) int {
	arr := heap.heapArray
	arrSize := len(arr)
	if arrSize == 0 {
		return 0
	}
	movedTop := arr[index]
	for index < arrSize/2 {
		var largerChild int
		leftChild := 2*index + 1
		rightChild := leftChild + 1
		if rightChild < arrSize && arr[leftChild].Key < arr[rightChild].Key {
			largerChild = rightChild
		} else {
			largerChild = leftChild
		}
		if movedTop.Key >= arr[largerChild].Key {
			break
		}
		arr[index] = arr[largerChild]
		index = largerChild
	}
	heap.heapArray[index] = movedTop
	return index
}
