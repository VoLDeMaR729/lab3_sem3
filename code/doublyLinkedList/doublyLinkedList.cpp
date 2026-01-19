#include <fstream>
#include "doublyLinkedList.h"

//конструктор
DoublyLinkedList::DoublyLinkedList() {
    head = nullptr;//пусто
    tail = nullptr;//пусто
}

//деструктор
DoublyLinkedList::~DoublyLinkedList() {
    while (head) deleteHead();//чистим пока есть
}

//вспомогательный поиск узла
DNode* DoublyLinkedList::searchNode(string val) const {
    DNode* curr = head;//итератор
    while (curr) {//пока есть элементы
        if (curr->val == val) return curr;//нашли
        curr = curr->next;//дальше
    }
    return nullptr;//не нашли
}

//добавить в голову
void DoublyLinkedList::pushHead(string val) {
    DNode* newNode = new DNode(val);//создаем
    if (!head) {//если пуст
        head = tail = newNode;//он единственный
    } else {
        newNode->next = head;//связь вперед
        head->prev = newNode;//связь назад
        head = newNode;//обновляем голову
    }
}

//добавить в хвост
void DoublyLinkedList::pushTail(string val) {
    DNode* newNode = new DNode(val);//создаем
    if (!tail) {//если пуст
        head = tail = newNode;//он единственный
    } else {
        tail->next = newNode;//связь вперед
        newNode->prev = tail;//связь назад
        tail = newNode;//обновляем хвост
    }
}

//вставка после
void DoublyLinkedList::insertAfter(string target, string val) {
    DNode* node = searchNode(target);//ищем цель
    if (!node) return;//не нашли
    
    if (node == tail) {//если это хвост
        pushTail(val);//просто в конец
        return;
    }
    
    DNode* newNode = new DNode(val);//создаем
    newNode->next = node->next;//связь 1
    newNode->prev = node;//связь 2
    node->next->prev = newNode;//связь 3
    node->next = newNode;//связь 4
}

//вставка до
void DoublyLinkedList::insertBefore(string target, string val) {
    DNode* node = searchNode(target);//ищем
    if (!node) return;//не нашли
    
    if (node == head) {//если голова
        pushHead(val);//просто в начало
        return;
    }
    
    DNode* newNode = new DNode(val);//создаем
    newNode->prev = node->prev;//связь 1
    newNode->next = node;//связь 2
    node->prev->next = newNode;//связь 3
    node->prev = newNode;//связь 4
}

//удалить голову
void DoublyLinkedList::deleteHead() {
    if (!head) return;//пусто
    DNode* temp = head;//запоминаем
    head = head->next;//сдвигаем
    
    if (head) head->prev = nullptr;//обнуляем пред
    else tail = nullptr;//список опустел
    
    delete temp;
}

//удалить хвост
void DoublyLinkedList::deleteTail() {
    if (!tail) return;//пусто
    DNode* temp = tail;//запоминаем
    tail = tail->prev;//сдвигаем
    
    if (tail) tail->next = nullptr;//обнуляем след
    else head = nullptr;//список опустел
    
    delete temp;
}

//удалить по значению
void DoublyLinkedList::deleteByValue(string val) {
    DNode* node = searchNode(val);//ищем
    if (!node) return;//не нашли
    
    if (node == head) deleteHead();//если голова
    else if (node == tail) deleteTail();//если хвост
    else {
        node->prev->next = node->next;//обход слева
        node->next->prev = node->prev;//обход справа
        delete node;
    }
}

//удалить после
void DoublyLinkedList::deleteAfter(string target) {
    DNode* node = searchNode(target);//ищем
    if (!node || !node->next) return;//нечего удалять
    
    DNode* toDel = node->next;//кандидат
    if (toDel == tail) deleteTail();//если хвост
    else {
        node->next = toDel->next;//связь
        toDel->next->prev = node;//связь
        delete toDel;
    }
}

//удалить до
void DoublyLinkedList::deleteBefore(string target) {
    DNode* node = searchNode(target);//ищем
    if (!node || !node->prev) return;//нечего удалять
    
    DNode* toDel = node->prev;//кандидат
    if (toDel == head) deleteHead();//если голова
    else {
        node->prev = toDel->prev;//связь
        toDel->prev->next = node;//связь
        delete toDel;
    }
}

//поиск
bool DoublyLinkedList::find(string val) const {
    return searchNode(val) != nullptr;//рез
}

//печать
void DoublyLinkedList::print() const {
    cout << "DList: ";//префикс
    DNode* curr = head;//итератор
    while (curr) {//цикл
        cout << curr->val << " <-> ";//вывод
        curr = curr->next;//дальше
    }
    cout << "NULL" << endl;//конец
}

//печать с конца
void DoublyLinkedList::printReverse() const {
    cout << "DList (Rev): ";//префикс
    DNode* curr = tail;//с хвоста
    while (curr) {//цикл
        cout << curr->val << " <-> ";//вывод
        curr = curr->prev;//назад
    }
    cout << "NULL" << endl;//конец
}

//сериализация
bool DoublyLinkedList::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка
    
    int count = 0;//счетчик
    DNode* curr = head;//итератор
    while (curr) { count++; curr = curr->next; }//считаем
    
    file.write((char*)&count, sizeof(count));//пишем кол-во
    
    curr = head;//сброс
    while (curr) {//цикл записи
        size_t len = curr->val.length();//длина
        file.write((char*)&len, sizeof(len));//пишем длину
        file.write(curr->val.c_str(), len);//пишем данные
        curr = curr->next;//дальше
    }
    file.close();
    return true;
}

//десериализация
bool DoublyLinkedList::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка
    
    while (head) deleteHead();//чистим старое
    
    int count;//переменная
    file.read((char*)&count, sizeof(count));//читаем кол-во
    
    for (int i = 0; i < count; i++) {//цикл
        size_t len;//длина
        file.read((char*)&len, sizeof(len));//читаем длину
        char* buf = new char[len + 1];//буфер
        file.read(buf, len);//читаем данные
        buf[len] = '\0';
        pushTail(string(buf));//вставляем
        delete[] buf;//чистим
    }
    file.close();
    return true;
}
