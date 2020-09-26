package main

import (
	"log"
	"net"
)

// SendCmd ...
type SendCmd struct {
	DestAddr   string `short:"d" long:"dest-address" description:"Destination address" required:"yes"`
	SourceAddr string `short:"s" long:"source-port" default:":4711" description:"Source port for packet"`
}

func init() {
	parser.AddCommand(
		"send",
		"Send UDP packets",
		"Send UDP packets",
		&SendCmd{})
}

// Execute the run command
func (r *SendCmd) Execute(args []string) error {
	if len(args) == 0 {
		log.Fatalf("Nothing to send")
	}

	log.Printf("Sending '%s'", args[0])

	SourceAddress, err := net.ResolveUDPAddr("udp", r.SourceAddr)
	if err != nil {
		log.Fatalf("Unable to resolve source address: %v", err)
	}
	DestAddress, err := net.ResolveUDPAddr("udp", r.DestAddr)
	if err != nil {
		log.Fatalf("Unable to resolve destination address: %v", err)
	}

	c, err := net.DialUDP("udp", SourceAddress, DestAddress)
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}

	_, err = c.Write([]byte(args[0]))
	return err
}
