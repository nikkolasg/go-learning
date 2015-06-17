package main

import (
    "os"
    "fmt"
    "net"
    "io/ioutil"
    "time"
)


func main() {
    service := ":1400"
    tcpAddr,err := net.ResolveTCPAddr("tcp",service)
    if err != nil {
        fmt.Fprintf(os.Stderr,"Error resolving service.\n%s\nAbort.",err.Error())
        os.Exit(1)
    }
    listener,err := net.ListenTCP("tcp",tcpAddr)
    if err != nil {
        fmt.Fprintf(os.Stderr,"Error listening to address.\n%s\nAbort.",err.Error())
        os.Exit(1)
    }
    msgchan := make(chan string)
    go logMsg(msgchan)
    for {
    con,err := listener.Accept()
        if err != nil {
             fmt.Fprintf(os.Stderr,"Error while Accepting connection.\n%s\nAbort.",err.Error())
            os.Exit(1)
        }
    go handleConnection(con,msgchan)
    }
    }

func handleConnection(con net.Conn,msgChan chan <- string) {
    defer con.Close()
    fmt.Printf("Received Connection from : %s\n",con.RemoteAddr().String())
    buf := make([]byte,4096)
    n,err := con.Read(buf)
    if err != nil {
        fmt.Fprintf(os.Stderr,"Error while reading socket.\n%s\nAbort.",err.Error())
        os.Exit(1)
    }
    msgChan <- string(buf[:n])
    ts := time.Now().String() + " (echo) : " + string(val)
    i,err := con.Write([]byte(ts))
    if err != nil {
        fmt.Fprintf(os.Stderr,"Error while sending back echoing.\n%s\nAbort.",err.Error())
        os.Exit(1)
    }
    i +=1

}

func logMsg(msgChan <- chan string) {
    for msg := range msgChan {
        fmt.Printf("Message Received : %s\n",msg)
    }
}
