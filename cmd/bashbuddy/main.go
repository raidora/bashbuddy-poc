package main

import (
	"bashbuddy/internal/app"
	"flag"
	"fmt"
)

var runMode *string

func main() {
	printBanner()
	parseArgs()

	switch *runMode {
	case "client":
		app.RunClient("localhost:55555")
	case "host":
		app.RunHost(":55555")
	default:
		fmt.Println("Choose a mode you dingus! (app -m host) or (app -m client)")
	}
}

func printBanner() {
	fmt.Println("-----------------------------------------")
	fmt.Println("         Running bashbuddy 2000          ")
	fmt.Println("-----------------------------------------")
}

func parseArgs() {
	runMode = flag.String("m", "", "host/client")
	flag.Parse()
}
