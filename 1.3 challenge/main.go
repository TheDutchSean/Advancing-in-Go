// Your challenge is to implement the add function, which takes in an index and value to add to the linked list. Right now, this

package main

import (
	"errors"
	"fmt"
)

func main(){

	// pointer()
	// arrays()
	// slices()
	// maps()
	// structs()
	testCases := []struct {
		index int
		value string
	}{
	   {index: 0, value: "C"},
	   {index: 1, value: "A"},
	   {index: 2, value: "B"},
	   {index: 3, value: "D"},
	}
	dl := &DoublyLinkedList[string]{}
	dl.AddElements(testCases)
	dl.PrintForward()
	dl.PrintReverse()
	fmt.Println(dl)
}

type Node[T any] struct {
	value 	T
	next, prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}


// Add appends the given value at the given index.
// Returns an error in the case that the index exceeds the list size
func (l * DoublyLinkedList[T]) Add(index int, value T) error {

	currentSize := l.size

	if index > l.size {
		return errors.New("index exceeds list size")
	}

	l.size = currentSize + 1

	newElement := &Node[T] {
		value: value,
	}

	// list is empty
	if l.head == nil {
		l.head, l.tail = newElement, newElement
	}

	//  change head
	if index == 0 {
		newElement.next = l.head
		l.head.prev, l.head = newElement, newElement
		return nil
	}

	//  change tail
	if index == currentSize {
		newElement.prev = l.tail
		l.tail.next, l.tail = newElement, newElement
		return nil
	}

	// change index
	current := l.head
	for i := 1; i < index; i++{
		current = current.next;
	} 

	newElement.prev = current
	newElement.next = current.next
	current.next.prev, current.next = newElement, newElement

	return nil

}


func (l * DoublyLinkedList[T]) AddElements(elements []struct{
	index int
	value T
}) error {

	for _, e := range elements {
		if err := l.Add(e.index, e.value); err != nil {
			return err
		}
	}
	return nil
}


func (l * DoublyLinkedList[T]) PrintForward() string {
	if l.size == 0{
		return ""
	}
	current := l.head
	output := "HEAD"
	for current != nil {
		output = fmt.Sprintf("%s -> %v", output, current.value)
		current = current.next
	}

	return fmt.Sprintf("%s -> NULL", output)
}

func (l * DoublyLinkedList[T]) PrintReverse() string {
	if l.size == 0{
		return ""
	}
	current := l.tail
	output := "NULL"
	for current != nil {
		output = fmt.Sprintf("%s <- %v", output, current.value)
		current = current.prev
	}

	return fmt.Sprintf("%s <- HEAD", output)
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

