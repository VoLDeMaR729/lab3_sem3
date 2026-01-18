#pragma once
#include <string>
#include <iostream>

using namespace std;

//структура узла
struct DNode {
    string val;//значение
    DNode* next;//след
    DNode* prev;//пред
    DNode(string v) : val(v), next(nullptr), prev(nullptr) {}//констр
};

//класс двусвязного списка
class DoublyLinkedList {
private:
    DNode* head;//голова
    DNode* tail;//хвост

    DNode* searchNode(string val) const;//вспомогательный поиск

public:
    DoublyLinkedList();//констр
    ~DoublyLinkedList();//дестр

    void pushHead(string val);//в начало
    void pushTail(string val);//в конец
    void insertAfter(string target, string val);//вставка после
    void insertBefore(string target, string val);//вставка до
    
    void deleteHead();//уд голову
    void deleteTail();//уд хвост
    void deleteByValue(string val);//уд по значению
    void deleteAfter(string target);//уд после
    void deleteBefore(string target);//уд до

    bool find(string val) const;//поиск
    void print() const;//печать
    void printReverse() const;//печать с конца

    bool serialize(const string& filename) const;//сохранение
    bool deserialize(const string& filename);//загрузка
};
