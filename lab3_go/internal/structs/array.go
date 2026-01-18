package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

//структура динамического массива
type DynArray struct {
	Data []string `json:"data"` //данные массива (слайс) с тегом для json
}

// Конструктор
func NewDynArray() *DynArray {
	return &DynArray{Data: make([]string, 0)} //инициализируем пустой слайс
}

// Добавление элемента в конец
func (d *DynArray) Add(val string) {
	d.Data = append(d.Data, val) //встроенная функция расширения
}

// Вставка элемента по индексу
func (d *DynArray) Insert(idx int, val string) {
	if idx < 0 || idx > len(d.Data) { return } //проверка границ
	//раздвигаем слайс: берем часть до idx, добавляем val, потом добавляем остаток
	d.Data = append(d.Data[:idx], append([]string{val}, d.Data[idx:]...)...)
}

// Удаление элемента по индексу
func (d *DynArray) Remove(idx int) {
	if idx < 0 || idx >= len(d.Data) { return } //проверка границ
	//сшиваем часть до idx и часть после idx+1
	d.Data = append(d.Data[:idx], d.Data[idx+1:]...)
}

// Получение элемента
func (d *DynArray) Get(idx int) string {
	if idx < 0 || idx >= len(d.Data) { return "" } //если индекс неверен
	return d.Data[idx] //возвращаем значение
}

// Изменение элемента
func (d *DynArray) Set(idx int, val string) {
	if idx >= 0 && idx < len(d.Data) { d.Data[idx] = val } //если индекс ок, меняем
}

// Печать массива
func (d *DynArray) Print() {
	fmt.Print("Array: ", d.Data, "\n") //вывод всего слайса
}

// Сохранение в файл
func (d *DynArray) Save(filename string) error {
	file, err := os.Create(filename) //создаем файл
	if err != nil { return err } //если ошибка создания
	defer file.Close() //закроем при выходе
	return json.NewEncoder(file).Encode(d) //пишем json
}

// Загрузка из файла
func (d *DynArray) Load(filename string) error {
	file, err := os.Open(filename) //открываем файл
	if err != nil { return err } //если ошибка
	defer file.Close() //закроем потом
	return json.NewDecoder(file).Decode(d) //читаем json в структуру
}
