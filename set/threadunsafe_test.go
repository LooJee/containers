package set

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	dataSet := BuildSet[string]()

	assert.Equal(t, true, dataSet.IsEmpty())

	dataSet.Insert("hello", "world")

	dataSet.Range(func(data string) {
		fmt.Println(data)
	})

	assert.Equal(t, 2, dataSet.Size())
	assert.Equal(t, true, dataSet.Contains("hello"))
	assert.Equal(t, false, dataSet.IsEmpty())
	dataSet.Del("hello")
	assert.Equal(t, false, dataSet.Contains("hello"))
	assert.Equal(t, 1, dataSet.Size())

	dataSet.Range(func(data string) {
		fmt.Println(data)
	})

	otherSet := dataSet.Clone()
	assert.Equal(t, true, dataSet.Equal(otherSet))

	dataSet.Insert("hello")
	assert.Equal(t, false, dataSet.Equal(otherSet))

	dataSet.Clear()
	assert.Equal(t, true, dataSet.IsEmpty())
}

func TestSet_Union(t *testing.T) {
	s1 := BuildSet("1", "2", "3")
	s2 := BuildSet("1", "2", "4")

	unionSet := s1.Union(s2)

	if unionSet.Size() != 4 {
		t.Fatal("should be 4")
	}

	ss := BuildSet("1", "2", "3", "4")
	if !ss.Equal(unionSet) {
		t.Fatal("should equal")
	}
}

func TestSet_Diff(t *testing.T) {
	s1 := BuildSet("1", "2", "3")
	s2 := BuildSet("1", "2", "4")

	diffSet := s1.Diff(s2)

	if diffSet.Size() != 1 {
		t.Fatal("should be 1")
	}

	if !diffSet.Contains("3") {
		t.Fatal("should contains 3")
	}
}

func TestSet_Intersect(t *testing.T) {
	s1 := BuildSet("1", "2", "3")
	s2 := BuildSet("1", "2", "4")

	intersectSet := s1.Intersect(s2)

	if intersectSet.Contains("3") {
		t.Fatal("should not contains 3")
	}

	subSet := BuildSet("1", "2")
	if !subSet.Equal(intersectSet) {
		t.Fatal("should equal")
	}
}
