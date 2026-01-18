package structs

import (
	"os"
	"testing"
)

// Тест хеш-таблицы
func TestHash_Ops(t *testing.T) {
	h := NewHashTable(5) //маленький размер для коллизий
	h.Put("k1", "v1")//вставка
	h.Put("k2", "v2")
	
	// Обновление существующего
	h.Put("k1", "v1_new") 
	if h.Get("k1") != "v1_new" { t.Error("Update fail") } //проверка обновления

// Коллизия, Нужно подобрать ключи, дающие одинаковый хеш, или просто много добавить
	h.Put("k6", "v6") 
	
	if h.Get("k2") != "v2" { t.Error("Get fail") } //проверка
	if h.Get("miss") != "" { t.Error("Miss fail") } //проверка отсутствующего

	h.Remove("k1") //удаление
	if h.Get("k1") != "" { t.Error("Remove fail") } //проверка удаления
	
	h.Remove("miss") //удаление несуществующего (не должно падать)
	h.Print() //покрытие
}

// Тест файлов
func TestHash_File(t *testing.T) {
	f := "test_h.json" //файл
	h := NewHashTable(5) //создаем
	h.Put("key", "val") //пишем
	h.Save(f) //сохраняем

	h2 := NewHashTable(5) //новый
	h2.Load(f) //загружаем
	if h2.Get("key") != "val" { t.Error("Load fail") } //проверка
	os.Remove(f) //удаляем
	h2.Load("err.json") //ошибка
}

func BenchmarkHash_Put(b *testing.B) {
	h := NewHashTable(100) //инит
	for i := 0; i < b.N; i++ { //цикл
		h.Put("key", "val") //действие
	}
}
