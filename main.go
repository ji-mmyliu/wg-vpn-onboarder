package main

import (
	"os"
	"wg-vpn-onboarder/wgv/scripts"
)

func main() {
	switch os.Args[1] {
	case "server":
		switch os.Args[2] {
		case "new":
			scripts.SetupWireguardServer()
			break
		}
		break
	case "client":
		switch os.Args[2] {
		case "new":
			scripts.OnboardNewClient()
			break
		}
		break
	}
}
