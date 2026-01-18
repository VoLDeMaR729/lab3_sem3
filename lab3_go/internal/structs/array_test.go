package structs

import (
	"os"
	"testing"
)

// Тест основных операций массива
func TestDynArray_Operations(t *testing.T) {
	arr := NewDynArray() //создаем массив
	arr.Add("A")
	arr.Add("B")

	if arr.Get(0) != "A" { t.Error("Get(0) fail") } //проверка значения
	if arr.Get(1) != "B" { t.Error("Get(1) fail") } //проверка значения

	arr.Insert(1, "C") //вставка в середину: [A, C, B]
	if arr.Get(1) != "C" { t.Error("Insert fail") } //проверка вставки

	arr.Set(0, "Z") //замена значения: [Z, C, B]
	if arr.Get(0) != "Z" { t.Error("Set fail") } //проверка замены

	arr.Remove(1) //удаляем C -> [Z, B]
	if arr.Get(1) != "B" { t.Error("Remove fail") } //проверка сдвига

	// Тесты границ (чтобы зайти в if-ы ошибок)
	arr.Get(99); //выход за границы get
	arr.Set(99, "X"); //выход за границы set
	arr.Insert(99, "X"); //выход за границы insert
	arr.Remove(99); //выход за границы remove
	arr.Insert(-1, "X"); //отрицательный индекс
	
	arr.Print() //для покрытия функции Print
}

// Тест сохранения и загрузки
func TestDynArray_File(t *testing.T) {
	f := "test_arr.json" //имя файла
	a1 := NewDynArray() //создаем первый
	a1.Add("SaveMe") //добавляем данные
	a1.Save(f) //сохраняем в файл

	a2 := NewDynArray() //создаем второй
	a2.Load(f) //загружаем из файла
	if a2.Get(0) != "SaveMe" { t.Error("Load fail") } //сверяем данные
	os.Remove(f) //удаляем временный файл

	a2.Load("bad_file.json") //проверка ошибки открытия
}

// Бенчмарк добавления
func BenchmarkDynArray_Add(b *testing.B) {
	arr := NewDynArray() //инициализация
	for i := 0; i < b.N; i++ { //цикл бенчмарка
		arr.Add("val") //замеряем эту операцию
	}
}
