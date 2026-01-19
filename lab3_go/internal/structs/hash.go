package structs

import (
 "encoding/gob"
 "encoding/json"
 "fmt"
 "os"
)

type HashNode struct {
 Key  string
 Val  string
 Next *HashNode
}

type ChainingHash struct {
 Buckets  []*HashNode
 Capacity int
 Size     int
}

func NewHashTable(cap int) *ChainingHash {
 return &ChainingHash{
  Buckets:  make([]*HashNode, cap),
  Capacity: cap,
  Size:     0,
 }
}

func (h *ChainingHash) hashFunc(key string) int {
 hash := 0
 for _, c := range key {
  hash += int(c)
 }
 return hash % h.Capacity
}

func (h *ChainingHash) Put(key string, val string) {
 idx := h.hashFunc(key)
 curr := h.Buckets[idx]
 for curr != nil {
  if curr.Key == key {
   curr.Val = val
   return
  }
  curr = curr.Next
 }
 newNode := &HashNode{Key: key, Val: val, Next: h.Buckets[idx]}
 h.Buckets[idx] = newNode
 h.Size++
}

func (h *ChainingHash) Get(key string) string {
 idx := h.hashFunc(key)
 curr := h.Buckets[idx]
 for curr != nil {
  if curr.Key == key {
   return curr.Val
  }
  curr = curr.Next
 }
 return ""
}

// ДОБАВИЛ МЕТОД REMOVE
func (h *ChainingHash) Remove(key string) {
 idx := h.hashFunc(key)
 curr := h.Buckets[idx]
 var prev *HashNode

 for curr != nil {
  if curr.Key == key {
   if prev == nil {
    h.Buckets[idx] = curr.Next
   } else {
    prev.Next = curr.Next
   }
   h.Size--
   return
  }
  prev = curr
  curr = curr.Next
 }
}

func (h *ChainingHash) Print() {
 for i, bucket := range h.Buckets {
  if bucket != nil {
   fmt.Printf("[%d]: ", i)
   curr := bucket
   for curr != nil {
    fmt.Printf("{%s:%s} -> ", curr.Key, curr.Val)
    curr = curr.Next
   }
   fmt.Println("NULL")
  }
 }
}

func (h *ChainingHash) Save(filename string) error {
 file, err := os.Create(filename)
 if err != nil {
  return err
 }
 defer file.Close()
 dump := make(map[string]string)
 for _, node := range h.Buckets {
  curr := node
  for curr != nil {
   dump[curr.Key] = curr.Val
   curr = curr.Next
  }
 }
 encoder := json.NewEncoder(file)
 encoder.SetIndent("", "  ")
 return encoder.Encode(dump)
}

func (h *ChainingHash) Load(filename string) error {
 file, err := os.Open(filename)
 if err != nil {
  return err
 }
 defer file.Close()
 var dump map[string]string
 decoder := json.NewDecoder(file)
 if err := decoder.Decode(&dump); err != nil {
  return err
 }
 h.Buckets = make([]*HashNode, h.Capacity)
 h.Size = 0
 for k, v := range dump {
  h.Put(k, v)
 }
 return nil
}

func (h *ChainingHash) SaveBinary(filename string) error {
 file, err := os.Create(filename)
 if err != nil {
  return err
 }
 defer file.Close()
 dump := make(map[string]string)
 for _, node := range h.Buckets {
  curr := node
  for curr != nil {
   dump[curr.Key] = curr.Val
   curr = curr.Next
  }
 }
 encoder := gob.NewEncoder(file)
 return encoder.Encode(dump)
}

func (h *ChainingHash) LoadBinary(filename string) error {
 file, err := os.Open(filename)
 if err != nil {
  return err
 }
 defer file.Close()
 var dump map[string]string
 decoder := gob.NewDecoder(file)
 if err := decoder.Decode(&dump); err != nil {
  return err
 }
 h.Buckets = make([]*HashNode, h.Capacity)
 h.Size = 0
 for k, v := range dump {
  h.Put(k, v)
 }
 return nil
}
