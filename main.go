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
	deleteLastElement(&element0)
	print(&element0)
	println("---------")
	addLastElement(&element0, 33)
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
		fmt.Println("Ck!")
		return
	}

	fmt.Println(root.value)
	if root.next != nil {
		print(root.next)
	}
}

// Удаляет последний элемент в списке
func deleteLastElement(root *Element) {
	if root == nil {
		fmt.Println("Bl!")
		return
	}

	if root.next != nil {
		if root.next.next == nil {
			root.next = nil
		} else {
			deleteLastElement(root.next)
		}
	}
}

// Добавляет элемент в конец списка со значением val
func addLastElement(root *Element, val int) {
	if root == nil {
		fmt.Println("AAA!")
		return
	}
	if root.next == nil {
		root.next = &Element{val, nil}
		return
	}
	addLastElement(root.next, val)
}

// Заменяет value элемента под номером num на val
func set(root *Element, num, val int) {
	setValue(root.next, 0, num, val)
}
func setValue(root *Element, i, num, val int) {
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

func reverse(list *Element) {
	if list == nil {
		fmt.Println("Error!")
		return
	}

	var currentKeeper *Element
	var nextKeeper *Element

	for list != nil {
		currentKeeper = list.next
		list.next = nextKeeper
		nextKeeper = list
		list = currentKeeper
	}
	return
}

/*
reverse
__________list: 1234     ck:nil     nk:nil
1 итерация
ck = ln   list: 1234     ck:234     nk:nil
ln = nk   list: 1        ck:234     nk:nil
nk = l    list: 1        ck:234     nk:1
l = ck    list: 234      ck:234     nk:1
2 итерация
ck = ln   list: 234      ck:34      nk:1
ln = nk   list: 21       ck:34      nk:1
nk = l    list: 21       ck:34      nk:21
l = ck    list: 34       ck:34      nk:21
3 итерация
ck = ln   list: 34       ck:4       nk:21
ln = nk   list: 321      ck:4       nk:21
nk = l    list: 321      ck:4       nk:321
l = ck    list: 4        ck:4       nk:321
4 итерация
ck = ln   list: 4        ck:nil     nk:321
ln = nk   list: 4321     ck:nil     nk:321
nk = l    list: 4321     ck:nil     nk:4321
l = ck    list: nil      ck:nil     nk:4321
*/
