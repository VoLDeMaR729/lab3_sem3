package structs

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	RED   = true
	BLACK = false
)

// RBNode - узел дерева
type RBNode struct {
	Key    int     `json:"key"`//ключ
	Color  bool    `json:"color"`// true=RED, false=BLACK
	Left   *RBNode `json:"left"`//левый сын
	Right  *RBNode `json:"right"`//правый сын
	Parent *RBNode `json:"-"`// Исключаем из JSON, чтобы не было циклов
}

// RBTree - обертка дерева
type RBTree struct {
	Root  *RBNode `json:"root"`//корень
	TNULL *RBNode `json:"-"`// Служебный узел (заглушка)
}

// NewRBTree - конструктор
func NewRBTree() *RBTree {
	tnull := &RBNode{Color: BLACK} //заглушка всегда черная
	// Важно: TNULL должен быть отдельным объектом
	return &RBTree{
		TNULL: tnull,
		Root:  tnull,
	}
}

//Вспомогательные функции (Приватные)

// Левый поворот
func (t *RBTree) leftRotate(x *RBNode) {
	y := x.Right //правый сын
	x.Right = y.Left //перекидываем левого сына Y к X
	if y.Left != t.TNULL {
		y.Left.Parent = x //обновляем родителя
	}
	y.Parent = x.Parent //связываем Y с родителем X
	if x.Parent == nil {
		t.Root = y //если корень
	} else if x == x.Parent.Left {
		x.Parent.Left = y //слева
	} else {
		x.Parent.Right = y //справа
	}
	y.Left = x //X становится левым сыном Y
	x.Parent = y //обновляем родителя X
}

// Правый поворот
func (t *RBTree) rightRotate(x *RBNode) {
	y := x.Left //левый сын
	x.Left = y.Right //перекидываем
	if y.Right != t.TNULL {
		y.Right.Parent = x //обновляем родителя
	}
	y.Parent = x.Parent //связываем с дедом
	if x.Parent == nil {
		t.Root = y //корень
	} else if x == x.Parent.Right {
		x.Parent.Right = y //справа
	} else {
		x.Parent.Left = y //слева
	}
	y.Right = x //X становится правым сыном Y
	x.Parent = y
}

// Балансировка вставки
func (t *RBTree) insertFix(k *RBNode) {
	var u *RBNode //дядя
	for k.Parent.Color == RED { //пока папа красный
		if k.Parent == k.Parent.Parent.Right { //папа справа
			u = k.Parent.Parent.Left //дядя слева
			if u.Color == RED { //случай 1: дядя красный
				u.Color = BLACK
				k.Parent.Color = BLACK
				k.Parent.Parent.Color = RED
				k = k.Parent.Parent //идем вверх
			} else {
				if k == k.Parent.Left { //случай 2: зигзаг
					k = k.Parent
					t.rightRotate(k)
				}
				//случай 3: линия
				k.Parent.Color = BLACK
				k.Parent.Parent.Color = RED
				t.leftRotate(k.Parent.Parent)
			}
		} else { //зеркальная ситуация
			u = k.Parent.Parent.Right //дядя справа
			if u.Color == RED { //случай 1
				u.Color = BLACK
				k.Parent.Color = BLACK
				k.Parent.Parent.Color = RED
				k = k.Parent.Parent
			} else {
				if k == k.Parent.Right { //случай 2
					k = k.Parent
					t.leftRotate(k)
				}
				//случай 3
				k.Parent.Color = BLACK
				k.Parent.Parent.Color = RED
				t.rightRotate(k.Parent.Parent)
			}
		}
		if k == t.Root { break } //дошли до корня
	}
	t.Root.Color = BLACK //корень всегда черный
}

// Перенос поддерева
func (t *RBTree) transplant(u, v *RBNode) {
	if u.Parent == nil {
		t.Root = v //корень
	} else if u == u.Parent.Left {
		u.Parent.Left = v //слева
	} else {
		u.Parent.Right = v //справа
	}
	v.Parent = u.Parent //обновляем родителя
}

// Минимум в поддереве
func (t *RBTree) minimum(node *RBNode) *RBNode {
	for node.Left != t.TNULL {
		node = node.Left //идем влево до упора
	}
	return node
}

// Балансировка удаления
func (t *RBTree) deleteFix(x *RBNode) {
	var s *RBNode //брат
	for x != t.Root && x.Color == BLACK { //пока x черный и не корень
		if x == x.Parent.Left { //мы слева
			s = x.Parent.Right //брат справа
			if s.Color == RED { //случай 1
				s.Color = BLACK
				x.Parent.Color = RED
				t.leftRotate(x.Parent)
				s = x.Parent.Right
			}
			if s.Left.Color == BLACK && s.Right.Color == BLACK { //случай 2
				s.Color = RED
				x = x.Parent
			} else {
				if s.Right.Color == BLACK { //случай 3
					s.Left.Color = BLACK
					s.Color = RED
					t.rightRotate(s)
					s = x.Parent.Right
				}
				s.Color = x.Parent.Color //случай 4
				x.Parent.Color = BLACK
				s.Right.Color = BLACK
				t.leftRotate(x.Parent)
				x = t.Root
			}
		} else { //зеркально
			s = x.Parent.Left
			if s.Color == RED {
				s.Color = BLACK
				x.Parent.Color = RED
				t.rightRotate(x.Parent)
				s = x.Parent.Left
			}
			if s.Right.Color == BLACK && s.Left.Color == BLACK {
				s.Color = RED
				x = x.Parent
			} else {
				if s.Left.Color == BLACK {
					s.Right.Color = BLACK
					s.Color = RED
					t.leftRotate(s)
					s = x.Parent.Left
				}
				s.Color = x.Parent.Color
				x.Parent.Color = BLACK
				s.Left.Color = BLACK
				t.rightRotate(x.Parent)
				x = t.Root
			}
		}
	}
	x.Color = BLACK //x стал черным
}

// Вспомогательный поиск
func (t *RBTree) searchHelper(node *RBNode, key int) *RBNode {
	if node == t.TNULL || key == node.Key {
		return node //нашли или тупик
	}
	if key < node.Key {
		return t.searchHelper(node.Left, key) //влево
	}
	return t.searchHelper(node.Right, key) //вправо
}

//Публичные методы

// Insert - вставка ключа
func (t *RBTree) Insert(key int) {
	node := &RBNode{Key: key, Color: RED, Left: t.TNULL, Right: t.TNULL, Parent: nil} //новый красный
	y := (*RBNode)(nil)
	x := t.Root

	for x != t.TNULL { //спуск вниз
		y = x
		if node.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	node.Parent = y //ставим родителя
	if y == nil {
		t.Root = node //корень
	} else if node.Key < y.Key {
		y.Left = node //левый сын
	} else {
		y.Right = node //правый сын
	}

	if node.Parent == nil {
		node.Color = BLACK //корень черный
		return
	}
	if node.Parent.Parent == nil {
		return
	}

	t.insertFix(node) //балансировка
}

//удаление ключа
func (t *RBTree) Delete(key int) {
	z := t.searchHelper(t.Root, key) //ищем узел
	if z == t.TNULL {
		fmt.Println("Key not found")
		return
	}

	var x, y *RBNode
	y = z
	yOrigColor := y.Color

	if z.Left == t.TNULL {
		x = z.Right
		t.transplant(z, z.Right) //замена правым
	} else if z.Right == t.TNULL {
		x = z.Left
		t.transplant(z, z.Left) //замена левым
	} else {
		y = t.minimum(z.Right) //ищем преемника
		yOrigColor = y.Color
		x = y.Right
		if y.Parent == z {
			x.Parent = y
		} else {
			t.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		t.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}
	if yOrigColor == BLACK {
		t.deleteFix(x) //балансировка если удалили черный
	}
}

// Search - поиск ключа
func (t *RBTree) Search(k int) bool {
	return t.searchHelper(t.Root, k) != t.TNULL
}

// Вспомогательная печать
func (t *RBTree) printHelper(root *RBNode, indent string, last bool) {
	if root != t.TNULL {
		fmt.Print(indent)
		if last {
			fmt.Print("R----") //правый
			indent += "   "
		} else {
			fmt.Print("L----") //левый
			indent += "|  "
		}
		color := "BLK"
		if root.Color == RED {
			color = "RED"
		}
		fmt.Printf("%d(%s)\n", root.Key, color)
		t.printHelper(root.Left, indent, false) //рек лево
		t.printHelper(root.Right, indent, true) //рек право
	}
}

// Print - красивый вывод дерева
func (t *RBTree) Print() {
	if t.Root == t.TNULL {
		fmt.Println("Tree is empty")
	} else {
		t.printHelper(t.Root, "", true)
	}
}

//Сериализация

// Save - сохранение в JSON
func (t *RBTree) Save(filename string) error {
	file, err := os.Create(filename) //файл
	if err != nil { return err }
	defer file.Close()
	return json.NewEncoder(file).Encode(t) //пишем
}

// Load - загрузка из JSON
func (t *RBTree) Load(filename string) error {
	file, err := os.Open(filename) //открываем
	if err != nil { return err }
	defer file.Close()

	// Сначала загружаем "сырую" структуру
	if err := json.NewDecoder(file).Decode(t); err != nil {
		return err
	}
//После загрузки из JSON поле Parent потеряно (мы его не сохраняли, чтобы не было циклов),и TNULL не настроен.
	t.TNULL = &RBNode{Color: BLACK}//Новый TNULL
	t.fixParentsAndTNULL(t.Root)//Рекурсивно чиним
	return nil
}

//восстанавливает связи Parent и заменяет nil на TNULL
func (t *RBTree) fixParentsAndTNULL(node *RBNode) {
	if node == nil {
		return
	}
	// Если узел был сохранен как лист (nil в json), делаем его TNULL
	if node.Left == nil {
		node.Left = t.TNULL
	} else {
		node.Left.Parent = node
		t.fixParentsAndTNULL(node.Left)
	}

	if node.Right == nil {
		node.Right = t.TNULL
	} else {
		node.Right.Parent = node
		t.fixParentsAndTNULL(node.Right)
	}
}
