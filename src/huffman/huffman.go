package huffman

import "container/heap"

type HuffmanNode struct {
	frequency int
	val       byte
	index     int
	left      *HuffmanNode
	right     *HuffmanNode
}

type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].frequency < pq[j].frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	var n int = len(*pq)
	item := x.(*HuffmanNode)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func GenerateFrequencies(input string) *[256]int {
	var frequency = [256]int{}
	for i := 0; i < len(input); i++ {
		frequency[input[i]]++
	}
	return &frequency
}

func BuildHeap(frequencies *[256]int) *PriorityQueue {
	var nodes PriorityQueue = make([]*HuffmanNode, 256)
	for i := 0; i < 256; i++ {
		nodes[i] = &HuffmanNode{
			frequency: frequencies[i],
			val:       uint8(i),
			index:     i,
			left:      nil,
			right:     nil,
		}
	}
	heap.Init(&nodes)
	return &nodes
}

func BuildTree(pq *PriorityQueue) *HuffmanNode {

}
