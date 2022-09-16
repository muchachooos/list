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
	/*print(&element0)
	println("---------")
	delete(&element0)
	print(&element0)
	println("---------")
	add(&element0, 33)
	print(&element0)
	println("---------")
	set(&element0, 1, 22)
	print(&element0)
	println("---------")*/
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
func delete(root *Element) {
	if root == nil {
		fmt.Println("Bl!")
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
		fmt.Println("AAA!")
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
ln = nk   list: 1nil     ck:234     nk:nil
nk = l    list: 1nil     ck:234     nk:1nil
l = ck    list: 234      ck:234     nk:1nil
2 итерация
ck = ln   list: 234      ck:34      nk:1nil
ln = nk   list: 21nil    ck:34      nk:1nil
nk = l    list: 21nil    ck:34      nk:21nil
l = ck    list: 34       ck:34      nk:21nil
3 итерация
ck = ln   list: 34       ck:4       nk:21nil
ln = nk   list: 321nil   ck:4       nk:21nil
nk = l    list: 321nil   ck:4       nk:321nil
l = ck    list: 4        ck:4       nk:321nil
4 итерация
ck = ln   list: 4        ck:nil     nk:321nil
ln = nk   list: 4321nil  ck:nil     nk:321nil
nk = l    list: 4321nil  ck:nil     nk:4321nil
l = ck    list: nil      ck:nil     nk:4321nil
*/
