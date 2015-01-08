package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
    "math/rand"
)

type Node struct {
    value int
    size int
    count int
    left *Node
    right *Node
}

func NewNode(value int) *Node {
    return &Node{
        value:value,
        size:1,
    }
}

func (self *Node) countLeft() (c int) {
    if self.left != nil {
        c = self.left.size
    }
    return
}

func (self *Node) countRight() (c int) {
    if self.right != nil {
        c = self.right.size
    }
    return
}

func (self *Node) balance() (b int) {
    b = self.countRight() - self.countLeft()
    return
}

func (self *Node) insert(n *Node) (top *Node){
    self.size++
    top = self
    if n.value < self.value {
        n.count += 1 + self.countRight()
        // Insert Left
        if self.left == nil {
            self.left = n
            return
        }
        self.left = self.left.insert(n)
        if rand.Intn(self.size) == 0 {
            top = self.rotateRight()
        }
    } else {
        // Insert Right
        if self.right == nil {
            self.right = n
            return
        }
        self.right = self.right.insert(n)
        if rand.Intn(self.size) == 0 {
            top = self.rotateLeft()
        }
    }
    return
}

func (self *Node) calcSize() {
    self.size = 1
    if self.left != nil {
        self.size += self.left.size
    }
    if self.right != nil {
        self.size += self.right.size
    }
}

/*
         X                         Y
       /  \                      /  \
      a     Y         ==>       X    c
          /  \                /  \
         b    c              a    b
*/
func (self *Node) rotateLeft() *Node {
    y := self.right
    self.right = y.left
    y.left = self

    y.size = self.size
    self.calcSize()
    return y
}

/*
         X                    Y
       /  \                 /  \
      Y    c     ==>       a    X
     / \                      /  \
    a   b                    b    c
*/
func (self *Node) rotateRight() *Node {
    y := self.left
    self.left = y.right
    y.right = self

    y.size = self.size
    self.calcSize()
    return y
}

func (self *Node) values() (v []int) {
    if self.left != nil {
        v = append(v, self.left.values()...)
    }
    v = append(v, self.value)
    if self.right != nil {
        v = append(v, self.right.values()...)
    }
    return
}

func solve(a []int) (ans int) {
    root := NewNode(a[0])
    for _,v := range a[1:] {
        node := NewNode(v)
        root = root.insert(node)
        ans += node.count
    }
    return
}

func printArray(a []int) {
    fmt.Println(strings.Trim(fmt.Sprint(a),"[]"))
}

func atoi(s string) (i int) {
    i,err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return
}

func strToArray(s string) (a []int) {
    for _,w := range strings.Split(s, " ") {
        a = append(a, atoi(w))
    }
    return
}

func countLessThanX(a []int, x int) (count int){
    for _,v := range a {
        if x < v {
            count++
        }
    }
    return
}

func countShifts(a []int) (count int) {
    for i:=1; i<len(a); i++ {
        count += countLessThanX(a[:i], a[i])
    }
    return
}

func main() {
    scanner := bufio.NewReader(os.Stdin)
    s,_ := scanner.ReadString('\n')
    s = strings.TrimSpace(s)
    cases := atoi(s)

    for c:=cases; c>0; c-- {
        s,_ = scanner.ReadString('\n')
        s = strings.TrimSpace(s)
        size := atoi(s)
        _ = size

        s,_ = scanner.ReadString('\n')
        s = strings.TrimSpace(s)
        a := strToArray(s)

        //n := countShifts(a)
        n := solve(a)
        fmt.Println(n)
    }
}
