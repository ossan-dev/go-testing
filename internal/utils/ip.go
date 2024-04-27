package utils

import "net"

// IsValidIpOrCidr checks whether an IP address or CIDR address is valid or not.
// It return a bool value that states it.
func IsValidIpOrCidr(addr string) bool {
	if net.ParseIP(addr) != nil {
		return true
	}
	if _, _, err := net.ParseCIDR(addr); err != nil {
		return false
	}
	return true
}
