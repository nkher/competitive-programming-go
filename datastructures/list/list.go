package list

type List interface {
	Add(val interface{})
	Remove(val interface{}) bool
	Contains(val interface{}) bool
	IndexOf(val interface{}) int
	Length() int
	Empty() bool
}





