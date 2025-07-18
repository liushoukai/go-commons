package netutil

import (
	"log"
	"net"
)

func GetServerIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("获取本地IP失败:%s", err.Error())
		return "127.0.0.1"
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.IsPrivate() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
