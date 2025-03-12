package common

import (
	"net"
)

func ParseIP(domain string) string {
	ips, err := net.LookupIP(domain)
	if err != nil {
		colorRedPrint := Colors(ColorRed)
		colorRedPrint.Printf("域名 %v 解析失败, %v\n", domain, err)
		return ""
	}

	return ips[0].String()

}
