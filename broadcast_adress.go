package wol

import "net"

// BroadcastAdress is an net.UDPAddr
type BroadcastAdress struct {
	adress *net.UDPAddr
}

// NewBroadcastAdressFrom creates a new BroadcastAdress for the given IPv4 adress string
func NewBroadcastAdressFrom(ipv4Adress string) (BroadcastAdress, error) {
	var broadcastAdress BroadcastAdress
	udpAddress, error := net.ResolveUDPAddr("udp", ipv4Adress+":9")
	if error != nil {
		return broadcastAdress, error
	}
	broadcastAdress.adress = udpAddress
	return broadcastAdress, nil
}
