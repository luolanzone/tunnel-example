package monitor

import (
	"log"
	"net"

	ipify "github.com/rdegges/go-ipify"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func GetPublicIP() (string, error) {
	ip, err := ipify.GetIp()
	if err != nil {
		return "", err
	} else {
		return ip, nil
	}
}
