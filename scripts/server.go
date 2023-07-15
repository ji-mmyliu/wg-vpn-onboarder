package scripts

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"
	"wg-vpn-onboarder/wgv/models"
	"wg-vpn-onboarder/wgv/util"
)

const INTERFACE_NAME string = "wg%d"
const WG_MAIN_DIR = "."
const INTERFACE_CONFIG string = "%s/%s/%s.conf"

func SetupWireguardServer() {
	fmt.Println("Creating new server interface")

	EnsureWireguardInstalled()

	interfaceId := 0
	var interfaceName string
	var configPath string

	// Find first interface ID that has not been used yet
	for ; ; interfaceId++ {
		interfaceName = fmt.Sprintf(INTERFACE_NAME, interfaceId)
		configPath = fmt.Sprintf(INTERFACE_CONFIG, WG_MAIN_DIR, interfaceName, interfaceName)
		if _, err := os.Stat(configPath); err != nil {
			break
		}
	}

	// Get preferences from user
	networkID := util.GetInput[uint](
		"Please enter a network ID (integer between 10-255 inclusive)",
		10,
		func(res uint) bool {
			return res >= 10 && res <= 255
		})

	// Generate network addresses
	networkAddressPrefix := fmt.Sprintf("10.%d.0", networkID)
	serverNetworkAddress := fmt.Sprintf("%s.1", networkAddressPrefix)
	// networkAddress := fmt.Sprintf("%s.0", networkAddressPrefix)

	// Generate main directory and server private and public keys
	interfaceMainDir := fmt.Sprintf("%s/%s", WG_MAIN_DIR, interfaceName)
	os.MkdirAll(interfaceMainDir, os.ModePerm)
	serverPrivateKey := strings.TrimSpace(util.ShellCommand("wg", "genkey"))

	writePrivateKey, _ := os.OpenFile(fmt.Sprintf("%s/server.private", interfaceMainDir), os.O_RDWR|os.O_CREATE, 0755)
	fmt.Fprint(writePrivateKey, serverPrivateKey)

	serverPublicKey := strings.TrimSpace(util.ShellCommandWithInput(fmt.Sprintf("%s/server.private", interfaceMainDir), "wg", "pubkey")) //, interfaceMainDir)) // | tee %s/server.public", interfaceMainDir, interfaceMainDir))
	writePublicKey, _ := os.OpenFile(fmt.Sprintf("%s/server.public", interfaceMainDir), os.O_RDWR|os.O_CREATE, 0755)
	fmt.Fprint(writePublicKey, serverPublicKey)

	// Format configuration file and write to designated location
	serverConfigWriter, _ := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE, 0755) // 0600

	// Read server config template file and prepare parser
	serverConfigData, _ := os.ReadFile("templates/server_config.conf")
	tmpl, _ := template.New("serverConfigParser").Parse(string(serverConfigData))

	// Prepare network information to be injected into server config template
	serverCreds := models.Creds{PublicKey: serverPublicKey, PrivateKey: serverPrivateKey}
	server := models.Server{Creds: serverCreds, ListenPort: 51280}
	network := models.Network{Address: serverNetworkAddress, Server: server, Clients: []models.Client{}}

	// Write server configuration data to JSON file
	serverJSON, _ := json.Marshal(network)
	serverJSONWriter, _ := os.OpenFile(fmt.Sprintf("%s/%s.json", interfaceMainDir, interfaceName), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	fmt.Fprint(serverJSONWriter, string(serverJSON))

	// Format and write configuration file
	tmpl.Execute(serverConfigWriter, &network)
}
