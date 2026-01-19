#include <fstream>
#include <vector>
#include "stack.h"

//конструктор
Stack::Stack() {
    topNode = nullptr;//пусто
}

//деструктор
Stack::~Stack() {
    while (topNode) pop();//чистим
}

//добавить наверх
void Stack::push(string val) {
    StNode* newNode = new StNode(val);//создаем
    newNode->next = topNode;//ссылается на старую вершину
    topNode = newNode;//новая вершина
}

//удалить сверху
string Stack::pop() {
    if (!topNode) return "";//пусто
    StNode* temp = topNode;//запоминаем
    string val = temp->val;//берем значение
    topNode = topNode->next;//сдвигаем вниз
    delete temp;
    return val;
}

//печать
void Stack::print() const {
    cout << "Stack: ";//префикс
    StNode* curr = topNode;//итератор
    while (curr) {//цикл
        cout << "[" << curr->val << "] ";//вывод
        curr = curr->next;//дальше
    }
    cout << endl;
}

//сериализация
bool Stack::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка
    
    //считаем кол-во
    int count = 0;
    StNode* curr = topNode;
    while (curr) { count++; curr = curr->next; }
    
    file.write((char*)&count, sizeof(count));//пишем размер
    
    curr = topNode;//сброс
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
bool Stack::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка
    
    while (topNode) pop();//чистим старое
    
    int count;//кол-во
    file.read((char*)&count, sizeof(count));//читаем
    
    vector<string> buffer;//временный вектор
    for (int i = 0; i < count; i++) {//цикл чтения
        size_t len;//длина
        file.read((char*)&len, sizeof(len));//читаем длину
        char* buf = new char[len + 1];//буфер
        file.read(buf, len);//читаем данные
        buf[len] = '\0';//терминатор
        buffer.push_back(string(buf));//в вектор
        delete[] buf;
    }
    
    //восстанавливаем порядок (с конца вектора), чтобы порядок в стеке остался прежним
    for (int i = count - 1; i >= 0; i--) {
        push(buffer[i]);//добавляем
    }
    
    file.close();
    return true;
}
