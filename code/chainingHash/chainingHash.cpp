#include "chainingHash.h"
#include <fstream>
#include <sstream>

//хеш-функция
int ChainingHash::hashFunc(string key) const {
    int hash = 0;
    for (char c : key) hash += c;//сумма кодов
    return hash % capacity;//остаток
}

//конструктор
ChainingHash::ChainingHash(int cap) {
    capacity = cap;//емкость
    size = 0;//размер
    buckets.resize(capacity, nullptr);//инит вектора
}

//деструктор
ChainingHash::~ChainingHash() {
    for (int i = 0; i < capacity; ++i) {
        HashNode* curr = buckets[i];
        while (curr) {
            HashNode* temp = curr;
            curr = curr->next;
            delete temp;//чистим узлы
        }
    }
}

//вставка
void ChainingHash::put(string key, string val) {
    int idx = hashFunc(key);//индекс
    HashNode* curr = buckets[idx];
    while (curr) {
        if (curr->key == key) {
            curr->val = val;//обновляем
            return;//выход
        }
        curr = curr->next;
    }
    //новый узел
    HashNode* newNode = new HashNode(key, val);
    newNode->next = buckets[idx];
    buckets[idx] = newNode;
    size++;
}

//получение
string ChainingHash::get(string key) const {
    int idx = hashFunc(key);//индекс
    HashNode* curr = buckets[idx];
    while (curr) {
        if (curr->key == key) return curr->val;//нашли
        curr = curr->next;
    }
    return "";//не нашли
}

//удаление
void ChainingHash::remove(string key) {
    int idx = hashFunc(key);//индекс
    HashNode* curr = buckets[idx];
    HashNode* prev = nullptr;
    while (curr) {
        if (curr->key == key) {
            if (prev) prev->next = curr->next;//середина
            else buckets[idx] = curr->next;//голова
            delete curr;//удаляем
            size--;
            return;//выход
        }
        prev = curr;
        curr = curr->next;
    }
}

//печать
void ChainingHash::print() const {
    for (int i = 0; i < capacity; ++i) {
        if (buckets[i]) {
            cout << "[" << i << "]: ";
            HashNode* curr = buckets[i];
            while (curr) {
                cout << "{" << curr->key << ":" << curr->val << "} -> ";
                curr = curr->next;
            }
            cout << "NULL" << endl;
        }
    }
}

// --- БИНАРНАЯ СЕРИАЛИЗАЦИЯ ---
void ChainingHash::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);//открываем
    if (file.is_open()) {
        file.write((char*)&capacity, sizeof(capacity));//емкость
        file.write((char*)&size, sizeof(size));//размер
        
        for (int i = 0; i < capacity; ++i) {HashNode* curr = buckets[i];
            while (curr) {
                int kLen = curr->key.length();
                file.write((char*)&kLen, sizeof(kLen));//длина ключа
                file.write(curr->key.c_str(), kLen);//ключ
                
                int vLen = curr->val.length();
                file.write((char*)&vLen, sizeof(vLen));//длина знач
                file.write(curr->val.c_str(), vLen);//значение
                
                curr = curr->next;
            }
        }
        file.close();
    }
}

void ChainingHash::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);//открываем
    if (file.is_open()) {
        //чистим старое
        for (int i = 0; i < capacity; ++i) {
            HashNode* curr = buckets[i];
            while (curr) {
                HashNode* temp = curr;
                curr = curr->next;
                delete temp;
            }
        }
        buckets.clear();
        
        file.read((char*)&capacity, sizeof(capacity));//емкость
        file.read((char*)&size, sizeof(size));//размер
        buckets.resize(capacity, nullptr);//ресайз
        
        for (int i = 0; i < size; ++i) {
            int kLen;
            file.read((char*)&kLen, sizeof(kLen));//длина ключа
            char* kBuf = new char[kLen + 1];
            file.read(kBuf, kLen);//ключ
            kBuf[kLen] = '\0';
            
            int vLen;
            file.read((char*)&vLen, sizeof(vLen));//длина знач
            char* vBuf = new char[vLen + 1];
            file.read(vBuf, vLen);//значение
            vBuf[vLen] = '\0';
            
            put(string(kBuf), string(vBuf));//вставка
            
            delete[] kBuf;
            delete[] vBuf;
        }
        file.close();
    }
}

// --- ТЕКСТОВАЯ СЕРИАЛИЗАЦИЯ (Файловая) ---
void ChainingHash::saveToText(const string& filename) const {
    ofstream file(filename);//открываем txt
    if (file.is_open()) {
        for (int i = 0; i < capacity; ++i) {
            HashNode* curr = buckets[i];
            while (curr) {
                //формат: ключ значение
                file << curr->key << " " << curr->val << "\n";
                curr = curr->next;
            }
        }
        file.close();
    }
}

void ChainingHash::loadFromText(const string& filename) {
    ifstream file(filename);//открываем txt
    if (file.is_open()) {
        //чистим текущую таблицу перед загрузкой (опционально)
        for (int i = 0; i < capacity; ++i) {
            HashNode* curr = buckets[i];
            while (curr) {
                HashNode* temp = curr;
                curr = curr->next;
                delete temp;
            }
            buckets[i] = nullptr; 
        }
        size = 0;

        string key, val;
        while (file >> key >> val) { //читаем пару слов
            put(key, val);//вставляем
        }
        file.close();
    }
}
