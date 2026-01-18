package structs

import (
	"os"
	"testing"
)

// Тест операций односвязного списка
func TestSList_Operations(t *testing.T) {
	l := NewSList() //создаем список

// Тест на пустом списке
	l.DelHead()
	l.DelTail()
	l.DelVal("X")

	// Тест с одним элементом
	l.AddHead("1") // [1]
	l.DelTail()    // []
	if l.Find("1") { t.Error("DelTail single fail") }

	l.AddTail("1") // [1]
	l.DelHead()    // []
	if l.Find("1") { t.Error("DelHead single fail") }

	l.AddHead("1") // [1]
	l.DelVal("1")  // []
	if l.Find("1") { t.Error("DelVal single fail") }

	// Тест с множеством элементов
	l.AddTail("1")
	l.AddTail("2")
	l.AddTail("3") // 1 -> 2 -> 3

	if !l.Find("2") { t.Error("Find fail") } //поиск

	l.DelHead() // 2 -> 3
	if l.Find("1") { t.Error("DelHead fail") }

	l.DelTail() // 2
	if l.Find("3") { t.Error("DelTail fail") }

	l.AddTail("3") // 2 -> 3
	l.DelVal("3")  // 2 (удаление хвоста по значению)
	if l.Find("3") { t.Error("DelVal tail fail") }
	
	l.DelVal("2") // удаление головы по значению
	if l.Find("2") { t.Error("DelVal head fail") }

	l.DelVal("NonExist") //удаление несуществующего

	l.AddHead("PrintMe")
	l.Print() //для покрытия
}

// Тест сериализации
func TestSList_File(t *testing.T) {
	f := "test_slist.json" //имя файла
	l := NewSList() //создаем
	l.AddHead("Data") //заполняем
	l.Save(f) //сохраняем

	l2 := NewSList() //новый список
	l2.Load(f) //загружаем
	if !l2.Find("Data") { t.Error("Load fail") } //проверка
	os.Remove(f) //чистим
	l2.Load("bad.json") //проверка ошибки
}

func BenchmarkSList_Add(b *testing.B) {
	l := NewSList() //создаем
	for i := 0; i < b.N; i++ { //цикл
		l.AddHead("val") //замер
	}
}
