package main

import (
    "crypto/rand"
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/sha1"
    "fmt"
    "os"
    "io"
    "io/ioutil"
)

func handleMessage(msgB []byte,other chan []byte) {
    str := string(msgB)
    splitted := strings.Split(str,":")
    if splitted[0] == "hash" {
        
    }
}

func handleBob(bob chan []byte) {

}
func main() {

    curve := elliptic.P256()

    //private := new(ecdsa.PrivateKey)
    private,err := ecdsa.GenerateKey(curve,rand.Reader)
    if err != nil {
        fmt.Fprintf(os.Stderr,"[-] Unable to generate ECDSA private/public keys. Abort\n")
        os.Exit(1)
    }

    var public ecdsa.PublicKey
    public = private.PublicKey

    fmt.Println("Public Key :\n%x\n",public)
    fmt.Println("Private Key:\n%x\n",private)

    msg := "I want to sign this message"
    fmt.Println("[+]Message to be signed (ECDSA with SHA2:\n%s\n",msg)

    sha := sha1.New()
    io.WriteString(sha,msg)
    hashed := sha1.Sum(nil)
    hashedSlice := hashed[0:len(hashed)]

    r,s,err := ecdsa.Sign(rand.Reader,private,hashedSlice)
    if err != nil {
        fmt.Fprintf(os.Stderr,"[-]Error while signing the message.Abort.")
        os.Exit(1)
    }

    fmt.Println("[+]Message hashed msg:\n%x\n",hashed)

    valid := ecdsa.Verify(&public,hashedSlice,r,s)
    if valid == true {
        fmt.Println("Message correctly verified !!! GREAT !")
        os.Exit(0)
    } else {
        fmt.Fprintf(os.Stderr,"[-] Message wrongly verified... snif.\n")
        os.Exit(1)
    }

}
