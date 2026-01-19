package structs

import "testing"

func TestAllStructures(t *testing.T) {
 // 1. Array
 arr := NewMassive(5)
 arr.PushBack("Test")
 if err := arr.SaveBinary("test_arr.bin"); err != nil { t.Error(err) }
 arr2 := NewMassive(1)
 if err := arr2.LoadBinary("test_arr.bin"); err != nil { t.Error(err) }
 if val, _ := arr2.Get(0); val != "Test" { t.Error("Array Bin fail") }

 // 2. SList
 sl := &SinglyLinkedList{}
 sl.PushTail("SVal")
 if err := sl.SaveBinary("test_sl.bin"); err != nil { t.Error(err) }
 sl2 := &SinglyLinkedList{}
 if err := sl2.LoadBinary("test_sl.bin"); err != nil { t.Error(err) }
 if !sl2.Find("SVal") { t.Error("SList Bin fail") }

 // 3. DList
 dl := NewDList()
 dl.AddTail("DVal")
 if err := dl.SaveBinary("test_dl.bin"); err != nil { t.Error(err) }
 dl2 := NewDList()
 if err := dl2.LoadBinary("test_dl.bin"); err != nil { t.Error(err) }
 if !dl2.Find("DVal") { t.Error("DList Bin fail") }

 // 4. Stack
 st := NewStack()
 st.Push("StackVal")
 if err := st.SaveBinary("test_st.bin"); err != nil { t.Error(err) }
 st2 := NewStack()
 if err := st2.LoadBinary("test_st.bin"); err != nil { t.Error(err) }
 if st2.Pop() != "StackVal" { t.Error("Stack Bin fail") }

 // 5. Queue
 qu := NewQueue()
 qu.Push("QVal")
 if err := qu.SaveBinary("test_qu.bin"); err != nil { t.Error(err) }
 qu2 := NewQueue()
 if err := qu2.LoadBinary("test_qu.bin"); err != nil { t.Error(err) }
 if qu2.Pop() != "QVal" { t.Error("Queue Bin fail") }

 // 6. Hash (Используем NewHashTable)
 hs := NewHashTable(10)
 hs.Put("Key", "HVal")
 if err := hs.SaveBinary("test_hs.bin"); err != nil { t.Error(err) }
 hs2 := NewHashTable(10)
 if err := hs2.LoadBinary("test_hs.bin"); err != nil { t.Error(err) }
 if hs2.Get("Key") != "HVal" { t.Error("Hash Bin fail") }

 // 7. Tree
 tr := NewRBTree()
 tr.Insert(100)
 tr.Delete(100) // Проверка delete
 tr.Insert(50)
 if err := tr.SaveBinary("test_tr.bin"); err != nil { t.Error(err) }
 tr2 := NewRBTree()
 if err := tr2.LoadBinary("test_tr.bin"); err != nil { t.Error(err) }
 if !tr2.Search(50) { t.Error("Tree Bin fail") }
}
