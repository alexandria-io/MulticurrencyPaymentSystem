package main

import (
	"fmt"

	"github.com/metacoin/multicurrency-payment-system"
)

func main() {

	// read configuration file
	fmt.Println("Reading configuration file (conf.json)...")
	conf := mucupa.ReadConfig("conf.json")
	fmt.Println("Starting listener on " + conf.Url + ":" + conf.Port)

	mucupa.MuxInit()

}
