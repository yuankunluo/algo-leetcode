package problem0332

import (
	"container/heap"
)

// MinHeap is to store destinations by it name, 
// sorted from small to large.
// We use stdlib container/heap interface.
type MinHeap []string

// Len is the length.
func (h MinHeap) Len() int {
	return len(h)
}

// Less define how to sort elements.
// This is Minheap, we use < operator to sort.
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap define how to swap element in heap.
func (h MinHeap) Swap(i, j int) {
	h[i],h[j] = h[j],h[i]
}

// Push define how to append element.
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

// Pop is the method to pop out the smallest element.
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x :=old[n-1]
	*h = old[0:n-1]
	return x
}


func findItinerary(tickets [][]string) []string {
	// First we define a map [city] => [city1, city2, city3]
	// to collect city's destinations by using MinHeap, which
	// sort destinations from small to large.
	tripsMap := make(map[string]*MinHeap)
	for _, t := range tickets {
		destinations, ok := tripsMap[t[0]]
		if !ok {
			h := make(MinHeap, 0)
			tripsMap[t[0]] = &h
			destinations = &h
		}
		heap.Push(destinations, t[1])
	 }
	 
	 // 
	 routes := make([]string, 0)

	 visit("JFK", tripsMap, &routes)

	return routes
}


func visit(city string, tripsMap map[string]*MinHeap, routes *[]string){
	destinations := tripsMap[city]
	for (*destinations).Len() != 0 {
		dest := heap.Pop(destinations)
		visit(dest.(string), tripsMap, routes)
	}
	*routes = append(*routes, city)
}

// reverse is a helper function to revers slice of strings
//
func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}