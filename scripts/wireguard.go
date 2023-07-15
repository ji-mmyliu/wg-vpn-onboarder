package scripts

import (
	"fmt"
	"os"
	"strings"
	"wg-vpn-onboarder/wgv/util"
)

func EnsureWireguardInstalled() {
	wireguardTestOutput := util.RunShellCommand("command", "-v", "wg")
	if len(wireguardTestOutput) > 0 {
		return
	}

	fmt.Print("Wireguard command 'wg' was not found. Install Wireguard? [Y/n]: ")
	var confirmInstall string
	fmt.Scan(&confirmInstall)
	if strings.ToUpper(confirmInstall) == "Y" {
		InstallWireguard()
	} else {
		fmt.Println("Aborting script.")
		os.Exit(0)
	}
}

func InstallWireguard() {
	util.RunShellCommand("sudo", "apt-get", "update", "-y")
	util.RunShellCommand("sudo", "apt-get", "install", "-y", "wireguard")
}
