package main

// TODO: solve using priority queue
import (
	"container/heap"
)

type Item struct {
	priority int
	index    int // index in the heap
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// priority queue can give the lowest or highest element depending on the use case and the implementation
	// We want Pop to give the lowest, priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pq := make(PriorityQueue, len(nums))

	for i, v := range nums {
		pq[i] = &Item{index: i, priority: v}
	}

	heap.Init(&pq)
	prev := heap.Pop(&pq).(*Item)

	currL := 0
	maxL := 0
	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Item)

		if curr.priority == prev.priority+1 {
			currL++
		} else if curr.priority == prev.priority {
			// nothing to do for this case
		} else {
			if currL+1 > maxL {
				maxL = currL + 1
			}
			currL = 0

		}
		prev = curr
	}

	if currL+1 > maxL {
		maxL = currL + 1
	}

	return maxL
}

func longestConsecutiveTLE(nums []int) int {
	m := make(map[int]int)

	nMax := -1000000000
	nMin := 1000000000

	lMax := -1
	for _, v := range nums {
		m[v] = 1
		if v > nMax {
			nMax = v
		}
		if v < nMin {
			nMin = v
		}
	}

	j := nMin
	l := 0
	for j <= nMax {
		if m[j] != 0 {
			l += m[j]
		} else {
			if l > lMax {
				lMax = l
			}
			l = 0
		}
		j++
	}

	if l > lMax {
		lMax = l
	}

	return lMax
}
