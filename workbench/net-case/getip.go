package main

import (
	"errors"
	"net"
)

func Ips() ([]string, error) {

	ips := []string{}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				ips = append(ips, ipv4.String())
			}
		}
	}
	if len(ips) == 0 {
		return nil, errors.New("no available ip")
	}
	return ips, nil
}
