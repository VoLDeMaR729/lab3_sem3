package structs

import (
	"os"
	"testing"
)

// Тест стека
func TestStack_Ops(t *testing.T) {
	s := NewStack() //создаем
	if s.Pop() != "" { t.Error("Empty pop fail") } //пустой поп

	s.Push("1") //кладем 1
	s.Push("2") //кладем 2

	if s.Pop() != "2" { t.Error("LIFO fail 1") } //должен быть 2 (последний)
	if s.Pop() != "1" { t.Error("LIFO fail 2") } //должен быть 1
	
	s.Print() //печать
}

// Тест файла
func TestStack_File(t *testing.T) {
	f := "test_s.json" //имя
	s := NewStack() //стэк
	s.Push("Y") //данные
	s.Save(f) //сохранить

	s2 := NewStack() //новый
	s2.Load(f) //загрузить
	if s2.Pop() != "Y" { t.Error("Load fail") } //проверка
	os.Remove(f) //очистка
	s2.Load("no.json") //ошибка
}

func BenchmarkStack_Push(b *testing.B) {
	s := NewStack() //инит
	for i := 0; i < b.N; i++ { //цикл
		s.Push("x") //действие
	}
}
