package scripts

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"strings"
	"wg-vpn-onboarder/wgv/models"
	"wg-vpn-onboarder/wgv/util"
)

func EnsureWireguardInstalled() {
	_, err := exec.LookPath("wg")
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

func GenerateKeyPair(prefix string) models.Creds {
	serverPrivateKey := strings.TrimSpace(util.ShellCommand("wg", "genkey"))

	writePrivateKey, _ := os.OpenFile(fmt.Sprintf("%s.private", prefix), os.O_RDWR|os.O_CREATE, 0755)
	fmt.Fprint(writePrivateKey, serverPrivateKey)

	serverPublicKey := strings.TrimSpace(util.ShellCommandWithInput(fmt.Sprintf("%s.private", prefix), "wg", "pubkey"))
	writePublicKey, _ := os.OpenFile(fmt.Sprintf("%s.public", prefix), os.O_RDWR|os.O_CREATE, 0755)
	fmt.Fprint(writePublicKey, serverPublicKey)

	return models.Creds{PrivateKey: serverPrivateKey, PublicKey: serverPublicKey}
}

func ListInterfacesFromDir(dir string) []fs.DirEntry {
	files, _ := os.ReadDir(dir)

	folders := []fs.DirEntry{}
	for _, fileInfo := range files {
		_, err := os.Stat(path.Join(dir, fileInfo.Name(), fmt.Sprintf("%s.conf", fileInfo.Name())))
		if fileInfo.IsDir() && !os.IsNotExist(err) {
			folders = append(folders, fileInfo)
		}
	}

	return folders
}
