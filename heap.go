package heap

import (
	"errors"
	"reflect"
)

func lessOrEqual(a, b any) bool {
	aValue := reflect.Indirect(reflect.ValueOf(a)).FieldByName("Value").Int()
	bValue := reflect.Indirect(reflect.ValueOf(b)).FieldByName("Value").Int()
	return aValue <= bValue
}

type Heap struct {
	size     int
	capacity int
	dataType reflect.Type
	nodes    reflect.Value
}

func NewHeap(capacity int, t any) *Heap {
	dataType := reflect.TypeOf(t)
	return &Heap{
		0,
		capacity,
		dataType,
		reflect.MakeSlice(reflect.SliceOf(dataType), capacity, capacity),
	}
}

func Len(self Heap) int {
	return self.size
}

func Cap(self Heap) int {
	return self.capacity
}

func Min(self Heap) (any, error) {
	if self.size == 0 {
		return -1, errors.New("Heap is empty. Can't get Min")
	}
	return self.nodes.Index(0).Interface(), nil
}

func ExtractMin(self Heap) (Heap, interface{}, error) {
	if self.size == 0 {
		return self, -1, errors.New("Heap is empty. Can't get Min")
	}
	reflectValue := self.nodes.Index(0).Interface()
	self = swap(self, 0, self.size-1)
	self.size--
	self = siftDown(self, 0)
	return self, reflectValue, nil
}

func Add(self Heap, element any) Heap {
	self.size++
	if self.size > self.capacity {
		self.capacity++
		self.nodes = reflect.Append(self.nodes, reflect.ValueOf(element))
	} else {
		self.nodes.Index(self.size - 1).Set(reflect.ValueOf(element))
	}
	self = siftUp(self, self.size-1)
	return self
}

func swap(self Heap, i, j int) Heap {
	buf := self.nodes.Index(i).Interface()
	self.nodes.Index(i).Set(self.nodes.Index(j))
	self.nodes.Index(j).Set(reflect.ValueOf(buf))
	return self
}

func siftDown(self Heap, i int) Heap {
	if i*2+1 >= self.size {
		return self
	}
	value := self.nodes.Index(i).Interface()
	left := self.nodes.Index(i*2 + 1).Interface()
	if i*2+2 == self.size {
		if lessOrEqual(value, left) {
			return self
		}
		self = swap(self, i, i*2+1)
		self = siftDown(self, i*2+1)
		return self
	}
	right := self.nodes.Index(i*2 + 2).Interface()
	if lessOrEqual(value, left) && lessOrEqual(value, right) {
		return self
	}
	if lessOrEqual(left, right) {
		self = swap(self, i, i*2+1)
		self = siftDown(self, i*2+1)
	} else {
		self = swap(self, i, i*2+2)
		self = siftDown(self, i*2+2)
	}
	return self
}

func siftUp(self Heap, i int) Heap {
	if i == 0 {
		return self
	}
	value := self.nodes.Index(i).Interface()
	parent := self.nodes.Index((i - 1) / 2).Interface()
	if lessOrEqual(parent, value) {
		return self
	}
	self = swap(self, i, (i-1)/2)
	self = siftUp(self, (i-1)/2)
	return self
}
