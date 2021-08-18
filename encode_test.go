package gojson

import (
	"fmt"
	"testing"
)

func makeParams(array ...interface{}) []byte {
	enc := NewEncoder()
	defer enc.Release()
	enc.EncodeArray(array)
	return enc.Bytes()
}

func makeSlice(array ...interface{}) []interface{} {
	return array
}

func TestEncoder(t *testing.T) {
	//fmt.Println(string(makeParams(100, true, "abc", []string{"see", "what"}, [][]int{[]int{1, 3, 5}, []int{9, 8}})))
	fmt.Println(string(makeParams(100, true, "abc", makeSlice("see", "what"), makeSlice(makeSlice(1, 3, 5), makeSlice(9, 8)))))
}
