package main

import (
	"flag"
	"fmt"
	"os"

	wifipw "github.com/g-lib/wifi-password"
)

func usage() {
	fmt.Fprintf(os.Stderr, `wifipw is a command to get the password of current wifi.
usage: [sudo] wifipw 
`)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, `
`)
	os.Exit(0)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	ssid, password, _ := wifipw.GetWIFIPassword()
	fmt.Printf("SSID:%s\nPassword:%s\n", ssid, password)
}
