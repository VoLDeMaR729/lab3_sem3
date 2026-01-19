package structs

import (
	"os"
	"testing"
)

// Тест двусвязного списка
func TestDList_Operations(t *testing.T) {
	d := NewDList() //создаем
	d.AddHead("A") //добавляем
	d.AddTail("B") //добавляем
	d.AddTail("C") // A <-> B <-> C
	
	if !d.Find("A") { t.Error("Find A fail") } //поиск
	
	d.DelVal("A") //удаление головы
	if d.Find("A") { t.Error("DelVal head fail") } //проверка

	d.DelVal("C") //удаление хвоста
	if d.Find("C") { t.Error("DelVal tail fail") } //проверка
	
	d.DelVal("B") //удаление последнего
	d.DelVal("Z") //удаление несуществующего

	d.Print() //покрытие Print
}

// Тест файлов
func TestDList_File(t *testing.T) {
	f := "test_dlist.json" //файл
	d := NewDList() //создаем
	d.AddTail("Data") //пишем
	d.Save(f) //сохраняем

	d2 := NewDList() //второй
	d2.Load(f) //загружаем
	if !d2.Find("Data") { t.Error("Load fail") } //сверяем
	os.Remove(f) //удаляем
	d2.Load("no.json") //ошибка
}

func BenchmarkDList_Add(b *testing.B) {
	d := NewDList() //инит
	for i := 0; i < b.N; i++ { //цикл
		d.AddTail("val") //замер
	}
}
