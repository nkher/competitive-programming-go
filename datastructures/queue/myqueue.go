package queue

type MyQueue interface {
	Insert(val interface{})
	Remove() interface{}
	Peek() interface{}
	Length() int
	Empty() bool
}