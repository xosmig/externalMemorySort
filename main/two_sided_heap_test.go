package main

import (
	"container/heap"
	"reflect"
	"sort"
	"testing"
)

func testHeapSort(t *testing.T, h heap.Interface, items []uint64) {
	sorted := make([]uint64, len(items))
	copy(sorted, items)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	for _, item := range items {
		heap.Push(h, item)
	}

	if h.Len() != len(items) {
		t.Errorf("Expected len: %v, actual: %v", len(items), h.Len())
	}

	var heapSorted []uint64
	for h.Len() > 0 {
		heapSorted = append(heapSorted, heap.Pop(h).(uint64))
	}

	if !reflect.DeepEqual(heapSorted, sorted) {
		t.Errorf("Invalid heap sort result. Expected: %v, actual: %v", sorted, heapSorted)
	}
}

func TestLeftHeap_HeapSort(t *testing.T) {
	items := []uint64{1, 7, 2, 5, 5, 12, 2, 4}
	leftHeap, _ := NewTwoSidedHeap(len(items))
	testHeapSort(t, leftHeap, items)
}

func TestRightHeap_HeapSort(t *testing.T) {
	items := []uint64{1, 7, 2, 5, 5, 12, 2, 4}
	_, rightHeap := NewTwoSidedHeap(len(items))
	testHeapSort(t, rightHeap, items)
}