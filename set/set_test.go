package set

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	dataSet, _ := Build()

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
	assert.Equal(t, nil, dataSet.Del("hello"))
	assert.Equal(t, false, dataSet.Contains("hello"))
	assert.Equal(t, 1, dataSet.Size())
	assert.EqualError(t, dataSet.Del(1), (&UnsuitableTypeErr{Want: "string", Got: "int"}).Error())

	dataSet.Range(func(data interface{}) {
		fmt.Println(data)
	})

	otherSet := dataSet.Clone()
	assert.Equal(t, true, dataSet.Equal(otherSet))

	dataSet.Insert("hello")
	assert.Equal(t, false, dataSet.Equal(otherSet))

	set2, _ := Build("1", "2", "3")

	unionSet, err := dataSet.Union(set2)
	if err != nil {
		t.Fatal(err)
	}
	unionSet.Range(func(data interface{}) {
		fmt.Println(data)
	})

	set3, _ := Build(1, 2, 3)
	_, err = dataSet.Union(set3)
	if err == nil {
		t.Fatal("should failed")
	} else {
		t.Log(err)
	}

	dataSet.Clear()
	assert.Equal(t, true, dataSet.IsEmpty())
}

func TestDiff(t *testing.T) {
	s1, _ := Build("1", "2", "3")
	s2, _ := Build("1", "2", "4")

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

func TestIntersect(t *testing.T) {
	s1, _ := Build("1", "2", "3")
	s2, _ := Build("1", "2", "4")

	intersectSet, err := s1.Intersect(s2)
	if err != nil {
		t.Fatal(err)
	}

	if intersectSet.Contains("3") {
		t.Fatal("should not contains 3")
	}

	if !intersectSet.Contains("1") {
		t.Fatal("should contains 1")
	}

	if !intersectSet.Contains("2") {
		t.Fatal("should contains 2")
	}
}
