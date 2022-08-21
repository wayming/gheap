package gheap_test

import (
	"fmt"
	"gheap"
	"reflect"
	"testing"
)

func TestInitIntHeap(t *testing.T) {
	testData := &gheap.IntHeap{10, 55, 18, 589, 29, 17, 3, 45, 333, 12, 1, 109}
	sortedData := []int{1, 3, 10, 12, 17, 18, 29, 45, 55, 109, 333, 589}

	h := gheap.Make(testData)
	len := testData.Len()
	result := make([]int, 0)
	for i := 0; i < len; i++ {
		result = append(result, h.Pop().(int))
	}
	fmt.Println(result)
	if !reflect.DeepEqual(result, sortedData) {
		t.Log(result, " is not the same as ", sortedData)
		t.Fail()
	}
}
