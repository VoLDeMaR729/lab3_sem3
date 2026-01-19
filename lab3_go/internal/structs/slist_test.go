package structs

import "testing"

func TestSList_Green(t *testing.T) {
 l := &SinglyLinkedList{}
 
 //заполняем чтобы циклы работали
 l.PushTail("A")
 l.PushTail("B")
 l.PushTail("C") //A->B->C
 
 //уд хвоста (цикл while станет зеленым)
 l.DeleteTail() 
 
 //уд значения (середина)
 l.PushTail("D")
 l.DeleteByValue("B") 
 
 //тест insert
 l.InsertAfter("A", "New")
 l.InsertBefore("D", "PreD")
 l.InsertBefore("A", "Head") //смена головы
 
 //тест delete after/before
 l.DeleteAfter("Head")
 l.DeleteBefore("D")
 l.DeleteBefore("Head") //пустой if
 
 //тест find
 l.Find("D")
 l.Find("Z") //не найдено (проход всего цикла)

 //граничные
 l2 := &SinglyLinkedList{}
 l2.DeleteHead()
 l2.DeleteTail()
 l2.DeleteByValue("X")
 l2.InsertAfter("X", "Y")
 l2.InsertBefore("X", "Y")
 l2.DeleteBefore("X")
 
 //печать
 l.Print()
}

func TestSList_File(t *testing.T) {
 l := &SinglyLinkedList{}
 l.PushHead("Data")
 l.Save("test_list.json")
 
 l2 := &SinglyLinkedList{}
 l2.Load("test_list.json")
 
 l.Load("bad.json")
}
