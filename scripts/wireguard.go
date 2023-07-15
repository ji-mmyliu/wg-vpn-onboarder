package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"wg-vpn-onboarder/wgv/util"
)

func EnsureWireguardInstalled() {
	_, err := exec.LookPath("wg")
	// wireguardTestOutput := util.ShellCommand("command", "-v", "wg")
	// if len(wireguardTestOutput) > 0 {
	// 	return
	// }
	if err == nil {
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
	util.ShellCommand("sudo", "apt-get", "update", "-y")
	util.ShellCommand("sudo", "apt-get", "install", "-y", "wireguard")
}
