package main

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/apparentlymart/go-cidr/cidr"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func netmask(m []byte) (string, error) {
	if len(m) != 4 {
		return "", errors.New("Not IP4")
	}
	return fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3]), nil
}

func printCIDRInfo(cidrStr string) {
	ip, network, err := net.ParseCIDR(cidrStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "CIDR parse error: %v | %v.\n", cidrStr, err)
		os.Exit(2)
	}

	fmt.Println()

	// Base info:
	fmt.Println("IP:         ", ip)
	fmt.Println("Network:    ", network)

	// Netmask:
	mask, err := netmask(network.Mask)
	if err == nil {
		fmt.Println("Netmask:    ", mask)
	}

	// Range:
	firstIP, lastIP := cidr.AddressRange(network)
	fmt.Println("CIDR Range: ", firstIP, " <-to-> ", lastIP)

	// Count:
	count := cidr.AddressCount(network)
	prnt := message.NewPrinter(language.English)
	prnt.Printf("Count:       %d\n", count)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Stderr.WriteString("CIDR(s) required as parameter.\n")
		os.Exit(1)
	}
	for _, cidr := range args {
		printCIDRInfo(cidr)
	}
}
