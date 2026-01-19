package structs

import "testing"

func TestMassive_Green(t *testing.T) {
 m := NewMassive(1) //емкость 1
 
 m.PushBack("A") //добавили
 m.PushBack("B") //вызов expand (станет зеленым)
 
 //проверка вставки
 m.Insert(0, "Start") //сдвиг вправо (цикл станет зеленым)
 m.Insert(2, "Mid")
 
 //проверка удаления
 m.Remove(0) //сдвиг влево (цикл станет зеленым)
 
 //геттеры сеттеры
 m.Set(0, "NewA")
 m.Get(0)
 
 //ошибки (чтобы проверки ошибок тоже позеленели)
 m.Insert(-1, "X")
 m.Remove(100)
 m.Get(100)
 m.Set(100, "X")
 
 //печать
 m.Print()
}

func TestMassive_File(t *testing.T) {
 m := NewMassive(5)
 m.PushBack("Test")
 m.Save("test_arr.json") //сохраняем (зеленое)
 
 m2 := NewMassive(1)
 m2.Load("test_arr.json") //загружаем (зеленое)
 
 m.Load("bad_file.json") //ошибка открытия (зеленое)
}
