#pragma once
#include <iostream>
#include <fstream>
#include <string>

using namespace std;

//цвета узла
enum Color { RED, BLACK };

//структура узла
struct NodeT {
    int data;//данные (ключ)
    Color color;//цвет
    NodeT *left, *right, *parent;//связи
    //конструктор
    NodeT(int d) : data(d), color(RED), left(nullptr), right(nullptr), parent(nullptr) {}
};

//класс красно-черного дерева
class RBTree {
private:
    NodeT *root;//корень
    NodeT *TNULL;//пустой лист (заглушка)

    void initializeTNULL();//инициализация заглушки
    void preOrderHelper(NodeT* node, ofstream& file) const;//для сохранения
    void deleteTree(NodeT* node);//очистка памяти
    void printHelper(NodeT* root, string indent, bool last) const;//печать

    //балансировка
    void leftRotate(NodeT* x);//левый поворот
    void rightRotate(NodeT* x);//правый поворот
    void insertFix(NodeT* k);//фикс после вставки
    void deleteFix(NodeT* x);//фикс после удаления
    void rbTransplant(NodeT* u, NodeT* v);//перенос узлов
    NodeT* minimum(NodeT* node);//минимум в поддереве

public:
    RBTree();//конструктор
    ~RBTree();//деструктор

    void insert(int key);//вставка
    void remove(int key);//удаление
    bool search(int key);//поиск
    void print() const;//вывод

    bool serialize(const string& filename) const;//сохранение
    bool deserialize(const string& filename);//загрузка
};
