package main

import (
	"fmt"
	"net"
)

func GetMac() (mac_addresses []string) {
	net_interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get network interfaces: %v", err)
		return nil
	}
	for _, net_interface := range net_interfaces {
		mac_address := net_interface.HardwareAddr.String()
		if len(mac_address) == 0 {
			continue
		}
		mac_addresses = append(mac_addresses, mac_address)
	}
	return mac_addresses
}

func main() {
	fmt.Printf("mac addresses: %q\n", GetMac())
}