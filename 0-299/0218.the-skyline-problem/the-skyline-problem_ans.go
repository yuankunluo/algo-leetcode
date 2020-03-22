package problem0218

import (
	"container/heap"
)

type IntHeap [][]int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][2] > h[j][2] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.([]int))
}
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func getSkyline(buildings [][]int) [][]int {
    points:=[][]int{}
    bheap:=IntHeap{}
    heap.Init(&bheap)
    updateRight:=func(){
        cur:=(heap.Pop(&bheap)).([]int)
        height:=0
        for len(bheap)>0 && cur[1]>=bheap[0][1]{
            heap.Pop(&bheap)
        }
        if len(bheap)>0{
            height=bheap[0][2]
        }
        if height==cur[2]{
            return 
        }
        if len(points)>0 && points[len(points)-1][0]==cur[1]{
            if height>points[len(points)-1][1]{
                points[len(points)-1][1]=height
            }
        }else{
            points=append(points,[]int{cur[1],height})
        }
    }
    for _,v:=range buildings{
        for len(bheap)>0 && bheap[0][1]<v[0]{
            updateRight()
        }
        if len(bheap)==0 || bheap[0][2]<v[2]{
            if len(points)>0 && points[len(points)-1][0]==v[0]{
                if v[2]>points[len(points)-1][1]{
                    points[len(points)-1][1]=v[2]
                }
            }else{
                points=append(points,[]int{v[0],v[2]})
            }
        }
        heap.Push(&bheap,v)
    }
    for len(bheap)>0{
        updateRight()
    }
    return points
}