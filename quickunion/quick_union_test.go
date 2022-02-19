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

func TestRoot(t *testing.T) {
	qu := &QuickUnion{
		ids: []int {0, 0, 0, 3},
	}

	cases := []struct {
		name string
		input int 
		expected int 
	} {
		{
			"Successfully find parent id",
			0,
			0,
		},
		{
			"Succssfully find child's parent",
			1,
			0,
		},
		{
			"Succssfully find orphan's id",
			3,
			3,
		},
	}

	for _, tt := range cases {
		got := qu.Root(tt.input)
		if got != tt.expected {
			t.Errorf("expected: %v but got %v", tt.expected, got)
		}
	}
}

func TestConnected(t *testing.T) {
	qu := &QuickUnion{
		ids: []int {0, 0, 2, 2, 3},
	}

	cases := []struct {
		name string
		input_a int 
		input_b int 
		expected bool 
	} {
		{
			"Successfully find connection between nodes",
			0,
			1,
			true,
		},
		{
			"Succsfully find unconnection between nodes",
			0,
			3,
			false,
		},
		{
			"Succsfully find connection between root and child of children",
			2,
			4,
			true,
		},
	}

	for _, tt := range cases {
		got := qu.Connected(tt.input_a, tt.input_b)
		if got != tt.expected {
			t.Errorf("expected: %v but got %v", tt.expected, got)
		}
	}
}

func TestUnion(t *testing.T) {
	qu := &QuickUnion{
		ids: []int {0, 0, 2, 2, 3},
	}

	cases := []struct {
		name string
		input_a int 
		input_b int 
		expected int
	} {
		{
			"Succssefully merge unconnected nodes",
			2, 
			0,
			0,
		},
		{
			"Succsfully merge unconnected nodes and its children",
			4,
			0,
			0,
		},
	}

	for _, tt := range cases {
		qu.Union(tt.input_a, tt.input_b)
		if qu.Root(tt.input_a) != tt.expected {
			t.Errorf("expected: %v but got %v", tt.expected, qu.Root(tt.input_a))
		}
	}

}