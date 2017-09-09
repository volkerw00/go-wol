package main

import (
	"bufio"
	"os"

	flags "github.com/jessevdk/go-flags"
)

var options struct {
	Mac         string `short:"m" long:"mac" description:"The MAC adress of the device to wake in format 12:34:56:78:9A:BC" required:"true"`
	BroadcastIP string `short:"b" long:"broadcastIp" description:"The broadcast IPv4 to advertise the magic packet to" required:"true"`
}

func parseArgs() error {
	parser := flags.NewParser(&options, flags.Default)
	_, error := parser.Parse()
	if error != nil {
		f := bufio.NewWriter(os.Stdout)
		defer f.Flush()
		parser.WriteHelp(f)
		return error
	}
	return nil
}
