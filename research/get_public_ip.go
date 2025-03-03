package main

import (
	"net/http"
	"net"
	"fmt"
)

func IsPublicIP(IP net.IP) bool {
    if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
        return false
    }
    if ip4 := IP.To4(); ip4 != nil {
        switch {
        case ip4[0] == 10:
            return false
        case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
            return false
        case ip4[0] == 192 && ip4[1] == 168:
            return false
        default:
            return true
        }
    }
    return false
}

func main() {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	fmt.Println(err)
	fmt.Println(req.Header.Get("X-Forwarded-For"))
}