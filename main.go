package main

import (
	"os"
	"wg-vpn-onboarder/wgv/scripts"
)

func main() {
	switch os.Args[1] {
	case "newserver":
		scripts.SetupWireguardServer()
		break
	case "newclient":
		scripts.OnboardClient()
		break
	}
}
