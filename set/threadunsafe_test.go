package set

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	dataSet, _ := BuildThreadUnsafe()

	assert.Equal(t, true, dataSet.IsEmpty())

	dataSet.Insert("hello", "world")

	dataSet.Range(func(data interface{}) {
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
	dataSet.Del("hello")
	assert.Equal(t, false, dataSet.Contains("hello"))
	assert.Equal(t, 1, dataSet.Size())

	dataSet.Range(func(data interface{}) {
		fmt.Println(data)
	})

	otherSet := dataSet.Clone()
	assert.Equal(t, true, dataSet.Equal(otherSet))

	dataSet.Insert("hello")
	assert.Equal(t, false, dataSet.Equal(otherSet))

	dataSet.Clear()
	assert.Equal(t, true, dataSet.IsEmpty())
}

func TestThreadUnsafe_Union(t *testing.T) {
	s1, _ := BuildThreadUnsafe("1", "2", "3")
	s2, _ := BuildThreadUnsafe("1", "2", "4")

	unionSet, err := s1.Union(s2)
	if err != nil {
		t.Fatal(err)
	}

	if unionSet.Size() != 4 {
		t.Fatal("should be 4")
	}

	ss, _ := BuildThreadUnsafe("1", "2", "3", "4")
	if !ss.Equal(unionSet) {
		t.Fatal("should equal")
	}

	set3, _ := BuildThreadUnsafe(1, 2, 3)
	_, err = s1.Union(set3)
	if err == nil {
		t.Fatal("should failed")
	} else {
		t.Log(err)
	}
}

func TestThreadUnsafe_Diff(t *testing.T) {
	s1, _ := BuildThreadUnsafe("1", "2", "3")
	s2, _ := BuildThreadUnsafe("1", "2", "4")

	diffSet, err := s1.Diff(s2)
	if err != nil {
		t.Fatal(err)
	}

	if diffSet.Size() != 1 {
		t.Fatal("should be 1")
	}

	if !diffSet.Contains("3") {
		t.Fatal("should contains 3")
	}
}

func TestThreadUnsafe_Intersect(t *testing.T) {
	s1, _ := BuildThreadUnsafe("1", "2", "3")
	s2, _ := BuildThreadUnsafe("1", "2", "4")

	intersectSet, err := s1.Intersect(s2)
	if err != nil {
		t.Fatal(err)
	}

	if intersectSet.Contains("3") {
		t.Fatal("should not contains 3")
	}

	subSet, _ := BuildThreadUnsafe("1", "2")
	if !subSet.Equal(intersectSet) {
		t.Fatal("should equal")
	}
}
