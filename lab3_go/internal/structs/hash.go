package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

// Узел хеш-таблицы (для цепочек)
type HashNode struct {
	Key   string    `json:"key"`//ключ
	Value string    `json:"value"`//значение
	Next  *HashNode `json:"next"`//след в цепочке
}

// Хеш-таблица
type HashTable struct {
	Buckets  []*HashNode `json:"buckets"`//массив списков
	Capacity int         `json:"capacity"`//емкость
}

// Конструктор
func NewHashTable(cap int) *HashTable {
	return &HashTable{
		Buckets:  make([]*HashNode, cap), //выделяем память
		Capacity: cap,
	}
}

// Хеш-функция
func (h *HashTable) hash(key string) int {
	sum := 0
	for _, c := range key {
		sum += int(c) //суммируем коды символов
	}
	return sum % h.Capacity //берем остаток от деления
}

// Вставка
func (h *HashTable) Put(key, val string) {
	idx := h.hash(key) //индекс корзины
	curr := h.Buckets[idx] //голова списка
	for curr != nil {
		if curr.Key == key { //если ключ уже есть
			curr.Value = val //обновляем
			return
		}
		curr = curr.Next //дальше
	}
	//если нет, добавляем в начало цепочки
	h.Buckets[idx] = &HashNode{Key: key, Value: val, Next: h.Buckets[idx]}
}

// Получение значения
func (h *HashTable) Get(key string) string {
	idx := h.hash(key) //индекс
	curr := h.Buckets[idx] //голова
	for curr != nil {
		if curr.Key == key {
			return curr.Value //нашли
		}
		curr = curr.Next //дальше
	}
	return "" //не нашли
}

// Удаление
func (h *HashTable) Remove(key string) {
	idx := h.hash(key) //индекс
	curr := h.Buckets[idx] //голова
	var prev *HashNode //предыдущий

	for curr != nil {
		if curr.Key == key { //нашли
			if prev == nil {
				h.Buckets[idx] = curr.Next //удаляем голову цепочки
			} else {
				prev.Next = curr.Next //связываем пред и след
			}
			return
		}
		prev = curr //шаг
		curr = curr.Next //шаг
	}
}

// Печать
func (h *HashTable) Print() {
	for i, b := range h.Buckets {
		fmt.Printf("[%d]: ", i)
		for b != nil { //проход по цепочке
			fmt.Printf("{%s:%s} -> ", b.Key, b.Value)
			b = b.Next
		}
		fmt.Println("nil")
	}
}

// Сохранение
func (h *HashTable) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(h)
}

// Загрузка
func (h *HashTable) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(h)
}
