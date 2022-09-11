package main

import (
	"fmt"
)

type Element struct {
	value int
	next  *Element
}

func main() {
	element3 := Element{
		3,
		nil,
	}
	element2 := Element{
		2,
		&element3,
	}
	element1 := Element{
		1,
		&element2,
	}
	element0 := Element{
		0,
		&element1,
	}
	print(&element0)
	println("---------")
	delete(&element0)
	print(&element0)
	println("---------")
	add(&element0, 33)
	print(&element0)
	println("---------")
	set(&element0, 1, 22)
	print(&element0)
	println("---------")
	reverse(&element0)
	print(&element3)
	println("---------")
}

func print(root *Element) {
	if root == nil {
		fmt.Println("Stop!")
		return
	}

	fmt.Println(root.value)
	if root.next != nil {
		print(root.next)
	}
}

// Удаляет последний элемент в списке
func delete(root *Element) {
	if root == nil {
		fmt.Println("Stop!")
		return
	}

	if root.next != nil {
		if root.next.next == nil {
			root.next = nil
		} else {
			delete(root.next)
		}
	}
}

// Добавляет элемент в конец списка со значением val
func add(root *Element, val int) {
	if root == nil {
		fmt.Println("Stop!")
		return
	}
	if root.next == nil {
		root.next = &Element{val, nil}
		return
	}
	add(root.next, val)
}

// Заменяет value элемента под номером num на val
func set(root *Element, num, val int) {
	setValue(root.next, 0, num, val)
}
func setValue(root *Element, i, num, val int) {
	s := func() {
		if root == nil {
			fmt.Println("Stop!")
			return
		}
		if num == i {
			root.value = val
		} else {
			i++
			setValue(root.next, i, num, val)
		}
	}
	s()
}

func reverse(root *Element) {
	if root == nil {
		fmt.Println("Error!")
		return
	}
	//o := root

	var j *Element = nil
	var i *Element = nil

	for root != nil {
		j = root.next
		root.next = i
		i = root
		root = j
	}
	return
}
