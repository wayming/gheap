package gheap

import "sort"

// Source - data source
type Source interface {
	sort.Interface
	Push(x any)
	Pop() any
}

// Heap min-heap
type Heap struct {
	data *Source
}

// Make an heap from data source
func Make(data *Source) Heap {
	var h Heap
	h.data = data
	return h
}

// Push an element into heap
func (heap *Heap) Push(x any) {

}

// Pop an element from heap
func (heap *Heap) Pop() any {

	return (*heap.data).Pop()
}
