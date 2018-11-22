package main

import (
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
)

var privateIPBlocks []*net.IPNet

const (
	isPvt  = "is-private"
	hasPvt = "has-private"
	isPub  = "is-public"
)

func init() {
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",
	} {
		_, block, _ := net.ParseCIDR(cidr)
		privateIPBlocks = append(privateIPBlocks, block)
	}
}

func isPvtIP(ip net.IP) bool {
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

func intersect(network *net.IPNet) bool {
	for _, block := range privateIPBlocks {
		if network.Contains(block.IP) || block.Contains(network.IP) {
			return true
		}
	}
	return false
}

func getType(network *net.IPNet) string {
	if intersect(network) {
		firstIP, lastIP := cidr.AddressRange(network)
		firstType := isPvtIP(firstIP)
		lastType := isPvtIP(lastIP)
		if firstType && lastType {
			return isPvt
		}
		return hasPvt
	}
	return isPub
}
