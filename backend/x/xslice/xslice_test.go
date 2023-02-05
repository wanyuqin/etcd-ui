package xslice

import (
	"fmt"
	"testing"
)

func TestUniqueSlice(t *testing.T) {
	s := []string{"a", "v", "a"}

	us := UniqueSlice(s)
	fmt.Printf("%v\n", us)

	i := []int64{1, 2, 1, 3}
	ui := UniqueSlice(i)
	fmt.Printf("%v\n", ui)
}
