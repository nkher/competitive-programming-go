package stack

type MyStack interface {
	Push(val interface{})
	Pop() interface{}
	Peek() interface{}
	Length() int
	Empty() bool
}