package main
import "fmt"
import "errors"
import "encoding/hex"

const b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func hex2base64(src []byte) []byte {
    // padding 
    for (len(src) % 3) != 0 {
        src = append(src,'0')
    }
    dst := make([]byte,len(src)*3/4)
    // take 3 character each time
    // then converr it to four characters
    for len(src) > 0 {
        // for byte to store the resulting base64 chars
        // character to translate
        var h1,h2,h3 = src[0],src[1],src[2]
        group24 := (int(h1) << 16) | (int(h2) << 8) | int(h3) 
        b1 := (group24 >> 18) & 0x3F
        b2 := (group24 >> 12) & 0x3F
        b3 := (group24 >> 6) & 0x3F
        b4 := (group24 >> 0) & 0x3F  
        
        //fmt.Printf("[h1,h2,h3] = [%d,%d,%d] || [b1,b2,b3,b4] = [%d,%d,%d,%d]\n",h1,h2,h3,b1,b2,b3,b4) 
        
        dst = append(dst,byte(b64[b1]))
        dst = append(dst,byte(b64[b2]))
        dst = append(dst,byte(b64[b3]))
        dst = append(dst,byte(b64[b4]))
        // the rest
        src = src[3:len(src)]

    }
    return dst
}

func fixedXOR(a,b []byte) ([]byte,error) {
    if len(a) != len(b) {
        return nil,errors.New("Buffer are not equal sized!")
    }
    c := make([]byte,len(a))
    for i:= 0; i < len(a); i++ {
        c[i] = a[i] ^ b[i]
    }
    return c,nil
}


func main() {
    input := []byte("Man is distinguished")
    out := hex2base64(input)
    fmt.Printf("1. input : %s ==> %s\n",string(input),string(out))
    input1 := "1c0111001f010100061a024b53535009181c"
    input1h,_ := hex.DecodeString(input1)
    input2 := "686974207468652062756c6c277320657965"
    input2h,_ := hex.DecodeString(input2)
    xored, _ := fixedXOR(input1h,input2h)
    xord := hex.EncodeToString(xored)
    fmt.Printf("2. output : %s\n",xord)
}

