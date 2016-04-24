package stack

type MyStack interface {
	Push(val interface{})
	Pop(val interface{}) interface{}
	Peek() interface{}
	Length() int
	Empty() bool
}