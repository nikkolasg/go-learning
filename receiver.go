package main

import "fmt"

type Walker interface {
    walk(nb int) bool
}

type Human struct {
    name string
}

type Robot struct {
    model string
}

func (h *Human) walk(nb int) bool {
    fmt.Println(h.name," has walked ", nb, " meters !")
    return true
}
func (r *Robot) walk(nb int) bool {
    fmt.Println(r.model," has walked ",nb," meters !")
    return true
}

func move(w Walker) {
    if w.walk(10) {
        fmt.Println("It has walked !!")
    } else {
        fmt.Println("It did not walked :/ ")
    }
}
func describe(h Human) {
    fmt.Println("Human named : ", h.name)
}
func describePointer(h *Human) {
    fmt.Println("Human pointer named : ", h.name)
}
func main () {
    h1 := Human{name: "Robert" }
    r1 := Robot{model: "R2D2" }
    move(&h1)
    move(&r1)
    describe(h1)
    describePointer(h1)
}
