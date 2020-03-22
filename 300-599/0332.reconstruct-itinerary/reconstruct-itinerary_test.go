package problem0332

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCase = [][]string{
	[]string{"MUC", "LHR"},
	[]string{"JFK", "MUC"},
	[]string{"SFO", "SJC"},
	[]string{"LHR", "SFO"},
}

var output = []string{"JFK", "MUC", "LHR", "SFO", "SJC"}

func TestMinHeap(t *testing.T) {
	var h MinHeap
	heap.Push(&h, "c")
	heap.Push(&h, "a")
	heap.Push(&h, "d")
	heap.Push(&h, "b")
	p := heap.Pop(&h)
	ast := assert.New(t)
	ast.Equal(p, "a")
}

func TestFindItinerary(t *testing.T) {
	result := findItinerary(testCase)
	ast := assert.New(t)
	ast.Equal(result, output)
}
