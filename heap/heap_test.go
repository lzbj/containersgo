package heap

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{7, -1, 3, 1, 2, 5, 8, 15}
	Sort(a)
	assert.Equal(t, a, []int{-1, 1, 2, 3, 5, 7, 8, 15})
}
