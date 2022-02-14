package queue

import (
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	cases := [] struct {
		name string
		givenQueue *LinkedStackOfString
		expected bool
	} {
		{
			name: "Successfully decide stack is empty",
			givenQueue: &LinkedStackOfString{},
			expected: true,
		},
		{
			name: "Successfully decide stack is not empty",
			givenQueue: &LinkedStackOfString{
				first: &Node{
					item: "first node",
					next: nil,
				},
			},
			expected: false,
		},
	}

	for _, tc := range cases {
		actual := tc.givenQueue.IsEmpty()
	
		if actual != tc.expected {
			t.Errorf("%s: expected: %v, but got: %v", tc.name, tc.expected, actual)
		}
	}
}

func TestPush(t *testing.T) {
	cases := []struct {
		name string 
		givenQueue *LinkedStackOfString
		input string
		expected *LinkedStackOfString
	} {
		{
			name: "Successfully push node when LinkedStackOfString is empty",
			givenQueue: &LinkedStackOfString{
				first: nil,
			},
			input: "s",
			expected: &LinkedStackOfString {
				first: &Node{
					item: "s",
					next: nil,
				},
			},
		},
	}

	for _, tc := range cases {
		tc.givenQueue.Push(tc.input)

		if !reflect.DeepEqual(tc.givenQueue, tc.expected) {
			t.Errorf("%s expected: %v, but got: %v", tc.name, tc.expected, tc.givenQueue)
		}
	}

}

func TestPop(t *testing.T) {
	cases := []struct {
		name string 
		givenQueue *LinkedStackOfString
		expected *LinkedStackOfString
	} {
		{
			"Successfully pop element when queue has enough elements",
			&LinkedStackOfString{
				first: &Node{
					item: "s",
					next: &Node{
						item: "v",
						next: nil,
					},
				},
			},
			&LinkedStackOfString{
				first: &Node{
					item: "v",
					next: nil,
				},
			},
		},
	}

	for _, tc := range cases {
		tc.givenQueue.Pop()


		if !reflect.DeepEqual(tc.givenQueue, tc.expected) {
			t.Errorf("%s expected %v but got %v", tc.name, tc.expected, tc.givenQueue)
		}
	}
}