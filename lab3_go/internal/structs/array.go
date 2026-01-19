package structs

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Massive struct {
	Data     []string
	Size     int
	Capacity int
}

func NewMassive(initialCapacity int) *Massive {
	if initialCapacity <= 0 {
		initialCapacity = 1
	}
	return &Massive{
		Data:     make([]string, initialCapacity),
		Size:     0,
		Capacity: initialCapacity,
	}
}

func (m *Massive) expand() {
	newCapacity := m.Capacity * 2
	newData := make([]string, newCapacity)
	for i := 0; i < m.Size; i++ {
		newData[i] = m.Data[i]
	}
	m.Data = newData
	m.Capacity = newCapacity
}

func (m *Massive) PushBack(val string) {
	if m.Size == m.Capacity {
		m.expand()
	}
	m.Data[m.Size] = val
	m.Size++
}

func (m *Massive) Insert(index int, val string) error {
	if index < 0 || index > m.Size {
		return errors.New("bounds error")
	}
	if m.Size == m.Capacity {
		m.expand()
	}
	for i := m.Size; i > index; i-- {
		m.Data[i] = m.Data[i-1]
	}
	m.Data[index] = val
	m.Size++
	return nil
}

func (m *Massive) Remove(index int) error {
	if index < 0 || index >= m.Size {
		return errors.New("bounds error")
	}
	for i := index; i < m.Size-1; i++ {
		m.Data[i] = m.Data[i+1]
	}
	m.Data[m.Size-1] = ""
	m.Size--
	return nil
}

func (m *Massive) Get(index int) (string, error) {
	if index < 0 || index >= m.Size {
		return "", errors.New("bounds error")
	}
	return m.Data[index], nil
}

func (m *Massive) Set(index int, val string) error {
	if index < 0 || index >= m.Size {
		return errors.New("bounds error")
	}
	m.Data[index] = val
	return nil
}

func (m *Massive) Print() {
	fmt.Print("Arr: ")
	for i := 0; i < m.Size; i++ {
		fmt.Print(m.Data[i], " ")
	}
	fmt.Println()
}

func (m *Massive) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(m)
}

func (m *Massive) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(m)
}

func (m *Massive) SaveBinary(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	return encoder.Encode(m)
}

func (m *Massive) LoadBinary(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	return decoder.Decode(m)
}
