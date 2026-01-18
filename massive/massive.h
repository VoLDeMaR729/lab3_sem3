#pragma once
#include <string>
#include <iostream>

using namespace std;

//класс динамического массива строк
class Massive {
private:
    string* data;//указатель на массив
    int size;//текущий размер
    int capacity;//вместимость

    void expand();//расширение памяти

public:
    Massive();//конструктор
    Massive(int cap);//конструктор с размером
    ~Massive();//деструктор

    int length() const;//получить длину
    
    void pushBack(string value);//добавить в конец
    void insert(int index, string value);//вставить по индексу
    string get(int index) const;//получить элемент
    void remove(int index);//удалить по индексу
    void set(int index, string value);//заменить
    void print() const;//вывод
    bool serialize(const string& filename) const;//сохранение
    bool deserialize(const string& filename);//загрузка
};
