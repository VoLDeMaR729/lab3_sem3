#include <iostream>
#include <fstream> 
#include "chainingHash.h"

using namespace std;

//конструктор
ChainingHash::ChainingHash(int cap) : capacity(cap), size(0) {
    buckets.resize(capacity, nullptr);//выделяем память под корзины
}

//деструктор
ChainingHash::~ChainingHash() {
    for (auto node : buckets) {//проход по всем корзинам
        while (node) {//пока список не пуст
            NodeH* temp = node;//запоминаем текущий
            node = node->next;//сдвигаем указатель
            delete temp;
        }
    }
}

//хеш-функция (алгоритм djb2)
size_t ChainingHash::hashFunc(const string& key) {
    size_t hash = 5381;//начальное значение
    for (char c : key) {//проход по символам
        hash = ((hash << 5) + hash) + c;//сдвиг и сложение
    }
    return hash % capacity;//возвращаем индекс
}

//добавление или обновление
void ChainingHash::put(string key, string value) {
    size_t idx = hashFunc(key);//вычисляем индекс
    NodeH* curr = buckets[idx];//берем голову списка

    //проверяем наличие ключа
    while (curr) {//пока есть узлы
        if (curr->key == key) {//если нашли ключ
            curr->value = value;//обновляем значение
            return;//выход
        }
        curr = curr->next;//идем дальше
    }

    //ключ не найден, создаем новый
    NodeH* newNode = new NodeH(key, value);//создаем узел
    newNode->next = buckets[idx];//вставляем в начало
    buckets[idx] = newNode;//обновляем голову
    size++;
}

//получение значения
string ChainingHash::get(string key) {
    size_t idx = hashFunc(key);//вычисляем индекс
    NodeH* curr = buckets[idx];//берем голову списка

    while (curr) {//проход по цепочке
        if (curr->key == key) return curr->value;//нашли возвращаем
        curr = curr->next;//дальше
    }
    return "";//не нашли
}

//удаление элемента
void ChainingHash::remove(string key) {
    size_t idx = hashFunc(key);//вычисляем индекс
    NodeH* curr = buckets[idx];//текущий узел
    NodeH* prev = nullptr;//предыдущий узел

    while (curr) {//проход по цепочке
        if (curr->key == key) {//нашли ключ
            if (prev) {//если не голова
                prev->next = curr->next;//связываем пред и след
            } else {//если голова
                buckets[idx] = curr->next;//сдвигаем голову
            }
            delete curr;//удаляем узел
            size--;//уменьшаем размер
            return;//выход
        }
        prev = curr;//шаг вперед
        curr = curr->next;//шаг вперед
    }
}

//печать таблицы
void ChainingHash::print() const {
    for (int i = 0; i < capacity; i++) {//цикл по корзинам
        cout << "[" << i << "]: ";//вывод индекса
        NodeH* curr = buckets[i];//текущий узел
        while (curr) {//проход по списку
            cout << "{" << curr->key << ":" << curr->value << "} -> ";//вывод пары
            curr = curr->next;
        }
        cout << "NULL" << endl;//конец строки
    }
}

//сохранение в файл
bool ChainingHash::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка

    file.write((char*)&capacity, sizeof(capacity));//пишем емкость
    file.write((char*)&size, sizeof(size));//пишем размер

    for (int i = 0; i < capacity; i++) {//цикл по корзинам
        NodeH* curr = buckets[i];//голова
        
        //считаем длину цепочки
        int chainLen = 0;//счетчик
        NodeH* temp = curr;//временный
        while (temp) {//пока есть
            chainLen++;//плюс один
            temp = temp->next;//дальше
        }
        file.write((char*)&chainLen, sizeof(chainLen));//пишем длину цепочки

        //записываем элементы
        while (curr) {//проход по цепочке
            size_t kLen = curr->key.length();//длина ключа
            file.write((char*)&kLen, sizeof(kLen));//пишем длину
            file.write(curr->key.c_str(), kLen);//пишем ключ

            size_t vLen = curr->value.length();//длина значения
            file.write((char*)&vLen, sizeof(vLen));//пишем длину
            file.write(curr->value.c_str(), vLen);//пишем значение

            curr = curr->next;//дальше
        }
    }
    file.close();
    return true;
}

//загрузка из файла
bool ChainingHash::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка

    //очищаем текущую таблицу
    for (auto node : buckets) {//проход по корзинам
        while (node) {//пока есть узлы
            NodeH* temp = node;//запоминаем
            node = node->next;//сдвигаем
            delete temp;
        }
    }
    buckets.clear();//очищаем вектор

    file.read((char*)&capacity, sizeof(capacity));//читаем емкость
    file.read((char*)&size, sizeof(size));//читаем размер
    buckets.resize(capacity, nullptr);//ресайз вектора

    for (int i = 0; i < capacity; i++) {//цикл по корзинам
        int chainLen;//длина цепочки
        file.read((char*)&chainLen, sizeof(chainLen));//читаем длину

        //читаем цепочку
        for (int j = 0; j < chainLen; j++) {//цикл по узлам
            size_t len;//переменная длины
            
            //читаем ключ
            file.read((char*)&len, sizeof(len));//читаем длину
            char* kBuf = new char[len + 1];//буфер
            file.read(kBuf, len);//читаем данные
            kBuf[len] = '\0';//терминатор
            string key(kBuf);//создаем строку
            delete[] kBuf;

            //читаем значение
            file.read((char*)&len, sizeof(len));//читаем длину
            char* vBuf = new char[len + 1];//буфер
            file.read(vBuf, len);//читаем данные
            vBuf[len] = '\0';//терминатор
            string val(vBuf);//создаем строку
            delete[] vBuf;

            //восстанавливаем узел
            NodeH* newNode = new NodeH(key, val);//создаем
            newNode->next = buckets[i];//вставляем в начало
            buckets[i] = newNode;//обновляем голову
        }
    }
    file.close();
    return true;
}
