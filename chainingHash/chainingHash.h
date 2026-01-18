#pragma once
#include <string>
#include <vector>
#include <iostream>

using namespace std;

//структура узла для хеш-таблицы
struct NodeH {
    string key;//ключ
    string value;//значение
    NodeH* next;//указатель на следующий узел
    //конструктор узла
    NodeH(string k, string v) : key(k), value(v), next(nullptr) {}
};

//класс хеш-таблицы
class ChainingHash {
private:
    vector<NodeH*> buckets;//вектор списков
    int capacity;//количество корзин
    int size;//количество элементов

    size_t hashFunc(const string& key);//хеш-функция

public:
    ChainingHash(int cap = 10);//конструктор
    ~ChainingHash();//деструктор

    void put(string key, string value);//вставка
    string get(string key);//получение
    void remove(string key);//удаление
    void print() const;//печать

    bool serialize(const string& filename) const;//сохранение
    bool deserialize(const string& filename);//загрузка
};
