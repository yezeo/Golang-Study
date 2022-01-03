package main
import (
	"fmt"
	"container/list"
)
type Queue struct {  // queue 구조체 정의
	v *list.List
}
func (q *Queue) Push(val interface{}){
	q.v.PushBack(val)
}
func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil //요소 없을 경우
}
func NewQueue() *Queue {
	return &Queue{ list.New() }
}
func main(){
	queue := NewQueue()

	for i:=1; i<5; i++ {
		queue.Push(i)  // {1,2,3,4,5}
	}
	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v) // 1->2->3->4->5
		v = queue.Pop()
	}
}