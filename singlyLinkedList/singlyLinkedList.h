#pragma once
#include <string>
#include <iostream>

using namespace std;

//структура узла
struct SNode {
    string val;//значение
    SNode* next;//указатель на след
    SNode(string v) : val(v), next(nullptr) {}//конструктор
};

//класс односвязного списка
class SinglyLinkedList {
private:
    SNode* head;//голова списка

public:
    SinglyLinkedList();//конструктор
    ~SinglyLinkedList();//деструктор

    void pushHead(string val);//в начало
    void pushTail(string val);//в конец
    void deleteHead();//удалить голову
    void deleteTail();//удалить хвост
    void deleteByValue(string val);//удалить по значению
    bool find(string val) const;//поиск
    void print() const;//печать
    
    bool serialize(const string& filename) const;//сохранение
    bool deserialize(const string& filename);//загрузка
};
