package gheap

import (
	"fmt"
	"sort"
)

// Source - data source
type Source interface {
	sort.Interface
	PushBack(x any)
	PopBack() any
}

// Heap min-heap
// Heap min-heap
type Heap struct {
	data Source
}

// Make an heap from data source
func Make(data Source) Heap {
	var h Heap
	h.data = data

	for i := (data.Len() - 1) / 2; i >= 0; i-- {
		h.pushDown(i, data.Len()-1)
	}
	return h
}

// Push an element into heap
func (heap *Heap) Push(element any) {
	heap.data.PushBack(element)
	heap.popUp(heap.data.Len() - 1)
}

// Pop an element from heap
func (heap *Heap) Pop() any {
	if heap.data.Len() == 0 {
		return nil
	}
	heap.data.Swap(0, heap.data.Len()-1)
	popKey := heap.data.PopBack()
	heap.pushDown(0, heap.data.Len()-1)
	return popKey
}

func (heap *Heap) popUp(currIdx int) {
	parentIdx := (currIdx - 1) / 2
	fmt.Println("popUp: curr=", currIdx, " parent=", parentIdx)
	if heap.data.Less(currIdx, parentIdx) {
		// current key is less than its parent
		heap.data.Swap(currIdx, parentIdx)
		if parentIdx != 0 {
			heap.popUp(parentIdx)
		}
	}
}

func (heap *Heap) pushDown(currIdx int, lastIdx int) {
	leftIdx := 2*currIdx + 1
	rightIdx := 2*currIdx + 2
	minChildIdx := leftIdx
	fmt.Println("pushDown: curr=", currIdx, " last=", lastIdx, " left=", leftIdx, " right=", rightIdx)
	if leftIdx > lastIdx {
		// No left or right child
		return
	} else if rightIdx <= lastIdx {
		// Has both left and right child
		if heap.data.Less(rightIdx, leftIdx) {
			// Right child key is less than the left child key
			minChildIdx = rightIdx
		}
	}

	if heap.data.Less(minChildIdx, currIdx) {
		heap.data.Swap(minChildIdx, currIdx)
		heap.pushDown(minChildIdx, lastIdx)
	}
}
