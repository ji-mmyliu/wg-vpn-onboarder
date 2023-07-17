package main

import (
	"fmt"
	"os"
	"strings"
	"wg-vpn-onboarder/wgv/scripts"
)

func main() {
	if len(os.Args) == 1 || strings.ToLower(os.Args[1]) == "help" {
		helpText, _ := os.ReadFile("templates/help.txt")
		fmt.Println(string(helpText))
		return
	}

	switch os.Args[1] {
	case "server":
		switch os.Args[2] {
		case "new":
			scripts.SetupWireguardServer()
			break
		case "up":
			scripts.EnableServer()
		case "down":
			scripts.DisableServer()
		}
		break
	case "client":
		switch os.Args[2] {
		case "new":
			scripts.OnboardNewClient()
			break
		case "connect":
			scripts.GetClientConfig()
			break
		}
		break
	case "version":
		versionText, _ := os.ReadFile("templates/version.txt")
		fmt.Println(string(versionText))
		break
	}

}
