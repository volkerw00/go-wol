package wol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

// MagicPacket is the representation of a magic packet.
// It is made up of an header containing 6 bytes of 0xFF
// and 16 repetitions of the MAC the packet is destined for.
type MagicPacket struct {
	header  [6]byte
	payload [16]MacAdress
}

// NewMagicPacket constructs a new magic packet for the given MAC
func NewMagicPacket(macAddress MacAdress) *MagicPacket {
	var packet MagicPacket
	for i := range packet.header {
		packet.header[i] = 0xFF
	}
	for i := range packet.payload {
		packet.payload[i] = macAddress
	}
	return &packet
}

// Send sends this magic packet to the given broadcast adress
func (m MagicPacket) Send(broadcastAdress BroadcastAdress) error {
	var buf bytes.Buffer

	binary.Write(&buf, binary.BigEndian, m)

	var localAddress *net.UDPAddr
	connection, error := net.DialUDP("udp", localAddress, broadcastAdress.adress)
	if error != nil {
		return error
	}
	defer connection.Close()

	bytesWritten, error := connection.Write(buf.Bytes())
	if error != nil {
		return error
	} else if bytesWritten != 102 {
		log.Panic(fmt.Sprintf("Warning: %d bytes written, %d expected!\n", bytesWritten, 102))
	}

	return nil
}
