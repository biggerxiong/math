package pq

import (
	"fmt"
)

// Item 是优先队列中包含的元素。
type Item struct {
	From int
	To   int

	Priority float64 // 元素在队列中的优先级。
	// 元素的索引可以用于更新操作，它由 heap.Interface 定义的方法维护。
	index int // 元素在堆中的索引。
}

func (i *Item) String() string {
	return fmt.Sprintf("Item{From:%d, To:%d, Priority:%f}", i.From, i.To, i.Priority)
}

// HeapItem 一个实现了 heap.Interface 接口的优先队列，队列中包含任意多个 Item 结构。
type HeapItem []*Item

func (pq HeapItem) Len() int { return len(pq) }

func (pq HeapItem) Less(i, j int) bool {
	// 我们希望 Pop 返回的是最小值而不是最大值，
	// 因此这里使用小于号进行对比。
	return pq[i].Priority < pq[j].Priority
}

func (pq HeapItem) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *HeapItem) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *HeapItem) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}

// 更新函数会修改队列中指定元素的优先级以及值。
// func (pq *HeapItem) update(item *Item, value string, Priority float64) {
// 	item.value = value
// 	item.Priority = Priority
// 	heap.Fix(pq, item.index)
// }
