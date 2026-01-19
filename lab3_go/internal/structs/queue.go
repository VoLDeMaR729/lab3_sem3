package structs

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type Queue struct {
	Head *SNode
	Tail *SNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(val string) {
	newNode := &SNode{Val: val}
	if q.Tail == nil {
		q.Head = newNode
		q.Tail = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = newNode
	}
}

func (q *Queue) Pop() string {
	if q.Head == nil {
		return ""
	}
	val := q.Head.Val
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	return val
}

func (q *Queue) Print() {
	fmt.Print("Queue: ")
	curr := q.Head
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Println("NULL")
}

func (q *Queue) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []string
	curr := q.Head
	for curr != nil {
		data = append(data, curr.Val)
		curr = curr.Next
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (q *Queue) Load(filename string) error {
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
	q.Head = nil
	q.Tail = nil
	for _, v := range data {
		q.Push(v)
	}
	return nil
}

func (q *Queue) SaveBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []string
	curr := q.Head
	for curr != nil {
		data = append(data, curr.Val)
		curr = curr.Next
	}
	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

func (q *Queue) LoadBinary(filename string) error {
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
	q.Head = nil
	q.Tail = nil
	for _, v := range data {
		q.Push(v)
	}
	return nil
}
