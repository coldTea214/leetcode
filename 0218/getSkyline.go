package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 扫描线问题
type pair struct {
	right, height int
}

type priorityQueue []pair

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].height > pq[j].height }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(v interface{}) { *pq = append(*pq, v.(pair)) }
func (pq *priorityQueue) Pop() interface{}   { a := *pq; v := a[len(a)-1]; *pq = a[:len(a)-1]; return v }

func getSkyline(buildings [][]int) (ans [][]int) {
	n := len(buildings)
	boundaries := make([]int, 0, n*2)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}
	sort.Ints(boundaries)

	idx := 0
	pq := priorityQueue{}
	for _, boundary := range boundaries {
		fmt.Println(boundary)
		for idx < n && buildings[idx][0] == boundary {
			heap.Push(&pq, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}
		fmt.Println(pq)
		for len(pq) > 0 && pq[0].right <= boundary {
			heap.Pop(&pq)
		}
		fmt.Println(pq)

		maxHeight := 0
		if len(pq) > 0 {
			maxHeight = pq[0].height
		}
		if len(ans) == 0 || maxHeight != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxHeight})
		}
		fmt.Println(ans)
		fmt.Println()
	}
	return
}

func main() {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	fmt.Println(buildings)
	ans := getSkyline(buildings)
	fmt.Println(ans)
	fmt.Println(`[[2 10] [3 15] [7 12] [12 0] [15 10] [20 8] [24 0]]`)
}
