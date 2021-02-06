package sort

import "errors"

type MinHeap interface {
	// Sort is an efficient way to prioritize a queue where there is a requirement of either getting the largest or
	// smallest item.
	Sort() []int
}

type minHeap struct {
	MinHeap
	Items []int
}

func NewMinHeap(items []int) (MinHeap, error) {
	if len(items) == 0 {
		return nil, errors.New("error: no such items were found")
	}

	return &minHeap{
		Items: items,
	}, nil
}

func (m *minHeap) isLeaf(index int, size int) bool {
	return index >= (size/2) && index <= size
}

func (m *minHeap) swap(first, second int) {
	temp := m.Items[first]
	m.Items[first] = m.Items[second]
	m.Items[second] = temp
}

func (m *minHeap) findSmallestChild(current int, size int) int {
	smallest := current

	leftChildIndex := 2*current + 1
	if leftChildIndex < size && m.Items[leftChildIndex] < m.Items[smallest] {
		smallest = leftChildIndex
	}

	rightChildIndex := 2*current + 2
	if rightChildIndex < size && m.Items[rightChildIndex] < m.Items[smallest] {
		smallest = rightChildIndex
	}

	return smallest
}

func (m *minHeap) heapify(current int, size int) {
	if m.isLeaf(current, size) {
		return
	}

	smallest := m.findSmallestChild(current, size)

	if smallest != current {
		m.swap(current, smallest)
		m.heapify(smallest, size)
	}
}

func (m *minHeap) build(size int) {
	for index := size/2 - 1; index >= 0; index-- {
		m.heapify(index, size)
	}
}

func (m *minHeap) Sort() []int {
	size := len(m.Items)

	m.build(size)

	for i := size - 1; i > 0; i-- {
		m.swap(0, i)
		m.heapify(0, i)
	}

	return m.Items
}
