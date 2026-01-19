#pragma once
#include <string>
#include <iostream>

using namespace std;

//структура узла стека
struct StNode {
    string val;//значение
    StNode* next;//указатель на след
    StNode(string v) : val(v), next(nullptr) {}//конструктор
};

//класс стека
class Stack {
private:
    StNode* topNode;//вершина стека

public:
    Stack();//конструктор
    ~Stack();//деструктор

    void push(string val);//добавить
    string pop();//удалить
    void print() const;
    
    bool serialize(const string& filename) const;//сохранение
    bool deserialize(const string& filename);//загрузка
};
