package main
import (
	"net"
	"strconv"
)

func cidrToMask(value string) string {
	result := value
	// If cidr is in range convert
	cidr, _ := strconv.Atoi(value)
	if (cidr >= 1 && cidr <= 32) {
		// Int to mask
		mask := net.CIDRMask(cidr, 32)
		// Mask to IP address with proper format
		result = net.IP(mask).String()
		return result
	}
	result = "Invalid"
	return result
}

func maskToCidr(value string) string {
	// if IP is not valid return Invalid
	valid := ipv4Validation(value)
	result := value

	if valid == false {
		result = "Invalid"
		return result
	}
	// Convert mask to CIDR
	ip := net.ParseIP(value)
	// IP to 4 byte
	addr := ip.To4()
	// 4 byte to Int
	cidr, _ := net.IPv4Mask(addr[0], addr[1], addr[2], addr[3]).Size()
	// Int to String
	result = strconv.Itoa(cidr)

	return result
}

func ipv4Validation(value string) bool {
	ip := net.ParseIP(value)
    var status bool
	status = true

	if ip == nil {
		status = false
	}

	return status
}
