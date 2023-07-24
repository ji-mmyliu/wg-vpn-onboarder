package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ji-mmyliu/wg-vpn-onboarder/scripts"
	"github.com/ji-mmyliu/wg-vpn-onboarder/templates"
)

func main() {
	if len(os.Args) == 1 || strings.ToLower(os.Args[1]) == "help" {
		helpText, _ := templates.Asset("templates/help.txt")
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
			interfaceName, wgMainDir := scripts.ChooseExistingInterface()
			scripts.EnableServer(interfaceName, wgMainDir)
		case "down":
			interfaceName, wgMainDir := scripts.ChooseExistingInterface()
			scripts.DisableServer(interfaceName, wgMainDir)
		case "restart":
			interfaceName, wgMainDir := scripts.ChooseExistingInterface()
			scripts.DisableServer(interfaceName, wgMainDir)
			scripts.EnableServer(interfaceName, wgMainDir)
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
		versionText, _ := templates.Asset("templates/version.txt")
		fmt.Println(string(versionText))
		break
	}

}
