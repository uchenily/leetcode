package leetcode

import "container/heap"

// 优先队列(二叉堆, 最小堆实现, 拷贝自go源码/src/container/heap/example_pq_test.go, 略有修改)

// Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// 未使用
// // update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, value int, priority int) {
// 	item.value = value
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	pq := make(PriorityQueue, len(m))
	i := 0
	for k, v := range m {
		pq[i] = &Item{
			value:    k,
			priority: v,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	result := make([]int, 0)
	for j := 0; j < k; j++ {
		result = append(result, heap.Pop(&pq).(*Item).value)
	}
	return result
}
