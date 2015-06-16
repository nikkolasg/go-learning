package main
import "fmt"
func main() {
    nb := 4
    str := "Man"
    strbyte := []byte(str)
    fmt.Printf("String %s[2] = %s\n",str,string(str[2]))
    fmt.Printf("String byte %v[:3] = %v\n",strbyte,strbyte[:3])
    fmt.Printf("Nb %d << 2 = %d, >> 2 = %d\n",nb,(nb<<2),(nb>>2))
}
