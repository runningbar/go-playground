// 索引数组实现双向链表
package main 

import (
	"fmt"
)

func main() {
	l := links{
		head: -1,
		tail: -1,
	}
	fmt.Println("initial:", l)
	l.append("a")
	l.append("b")
	l.append("c")
	fmt.Println("after add:", l)
	l.delete(0)
	fmt.Println("after delete:", l)
}

type item struct {
	pre int
	bhd int
	value string
}

type links struct {
	list []item
	head int //指向链表第一个元素
	tail int //指向链表最后一个元素
}

func (l links) String() string{
	out := fmt.Sprintf("head = %v, tail = %v, list = %v", l.head, l.tail, l.list)
	return out
}

func (l *links) append(v string) {
	for index := range l.list {
		if l.list[index].pre == -2 && l.list[index].bhd == -2 {
			l.list[index].value = v
			l.list[l.tail].bhd = index
			l.list[index].pre = l.tail
			l.tail = index
			return
		}
	}
	itm := item{
		pre: -1,
		bhd: -1,
		value: v,
	}
	l.list = append(l.list, itm)
	index := len(l.list) - 1
	if l.tail != -1 {
		l.list[l.tail].bhd = index
	}
	if l.head == -1 {
		l.head = index
	}
	l.list[index].pre = l.tail
	l.tail = index
}

func (l *links) delete(i int) {
	if l.list[i].pre != -1 {
		l.list[l.list[i].pre].bhd = l.list[i].bhd
	} else {
		l.head = l.list[i].bhd
	}

	if l.list[i].bhd != -1 {
		l.list[l.list[i].bhd].pre = l.list[i].pre
	} else {
		l.tail = l.list[i].pre
	}

	l.list[i].pre = -2
	l.list[i].bhd = -2
}

//output
/* 
initial: head = -1, tail = -1, list = []
after add: head = 0, tail = 2, list = [{-1 1 a} {0 2 b} {1 -1 c}]
after delete: head = 1, tail = 2, list = [{-2 -2 a} {-1 2 b} {1 -1 c}]
 */
