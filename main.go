package main

import (
	"log"
	"markets/cmd"
)

func main() {
	log.Fatal(cmd.HTTPServer())
}
