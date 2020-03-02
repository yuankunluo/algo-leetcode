package problem0692

import (
	"container/heap"
)

type Item struct {
	value    string
	priority int
}

type PriorityQueue []Item

func (q PriorityQueue) Len() int {
	return len(q)
}

func (q PriorityQueue) Less(i, j int) bool {
	if q[i].priority == q[j].priority {
		return q[i].value > q[j].value
	}
	return q[i].priority < q[j].priority
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	*q = append(*q, item)
}

func (q *PriorityQueue) Pop() interface{} {
	old := *q
	item := old[len(old)-1]
	*q = old[0 : len(old)-1]
	return item
}

func topKFrequent(words []string, k int) []string {
	wordMap := make(map[string]int, 0)
	for _, word := range words {
		wordMap[word]++
	}
	q := make(PriorityQueue, 0)
	heap.Init(&q)
	for w, c := range wordMap {
		item := Item{
			priority: c,
			value:    w,
		}
		heap.Push(&q, item)
		if q.Len() > k {
			heap.Pop(&q)
		}
	}
	ans := make([]string, 0)
	for q.Len() > 0 {
		item := heap.Pop(&q).(Item)
		ans = append([]string{item.value}, ans...)
	}

	return ans
}
