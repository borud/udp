package main

import (
	"log"
	"net"
	"strings"
)

// ListenCmd ...
type ListenCmd struct {
	ListenAddress string `short:"l" long:"listen-address" default:":4711" description:"Listen address" required:"yes"`
	BufferSize    int    `short:"b" long:"buffer-size" default:"1024" description:"Buffer size for incoming packets" `
}

func init() {
	parser.AddCommand(
		"listen",
		"Listen for incoming UDP packets",
		"Listen for incoming UDP packets",
		&ListenCmd{})
}

// Execute the run command
func (r *ListenCmd) Execute(args []string) error {
	pc, err := net.ListenPacket("udp", r.ListenAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	log.Printf("| Listening on %s", r.ListenAddress)

	buf := make([]byte, r.BufferSize)
	for {
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		log.Printf("[%s] '%s'", addr, strings.TrimSuffix(string(buf[:n]), "\n"))
	}
}
