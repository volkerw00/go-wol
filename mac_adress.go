package wol

import (
	"net"
)

// MacAdress is made up of 6 bytes
type MacAdress [6]byte

// NewMacAdressFrom creates a MacAdress struct for the given input MAC string.
// Since it is parsed via net.ParseMAC it should conform to that interface.
func NewMacAdressFrom(mac string) (MacAdress, error) {

	var macAddress MacAdress
	macHardwareAdress, error := net.ParseMAC(mac)
	if error != nil {
		return macAddress, error
	}

	for i := range macAddress {
		macAddress[i] = macHardwareAdress[i]
	}
	return macAddress, nil
}
