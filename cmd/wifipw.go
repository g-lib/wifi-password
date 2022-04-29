package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	wifipw "github.com/g-lib/wifi-password"
)

func usage() {
	fmt.Fprintf(os.Stderr, `wifipw is a command to get the password of current WIFI.
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
	ssid, password, err := wifipw.GetWIFIPassword()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("WIFI名称标识(SSID):%s\n密码(Password):%s\n", ssid, password)
	}
}
