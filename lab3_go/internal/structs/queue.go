package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

// Структура очереди
type Queue struct {
	Data []string `json:"data"` //храним как массив для простоты
}

// Конструктор
func NewQueue() *Queue { return &Queue{Data: make([]string, 0)} }

// Добавить элемент
func (q *Queue) Push(val string) {
	q.Data = append(q.Data, val) //в конец слайса
}

// Извлечь элемент
func (q *Queue) Pop() string {
	if len(q.Data) == 0 { return "" } //пусто
	val := q.Data[0] //берем первый
	q.Data = q.Data[1:] //отрезаем голову слайса
	return val
}

// Печать
func (q *Queue) Print() {
	fmt.Println("Queue:", q.Data)
}

// Сохранение
func (q *Queue) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil { return err }
	defer file.Close()
	return json.NewEncoder(file).Encode(q)
}

// Загрузка
func (q *Queue) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil { return err }
	defer file.Close()
	return json.NewDecoder(file).Decode(q)
}
