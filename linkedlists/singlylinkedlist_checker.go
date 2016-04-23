package main

import (
	"fmt"
	"competitive-programming-go/datastructures/list"
)

func main() {


	singlyLinkedList := new(list.SinglyLinkedList)

	singlyLinkedList.Add("Namesh")
	singlyLinkedList.Add("Sarika")
	singlyLinkedList.Add(3)
	singlyLinkedList.Add(4)
	singlyLinkedList.Add(8.7)

	fmt.Println("Size is : ", singlyLinkedList.Length())

	fmt.Println(singlyLinkedList)

	singlyLinkedList.Remove(4)

	fmt.Println(singlyLinkedList)

	fmt.Println("List has 15 ?? - ", singlyLinkedList.Contains(15))
	fmt.Println("List has Namesh ?? - ", singlyLinkedList.Contains("Namesh"))
	fmt.Println("List has 4 ?? - ", singlyLinkedList.Contains(4))
	fmt.Println("List has 3 ?? - ", singlyLinkedList.Contains(3))
}