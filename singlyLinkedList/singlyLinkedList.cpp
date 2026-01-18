#include <fstream>
#include "singlyLinkedList.h"

//конструктор
SinglyLinkedList::SinglyLinkedList() {
    head = nullptr;//изначально пуст
}

//деструктор
SinglyLinkedList::~SinglyLinkedList() {
    while (head) deleteHead();//чистим память
}

//добавить в голову
void SinglyLinkedList::pushHead(string val) {
    SNode* newNode = new SNode(val);//создаем узел
    newNode->next = head;//новый ссылается на старую голову
    head = newNode;//обновляем голову
}

//добавить в хвост
void SinglyLinkedList::pushTail(string val) {
    SNode* newNode = new SNode(val);//создаем
    if (!head) {//если список пуст
        head = newNode;//он голова
        return;//выход
    }
    SNode* curr = head;//итератор
    while (curr->next) curr = curr->next;//ищем последний
    curr->next = newNode;//цепляем
}

//удалить голову
void SinglyLinkedList::deleteHead() {
    if (!head) return;//пусто
    SNode* temp = head;//запоминаем
    head = head->next;//сдвигаем
    delete temp;//удаляем
}

//удалить хвост
void SinglyLinkedList::deleteTail() {
    if (!head) return;//пусто
    if (!head->next) {//один элемент
        delete head;//удаляем
        head = nullptr;//обнуляем
        return;//выход
    }
    SNode* curr = head;//итератор
    while (curr->next->next) curr = curr->next;//до пред-последнего
    delete curr->next;//удаляем последний
    curr->next = nullptr;//обнуляем хвост
}

//удалить по значению
void SinglyLinkedList::deleteByValue(string val) {
    if (!head) return;//пусто
    if (head->val == val) {//если в голове
        deleteHead();//удаляем голову
        return;//выход
    }
    SNode* curr = head;//итератор
    while (curr->next && curr->next->val != val) {
        curr = curr->next;//ищем
    }
    if (curr->next) {//нашли
        SNode* toDel = curr->next;//кандидат
        curr->next = toDel->next;//перекидываем связь
        delete toDel;//удаляем
    }
}

//поиск
bool SinglyLinkedList::find(string val) const {
    SNode* curr = head;//итератор
    while (curr) {//пока есть элементы
        if (curr->val == val) return true;//нашли
        curr = curr->next;//дальше
    }
    return false;//не нашли
}

//печать
void SinglyLinkedList::print() const {
    SNode* curr = head;//итератор
    cout << "SList: ";//префикс
    while (curr) {//цикл
        cout << curr->val << " -> ";//вывод
        curr = curr->next;//дальше
    }
    cout << "NULL" << endl;//конец
}

//сериализация
bool SinglyLinkedList::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка
    
    //считаем размер
    int count = 0;
    SNode* curr = head;
    while (curr) { count++; curr = curr->next; }
    
    file.write((char*)&count, sizeof(count));//пишем кол-во узлов
    
    curr = head;
    while (curr) {
        size_t len = curr->val.length();//длина строки
        file.write((char*)&len, sizeof(len));//пишем длину
        file.write(curr->val.c_str(), len);//пишем данные
        curr = curr->next;//дальше
    }
    
    file.close();//закрываем
    return true;//успех
}

//десериализация
bool SinglyLinkedList::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка
    
    while (head) deleteHead();//чистим текущий
    
    int count;
    file.read((char*)&count, sizeof(count));//читаем кол-во
    
    for (int i = 0; i < count; i++) {
        size_t len;
        file.read((char*)&len, sizeof(len));//читаем длину
        char* buf = new char[len + 1];//буфер
        file.read(buf, len);//читаем данные
        buf[len] = '\0';//терминатор
        pushTail(string(buf));//добавляем в конец
        delete[] buf;//чистим буфер
    }
    
    file.close();
    return true;
}
