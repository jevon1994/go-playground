package tree

import "testing"

func TestSet(t *testing.T) {
	sets := make([]*Set, max)
	for i := 0; i < max; i++ {
		sets[i] = NewSet(i+1, -1)
	}

}
