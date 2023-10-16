package main

import (
	"awesomeProject/heap"
	"fmt"
	"reflect"
)

type node struct {
	Value  int
	row    int
	column int
}

func toNode(a any) node {
	value := reflect.ValueOf(a).FieldByName("Value").Int()
	row := reflect.ValueOf(a).FieldByName("row").Int()
	column := reflect.ValueOf(a).FieldByName("column").Int()
	return node{int(value), int(row), int(column)}
}

func mergeSLices(nums [][]int) []int {
	capacity := 0
	for _, e := range nums {
		capacity += len(e)
	}
	merged := make([]int, 0, capacity)
	h := *heap.NewHeap(len(nums), node{})
	for i := range nums {
		h = heap.Add(h, node{nums[i][0], i, 0})
	}
	for heap.Len(h) > 0 {
		var reflectMin any
		h, reflectMin, _ = heap.ExtractMin(h)
		min := toNode(reflectMin)
		fmt.Println(min, heap.Len(h))
		merged = append(merged, min.Value)
		if min.column != len(nums[min.row])-1 {
			h = heap.Add(h, node{nums[min.row][min.column+1], min.row, min.column + 1})
		}
	}
	return merged
}

func main() {
	fmt.Println(mergeSLices([][]int{{1, 2, 3}, {6, 7, 9}, {1, 2}, {0, 10}}))
}
