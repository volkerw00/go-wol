package main

import (
	"fmt"
	"os"

	wol "github.com/voowoo/go-wol"
)

func sendMagicPacket(macAdress wol.MacAdress, broadcastAdress wol.BroadcastAdress) {
	var magicPacket = wol.NewMagicPacket(macAdress)
	magicPacket.Send(broadcastAdress)
}

func main() {

	error := parseArgs()
	if error != nil {
		os.Exit(1)
	}

	macAdress, error := wol.NewMacAdressFrom(options.Mac)
	if error != nil {
		fmt.Printf("Failed to parse %s as a MacAdress", options.Mac)
		os.Exit(1)
	}
	broadcastAdress, error := wol.NewBroadcastAdressFrom(options.BroadcastIP)
	if error != nil {
		fmt.Printf("Failed to parse %s as a broadcast adress", options.BroadcastIP)
		os.Exit(1)
	}

	sendMagicPacket(macAdress, broadcastAdress)

	os.Exit(0)
}
