#include "singlyLinkedList.h"
#include <fstream>

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
    Node* newNode = new Node(val);//создаем узел
    newNode->next = head;//новый ссылается на старую голову
    head = newNode;//обновляем голову
}

//добавить в хвост
void SinglyLinkedList::pushTail(string val) {
    Node* newNode = new Node(val);//создаем
    if (!head) {//если список пуст
        head = newNode;//он голова
        return;//выход
    }
    Node* curr = head;//итератор
    while (curr->next) curr = curr->next;//ищем последний
    curr->next = newNode;//цепляем
}

//вставка после
void SinglyLinkedList::insertAfter(string target, string val) {
    Node* curr = head;//итератор
    while (curr && curr->val != target) {//ищем цель
        curr = curr->next;//дальше
    }
    if (curr) {//нашли
        Node* newNode = new Node(val);//создаем
        newNode->next = curr->next;//цепляем хвост
        curr->next = newNode;//цепляем к текущему
    }
}

//вставка до
void SinglyLinkedList::insertBefore(string target, string val) {
    if (!head) return;//пусто
    if (head->val == target) {//если цель в голове
        pushHead(val);//вставляем в начало
        return;//выход
    }
    Node* curr = head;//итератор
    while (curr->next && curr->next->val != target) {//ищем пред-цель
        curr = curr->next;//дальше
    }
    if (curr->next) {//нашли
        Node* newNode = new Node(val);//создаем
        newNode->next = curr->next;//цепляем
        curr->next = newNode;//вставляем
    }
}

//удалить голову
void SinglyLinkedList::deleteHead() {
    if (!head) return;//пусто
    Node* temp = head;//запоминаем
    head = head->next;//сдвигаем
    delete temp;//удаляем
}

//удалить хвост
void SinglyLinkedList::deleteTail() {
    if (!head) return;//пусто
    if (!head->next) {//один элемент
        deleteHead();//удаляем голову
        return;//выход
    }
    Node* curr = head;//итератор
    while (curr->next->next) curr = curr->next;//до пред-последнего
    delete curr->next;//удаляем последний
    curr->next = nullptr;//обнуляем хвост
}

//удалить по значению
void SinglyLinkedList::deleteByValue(string val) {
    if (!head) return;//пусто
    if (head->val == val) {//если в голове
        deleteHead();//удаляем
        return;//выход
    }
    Node* curr = head;//итератор
    while (curr->next && curr->next->val != val) {
        curr = curr->next;//ищем
    }
    if (curr->next) {//нашли
        Node* toDel = curr->next;//кандидат
        curr->next = toDel->next;//перекидываем
        delete toDel;//удаляем
    }
}

//удалить после
void SinglyLinkedList::deleteAfter(string target) {
    Node* curr = head;//итератор
    while (curr && curr->val != target) {//ищем цель
        curr = curr->next;//дальше
    }
    if (curr && curr->next) {//нашли и есть след
        Node* toDel = curr->next;//кандидат
        curr->next = toDel->next;//перекидываем
        delete toDel;//удаляем
    }
}

//удалить до
void SinglyLinkedList::deleteBefore(string target) {
    if (!head || !head->next) return;//мало элементов
    if (head->next->val == target) {//если цель вторая
        deleteHead();//удаляем голову
        return;//выход
    }
    Node* curr = head;//итератор
    while (curr->next->next && curr->next->next->val != target) {//ищем за два шага
        curr = curr->next;//дальше
    }
    if (curr->next->next) {//нашли
        Node* toDel = curr->next;//кандидат
        curr->next = toDel->next;//связываем
        delete toDel;//удаляем
    }
}

//поиск
bool SinglyLinkedList::find(string val) const {
    Node* curr = head;//итератор
    while (curr) {//пока есть
        if (curr->val == val) return true;//нашли
        curr = curr->next;//дальше
    }
    return false;//не нашли
}

//печать
void SinglyLinkedList::print() const {
    Node* curr = head;//итератор
    cout << "SList: ";//префикс
    while (curr) {//цикл
        cout << curr->val << " -> ";//вывод
        curr = curr->next;//дальше
    }
    cout << "NULL" << endl;//конец
}

//сериализация
void SinglyLinkedList::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);//открываем
    if (file.is_open()) {
        Node* curr = head;
        while (curr) {
            int len = curr->val.length();//длина
            file.write((char*)&len, sizeof(len));//пишем длину
            file.write(curr->val.c_str(), len);//пишем данные
            curr = curr->next;//дальше
        }
        file.close();//закрываем
    }
}

//десериализация
void SinglyLinkedList::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);//открываем
    if (file.is_open()) {
        while (head) deleteHead();//чистим
        
        while (file.peek() != EOF) {//пока не конец
            int len;
            if (!file.read((char*)&len, sizeof(len))) break;//читаем длину
            char* buf = new char[len + 1];//буфер
            file.read(buf, len);//читаем данные
            buf[len] = '\0';//терминатор
            pushTail(string(buf));//добавляем
            delete[] buf;//чистим
        }
        file.close();
    }
}
