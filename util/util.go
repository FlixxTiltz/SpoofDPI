package util

import (
    "net"
    "log"
    "io"
    "strings"
    "github.com/babolivier/go-doh-client"
)


func ReadMessage(conn net.Conn)([]byte, error) {
    buf := make([]byte, 0, 4096) // big buffer
    tmp := make([]byte, 1024)     // using small tmo buffer for demonstrating
    for {
        n, err := conn.Read(tmp)
        if err != nil {
            if err != io.EOF {
                log.Fatal("ReadRequest error:", err)
            }
            return nil, err
        }
        buf = append(buf, tmp[:n]...)

        if n < 1024 {
            break
        }
    }

    return buf, nil
}

func ExtractDomain(message *[]byte) (string) {
    i := 0
    for ; i < len(*message); i++ {
        if (*message)[i] == '\n' {
            i++
            break;
        }
    }

    for ; i < len(*message); i++ {
        if (*message)[i] == ' ' {
            i++
            break;
        }
    }

    j := i
    for ; j < len(*message); j++ {
        if (*message)[j] == '\n' {
            break;
        }
    }

    domain := strings.Split(string((*message)[i:j]), ":")[0]

    return strings.TrimSpace(domain)
}

func DnsLookupOverHttps(dns string, domain string)(string, error) {
    // Perform a A lookup on example.com
    resolver := doh.Resolver{
        Host:  dns, // Change this with your favourite DoH-compliant resolver.
        Class: doh.IN,
    }

    log.Println(domain)
    a, _, err := resolver.LookupA(domain)
    if err != nil {
        log.Fatal("Error looking up dns. ", err)
        return "", err
    }

    ip := a[0].IP4 

    return ip, nil
}

func ExtractMethod(message *[]byte) (string) {
    i := 0
    for ; i < len(*message); i++ {
        if (*message)[i] == ' ' {
            break;
        }
    }

    method := strings.TrimSpace(string((*message)[:i]))
    log.Println(method)

    return strings.ToUpper(method)
}