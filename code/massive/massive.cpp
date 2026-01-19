#include <fstream>
#include "massive.h"

//конструктор по умолчанию
Massive::Massive() {
    size = 0;//размер 0
    capacity = 10;//начальная емкость
    data = new string[capacity];//выделяем память
}

//конструктор с заданным размером
Massive::Massive(int cap) {
    size = 0;//размер 0
    capacity = cap;//емкость
    if (capacity < 1) capacity = 1;//защита от дурака
    data = new string[capacity];//память
}

//деструктор
Massive::~Massive() {
    delete[] data;//удаляем массив
}

//вспомогательная функция расширения памяти
void Massive::expand() {
    capacity *= 2;//увеличиваем в 2 раза
    string* newData = new string[capacity];//новая память
    
    //копируем старые данные
    for (int i = 0; i < size; i++) {
        newData[i] = data[i];//копирование
    }

    delete[] data;//удаляем старое
    data = newData;//переключаем указатель
}

//получить текущее количество элементов
int Massive::length() const {
    return size;//возврат размера
}

//добавление элемента в конец
void Massive::pushBack(string value) {
    if (size >= capacity) {
        expand();//если места нет расширяем
    }
    data[size] = value;//записываем
    size++;//увеличиваем счетчик
}

//вставка элемента по индексу
void Massive::insert(int index, string value) {
    if (index < 0 || index > size) return;//проверка границ

    if (size >= capacity) {
        expand();//проверка места
    }

    //сдвигаем элементы вправо
    for (int i = size; i > index; i--) {
        data[i] = data[i - 1];//сдвиг
    }

    data[index] = value;//вставляем
    size++;//плюс размер
}

//получение элемента по индексу
string Massive::get(int index) const {
    if (index < 0 || index >= size) return "";//проверка
    return data[index];//возврат
}

//удаление элемента по индексу
void Massive::remove(int index) {
    if (index < 0 || index >= size) return;//проверка

    //сдвигаем элементы влево
    for (int i = index; i < size - 1; i++) {
        data[i] = data[i + 1];//сдвиг
    }

    size--;//минус размер
}

//замена значения по индексу
void Massive::set(int index, string value) {
    if (index >= 0 && index < size) {
        data[index] = value;//запись
    }
}

//вывод массива в консоль
void Massive::print() const {
    cout << "Massive: [ ";//скобка
    for (int i = 0; i < size; i++) {
        cout << data[i] << " ";//вывод элемента
    }
    cout << "]" << endl;//конец
}

//сохранение в файл
bool Massive::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);//открываем бинарно
    if (!file.is_open()) return false;//ошибка открытия
    
    file.write((char*)&size, sizeof(size));//пишем размер
    file.write((char*)&capacity, sizeof(capacity));//пишем емкость

    for (int i = 0; i < size; i++) {
        size_t len = data[i].length();//длина строки
        file.write((char*)&len, sizeof(len));//пишем длину
        file.write(data[i].c_str(), len);//пишем данные
    }
    
    file.close();//закрываем
    return true;
}

//загрузка из файла
bool Massive::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка

    delete[] data;//чистим старое
    
    file.read((char*)&size, sizeof(size));//читаем размер
    file.read((char*)&capacity, sizeof(capacity));//читаем емкость
    
    data = new string[capacity];//выделяем память

    for (int i = 0; i < size; i++) {
        size_t len;//переменная длины
        file.read((char*)&len, sizeof(len));//читаем длину
        
        char* buf = new char[len + 1];//буфер
        file.read(buf, len);//читаем символы
        buf[len] = '\0';
        
        data[i] = string(buf);//преобразуем
        delete[] buf;//чистим буфер
    }
    
    file.close();//закрываем
    return true;
}
