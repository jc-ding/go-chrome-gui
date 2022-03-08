package utils

import (
	"net"
	"strconv"
)

func GeneratorPort() string {
	p := 8085
	for {
		err := CheckPorts(strconv.Itoa(p))
		if err != nil {
			p++
		} else {
			return strconv.Itoa(p)
		}
	}
}

func CheckPorts(port string) error {
	var err error

	tcpAddress, err := net.ResolveTCPAddr("tcp4", ":"+port)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", tcpAddress)
	if err != nil {
		return err
	} else {
		listener.Close()
	}

	return nil
}
