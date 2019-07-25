package main

import (
	"flag"
	"xi-iot-cli/xi-iot/cmd"
)

func main() {
	flag.Parse()
	cmd.Execute()
}
