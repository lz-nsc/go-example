package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int
type any = interface{}

func (m MinHeap) Less(i, j int) bool {
	return m[i] <= m[j]
}

func (m MinHeap) Len() int {
	return len(m)
}

// Slice uses pointer to point to its data, when we change the data in slice
// we changed the data the pointer pointing to, the slice itself is not changed,
// so we don't need to use *MinHeap here
func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Assign a new MinHeap, so we use pointer here
func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}
func (m *MinHeap) Pop() any {
	size := len(*m)
	last := (*m)[size-1]
	*m = (*m)[:size-1]
	return last
}

func main() {
	fmt.Println("[Heap] Start heap tast, creating original array ...")
	mh := &MinHeap{5, 3, 7, 8, 2, 1}
	fmt.Printf("[Heap] Original array :%v\n", *mh)
	//Init
	fmt.Println("[Heap] Start to build minHeap...")
	heap.Init(mh)
	fmt.Printf("[Heap] MinHeap:%v\n", *mh)
	//Pop
	first := heap.Pop(mh)
	fmt.Printf("[Heap] Pop:%d MinHeap:%v\n", first, *mh)
	//Push
	heap.Push(mh, 4)
	fmt.Printf("[Heap] Push:%d MinHeap:%v\n", 4, *mh)
	//Fix(down)
	fmt.Printf("[Heap] Changing one node of minHeap %v..\n", *mh)
	(*mh)[2] = 6
	fmt.Printf("[Heap] MinHeap will be fixed:%v\n", *mh)
	heap.Fix(mh, 2)
	fmt.Printf("[Heap] MinHeap after fix:%v\n", *mh)
	//Fix(up)
	fmt.Printf("[Heap] Changing one node of minHeap %v..\n", *mh)
	(*mh)[1] = 1
	heap.Fix(mh, 1)
	fmt.Printf("[Heap] MinHeap after fix:%v\n", *mh)
	//Remove
	fmt.Printf("[Heap] Removing one node from minHeap %v..\n", *mh)
	elem := heap.Remove(mh, 1)
	fmt.Printf("[Heap] MinHeap after removing element %d :%v\n", elem, *mh)
}
