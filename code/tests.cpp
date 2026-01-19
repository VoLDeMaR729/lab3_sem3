#include <gtest/gtest.h> //библиотека тестов
#include <string>
#include <cstdlib> //для рандома
#include <vector>  //вектор
#include <algorithm> //алгоритмы

// Подключаем структуры
#include "massive/massive.h" //массив
#include "singlyLinkedList/singlyLinkedList.h" //односвязный
#include "doublyLinkedList/doublyLinkedList.h" //двусвязный
#include "stack/stack.h" //стек
#include "queue/queue.h" //очередь
#include "chainingHash/chainingHash.h" //хеш
#include "tree/tree.h" //дерево

using namespace std;

// --- МАССИВ ---
TEST(MassiveTest, FullTest) {
    Massive m(2); //создаем массив
    m.pushBack("A"); //добавляем
    m.pushBack("B"); //добавляем
    m.pushBack("C"); //расширение
    EXPECT_EQ(m.get(0), "A"); //проверка
    m.insert(1, "Ins"); //вставка
    EXPECT_EQ(m.get(1), "Ins"); //проверка
    m.set(0, "Z"); //замена
    m.remove(1); //удаление
    // Граничные случаи
    m.get(-1); //ошибка индекса
    m.get(100); //выход за границы
    m.set(100, "X"); //сет за границей
    m.remove(100); //удаление за границей
    m.insert(100, "X"); //вставка за границей
    m.print(); //печать
}

TEST(MassiveTest, Serialization) {
    Massive m1(5); //создаем
    m1.pushBack("SaveMe"); //данные
    m1.serialize("test_arr.bin"); //сохраняем
    Massive m2(5); //второй
    m2.deserialize("test_arr.bin"); //загружаем
    EXPECT_EQ(m2.get(0), "SaveMe"); //проверка
}

// --- ОДНОСВЯЗНЫЙ СПИСОК (ПОЛНЫЙ ТЕСТ) ---
TEST(SListTest, FullTest) {
    SinglyLinkedList l; //создаем
    
    // 1. Тест удаления хвоста с циклом
    l.pushTail("A"); //добавляем
    l.pushTail("B"); //добавляем
    l.pushTail("C"); //список A->B->C
    l.deleteTail();  //удаляем хвост (цикл)
    EXPECT_FALSE(l.find("C")); //проверка
    EXPECT_TRUE(l.find("B")); //проверка

    // 2. Тест удаления головы по значению
    l.deleteByValue("A"); //удаляем голову
    EXPECT_FALSE(l.find("A")); //проверка
    EXPECT_TRUE(l.find("B")); //B - новая голова

    // 3. Тест удаления из середины
    l.pushTail("D"); //добавляем
    l.pushTail("E"); //список B->D->E
    l.deleteByValue("D"); //удаляем середину
    EXPECT_FALSE(l.find("D")); //проверка
    EXPECT_TRUE(l.find("B")); //проверка
    EXPECT_TRUE(l.find("E")); //проверка

    // 4. Тест удаления хвоста по значению
    l.deleteByValue("E"); //удаляем хвост
    EXPECT_FALSE(l.find("E")); //проверка

    // 5. Несуществующий элемент
    l.deleteByValue("Z"); //проход цикла без удаления

    // 6. Вставка и удаление (добить покрытие)
    l.pushHead("1"); //в голову
    l.pushTail("3"); //в хвост
    l.insertAfter("1", "2"); //вставка после
    l.insertBefore("3", "2.5"); //вставка до
    
    l.deleteAfter("1"); //удаление после
    l.deleteBefore("3"); //удаление до

    // Очистка
    l.deleteHead(); //уд голову
    l.deleteTail(); //уд хвост
    
    // Пустой список
    SinglyLinkedList lEmpty; //пустой
    lEmpty.deleteHead(); //из пустого
    lEmpty.deleteTail(); //из пустого
    lEmpty.deleteByValue("A"); //из пустого
    
    l.print(); //печать
}

TEST(SListTest, Serialization) {
    SinglyLinkedList l; //создаем
    l.pushHead("Data"); //данные
    l.serialize("test_slist.bin"); //сохраняем
    SinglyLinkedList l2; //новый
    l2.deserialize("test_slist.bin"); //загружаем
    EXPECT_TRUE(l2.find("Data")); //проверка
}

// --- ДВУСВЯЗНЫЙ СПИСОК ---
TEST(DListTest, FullTest) {
    DoublyLinkedList d; //создаем
    d.pushHead("1"); //голова
    d.pushTail("3"); //хвост
    d.insertAfter("1", "2"); //после
    d.insertBefore("3", "2.5"); //до
    d.deleteAfter("1"); //уд после
    d.deleteBefore("3"); //уд до
    d.deleteHead(); //уд голову
    d.deleteTail(); //уд хвост
    d.deleteByValue("3"); //уд значение
    d.insertAfter("99", "X"); //ошибка
    d.insertBefore("99", "X"); //ошибка
    d.deleteAfter("99"); //ошибка
    d.deleteBefore("99"); //ошибка
    d.print(); //печать
    d.printReverse(); //печать с конца
}

TEST(DListTest, Serialization) {
    DoublyLinkedList d; //создаем
    d.pushHead("Data"); //данные
    d.serialize("test_dlist.bin"); //сохраняем
    DoublyLinkedList d2; //новый
    d2.deserialize("test_dlist.bin"); //загружаем
    EXPECT_TRUE(d2.find("Data")); //проверка
}

// --- СТЕК ---
TEST(StackTest, FullTest) {
    Stack s; //создаем
    s.push("1"); //пуш
    EXPECT_EQ(s.pop(), "1"); //поп
    EXPECT_EQ(s.pop(), ""); //пусто
    s.print(); //печать
}

TEST(StackTest, Serialization) {
    Stack s; //создаем
    s.push("Y"); //данные
    s.serialize("test_stack.bin"); //сохраняем
    Stack s2; //новый
    s2.deserialize("test_stack.bin"); //загружаем
    EXPECT_EQ(s2.pop(), "Y"); //проверка
}

// --- ОЧЕРЕДЬ ---
TEST(QueueTest, FullTest) {
    Queue q; //создаем
    q.push("A"); //пуш
    EXPECT_EQ(q.pop(), "A"); //поп
    EXPECT_EQ(q.pop(), ""); //пусто
    q.print(); //печать
}

TEST(QueueTest, Serialization) {
    Queue q; //создаем
    q.push("X"); //данные
    q.serialize("test_queue.bin"); //сохраняем
    Queue q2; //новая
    q2.deserialize("test_queue.bin"); //загружаем
    EXPECT_EQ(q2.pop(), "X"); //проверка
}

// --- ХЕШ-ТАБЛИЦА ---
TEST(HashTest, FullTest) {
    ChainingHash h(3); //создаем
    h.put("k1", "v1"); //кладем
    h.put("k2", "v2"); //кладем
    h.put("k1", "v1_new"); //обновляем
    EXPECT_EQ(h.get("k1"), "v1_new"); //проверка
    h.remove("k1"); //удаляем
    h.remove("miss"); //удаляем несущ
    h.print(); //печать
}

TEST(HashTest, Serialization) {
    ChainingHash h(10); 
    h.put("key1", "val1"); 
    h.put("key2", "val2");

    // 1. Бинарная
    h.serialize("test_hash.bin"); 
    ChainingHash h2(10); 
    h2.deserialize("test_hash.bin"); 
    EXPECT_EQ(h2.get("key1"), "val1"); 

    // 2. Текстовая (Файловая)
    h.saveToText("test_hash.txt"); //сохраняем в txt
    ChainingHash h3(10);
    h3.loadFromText("test_hash.txt"); //загружаем из txt
    EXPECT_EQ(h3.get("key1"), "val1"); //проверка
    EXPECT_EQ(h3.get("key2"), "val2"); //проверка
}
// --- ДЕРЕВО ---
TEST(TreeTest, FullTest) {
    RBTree t; //создаем
    t.insert(10); t.insert(20); t.insert(5); //вставка
    EXPECT_TRUE(t.search(10)); //поиск
    t.remove(20); t.remove(10); t.remove(99); //удаление
    t.print(); //печать
}

TEST(TreeTest, StressTest) {
    RBTree t; //создаем
    vector<int> keys; //ключи
    for(int i=0; i<500; ++i) { //цикл
        int val = rand() % 10000; //рандом
        t.insert(val); //вставка
        keys.push_back(val); //запомнить
    }
    for(int val : keys) t.remove(val); //удаление всего
}

TEST(TreeTest, Serialization) {
    RBTree t; //создаем
    t.insert(100); //данные
    t.serialize("test_tree.bin"); //сохраняем
    RBTree t2; //новое
    t2.deserialize("test_tree.bin"); //загружаем
    EXPECT_TRUE(t2.search(100)); //проверка
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv); //инит тестов
    return RUN_ALL_TESTS(); //запуск
}
