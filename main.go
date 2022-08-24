package main

import (
	"fmt"
)

var i int

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
		&element4,
	}

	element1 := Element{
		1,
		&element3,
	}

	element0 := Element{
		0,
		&element2,
	}
	print(&element0)
	set(&element0, 3, 5)
}

// Заменяет value элемента под номером num на val
func set(root *Element, num int, val int) {
	if root == nil {
		fmt.Println("")
		return
	}

	if num == i {
		root.value = val
	} else {
		i++
		set(root.next, num, val)
	}
	i = 0
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

// Заменяет значение элемента под номером number на val
func reverse(root *Element) {

}
