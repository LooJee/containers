package set

import (
	"fmt"
	"testing"
)

func TestSet_Insert(t *testing.T) {
	dataSet := NewSet()

	dataSet.Insert("hello")
	dataSet.Insert("world")

	dataSet.Iter(func(data interface{}) {
		fmt.Println(data)
	})

	if err := dataSet.Insert(1); err == nil {
		t.Fatal("should failed")
	} else {
		t.Log(err)
	}
}
