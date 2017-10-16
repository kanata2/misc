package main

import (
	"errors"
	"fmt"
	"sync"
)

type Stack struct {
	mu    *sync.Mutex
	buf   []string
	limit int
}

func main() {
	s, _ := NewStack(10)
	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push("dataD")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func NewStack(limit int) (*Stack, error) {
	if limit < 0 {
		return nil, errors.New("limit is a positive number")
	}
	return &Stack{
		mu:    &sync.Mutex{},
		buf:   []string{},
		limit: limit,
	}, nil
}

func (s *Stack) Pop() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.buf) == 0 {
		return ""
	}
	l := s.buf[len(s.buf)-1]
	s.buf = s.buf[:len(s.buf)-1]
	return l
}

func (s *Stack) Push(ss string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.buf) == s.limit {
		s.buf = s.buf[1:]
	}
	s.buf = append(s.buf, ss)
}
