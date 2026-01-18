#include <fstream>
#include "queue.h"

//конструктор
Queue::Queue() {
    head = nullptr;//пусто
    tail = nullptr;//пусто
    size = 0;//0 элементов
}

//деструктор
Queue::~Queue() {
    while (head) pop();//чистим пока есть
}

//добавить элемент
void Queue::push(string val) {
    QNode* newNode = new QNode(val);//создаем
    if (tail) {//если есть хвост
        tail->next = newNode;//цепляем
        tail = newNode;//обновляем хвост
    } else {
        head = tail = newNode;//первый элемент
    }
    size++;
}

//удалить элемент
string Queue::pop() {
    if (!head) return "";//пусто
    QNode* temp = head;//запоминаем
    string val = head->val;//берем значение
    head = head->next;//сдвигаем
    if (!head) tail = nullptr;//если опустела
    delete temp;//удаляем
    size--;//минус один
    return val;
}

//печать
void Queue::print() const {
    cout << "Queue: ";//префикс
    QNode* curr = head;//итератор
    while (curr) {//цикл
        cout << "[" << curr->val << "] ";//вывод
        curr = curr->next;//дальше
    }
    cout << endl;//конец
}

//размер
int Queue::length() const {
    return size;//возврат
}

//сериализация
bool Queue::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка
    
    file.write((char*)&size, sizeof(size));//пишем размер
    
    QNode* curr = head;//итератор
    while (curr) {//цикл
        size_t len = curr->val.length();//длина
        file.write((char*)&len, sizeof(len));//пишем длину
        file.write(curr->val.c_str(), len);//пишем строку
        curr = curr->next;//дальше
    }
    file.close();
    return true;
}

//десериализация
bool Queue::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка
    
    while (head) pop();//чистим старое
    
    int count;//кол-во
    file.read((char*)&count, sizeof(count));//читаем
    
    for (int i = 0; i < count; i++) {//цикл
        size_t len;//длина
        file.read((char*)&len, sizeof(len));//читаем
        char* buf = new char[len + 1];//буфер
        file.read(buf, len);//читаем данные
        buf[len] = '\0';//терминатор
        push(string(buf));//добавляем
        delete[] buf;//чистим
    }
    file.close();
    return true;
}
