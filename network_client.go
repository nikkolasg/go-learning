package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
)

func usage() string {
    return "[*] Usage : <IP address> <Port>\n"
}

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr,usage())
        os.Exit(1)
    }

    name := os.Args[1]
    strport := os.Args[2]
    ip := net.ParseIP(name)
    if ip == nil {
        fmt.Fprintf(os.Stderr,"[-] Wrong IP format. Abort")
        os.Exit(1)
    }

    if _,err := strconv.ParseInt(strport,0,64); err != nil  {
        fmt.Fprintf(os.Stderr,"[-] Wrong Port %s.%s\nAbort\n",strport,err.Error())
        os.Exit(1)
    }

    tcpaddr :=  name + ":" + strport
    fmt.Printf("[+] TCP address is %s\n",tcpaddr)
    //addr := ip
    //mask := addr.DefaultMask()
    //network := addr.Mask(mask)
    //ones, bits := mask.Size()
    //fmt.Println("Address is ", addr.String(),
    //" Default mask length is ", bits,
    //"Leading ones count is ", ones,
    //"Mask is (hex) ", mask.String(),
    //" Network is ", network.String())

    tcp,err := net.ResolveTCPAddr("tcp",tcpaddr)
    if err != nil {
        fmt.Fprintf(os.Stderr,"[-] Can not resolv to %s.\n[-]%s\nAbort.",tcp.String(),err.Error())
        os.Exit(1)
    }

    tcpconn,err := net.DialTCP("tcp",nil,tcp)
    if err != nil  {
        fmt.Fprintf(os.Stderr,"[-] Can not connect to %s.\n[-]%s\nAbort.\n",tcp.String(),err.Error())
        os.Exit(1)
    }
    _,err := tcpconn.Write([]byte("hello"))
    if err != nil {
        fmt.Fprintf(os.Stderr,"[-] Can not write.\n%s\nAbort",err.Error())
        os.Exit(1)
    }

    readBuf := bufio.NewReader(tcpconn)
    str,err :=  readbuf.readString('\n')
    if err != nil {
        fmt.Printf(os.Stderr,"[-] Error while receiving data.\n%s\nAbort",err.Error())
        tcpconn.Close()
    }
    if len(str) > 0 {
        fmt.Printf("Received from server : %s\n",str)
    }
    tcpconn.Close()
    os.Exit(0)
}
