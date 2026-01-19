#include <iostream>
#include <string>
#include <vector>

//подключаем все наши заголовки
#include "massive/massive.h"
#include "singlyLinkedList/singlyLinkedList.h"
#include "doublyLinkedList/doublyLinkedList.h"
#include "queue/queue.h"
#include "stack/stack.h"
#include "chainingHash/chainingHash.h"
#include "tree/tree.h"

using namespace std;

void showHelp() {
    cout << "=== СПИСОК КОМАНД ===\n";
    cout << "МАССИВ: M_ADD <val>, M_INS <idx> <val>, M_DEL <idx>, M_GET <idx>, M_SET <idx> <val>, M_PRINT\n";
    cout << "С.СПИСОК: S_ADDH <val>, S_ADDT <val>, S_DELH, S_DELT, S_DELV <val>, S_FIND <val>, S_PRINT\n";
    cout << "Д.СПИСОК: D_ADDH <val>, D_ADDT <val>, D_DELH, D_DELT, D_DELV <val>, D_FIND <val>, D_PRINT\n";
    cout << "ОЧЕРЕДЬ: Q_PUSH <val>, Q_POP, Q_PRINT\n";
    cout << "СТЕК: ST_PUSH <val>, ST_POP, ST_PRINT\n";
    cout << "ХЕШ: H_PUT <key> <val>, H_GET <key>, H_DEL <key>, H_PRINT\n";
    cout << "ДЕРЕВО: T_INS <key>, T_DEL <key>, T_FIND <key>, T_PRINT\n";
    cout << "ОБЩИЕ: SAVE <file>, LOAD <file>, EXIT\n";
}

int main() {
    //инициализация всех структур
    Massive m;
    SinglyLinkedList s;
    DoublyLinkedList d;
    Queue q;
    Stack st;
    ChainingHash h;
    RBTree t;

    string cmd;
    string val, val2;
    int idx, key;

    cout << "Лабораторная работа готова. Введите HELP для списка команд.\n";

    while (cin >> cmd) {
        if (cmd == "EXIT") break;
        else if (cmd == "HELP") showHelp();

        // --- МАССИВ ---
        else if (cmd == "M_ADD") { cin >> val; m.pushBack(val); }
        else if (cmd == "M_INS") { cin >> idx >> val; m.insert(idx, val); }
        else if (cmd == "M_DEL") { cin >> idx; m.remove(idx); }
        else if (cmd == "M_SET") { cin >> idx >> val; m.set(idx, val); }
        else if (cmd == "M_GET") { cin >> idx; cout << m.get(idx) << endl; }
        else if (cmd == "M_PRINT") { m.print(); }

        // --- ОДНОСВЯЗНЫЙ ---
        else if (cmd == "S_ADDH") { cin >> val; s.pushHead(val); }
        else if (cmd == "S_ADDT") { cin >> val; s.pushTail(val); }
        else if (cmd == "S_DELH") { s.deleteHead(); }
        else if (cmd == "S_DELT") { s.deleteTail(); }
        else if (cmd == "S_DELV") { cin >> val; s.deleteByValue(val); }
        else if (cmd == "S_FIND") { cin >> val; cout << (s.find(val) ? "YES" : "NO") << endl; }
        else if (cmd == "S_PRINT") { s.print(); }

        // --- ДВУСВЯЗНЫЙ ---
        else if (cmd == "D_ADDH") { cin >> val; d.pushHead(val); }
        else if (cmd == "D_ADDT") { cin >> val; d.pushTail(val); }
        else if (cmd == "D_DELH") { d.deleteHead(); }
        else if (cmd == "D_DELT") { d.deleteTail(); }
        else if (cmd == "D_DELV") { cin >> val; d.deleteByValue(val); }
        else if (cmd == "D_FIND") { cin >> val; cout << (d.find(val) ? "YES" : "NO") << endl; }
        else if (cmd == "D_PRINT") { d.print(); }

        // --- ОЧЕРЕДЬ ---
        else if (cmd == "Q_PUSH") { cin >> val; q.push(val); }
        else if (cmd == "Q_POP") { cout << "Popped: " << q.pop() << endl; }
        else if (cmd == "Q_PRINT") { q.print(); }

        // --- СТЕК ---
        else if (cmd == "ST_PUSH") { cin >> val; st.push(val); }
        else if (cmd == "ST_POP") { cout << "Popped: " << st.pop() << endl; }
        else if (cmd == "ST_PRINT") { st.print(); }

        // --- ХЕШ ---
        else if (cmd == "H_PUT") { cin >> val >> val2; h.put(val, val2); } // key value
        else if (cmd == "H_GET") { cin >> val; cout << h.get(val) << endl; }
        else if (cmd == "H_DEL") { cin >> val; h.remove(val); }
        else if (cmd == "H_PRINT") { h.print(); }

        // --- ДЕРЕВО ---
        else if (cmd == "T_INS") { cin >> key; t.insert(key); }
        else if (cmd == "T_DEL") { cin >> key; t.remove(key); }
        else if (cmd == "T_FIND") { cin >> key; cout << (t.search(key) ? "YES" : "NO") << endl; }
        else if (cmd == "T_PRINT") { t.print(); }

        // --- СОХРАНЕНИЕ/ЗАГРУЗКА (Пример для массива, можно расширить) ---
        else if (cmd == "SAVE") { 
            cin >> val; 
            m.serialize(val + "_arr.bin"); 
            s.serialize(val + "_slist.bin");
            d.serialize(val + "_dlist.bin");
            q.serialize(val + "_queue.bin");
            st.serialize(val + "_stack.bin");
            h.serialize(val + "_hash.bin");
            t.serialize(val + "_tree.bin");
            cout << "Все структуры сохранены с префиксом " << val << endl;
        }
        else if (cmd == "LOAD") { 
            cin >> val; 
            m.deserialize(val + "_arr.bin"); 
            s.deserialize(val + "_slist.bin");
            d.deserialize(val + "_dlist.bin");
            q.deserialize(val + "_queue.bin");
            st.deserialize(val + "_stack.bin");
            h.deserialize(val + "_hash.bin");
            t.deserialize(val + "_tree.bin");
            cout << "Все структуры загружены с префиксом " << val << endl;
        }

        else { cout << "Неизвестная команда\n"; }
    }
    return 0;
}
