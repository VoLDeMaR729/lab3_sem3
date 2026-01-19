#include <benchmark/benchmark.h>
#include <string>
#include <iostream>

//Подключаем структуры
#include "massive/massive.h"
#include "singlyLinkedList/singlyLinkedList.h"
#include "doublyLinkedList/doublyLinkedList.h"
#include "stack/stack.h"
#include "queue/queue.h"
#include "chainingHash/chainingHash.h"
#include "tree/tree.h"

using namespace std;

// Бенчмарк для массива - добавление в конец
static void BM_Massive_PushBack(benchmark::State& state) {
    for (auto _ : state) { //цикл работы бенчмарка
        Massive m(state.range(0) + 10); //создаем массив с запасом
        for (int i = 0; i < state.range(0); i++) { //проход по элементам
            m.pushBack("elem_" + to_string(i)); //добавляем в конец
        }
    }
}
// Регистрируем тест на диапазон от 1000 до 100000 элементов
BENCHMARK(BM_Massive_PushBack)->Range(1000, 100000);

// Бенчмарк для односвязного списка - вставка в хвост
static void BM_SList_PushTail(benchmark::State& state) {
    for (auto _ : state) { //цикл замеров
        SinglyLinkedList list; //создаем список
        for (int i = 0; i < state.range(0); i++) { //цикл вставки
            list.pushTail("elem_" + to_string(i)); //вставка в хвост
        }
    }
}
BENCHMARK(BM_SList_PushTail)->Range(1000, 100000); //диапазон

// Бенчмарк для двусвязного списка - вставка в хвост
static void BM_DList_PushTail(benchmark::State& state) {
    for (auto _ : state) { //цикл замеров
        DoublyLinkedList list; //создаем список
        for (int i = 0; i < state.range(0); i++) { //цикл вставки
            list.pushTail("elem_" + to_string(i)); //вставка в конец
        }
    }
}
BENCHMARK(BM_DList_PushTail)->Range(1000, 100000); //диапазон

// Бенчмарк для стека - добавление элементов
static void BM_Stack_Push(benchmark::State& state) {
    for (auto _ : state) { //цикл замеров
        Stack st; //создаем стек
        for (int i = 0; i < state.range(0); i++) { //цикл пуша
            st.push("elem_" + to_string(i)); //кладем на вершину
        }
    }
}
BENCHMARK(BM_Stack_Push)->Range(1000, 100000); //диапазон

// Бенчмарк для очереди - добавление элементов
static void BM_Queue_Push(benchmark::State& state) {
    for (auto _ : state) { //цикл замеров
        Queue q; //создаем очередь
        for (int i = 0; i < state.range(0); i++) { //цикл добавления
            q.push("elem_" + to_string(i)); //добавляем в очередь
        }
    }
}
BENCHMARK(BM_Queue_Push)->Range(1000, 100000); //диапазон

// Бенчмарк хеш-таблицы - вставка пар ключ-значение
static void BM_Hash_Put(benchmark::State& state) {
    for (auto _ : state) { //цикл замеров
        // Создаем хеш с размером range, чтобы уменьшить коллизии
        ChainingHash h(state.range(0)); 
        for (int i = 0; i < state.range(0); i++) { //цикл вставки
            string key = "key_" + to_string(i); //генерируем ключ
            string val = "val_" + to_string(i); //генерируем значение
            h.put(key, val); //вставка пары
        }
    }
}
BENCHMARK(BM_Hash_Put)->Range(1000, 100000); //диапазон

// Бенчмарк дерева - вставка узлов
static void BM_Tree_Insert(benchmark::State& state) {
    for (auto _ : state) { //цикл замеров
        RBTree t; //создаем дерево
        for (int i = 0; i < state.range(0); i++) { //цикл вставки
            // Вставляем i последовательно
            t.insert(i); //вставка ключа
        }
    }
}
BENCHMARK(BM_Tree_Insert)->Range(1000, 100000); //диапазон

// Заменяет int main(), запускает все зарегистрированные бенчмарки
BENCHMARK_MAIN();
