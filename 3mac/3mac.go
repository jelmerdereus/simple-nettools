//3mac.go ~ v0.1 ~ Jelmer de Reus

package main

import (
        "fmt"
        "net"
        "os"
        "strings"
)

const usage = `3mac ~ v0.1 ~ Jelmer de Reus

Usage: 3mac <<MACADDRESS IN ANY FORMAT>>

Output: ABCDABCDABCD  ab:cd:ab:cd:ab:cd  abcd.abcd.abcd`

func main() {
        var hwaddr net.HardwareAddr
        var err error
        if len(os.Args) != 2 {
                fmt.Println("\n", usage)
                os.Exit(0)
        }
        if len(os.Args[1]) == 12 {
                hwaddr, err = net.ParseMAC((os.Args[1][:4] + "." + os.Args[1][4:8] + "." + os.Args[1][8:]))
                if err != nil {
                        fmt.Printf("\nError parsing the MAC address:  %s\n", err.Error())
                        os.Exit(1)
                }
        } else {
                hwaddr, err = net.ParseMAC(os.Args[1])
                if err != nil {
                        fmt.Printf("\nError parsing the MAC address:  %s\n", err.Error())
                        os.Exit(1)
                }
        }
        chars := strings.Replace(hwaddr.String(), ":", "", -1)
        flavour_two := strings.ToUpper(chars)
        flavour_three := chars[:4] + "." + chars[4:8] + "." + chars[8:]
        fmt.Printf("\n%s  %s  %s\n", hwaddr.String(), flavour_two, flavour_three)
}
