package main

import "fmt"
import "time"
import "strconv"

const constStr string = "My First Const String"
const constInt  = 56878

type person struct {
    name string
    age int
}
type baby struct {
    name string
    age int
}

func main() {
    fmt.Println("Hello YouShka")
    var mystr string =  "my First Declared String"
    var inferedStr = "My First Infered String"
    shortStr := "Short syntax infered style String"
    fmt.Println(mystr)
    fmt.Println(inferedStr)
    fmt.Println(shortStr)
    fmt.Println(constStr)
    fmt.Println(constInt)

    for i := 0; i < 2; i++ {
        fmt.Println("Long syntax For")
    }
    i := 2
    for i < 4 {
        fmt.Println("Short syntax ;)")
        i = i + 1
    }
    fmt.Print(fmt.Sprintf("i = %d\n",i))
    i += 2
    fmt.Print(fmt.Sprintf("i = %d\n",i))
    if i+=1; i < 8 {
        fmt.Println("first if syntax with statement inside !!!")
    }
    j := 1
    switch time.Now().Weekday() {
        case time.Friday, time.Saturday:
            fmt.Println("Hey we are Friday!")
        default:
            fmt.Println("Default case")
    }
    switch {
    case j == 2:
        fmt.Println("Variable case !")
    }
    /* ARRAY */
    arr := [1]int{10}
    fmt.Println("int arr ",arr," len = ", len(arr))
    strarr := make([]string,2)
    fmt.Println("slice ",strarr," cap = ", cap(arr))
    strarr2 := []string {"first","string","slice"}
    fmt.Println("slice2 ",strarr2," [1:",len(strarr2),"] ",strarr2[1:len(strarr2)])
    fmt.Println("slice2: [1:len()] = ",strarr2[1:len(strarr2)])
    fmt.Println("slice2: [1:2] = " , strarr2[1:2])
    fmt.Println("slice2: [1:3] = ", strarr2[1:3])

    /* MAP */
    map1 := make(map[string]int)
    fmt.Println("map1 : ",map1)
    map2 := map[string]int{"1":1,"2":2}
    fmt.Println("map2 : ",map2)
    value,present := map2["1"]
    fmt.Println("map2[\"1\"] : presnet? ",present," value : ", value)
    sum := 0
    for _,v := range map2 {
        sum += v
    }
    fmt.Println("sum map2: ",sum)
    sum = 0 
    ind := 0
    for i,v := range arr {
        sum += v
        ind += i
    }
    fmt.Println("arr sum : ",sum, " ind : " , ind)
    /* possible range over string "string" */
    /* FUNCTION */
    fmt.Println("######## FUNCTION ###########")
    suming,bo := add(4,5,6)
    fmt.Println("add 4,5,6 : ",suming," 2nd return : ",bo)
    arr3  := []int{5,6}
    fmt.Println("add2 4,[5,6] : ",add2(4,arr3 ...))
    fmt.Println("anonymous with type mult : ",apply(mult,3,4))
    fmt.Println("anonymous declared: ",apply(func (a,b int) int { return a/b },8,2))
    /* POINTERS */
    value2 := 4
    fmt.Println("pointer -1 : ",minusone(&value2)," and orig value =  ",value2)
    /* STRUCT */
    p1 := person{age: 40,name: "moha"}
    b1 := baby{name: "baby",age: 2}
    pb1 := &b1
    fmt.Println("struct p1 : ", p1)
    pp1 := &p1
    fmt.Println("struct pointer p1 : ", pp1.age)
    fmt.Println("struct before receiver ", pp1)
    pp1.older()
    fmt.Println("after : ",pp1)
    /* INTERFACES */
    fmt.Println("interface human:present() : ", pp1.present())
    /* works because receiver pointer type acccept for pointer and value */
    fmt.Println("baby human:present() : ", b1.present())
    
    fmt.Println("interface method present(): ", presentation(pb1))
    //fmt.Println("interface method pesentation2: ",presentation2(pb1))
}

type human interface {
    present() string
}

func (p *person) present() string {
    return "My name is " + p.name + " of " +strconv.Itoa(p.age) +  "years old"
}
/** cannot call with value directly for interface i think */
func (b *baby) present() string {
    return "Baby name is " + b.name + " of " + strconv.Itoa(b.age)
}

func presentation(h human) string {
    return h.present()
}

func (p *person) older() {
    p.age += 1
}
func minusone(a * int) int {
    return *a - 1
}
func apply(a bin,arg1,arg2 int) int {
    return a(arg1,arg2)
}
type bin func(int,int) int
func mult(a,b int) int {
    return a * b
}

func add(a,b,c int) (int,bool) {
    return a + b + c,true
}
func add2(a int,b ...int) int {
    sum := a
    for _,v := range b {
        sum += v
    }
    return sum
}
