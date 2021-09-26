package util

import (
	"container/list"
	"log"
)

type Stack struct {
	data *list.List
	size int
}

func NewSatck() *Stack {
	q := new(Stack)
	q.init()
	return q
}

func (q *Stack) init() {
	q.data = list.New()
}

func (q *Stack) Size() int {
	return q.size
}

func (q *Stack) Empty() bool {
	return q.size == 0
}

func (q *Stack) Top() string {
	return q.data.Back().Value.(string)
}

func (q *Stack) PData() {
	s := ""
	st := q.data.Front()
	for {
		if st==nil{
			break
		}
		s += " "+st.Value.(string)
		st = st.Next()
	}
	log.Println(s)
}

func (q *Stack) Push(value string) {
	q.data.PushBack(value)
	q.size++
}

func (q *Stack) Pop() string{
	if q.size > 0 {
		tmp := q.data.Back()
		q.data.Remove(tmp)
		q.size--
		return tmp.Value.(string)
	}
	return ""
}
