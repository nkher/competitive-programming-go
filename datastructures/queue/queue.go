package queue 

import (
	"fmt"
)

type QueueNode struct {
	val interface{}
	above *QueueNode
	below *QueueNode
}

type Queue struct {
	size int
	Head *QueueNode
	Tail *QueueNode
}

func (queue *Queue) Insert(val interface{}) {
	newNode := &QueueNode{val, nil, nil}
	if queue.size == 0 {
		queue.Head, queue.Tail = newNode, newNode
	} else {
		queue.Tail.above = newNode
		newNode.below = queue.Tail
		queue.Tail = newNode
	}	
	queue.size++
}

func (queue *Queue) Remove() interface{} {
	if queue.Head != nil && queue.size > 0 {
		val := queue.Head.val
		queue.Head = queue.Head.below
		queue.size--
		if queue.size == 0 {
			queue.Tail = nil
		}
		return val
	}
	return nil
}

func (queue *Queue) Peek() interface{} {
	if queue.Head != nil && queue.size > 0 {
		return queue.Head.val
	}
	return nil
}

func (queue *Queue) Length() int {
	return queue.size;
}

func (queue *Queue) Empty() bool {
	return (queue.Head == nil && queue.Tail == nil && queue.size == 0)
}

func (queue *Queue) String() string {
	temp := queue.Head
	str := "[ "
	for temp != nil {
		str += fmt.Sprintf("%v", temp.val) + " "
		temp = temp.above
	}
	str += "]"
	return str
}