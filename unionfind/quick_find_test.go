package quickfind

import (
	"reflect"
	"testing"
)

func TestInitializeQuickFind(t *testing.T) {
	cases := []struct {
		name string
		num int
		expected []int 
	} {
		{
			name: "Successfully initialize quick-find",
			num: 5,
			expected: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range cases {
		qf := New(tc.num)

		actual := qf.ids

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("%s: expected %v, but got: %v", tc.name, tc.expected, actual)
		}
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		name string
		num int
		target int
		merged int
		expected []int
	} {
		{
			name: "Successfully elements merged",
			num: 4,
			target: 0,
			merged: 1,
			expected: []int{0, 0, 2, 3},
		},
	}

	for _, tc := range cases {
		qf := New(tc.num)

		qf.Union(1, 0)

		actual := qf.ids

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("%s: expected %v, but got: %v", tc.name, tc.expected, actual)
		}
	}
} 

func TestConnected(t *testing.T) {
	cases := []struct {
		name string
		num int
		a int 
		b int 
		expected bool
	} {
		{
			name: "Successfully connected",
			num: 4,
			a: 0, 
			b: 1,
			expected: true,
		},
	}
	
	for _, tc := range cases {
		qf := New(tc.num)

		qf.Union(1, 0)

		actual := qf.Connected(0, 1)

		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("%s: expected %v, but got: %v", tc.name, tc.expected, actual)
		}
	}
}