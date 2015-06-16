package main

import "fmt"

type Walker interface {
    walk(nb int) bool
    present()
}

type Human struct {
    name string
    m int
}

type Robot struct {
    model string
    m int
}
type Humanoid struct {
    Human
    age int
}
func (h * Humanoid) introduction() {
    h.present()
    fmt.Println("AND is also ",h.age," years old !!")
}
func (h *Human) walk(nb int) bool {
    fmt.Printf("%s (%p) has walked %d meters !\n",h.name,h,nb)
    h.m += nb
    return true
}
func (h *Human) present() {
    fmt.Printf("Human %s is not at %d meters\n",h.name,h.m)
}
func (r * Robot) present() {
    fmt.Printf("Robot %s is now at %d meters\n",r.model,r.m)
}
func (r *Robot) walk(nb int) bool {
    fmt.Printf("%s (%p) has walked %d meters !\n",r.model,r,nb)
    r.m += nb
    return true
}

func move(w Walker) {
    if w.walk(10) {
        w.present()
    } else {
        fmt.Println("It (%p)did not walked :/\n ",&w)
    }
}
func describe(h Human) {
    fmt.Printf("Human named : %s (%p)\n",h.name,&h)
}

type MathFunc func(a,b int) int
func (f MathFunc) Apply(a,b int) int  {
    return f(a,b)
}
func add(a,b int) int {
    return a + b
}
func mult(a,b int) int  {
    return a * b
}

/*
func describePointer(h *Human) {
    fmt.Println("Human pointer named : ", h.name)
}
*/
func main () {
    h1 := Human{name: "Robert" }
    r1 := Robot{model: "R2D2" }
    h2 := Humanoid{age: 45,Human: h1}
    move(&h1)
    move(&r1)
    describe(h1)
    // embeeded type is directly abaiable on outer type
    h2.present()
    fmt.Println("########################")
    f1 := MathFunc(add)
    f2 := MathFunc(mult)
    fmt.Println("Applying apply on func add: ",f1.Apply(3,4))
    fmt.Println("Applying apply on func mult: ", f2.Apply(3,4))
}
