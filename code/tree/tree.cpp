#include <iostream>
#include <fstream>
#include "tree.h"

using namespace std;

//инициализация nil узла
void RBTree::initializeTNULL() {
    TNULL = new NodeT(0);//создаем
    TNULL->color = BLACK;//всегда черный
    TNULL->left = nullptr;//нет детей
    TNULL->right = nullptr;//нет детей
}

//конструктор
RBTree::RBTree() {
    initializeTNULL();//созд заглушку
    root = TNULL;//корень это заглушка
}

//рекурсивное удаление
void RBTree::deleteTree(NodeT* node) {
    if (node != TNULL) {//если не заглушка
        deleteTree(node->left);//лево
        deleteTree(node->right);//право
        delete node;
    }
}

//деструктор
RBTree::~RBTree() {
    deleteTree(root);//чистим дерево
    delete TNULL;//чистим заглушку
}

//левый поворот
void RBTree::leftRotate(NodeT* x) {
    NodeT* y = x->right;//правый сын
    x->right = y->left;//перекидываем
    if (y->left != TNULL) y->left->parent = x;//обновляем родителя
    y->parent = x->parent;//связываем с дедом
    if (x->parent == nullptr) root = y;//если корень
    else if (x == x->parent->left) x->parent->left = y;//левый
    else x->parent->right = y;//правый
    y->left = x;//ставим x слева
    x->parent = y;//обновляем родителя
}

//правый поворот
void RBTree::rightRotate(NodeT* x) {
    NodeT* y = x->left;//левый сын
    x->left = y->right;//перекидываем
    if (y->right != TNULL) y->right->parent = x;//обновляем
    y->parent = x->parent;//связываем
    if (x->parent == nullptr) root = y;//корень
    else if (x == x->parent->right) x->parent->right = y;//правый
    else x->parent->left = y;//левый
    y->right = x;//ставим x справа
    x->parent = y;//обновляем
}

//балансировка вставки
void RBTree::insertFix(NodeT* k) {
    NodeT* u;//дядя
    while (k->parent->color == RED) {//пока папа красный
        if (k->parent == k->parent->parent->right) {//папа справа
            u = k->parent->parent->left;//дядя слева
            if (u->color == RED) {//случай 1
                u->color = BLACK;//красим
                k->parent->color = BLACK;//красим
                k->parent->parent->color = RED;//дед красный
                k = k->parent->parent;//идем вверх
            } else {
                if (k == k->parent->left) {//случай 2
                    k = k->parent;//вверх
                    rightRotate(k);//поворот
                }
                k->parent->color = BLACK;//случай 3
                k->parent->parent->color = RED;//дед красный
                leftRotate(k->parent->parent);//поворот
            }
        } else {//зеркальная ситуация
            u = k->parent->parent->right;//дядя справа
            if (u->color == RED) {//случай 1
                u->color = BLACK;//красим
                k->parent->color = BLACK;//красим
                k->parent->parent->color = RED;//дед красный
                k = k->parent->parent;//вверх
            } else {
                if (k == k->parent->right) {//случай 2
                    k = k->parent;//вверх
                    leftRotate(k);//поворот
                }
                k->parent->color = BLACK;//случай 3
                k->parent->parent->color = RED;//дед красный
                rightRotate(k->parent->parent);//поворот
            }
        }
        if (k == root) break;//если дошли до корня
    }
    root->color = BLACK;//корень всегда черный
}

//вставка
void RBTree::insert(int key) {
    NodeT* node = new NodeT(key);//создаем
    node->parent = nullptr;//нет родителя
    node->data = key;//данные
    node->left = TNULL;//заглушка
    node->right = TNULL;//заглушка
    node->color = RED;//новый красный

    NodeT* y = nullptr;//будущий родитель
    NodeT* x = root;//текущий

    while (x != TNULL) {//спуск вниз
        y = x;//сохраняем
        if (node->data < x->data) x = x->left;//лево
        else x = x->right;//право
    }

    node->parent = y;//ставим родителя
    if (y == nullptr) root = node;//если дерево пустое
    else if (node->data < y->data) y->left = node;//левый сын
    else y->right = node;//правый сын

    if (node->parent == nullptr) {//корень
        node->color = BLACK;//черный
        return;
    }
    if (node->parent->parent == nullptr) return;//нет деда

    insertFix(node);//балансировка
}

//трансплантация (замена поддерева)
void RBTree::rbTransplant(NodeT* u, NodeT* v) {
    if (u->parent == nullptr) root = v;//корень
    else if (u == u->parent->left) u->parent->left = v;//левый
    else u->parent->right = v;//правый
    v->parent = u->parent;//родитель
}

//поиск минимума
NodeT* RBTree::minimum(NodeT* node) {
    while (node->left != TNULL) node = node->left;//до упора влево
    return node;
}

//балансировка удаления
void RBTree::deleteFix(NodeT* x) {
    NodeT* s;//брат
    while (x != root && x->color == BLACK) {//пока черный
        if (x == x->parent->left) {//слева
            s = x->parent->right;//брат справа
            if (s->color == RED) {//случай 1
                s->color = BLACK;//красим
                x->parent->color = RED;//красим
                leftRotate(x->parent);//поворот
                s = x->parent->right;//обновляем брата
            }
            if (s->left->color == BLACK && s->right->color == BLACK) {//случай 2
                s->color = RED;//красим
                x = x->parent;//вверх
            } else {
                if (s->right->color == BLACK) {//случай 3
                    s->left->color = BLACK;//красим
                    s->color = RED;//красим
                    rightRotate(s);//поворот
                    s = x->parent->right;//обновляем
                }
                s->color = x->parent->color;//случай 4
                x->parent->color = BLACK;//красим
                s->right->color = BLACK;//красим
                leftRotate(x->parent);//поворот
                x = root;//выход
            }
        } else {//зеркально
            s = x->parent->left;//брат слева
            if (s->color == RED) {//случай 1
                s->color = BLACK;//красим
                x->parent->color = RED;//красим
                rightRotate(x->parent);//поворот
                s = x->parent->left;//обновляем
            }
            if (s->right->color == BLACK && s->left->color == BLACK) {//случай 2
                s->color = RED;//красим
                x = x->parent;//вверх
            } else {
                if (s->left->color == BLACK) {//случай 3
                    s->right->color = BLACK;//красим
                    s->color = RED;//красим
                    leftRotate(s);//поворот
                    s = x->parent->left;//обновляем
                }
                s->color = x->parent->color;//случай 4
                x->parent->color = BLACK;//красим
                s->left->color = BLACK;//красим
                rightRotate(x->parent);//поворот
                x = root;//выход
            }
        }
    }
    x->color = BLACK;//красим в черный
}

//удаление
void RBTree::remove(int key) {
    NodeT* z = TNULL;//узел для удаления
    NodeT *x, *y;//временные
    NodeT* node = root;//поиск
    while (node != TNULL) {
        if (node->data == key) z = node;//нашли
        if (node->data <= key) node = node->right;//право
        else node = node->left;//лево
    }

    if (z == TNULL) return;//не нашли
    y = z;//сохраняем
    Color y_original_color = y->color;//цвет
    if (z->left == TNULL) {//нет левого
        x = z->right;//правый
        rbTransplant(z, z->right);//заменяем
    } else if (z->right == TNULL) {//нет правого
        x = z->left;//левый
        rbTransplant(z, z->left);//заменяем
    } else {//два ребенка
        y = minimum(z->right);//ищем замену
        y_original_color = y->color;//цвет
        x = y->right;//правый сына замены
        if (y->parent == z) x->parent = y;//если близко
        else {
            rbTransplant(y, y->right);//перенос
            y->right = z->right;//связь
            y->right->parent = y;//родитель
        }
        rbTransplant(z, y);//перенос
        y->left = z->left;//связь
        y->left->parent = y;//родитель
        y->color = z->color;//красим
    }
    delete z;//удаляем из памяти
    if (y_original_color == BLACK) deleteFix(x);//баланс
}

//поиск (публичный)
bool RBTree::search(int key) {
    NodeT* curr = root;//корень
    while (curr != TNULL) {
        if (curr->data == key) return true;//нашли
        if (key < curr->data) curr = curr->left;//лево
        else curr = curr->right;//право
    }
    return false;//не нашли
}

//помощник печати
void RBTree::printHelper(NodeT* root, string indent, bool last) const {
    if (root != TNULL) {//если не пустой
        cout << indent;//отступ
        if (last) {
            cout << "R----";//правый
            indent += "   ";
        } else {
            cout << "L----";//левый
            indent += "|  ";
        }
        string sColor = (root->color == RED) ? "RED" : "BLK";//строка цвета
        cout << root->data << "(" << sColor << ")" << endl;//вывод
        printHelper(root->left, indent, false);//левый рек
        printHelper(root->right, indent, true);//правый рек
    }
}

//печать
void RBTree::print() const {
    if (root != TNULL) printHelper(root, "", true);//запуск
    else cout << "Tree is empty" << endl;//пусто
}

//помощник сохранения (PreOrder)
void RBTree::preOrderHelper(NodeT* node, ofstream& file) const {
    if (node != TNULL) {//если есть
        file.write((char*)&node->data, sizeof(int));//пишем ключ
        preOrderHelper(node->left, file);//лево
        preOrderHelper(node->right, file);//право
    }
}

//сохранение
bool RBTree::serialize(const string& filename) const {
    ofstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка
    preOrderHelper(root, file);//запуск обхода
    file.close();
    return true;
}

//загрузка
bool RBTree::deserialize(const string& filename) {
    ifstream file(filename, ios::binary);//открываем
    if (!file.is_open()) return false;//ошибка
    
    //очистка текущего
    deleteTree(root);
    root = TNULL;

    //читаем пока есть данные
    int val;
    while (file.read((char*)&val, sizeof(int))) {
        insert(val);//вставляем
    }
    
    file.close();
    return true;
}
