package convert
// These functions need to be implemented
func cidrToMask(value string) string {
	return routeGetCidrToMask(value)
}

func maskToCidr(value string) string {
	return value
}

func ipv4Validation(value string) bool {
	return true
}
