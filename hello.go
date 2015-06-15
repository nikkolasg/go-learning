package main

import "fmt"
import "time"
import "strconv"
import "errors"
import "sort"
import "os"
import "strings"
import r "regexp"
import "math/rand"

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
    /* ERRORS */
    if v,e := raiser(4); e != nil{
        fmt.Printf("Raised error : %s\n",e)
    }else {
        fmt.Println("No error raised with value %d\n",v)
    }
    /* SORING */
    strs := []string{"dude","i","am","allright"}
    sort.Strings(strs)
    // C
    fmt.Println(strs)
    /* FILE */
    if f,err := os.Create("test.txt"); err != nil {
        panic(err)
    } else  {
        defer f.Close()
        fmt.Fprintln(f,"test data go")
        fmt.Println("Data written to file !!")
    }
    /* Function argument */
    ff := func (val string) bool { 
            if i := strings.Index(val,"l"); i != -1 {
                return true
            }
        return false 
    }
    if PIsTrue("hello",ff) {
        fmt.Println("Function argument working with hello")
    } else {
        fmt.Println("Problem with function arugment ??")
    }
    /* REGEXP */
    str4 := "Peach Peer Punch Apple"
    reg,_ := r.Compile("P([a-z]+)h")
    fmt.Println(reg.FindAllString(str4,-1))
    /* RANDOM */
     // < 100
     s1 := rand.NewSource(56)
    r1 := rand.New(s1)
    fmt.Println("Random #1 : " ,rand.Intn(100))
    fmt.Println("Random seeded : ",r1.Intn(50))
}

func PIsTrue(val string,f func(string) bool) bool {
    for _,v := range val {
        if f(string(v)) {
            return true
        }
    }
    return false
}
func raiser(v int) (int,error) {
    if v >= 2 {
        return -1,errors.New("value > 2... ERROR")
    }
    return v,nil
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
