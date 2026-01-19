package structs

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type Stack struct {
	Head *SNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val string) {
	newNode := &SNode{Val: val, Next: s.Head}
	s.Head = newNode
}

func (s *Stack) Pop() string {
	if s.Head == nil {
		return ""
	}
	val := s.Head.Val
	s.Head = s.Head.Next
	return val
}

func (s *Stack) Print() {
	fmt.Print("Stack: ")
	curr := s.Head
	for curr != nil {
		fmt.Print(curr.Val, " -> ")
		curr = curr.Next
	}
	fmt.Println("NULL")
}

func (s *Stack) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []string
	curr := s.Head
	for curr != nil {
		data = append(data, curr.Val)
		curr = curr.Next
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (s *Stack) Load(filename string) error {
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
	s.Head = nil
	for i := len(data) - 1; i >= 0; i-- {
		s.Push(data[i])
	}
	return nil
}

func (s *Stack) SaveBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	var data []string
	curr := s.Head
	for curr != nil {
		data = append(data, curr.Val)
		curr = curr.Next
	}
	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

func (s *Stack) LoadBinary(filename string) error {
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
	s.Head = nil
	for i := len(data) - 1; i >= 0; i-- {
		s.Push(data[i])
	}
	return nil
}
