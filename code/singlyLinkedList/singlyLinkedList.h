#pragma once
#include <string>
#include <iostream>

using namespace std;

//структура узла
struct Node {
    string val;//значение
    Node* next;//след
    Node(string v) : val(v), next(nullptr) {}//констр
};

//класс списка
class SinglyLinkedList {
private:
    Node* head;//голова
public:
    SinglyLinkedList();//констр
    ~SinglyLinkedList();//дестр

    void pushHead(string val);//в начало
    void pushTail(string val);//в конец
    
    //новые методы
    void insertAfter(string target, string val);//вставка после
    void insertBefore(string target, string val);//вставка до
    void deleteAfter(string target);//уд после
    void deleteBefore(string target);//уд до

    void deleteHead();//уд голову
    void deleteTail();//уд хвост
    void deleteByValue(string val);//уд по значению

    bool find(string val) const;//поиск
    void print() const;//печать

    void serialize(const string& filename) const;//сохранение
    void deserialize(const string& filename);//загрузка
};


