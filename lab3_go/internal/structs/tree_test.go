package structs

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

// Базовый тест операций (ручные сценарии)
func TestTree_Ops(t *testing.T) {
	tr := NewRBTree() //создаем дерево

// 1. Вставка с вызовом поворотов
	keys := []int{10, 20, 30, 15, 25, 5, 1}
	for _, k := range keys {
		tr.Insert(k) //вставка элемента
	}
// 2. Поиск
	if !tr.Search(10) { t.Error("Search 10 fail") } //поиск корня
	if !tr.Search(15) { t.Error("Search 15 fail") } //поиск внутреннего узла
	if tr.Search(999) { t.Error("Search non-exist fail") } //поиск несуществующего
// 3. Удаление
	tr.Delete(1)//удаление листа
	tr.Delete(30)//удаление узла с одним ребенком
	tr.Delete(10) //удаление корня
	
	if tr.Search(10) { t.Error("Delete root fail") } //проверка
	tr.Delete(999) //удаление несуществующего (не должно падать)

	tr.Print() //покрытие функции печати
}

// Stress Test - вставка и удаление большого количества случайных данных
// Это гарантирует покрытие сложных случаев балансировки (все cases в fix-функциях)
func TestTree_RandomCoverage(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) //инициализация рандома
	tr := NewRBTree() //новое дерево
	count := 1000 //количество элементов

// Слайс для запоминания вставленных ключей
	var added []int

	// 1. Вставка случайных чисел
	for i := 0; i < count; i++ {
		k := rand.Intn(100000) //случайное число
		tr.Insert(k) //вставка
		added = append(added, k) //запоминаем
	}

// 2. Проверка, что они все находятся
	for _, k := range added {
		if !tr.Search(k) {
			t.Errorf("Lost key %d after insert", k)
		}
	}
// 3. Перемешиваем порядок удаления (чтобы ломать структуру хаотично)
	rand.Shuffle(len(added), func(i, j int) {
		added[i], added[j] = added[j], added[i]
	})

//4. Удаление всех элементов
	for _, k := range added {
		tr.Delete(k) //удаление
	}

// 5. Дерево должно быть пустым (TNULL). Пробуем найти первый добавленный элемент, его быть не должно
	if len(added) > 0 && tr.Search(added[0]) {
		t.Error("Tree not empty after stress delete")
	}
}

// Тест сохранения и загрузки
func TestTree_File(t *testing.T) {
	f := "test_t.json" //имя файла
	tr := NewRBTree() //дерево
	tr.Insert(100) //данные
	tr.Insert(50)
	tr.Insert(150)

	err := tr.Save(f) //сохраняем
	if err != nil { t.Errorf("Save fail: %v", err) }

	tr2 := NewRBTree() //новое
	err = tr2.Load(f) //загружаем
	if err != nil { t.Errorf("Load fail: %v", err) }

	// Проверка целостности
	if !tr2.Search(100) { t.Error("Root lost") } 
	if !tr2.Search(50) { t.Error("Left lost") }
	
	// Проверка восстановления TNULL
	tr2.Insert(200) //если TNULL сломан, тут будет panic

	os.Remove(f) //удаляем файл
	tr2.Load("bad.json") //ошибка загрузки
}

func BenchmarkTree_Insert(b *testing.B) {
	tr := NewRBTree() //инит
	for i := 0; i < b.N; i++ { //цикл бенчмарка
		tr.Insert(i) //замер вставки
	}
}
