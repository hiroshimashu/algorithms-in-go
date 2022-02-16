package quickunion

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	num_ids := 5
	expected := &QuickUnion{
		ids: []int {0, 1, 2, 3, 4},
	}

	got := New(num_ids)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v but got:%v", expected ,got)
	}
}