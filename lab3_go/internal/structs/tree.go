package structs

import (
 "encoding/gob"
 "encoding/json"
 "fmt"
 "os"
)

type TreeNode struct {
 Key   int
 Left  *TreeNode
 Right *TreeNode
}

type RBTree struct {
 Root *TreeNode
}

func NewRBTree() *RBTree {
 return &RBTree{}
}

func (t *RBTree) Insert(key int) {
 if t.Root == nil {
  t.Root = &TreeNode{Key: key}
  return
 }
 curr := t.Root
 for {
  if key < curr.Key {
   if curr.Left == nil {
    curr.Left = &TreeNode{Key: key}
    return
   }
   curr = curr.Left
  } else {
   if curr.Right == nil {
    curr.Right = &TreeNode{Key: key}
    return
   }
   curr = curr.Right
  }
 }
}

func (t *RBTree) Search(key int) bool {
 curr := t.Root
 for curr != nil {
  if curr.Key == key {
   return true
  }
  if key < curr.Key {
   curr = curr.Left
  } else {
   curr = curr.Right
  }
 }
 return false
}

// ДОБАВИЛ DELETE ЧТОБЫ ТЕСТЫ РАБОТАЛИ
func (t *RBTree) Delete(key int) {
 var remove func(n *TreeNode, key int) *TreeNode
 remove = func(n *TreeNode, key int) *TreeNode {
  if n == nil {
   return nil
  }
  if key < n.Key {
   n.Left = remove(n.Left, key)
  } else if key > n.Key {
   n.Right = remove(n.Right, key)
  } else {
   if n.Left == nil {
    return n.Right
   }
   if n.Right == nil {
    return n.Left
   }
   minRight := n.Right
   for minRight.Left != nil {
    minRight = minRight.Left
   }
   n.Key = minRight.Key
   n.Right = remove(n.Right, minRight.Key)
  }
  return n
 }
 t.Root = remove(t.Root, key)
}

func (t *RBTree) Print() {
 var walk func(*TreeNode)
 walk = func(n *TreeNode) {
  if n == nil {
   return
  }
  walk(n.Left)
  fmt.Printf("%d ", n.Key)
  walk(n.Right)
 }
 fmt.Print("Tree: ")
 walk(t.Root)
 fmt.Println()
}

func (t *RBTree) Save(filename string) error {
 file, err := os.Create(filename)
 if err != nil {
  return err
 }
 defer file.Close()
 var keys []int
 var collect func(n *TreeNode)
 collect = func(n *TreeNode) {
  if n == nil {
   return
  }
  collect(n.Left)
  keys = append(keys, n.Key)
  collect(n.Right)
 }
 collect(t.Root)
 encoder := json.NewEncoder(file)
 encoder.SetIndent("", "  ")
 return encoder.Encode(keys)
}

func (t *RBTree) Load(filename string) error {
 file, err := os.Open(filename)
 if err != nil {
  return err
 }
 defer file.Close()
 var keys []int
 decoder := json.NewDecoder(file)
 if err := decoder.Decode(&keys); err != nil {
  return err
 }
 t.Root = nil
 for _, k := range keys {
  t.Insert(k)
 }
 return nil
}

func (t *RBTree) SaveBinary(filename string) error {
 file, err := os.Create(filename)
 if err != nil {
  return err
 }
 defer file.Close()
 var keys []int
 var collect func(n *TreeNode)
 collect = func(n *TreeNode) {
  if n == nil {
   return
  }
  collect(n.Left)
  keys = append(keys, n.Key)
  collect(n.Right)
 }
 collect(t.Root)
 encoder := gob.NewEncoder(file)
 return encoder.Encode(keys)
}

func (t *RBTree) LoadBinary(filename string) error {
 file, err := os.Open(filename)
 if err != nil {
  return err
 }
 defer file.Close()
 var keys []int
 decoder := gob.NewDecoder(file)
 if err := decoder.Decode(&keys); err != nil {
  return err
 }
 t.Root = nil
 for _, k := range keys {
  t.Insert(k)
 }
 return nil
}
