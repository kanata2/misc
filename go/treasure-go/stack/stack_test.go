package main

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	t.Parallel()

	cases := []struct {
		current  []string
		expected string
	}{
		{current: []string{}, expected: ""},
		{current: []string{"C", "B", "A"}, expected: "A"},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			s := &Stack{mu: &sync.Mutex{}, buf: tc.current, limit: 10}
			if got := s.Pop(); got != tc.expected {
				t.Errorf("got %s but want %s", got, tc.expected)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	cases := []struct {
		stack    *Stack
		in       string
		expected []string
	}{
		{
			stack: &Stack{
				buf:   []string{"a", "b", "c"},
				limit: 10,
			},
			in:       "d",
			expected: []string{"a", "b", "c", "d"},
		},
		{
			stack: &Stack{
				buf:   []string{"a", "b", "c"},
				limit: 3,
			},
			in:       "d",
			expected: []string{"b", "c", "d"},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tc.stack.mu = &sync.Mutex{}
			if tc.stack.Push(tc.in); !reflect.DeepEqual(tc.stack.buf, tc.expected) {
				t.Errorf("got %v, but want %v", tc.stack.buf, tc.expected)
			}
		})
	}
}
