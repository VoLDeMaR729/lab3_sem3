#pragma once
#include <string>
#include <iostream>

using namespace std;

//структура узла очереди
struct QNode {
    string val;//значение
    QNode* next;//указатель на след
    QNode(string v) : val(v), next(nullptr) {}//конструктор
};

//класс очереди
class Queue {
private:
    QNode* head;//голова (отсюда забираем)
    QNode* tail;//хвост (сюда добавляем)
    int size;//размер

public:
    Queue();//конструктор
    ~Queue();//деструктор

    void push(string val);//добавить
    string pop();//удалить
    void print() const;
    int length() const;//узнать размер
    
    bool serialize(const string& filename) const;//сохранение
    bool deserialize(const string& filename);//загрузка
};
