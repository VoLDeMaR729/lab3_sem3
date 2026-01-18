package main

import (
	"bufio"
	"fmt"
	"lab3_go/internal/structs"
	"os"
	"strconv"
	"strings"
)

func showHelp() {
	fmt.Println("=== СПИСОК КОМАНД ===")
	fmt.Println("МАССИВ:    M_ADD <val>, M_INS <idx> <val>, M_DEL <idx>, M_GET <idx>, M_SET <idx> <val>, M_PRINT")
	fmt.Println("           M_SAVE <file>, M_LOAD <file>")
	fmt.Println("С.СПИСОК:  S_ADDH <val>, S_ADDT <val>, S_DELH, S_DELT, S_DELV <val>, S_FIND <val>, S_PRINT")
	fmt.Println("           S_SAVE <file>, S_LOAD <file>")
	fmt.Println("Д.СПИСОК:  D_ADDH <val>, D_ADDT <val>, D_DELV <val>, D_FIND <val>, D_PRINT")
	fmt.Println("           D_SAVE <file>, D_LOAD <file>")
	fmt.Println("ОЧЕРЕДЬ:   Q_PUSH <val>, Q_POP, Q_PRINT, Q_SAVE <file>, Q_LOAD <file>")
	fmt.Println("СТЕК:      ST_PUSH <val>, ST_POP, ST_PRINT, ST_SAVE <file>, ST_LOAD <file>")
	fmt.Println("ХЕШ:       H_PUT <key> <val>, H_GET <key>, H_DEL <key>, H_PRINT, H_SAVE <file>, H_LOAD <file>")
	fmt.Println("ДЕРЕВО:    T_INS <key>, T_FIND <key>, T_PRINT, T_SAVE <file>, T_LOAD <file>")
	fmt.Println("ОБЩИЕ:     EXIT")
}

func main() {
	// Инициализация структур
	m := structs.NewDynArray() // Массив
	s := structs.NewSList()// Односвязный список
	d := structs.NewDList()// Двусвязный список
	q := structs.NewQueue()// Очередь
	st := structs.NewStack()// Стек
	h := structs.NewHashTable(10)//Хеш-таблица
	t := structs.NewRBTree()//Дерево

	scanner := bufio.NewScanner(os.Stdin) //сканер ввода
	fmt.Println("Go Lab Ready. Введите HELP для списка команд.")

	for {
		fmt.Print("\n> ") //приглашение
		if !scanner.Scan() {
			break 
		}
		line := scanner.Text() //читаем строку
		parts := strings.Fields(line) //делим на слова
		if len(parts) == 0 {
			continue
		}

		cmd := strings.ToUpper(parts[0]) //команда капсом

		switch cmd {
		case "EXIT":
			return
		case "HELP":
			showHelp()
//МАССИВ
		case "M_ADD":
			if len(parts) > 1 {
				m.Add(parts[1])
				fmt.Println("Добавлено")
			}
		case "M_INS":
			if len(parts) > 2 {
				idx, _ := strconv.Atoi(parts[1])
				m.Insert(idx, parts[2])
				fmt.Println("Вставлено")
			}
		case "M_DEL":
			if len(parts) > 1 {
				idx, _ := strconv.Atoi(parts[1])
				m.Remove(idx)
				fmt.Println("Удалено")
			}
		case "M_GET":
			if len(parts) > 1 {
				idx, _ := strconv.Atoi(parts[1])
				fmt.Println("Значение:", m.Get(idx))
			}
		case "M_SET":
			if len(parts) > 2 {
				idx, _ := strconv.Atoi(parts[1])
				m.Set(idx, parts[2])
				fmt.Println("Изменено")
			}
		case "M_PRINT":
			m.Print()
		case "M_SAVE":
			if len(parts) > 1 {
				if err := m.Save(parts[1]); err == nil {
					fmt.Println("Сохранено")
				} else {
					fmt.Println("Ошибка:", err)
				}
			}
		case "M_LOAD":
			if len(parts) > 1 {
				if err := m.Load(parts[1]); err == nil {
					fmt.Println("Загружено")
				} else {
					fmt.Println("Ошибка:", err)
				}
			}

//ОДНОСВЯЗНЫЙ СПИСОК
		case "S_ADDH":
			if len(parts) > 1 {
				s.AddHead(parts[1])
				fmt.Println("Добавлено в голову")
			}
		case "S_ADDT":
			if len(parts) > 1 {
				s.AddTail(parts[1])
				fmt.Println("Добавлено в хвост")
			}
		case "S_DELH":
			s.DelHead()
			fmt.Println("Голова удалена")
		case "S_DELT":
			s.DelTail()
			fmt.Println("Хвост удален")
		case "S_DELV":
			if len(parts) > 1 {
				s.DelVal(parts[1])
				fmt.Println("Удалено по значению")
			}
		case "S_FIND":
			if len(parts) > 1 {
				if s.Find(parts[1]) {
					fmt.Println("Найдено")
				} else {
					fmt.Println("Не найдено")
				}
			}
		case "S_PRINT":
			s.Print()
		case "S_SAVE":
			if len(parts) > 1 {
				s.Save(parts[1])
				fmt.Println("Сохранено")
			}
		case "S_LOAD":
			if len(parts) > 1 {
				s.Load(parts[1])
				fmt.Println("Загружено")
			}

//ДВУСВЯЗНЫЙ СПИСОК
		case "D_ADDH":
			if len(parts) > 1 {
				d.AddHead(parts[1])
				fmt.Println("Добавлено в голову")
			}
		case "D_ADDT":
			if len(parts) > 1 {
				d.AddTail(parts[1])
				fmt.Println("Добавлено в хвост")
			}
		case "D_DELV":
			if len(parts) > 1 {
				d.DelVal(parts[1])
				fmt.Println("Удалено")
			}
		case "D_FIND":
			if len(parts) > 1 {
				if d.Find(parts[1]) {
					fmt.Println("Найдено")
				} else {
					fmt.Println("Не найдено")
				}
			}
		case "D_PRINT":
			d.Print()
		case "D_SAVE":
			if len(parts) > 1 {
				d.Save(parts[1])
				fmt.Println("Сохранено")
			}
		case "D_LOAD":
			if len(parts) > 1 {
				d.Load(parts[1])
				fmt.Println("Загружено")
			}

//ОЧЕРЕДЬ
		case "Q_PUSH":
			if len(parts) > 1 {
				q.Push(parts[1])
				fmt.Println("Добавлено")
			}
		case "Q_POP":
			fmt.Println("Извлечено:", q.Pop())
		case "Q_PRINT":
			q.Print()
		case "Q_SAVE":
			if len(parts) > 1 {
				q.Save(parts[1])
				fmt.Println("Сохранено")
			}
		case "Q_LOAD":
			if len(parts) > 1 {
				q.Load(parts[1])
				fmt.Println("Загружено")
			}

//СТЕК
		case "ST_PUSH":
			if len(parts) > 1 {
				st.Push(parts[1])
				fmt.Println("Добавлено")
			}
		case "ST_POP":
			fmt.Println("Извлечено:", st.Pop())
		case "ST_PRINT":
			st.Print()
		case "ST_SAVE":
			if len(parts) > 1 {
				st.Save(parts[1])
				fmt.Println("Сохранено")
			}
		case "ST_LOAD":
			if len(parts) > 1 {
				st.Load(parts[1])
				fmt.Println("Загружено")
			}

//ХЕШ-ТАБЛИЦА
		case "H_PUT":
			if len(parts) > 2 {
				h.Put(parts[1], parts[2])
				fmt.Println("Добавлено")
			}
		case "H_GET":
			if len(parts) > 1 {
				fmt.Println("Значение:", h.Get(parts[1]))
			}
		case "H_DEL":
			if len(parts) > 1 {
				h.Remove(parts[1])
				fmt.Println("Удалено")
			}
		case "H_PRINT":
			h.Print()
		case "H_SAVE":
			if len(parts) > 1 {
				h.Save(parts[1])
				fmt.Println("Сохранено")
			}
		case "H_LOAD":
			if len(parts) > 1 {
				h.Load(parts[1])
				fmt.Println("Загружено")
			}
//ДЕРЕВО
		case "T_INS":
			if len(parts) > 1 {
				k, _ := strconv.Atoi(parts[1])
				t.Insert(k)
				fmt.Println("Вставлено")
			}
		case "T_FIND":
			if len(parts) > 1 {
				k, _ := strconv.Atoi(parts[1])
				if t.Search(k) {
					fmt.Println("YES")
				} else {
					fmt.Println("NO")
				}
			}
		case "T_PRINT":
			t.Print()
		case "T_SAVE":
			if len(parts) > 1 {
				t.Save(parts[1])
				fmt.Println("Сохранено")
			}
		case "T_LOAD":
			if len(parts) > 1 {
				t.Load(parts[1])
				fmt.Println("Загружено")
			}

		default:
			fmt.Println("Неизвестная команда. Введите HELP.")
		}
	}
}
