#pragma once
#include <string>
#include <vector>
#include <list>
#include <iostream>

using namespace std;

//структура узла
struct HashNode {
    string key;//ключ
    string val;//значение
    HashNode* next;//след
    HashNode(string k, string v) : key(k), val(v), next(nullptr) {}//констр
};

//класс хеш-таблицы
class ChainingHash {
private:
    vector<HashNode*> buckets;//вектор списков
    int size;//размер
    int capacity;//емкость
    int hashFunc(string key) const;//хеш-функция

public:
    ChainingHash(int cap = 10);//констр
    ~ChainingHash();//дестр

    void put(string key, string val);//вставка
    string get(string key) const;//получение
    void remove(string key);//удаление
    void print() const;//печать

    //БИНАРНАЯ (для машины)
    void serialize(const string& filename) const;//сохранение bin
    void deserialize(const string& filename);//загрузка bin

    //ФАЙЛОВАЯ / ТЕКСТОВАЯ
    void saveToText(const string& filename) const;//сохранение txt
    void loadFromText(const string& filename);//загрузка txt
};
