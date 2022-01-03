package main

import (
	"fmt"
	"container/list"
)

type Stack struct {
	v *list.List
}
func NewStack() *Stack {
	return &Stack{ list.New() } 
}
func (s *Stack) Push(val interface{}){
	s.v.PushBack(val)
}
func (s *Stack) Pp() interface{} {
	back := s.v.Back() //Quee와 다른 부분, Back부터 remove
	if back != nil {
		return s.v.Remove(back)
	
	return nil 
}
func main(){   
	stack := NewStck()
	fr i:=1; i<5; i++{
		stack.Push(i)
	}
	val := stack.Pop()
	for val != nil {
		mt.Printf("%v -> ", val)  // 4->3->2->1
		val = stack.Pop()

	}
}