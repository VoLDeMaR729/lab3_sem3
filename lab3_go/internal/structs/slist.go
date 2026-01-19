package structs

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type SNode struct {
	Val  string
	Next *SNode
}

type SinglyLinkedList struct {
	Head *SNode
}

func (l *SinglyLinkedList) PushHead(val string) {
	newNode := &SNode{Val: val, Next: l.Head}
	l.Head = newNode
}

func (l *SinglyLinkedList) PushTail(val string) {
	newNode := &SNode{Val: val, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (l *SinglyLinkedList) InsertAfter(target string, val string) {
	curr := l.Head
	for curr != nil && curr.Val != target {
		curr = curr.Next
	}
	if curr != nil {
		newNode := &SNode{Val: val, Next: curr.Next}
		curr.Next = newNode
	}
}

func (l *SinglyLinkedList) InsertBefore(target string, val string) {
	if l.Head == nil {
		return
	}
	if l.Head.Val == target {
		l.PushHead(val)
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Val != target {
		curr = curr.Next
	}
	if curr.Next != nil {
		newNode := &SNode{Val: val, Next: curr.Next}
		curr.Next = newNode
	}
}

func (l *SinglyLinkedList) DeleteHead() {
	if l.Head != nil {
		l.Head = l.Head.Next
	}
}

func (l *SinglyLinkedList) DeleteTail() {
	if l.Head == nil {
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		return
	}
	curr := l.Head
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
}

func (l *SinglyLinkedList) DeleteByValue(val string) {
	if l.Head == nil {
		return
	}
	if l.Head.Val == val {
		l.DeleteHead()
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Val != val {
		curr = curr.Next
	}
	if curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

func (l *SinglyLinkedList) DeleteAfter(target string) {
	curr := l.Head
	for curr != nil && curr.Val != target {
		curr = curr.Next
	}
	if curr != nil && curr.Next != nil {
		curr.Next = curr.Next.Next
	}
}

func (l *SinglyLinkedList) DeleteBefore(target string) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}
	if l.Head.Next.Val == target {
		l.DeleteHead()
		return
	}
	curr := l.Head
	for curr.Next.Next != nil && curr.Next.Next.Val != target {
		curr = curr.Next
	}
	if curr.Next.Next != nil {
		curr.Next = curr.Next.Next
	}
}

func (l *SinglyLinkedList) Find(val string) bool {
	curr := l.Head
	for curr != nil {
		if curr.Val == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (l *SinglyLinkedList) Print() {
	fmt.Print("SList: ")
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Println("NULL")
}

func (l *SinglyLinkedList) Save(filename string) error {
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

func (l *SinglyLinkedList) Load(filename string) error {
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
	for _, v := range data {
		l.PushTail(v)
	}
	return nil
}

func (l *SinglyLinkedList) SaveBinary(filename string) error {
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

func (l *SinglyLinkedList) LoadBinary(filename string) error {
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
	for _, v := range data {
		l.PushTail(v)
	}
	return nil
}
