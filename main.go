package main

import (
	"fmt"
)

type Element struct {
	value int
	next  *Element
}

func main() {
	element4 := Element{
		4,
		nil,
	}

	element3 := Element{
		3,
		&element4,
	}

	element2 := Element{
		2,
		&element3,
	}

	element1 := Element{
		1,
		&element2,
	}
	print(&element1)
	add(&element1, 5)
	add(&element1, 6)
	print(&element1)
	delete(&element1)
	print(&element1)
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

// Заменяет значение элемента под номером num на val
func set(root *Element, num int, val int) {

}

// Заменяет значение элемента под номером number на val
func reverse(root *Element) {

}
