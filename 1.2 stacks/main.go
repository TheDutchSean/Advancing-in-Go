// A stack is a data structure that's related to the linked list. It has a last in, first out element access. 
// New elements are added to the top of the stack and the latest element is read first. 

// 

package main

import (
	"errors"
)

func main(){

	// pointer()
	// arrays()
	// slices()
	// maps()
	// structs()

}


// Generic stack implementation

type Stack[T any] struct {
	elements []T
}

// Push adds a new element to the top of the stack
func (s *Stack[T]) Push(element T) {
	// implement code
	s.elements = append(s.elements, element)
}

// IsEmpty returns a bool to indicate if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	// implement code
	return len(s.elements) == 0
}

//  Pop reads amd removes an element from the Stack or returns an error if the stack is empty
func (s *Stack[T]) Pop () (*T, error){
	// implement code
	if s.IsEmpty(){
		return nil, errors.New("cannot pop from empty stack")
	}
	
	// the last element is the one that gets read
	top := s.elements[len(s.elements) - 1]
	//  drop the read element
	s.elements = s.elements[:len(s.elements) - 1]
	// the function returns a pointer to the top element and a nil error
	return &top, nil
}

