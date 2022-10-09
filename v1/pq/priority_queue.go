package pq

import (
	"container/heap"
	"fmt"
)

type PriorityQueue struct {
	items HeapItem
	count int
}

func (pq *PriorityQueue) String() string {
	return fmt.Sprintf("PriorityQueue{count:%d, items:%v}", pq.count, pq.items)
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		items: make(HeapItem, 0),
	}
}

func (pq *PriorityQueue) Push(item *Item) {
	item.index = pq.count
	pq.count++
	heap.Push(&pq.items, item)
}

func (pq *PriorityQueue) Pop() *Item {
	return heap.Pop(&pq.items).(*Item)
}

func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}
