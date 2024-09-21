package main

import "fmt"

type Stack struct {
	stack []int
	size  int
}

func (s *Stack) init(length int) {
	s.stack = make([]int, 0)
	s.size = length
}

func (s *Stack) pop() (int, error) {

	if len(s.stack) > 0 {
		var value = s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return value, nil
	}
	return 0, fmt.Errorf("There are no elements in the stack to pop.")
}

func (s *Stack) push(i int) error {
	if len(s.stack) < s.size {
		s.stack = append(s.stack, i)
		return nil
	}
	return fmt.Errorf("Stack is filled.")
}

func StackDriver() {
	var s Stack
	s.init(10)
	for i := 1; i <= 11; i++ {
		err := s.push(i * i)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
	fmt.Println(s.stack)
	for i := 0; i < s.size+1; i++ {
		value, err := s.pop()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(value)
	}
	fmt.Println(s.stack)
}
