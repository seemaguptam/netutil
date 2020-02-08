package main

import (
    "fmt"
    "net"
)

func localAddresses() {
    ifaces, err := net.Interfaces()
    if err != nil {
        fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
        return
    }
    fmt.Printf("num of ifaces is %v\n", len(ifaces))
    for _, i := range ifaces {
        addrs, err := i.Addrs()
        if err != nil {
            fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
            continue
        }
        for _, a := range addrs {
	    fmt.Printf("%v  %v\n", i.Name, a)
        }
     }
}

func main() {
    localAddresses()
}
