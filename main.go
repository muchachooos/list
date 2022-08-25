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
	set(&element0, 1, 11)
	println("=========")
	print(&element0)
	set(&element0, 2, 22)
	println("=========")
	print(&element0)
}
func print(root *Element) {
	fmt.Println(root.value)
	if root.next != nil {
		print(root.next)
	}
}

// Удаляет последний элемент в списке
func delete(root *Element) {
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
}
