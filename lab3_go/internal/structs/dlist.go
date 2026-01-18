package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

type DList struct {
	Items []string `json:"items"` //храним как слайс для удобной JSON-сериализации
//Поля head и tail не экспортируем в JSON напрямую, восстановим их при Load
	head *DNode
	tail *DNode
}

// Узел двусвязного
type DNode struct {
	Val  string //значение
	Next *DNode //след
	Prev *DNode //пред
}

// Конструктор
func NewDList() *DList {
	return &DList{Items: make([]string, 0)} //инициализация
}

// Добавить в голову
func (l *DList) AddHead(val string) {
	// Обновляем слайс для JSON
	l.Items = append([]string{val}, l.Items...)

	n := &DNode{Val: val} //новый узел
	if l.head == nil { //если пуст
		l.head = n
		l.tail = n
	} else {
		n.Next = l.head //связь вперед
		l.head.Prev = n //связь назад
		l.head = n //новая голова
	}
}

// Добавить в хвост
func (l *DList) AddTail(val string) {
	// Обновляем слайс для JSON
	l.Items = append(l.Items, val)

	n := &DNode{Val: val} //новый узел
	if l.tail == nil { //если пуст
		l.head = n
		l.tail = n
	} else {
		l.tail.Next = n //связь вперед
		n.Prev = l.tail //связь назад
		l.tail = n //новый хвост
	}
}

// Удалить по значению
func (l *DList) DelVal(val string) {
	// 1. Удаляем из связного списка
	curr := l.head //итератор
	for curr != nil {
		if curr.Val == val { //нашли
			if curr.Prev != nil {
				curr.Prev.Next = curr.Next //обход слева
			} else {
				l.head = curr.Next //новая голова
			}

			if curr.Next != nil {
				curr.Next.Prev = curr.Prev //обход справа
			} else {
				l.tail = curr.Prev //новый хвост
			}
			break // Удаляем первое вхождение
		}
		curr = curr.Next //дальше
	}

	// 2. Удаляем из слайса (для JSON)
	for i, v := range l.Items {
		if v == val {
			l.Items = append(l.Items[:i], l.Items[i+1:]...) //вырезаем
			break
		}
	}
}

// Поиск значения
func (l *DList) Find(val string) bool {
	curr := l.head //итератор
	for curr != nil {
		if curr.Val == val {
			return true //нашли
		}
		curr = curr.Next
	}
	return false //не нашли
}

// Печать
func (l *DList) Print() {
	fmt.Print("DList: ")
	curr := l.head
	for curr != nil {
		fmt.Printf("%s <-> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

// Load загружает из JSON и восстанавливает связи
func (l *DList) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(l); err != nil { //читаем слайс
		return err
	}

	// Пересобираем связный список из загруженного слайса Items
	l.head = nil
	l.tail = nil
	tempItems := l.Items
	l.Items = make([]string, 0) // сброс, AddTail заполнит заново и корректно
	for _, v := range tempItems {
		l.AddTail(v) //восстанавливаем структуру
	}
	return nil
}

// Сохранение
func (l *DList) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(l)
}
