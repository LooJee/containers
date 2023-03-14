package set

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	dataSet := New(false)

	assert.Equal(t, true, dataSet.IsEmpty())

	dataSet.Insert("hello", "world")

	dataSet.Iter(func(data interface{}) {
		fmt.Println(data)
	})

	if err := dataSet.Insert(1); err == nil {
		t.Fatal("should failed")
	} else {
		t.Log(err)
	}

	assert.Equal(t, 2, dataSet.Size())
	assert.Equal(t, true, dataSet.Contains("hello"))
	assert.Equal(t, false, dataSet.Contains(1))
	assert.Equal(t, false, dataSet.IsEmpty())
	assert.Equal(t, nil, dataSet.Del("hello"))
	assert.Equal(t, false, dataSet.Contains("hello"))
	assert.Equal(t, 1, dataSet.Size())
	assert.EqualError(t, dataSet.Del(1), (&UnsuitableTypeErr{Want: "string", Got: "int"}).Error())

	dataSet.Iter(func(data interface{}) {
		fmt.Println(data)
	})

	otherSet := dataSet.Clone()
	assert.Equal(t, true, dataSet.Equal(otherSet))

	dataSet.Insert("hello")
	assert.Equal(t, false, dataSet.Equal(otherSet))

	dataSet.Clear()
	assert.Equal(t, true, dataSet.IsEmpty())
}

func BenchmarkThreadSafe(b *testing.B) {
	dataSet := NewThreadSafe()
	fn := func() {
		for i := 0; i < 100000; i++ {
			dataSet.Insert(i)
		}
	}
	for i := 0; i < b.N; i++ {
		fn()
	}
}
