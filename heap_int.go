package gheap

type IntHeap []int

func (heap IntHeap) Less(k1, k2 int) bool { return heap[k1] < heap[k2] }
func (heap IntHeap) Len() int             { return heap.Len() }
func (heap IntHeap) Swap(k1, k2 int)      { heap[k1], heap[k2] = heap[k2], heap[k1] }
func (heap *IntHeap) Push(element any)    { *heap = append(*heap, element.(int)) }
func (heap *IntHeap) Pop() any {
	h := *heap
	lastVal := h[h.Len()-1]
	*heap = h[0 : h.Len()-1]
	return lastVal
}
