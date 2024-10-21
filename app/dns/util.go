package dns

import "strings"

func encodeDomain(domain string) string {
	name, extension := splitDomain(domain)

	return string(uint8(len(name))) + name + string(uint8(len(extension))) + extension + "\x00"
}

func splitDomain(domain string) (string, string) {
	domainParts := strings.Split(domain, ".")
	if len(domainParts) == 2 {
		return domainParts[0], domainParts[1]
	}

	return "", ""
}
