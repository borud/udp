package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

// Options ...
type Options struct {
	DBFilename string `short:"d" long:"db" default:":memory:" description:"Filename of database or :memory: for in-memory database"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
