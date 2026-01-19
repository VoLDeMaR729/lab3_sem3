package structs

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type DNode struct {
	Val  string
	Next *DNode
	Prev *DNode
}

type DoublyLinkedList struct {
	Head *DNode
	Tail *DNode
}

func NewDList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (l *DoublyLinkedList) AddHead(val string) {
	newNode := &DNode{Val: val}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head.Prev = newNode
		l.Head = newNode
	}
}

func (l *DoublyLinkedList) PushHead(val string) {
	l.AddHead(val)
}

func (l *DoublyLinkedList) AddTail(val string) {
	newNode := &DNode{Val: val}
	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		newNode.Prev = l.Tail
		l.Tail = newNode
	}
}

func (l *DoublyLinkedList) PushTail(val string) {
	l.AddTail(val)
}

func (l *DoublyLinkedList) DeleteHead() {
	if l.Head == nil {
		return
	}
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		return
	}
	l.Head = l.Head.Next
	l.Head.Prev = nil
}

func (l *DoublyLinkedList) DeleteTail() {
	if l.Tail == nil {
		return
	}
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		return
	}
	l.Tail = l.Tail.Prev
	l.Tail.Next = nil
}

func (l *DoublyLinkedList) DelVal(val string) {
	curr := l.Head
	for curr != nil {
		if curr.Val == val {
			if curr == l.Head {
				l.DeleteHead()
				return
			}
			if curr == l.Tail {
				l.DeleteTail()
				return
			}
			curr.Prev.Next = curr.Next
			curr.Next.Prev = curr.Prev
			return
		}
		curr = curr.Next
	}
}

func (l *DoublyLinkedList) Find(val string) bool {
	curr := l.Head
	for curr != nil {
		if curr.Val == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (l *DoublyLinkedList) Print() {
	fmt.Print("DList: ")
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.Val, " <-> ")
		curr = curr.Next
	}
	fmt.Println("NULL")
}

func (l *DoublyLinkedList) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []string
	curr := l.Head
	for curr != nil {
		data = append(data, curr.Val)
		curr = curr.Next
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (l *DoublyLinkedList) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}
	l.Head = nil
	l.Tail = nil
	for _, v := range data {
		l.AddTail(v)
	}
	return nil
}

func (l *DoublyLinkedList) SaveBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []string
	curr := l.Head
	for curr != nil {
		data = append(data, curr.Val)
		curr = curr.Next
	}
	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

func (l *DoublyLinkedList) LoadBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []string
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}
	l.Head = nil
	l.Tail = nil
	for _, v := range data {
		l.AddTail(v)
	}
	return nil
}
