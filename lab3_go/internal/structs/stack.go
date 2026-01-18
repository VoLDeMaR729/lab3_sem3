package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

// Структура стека
type Stack struct {
	Data []string `json:"data"` //храним данные
}

// Конструктор
func NewStack() *Stack { return &Stack{Data: make([]string, 0)} }

// Добавить наверх
func (s *Stack) Push(val string) {
	s.Data = append(s.Data, val) //в конец слайса
}

// Снять сверху
func (s *Stack) Pop() string {
	if len(s.Data) == 0 { return "" } //пусто
	lastIdx := len(s.Data) - 1 //индекс последнего
	val := s.Data[lastIdx] //берем значение
	s.Data = s.Data[:lastIdx] //уменьшаем слайс
	return val
}

// Печать
func (s *Stack) Print() {
	fmt.Print("Stack (Top->Bottom): ")
	for i := len(s.Data)-1; i >= 0; i-- { //идем с конца
		fmt.Printf("[%s] ", s.Data[i])
	}
	fmt.Println()
}

// Сохранение
func (s *Stack) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil { return err }
	defer file.Close()
	return json.NewEncoder(file).Encode(s)
}

// Загрузка
func (s *Stack) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil { return err }
	defer file.Close()
	return json.NewDecoder(file).Decode(s)
}
