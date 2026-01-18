package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

// Структура узла
type SNode struct {
	Val  string `json:"val"`//значение
	Next *SNode `json:"next"` //указатель на следующий
}

// Структура списка
type SList struct {
	Head *SNode `json:"head"` //голова списка
}

// Конструктор
func NewSList() *SList { return &SList{} } //возвращаем пустой список

// Добавить в начало
func (l *SList) AddHead(val string) {
	l.Head = &SNode{Val: val, Next: l.Head} //новый узел указывает на старую голову
}

// Добавить в конец
func (l *SList) AddTail(val string) {
	if l.Head == nil { //если список пуст
		l.Head = &SNode{Val: val} //новый узел - голова
		return
	}
	curr := l.Head //итератор
	for curr.Next != nil { curr = curr.Next } //ищем последний
	curr.Next = &SNode{Val: val} //цепляем новый
}

// Удалить голову
func (l *SList) DelHead() {
	if l.Head != nil { l.Head = l.Head.Next } //сдвигаем голову вперед
}

// Удалить хвост
func (l *SList) DelTail() {
	if l.Head == nil { return } //пусто
	if l.Head.Next == nil { l.Head = nil; return } //один элемент
	curr := l.Head //итератор
	for curr.Next.Next != nil { curr = curr.Next } //ищем предпоследний
	curr.Next = nil //обнуляем связь с последним
}

// Удалить по значению
func (l *SList) DelVal(val string) {
	if l.Head == nil { return } //пусто
	if l.Head.Val == val { l.Head = l.Head.Next; return } //если в голове
	curr := l.Head //итератор
	for curr.Next != nil && curr.Next.Val != val { curr = curr.Next } //ищем
	if curr.Next != nil { curr.Next = curr.Next.Next } //перекидываем связь
}

// Поиск элемента
func (l *SList) Find(val string) bool {
	curr := l.Head //итератор
	for curr != nil { //пока есть элементы
		if curr.Val == val { return true } //нашли
		curr = curr.Next //дальше
	}
	return false //не нашли
}

// Печать
func (l *SList) Print() {
	fmt.Print("SList: ")
	curr := l.Head
	for curr != nil {
		fmt.Printf("%s -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

// Сохранение
func (l *SList) Save(filename string) error {
	file, err := os.Create(filename) //файл
	if err != nil { return err }
	defer file.Close()
	return json.NewEncoder(file).Encode(l) //пишем
}

// Загрузка
func (l *SList) Load(filename string) error {
	file, err := os.Open(filename) //открываем
	if err != nil { return err }
	defer file.Close()
	return json.NewDecoder(file).Decode(l) //читаем
}
