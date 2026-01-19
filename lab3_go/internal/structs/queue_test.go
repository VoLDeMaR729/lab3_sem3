package structs

import (
	"os"
	"testing"
)

// Тест очереди
func TestQueue_Ops(t *testing.T) {
	q := NewQueue() //создаем
	if q.Pop() != "" { t.Error("Empty pop fail") } //проверка пустой

	q.Push("1") //вставка 1
	q.Push("2") //вставка 2
	
	if q.Pop() != "1" { t.Error("FIFO fail 1") } //должен быть 1
	if q.Pop() != "2" { t.Error("FIFO fail 2") } //должен быть 2
	
	q.Print() //покрытие
}

// Тест сохранения
func TestQueue_File(t *testing.T) {
	f := "test_q.json" //файл
	q := NewQueue() //создаем
	q.Push("X") //пишем
	q.Save(f) //сейв

	q2 := NewQueue() //новый
	q2.Load(f) //лоад
	if q2.Pop() != "X" { t.Error("Load fail") } //чек
	os.Remove(f) //удаление
	q2.Load("err.json") //ошибка
}

// Бенчмарк
func BenchmarkQueue_Push(b *testing.B) {
	q := NewQueue() //инит
	for i := 0; i < b.N; i++ { //цикл
		q.Push("x") //действие
	}
}
