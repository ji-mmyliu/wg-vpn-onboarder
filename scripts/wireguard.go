package scripts

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/ji-mmyliu/wg-vpn-onboarder/models"
	"github.com/ji-mmyliu/wg-vpn-onboarder/util"
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

func ChooseExistingInterface() (string, string) {
	wgMainDir := util.GetInput[string](
		"(Optional) enter the custom directory you are using to store WireGuard configurations",
		WG_MAIN_DIR,
		func(res string) bool {
			_, err := os.Stat(res)
			return !os.IsNotExist(err)
		},
	)

	folders := ListInterfacesFromDir(wgMainDir)

	if len(folders) == 0 {
		log.Fatalf("The directory '%s' does not contain any configurations. Please run '%s server new' first.", wgMainDir, os.Args[0])
	}

	for idx, folderInfo := range folders {
		fmt.Printf("(%d) %s\n", idx+1, folderInfo.Name())
	}

	interfaceChoiceIdx := util.GetInput[int](
		"Please enter the number associated with the interface you would like to use",
		1,
		func(res int) bool {
			return res >= 1 && res <= len(folders)
		},
	)

	interfaceChoiceIdx-- // Convert back to 0-indexed

	interfaceName := folders[interfaceChoiceIdx].Name()

	return interfaceName, wgMainDir
}

func ChooseExistingClient(interfaceDir string, interfaceName string) string {
	interfaceJSONpath := path.Join(interfaceDir, fmt.Sprintf("%s.json", interfaceName))
	interfaceConfigData, err := os.ReadFile(interfaceJSONpath)
	if err != nil {
		log.Fatal("error occurred when reading server config JSON data:", err.Error())
	}
	var network models.Network
	if err := json.Unmarshal(interfaceConfigData, &network); err != nil {
		log.Fatal("error occurred when unmarshalling server config:", err.Error())
	}

	if len(network.Clients) == 0 {
		fmt.Fprintf(os.Stderr, "Error: interface %s does not yet have any clients. Please run '%s client new' first.\n", interfaceName, os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("\nAvailable clients on interface %s:\n", interfaceName)
	for idx, client := range network.Clients {
		fmt.Printf("(%d) Client %d %s\n", idx+1, client.ID, client.Nickname)
	}
	selectedClient := util.GetInput[int](
		"Please select a client number from above",
		len(network.Clients),
		func(res int) bool {
			return res >= 1 && res <= len(network.Clients)
		},
	)

	return path.Join(interfaceDir, fmt.Sprintf("client%d.conf", selectedClient))
}
