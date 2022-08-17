package main

import "fmt"

type Element struct {
	value       int
	nextElement *Element
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
}

func print(root *Element) {
	fmt.Println(root.value)
	if root.nextElement != nil {
		print(root.nextElement)
	}
}

// Удаляет последний элемент в списке
func delete(root *Element, value int) {
}

// Добавляет элемент в конец списка со знаяением value
func add(root *Element, value int) {
}

// Заменяет значение элемента под номером number на value
func set(root *Element, number int, value int) {
}

// Заменяет значение элемента под номером number на value
func reverse(root *Element) {
}
