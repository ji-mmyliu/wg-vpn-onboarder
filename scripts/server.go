package scripts

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/ji-mmyliu/wg-vpn-onboarder/models"
	"github.com/ji-mmyliu/wg-vpn-onboarder/util"
	templates "github.com/ji-mmyliu/wg-vpn-onboarder/wg-config-templates"
)

func SetupWireguardServer() {
	log.Println("Creating new server interface")

	EnsureWireguardInstalled()

	var interfaceName string
	var configPath string

	wgMainDir := util.GetInput[string](
		"(Optional) enter a custom directory for storing WireGuard configurations",
		WG_MAIN_DIR,
		func(res string) bool {
			_, err := os.Stat(res)
			return !os.IsNotExist(err)
		},
	)

	// Find first interface ID that has not been used yet
	availInterfaceId := 0
	for ; ; availInterfaceId++ {
		interfaceName = fmt.Sprintf(INTERFACE_NAME, availInterfaceId)
		configPath = fmt.Sprintf(INTERFACE_CONFIG, wgMainDir, interfaceName, interfaceName)
		if _, err := os.Stat(configPath); err != nil {
			break
		}
	}

	interfaceID := util.GetInput[int](
		"Please enter an interface ID",
		availInterfaceId,
		func(res int) bool {
			interfaceName = fmt.Sprintf(INTERFACE_NAME, res)
			configPath = fmt.Sprintf(INTERFACE_CONFIG, wgMainDir, interfaceName, interfaceName)
			if _, err := os.Stat(configPath); err != nil {
				return true
			}
			return false
		},
	)

	interfaceName = fmt.Sprintf(INTERFACE_NAME, interfaceID)

	// Get preferences from user
	networkID := util.GetInput[int](
		"Please enter a network ID (integer between 10-255 inclusive)",
		10,
		func(res int) bool {
			return res >= 10 && res <= 255
		},
	)

	endpoint := util.GetInput[string](
		"Please enter the server's public IPv4 address/endpoint",
		nil,
		func(res string) bool {
			return true
		},
	)

	listenPort := util.GetInput[int](
		"Please enter a listening port (integer)",
		51820,
		func(res int) bool {
			return res <= 65535
		},
	)

	isExitNode := strings.ToLower(util.GetInput[string](
		"Would you like this server to be an exit node? [Y/n]",
		nil,
		func(res string) bool {
			return strings.ToLower(res) == "y" || strings.ToLower(res) == "n"
		},
	)) == "y"

	dnsServer := util.GetInput[string](
		"(Optional) enter a custom DNS server address",
		"",
		func(res string) bool {
			return true
		},
	)

	// Generate network addresses
	networkAddressPrefix := fmt.Sprintf("10.%d.0", networkID)
	serverNetworkAddress := fmt.Sprintf("%s.1", networkAddressPrefix)
	networkAddress := fmt.Sprintf("%s.0", networkAddressPrefix)

	// Generate main directory and server private and public keys
	interfaceMainDir := fmt.Sprintf("%s/%s", wgMainDir, interfaceName)
	os.MkdirAll(interfaceMainDir, os.ModePerm)

	serverKeyPair := GenerateKeyPair(fmt.Sprintf("%s/server", interfaceMainDir))
	serverPublicKey := serverKeyPair.PublicKey
	serverPrivateKey := serverKeyPair.PrivateKey

	// Prepare to write server configuration
	serverConfigWriter, _ := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE, 0755) // 0600

	// Read server config template file and prepare parser
	serverConfigData, _ := templates.Asset("wg-config-templates/server_config.conf")
	tmpl, _ := template.New("serverConfigParser").Parse(string(serverConfigData))

	// Prepare network information to be injected into server config template
	serverCreds := models.Creds{
		PublicKey:  serverPublicKey,
		PrivateKey: serverPrivateKey,
	}
	server := models.Server{
		Creds:      serverCreds,
		ListenPort: uint(listenPort),
		Address:    serverNetworkAddress,
		Endpoint:   endpoint,
	}
	network := models.Network{
		InterfaceName: interfaceName,
		AddressPrefix: networkAddressPrefix,
		Address:       networkAddress,
		Server:        server,
		Clients:       []models.Client{},
		DnsServer:     dnsServer,
		IsExitNode:    isExitNode,
	}

	// Write server configuration data to JSON file
	serverJSON, _ := json.Marshal(network)
	serverJSONWriter, _ := os.OpenFile(fmt.Sprintf("%s/%s.json", interfaceMainDir, interfaceName), os.O_CREATE|os.O_RDWR, os.ModePerm)
	fmt.Fprint(serverJSONWriter, string(serverJSON))

	// Format and write configuration file
	tmpl.Execute(serverConfigWriter, &network)

	log.Printf("Successfully created new wireguard server interface %s!\n", interfaceName)
	if isExitNode {
		log.Printf("To enable this server as a VPN exit node, please run 'iptables -t nat -A POSTROUTING -o <public_interface> -j MASQUERADE' with the corresponding public network interface")
	}
}

func EnableServer(interfaceName string, wgMainDir string) {
	util.ShellCommand("sudo", "wg-quick", "up", path.Join(wgMainDir, interfaceName, fmt.Sprintf("%s.conf", interfaceName)))
	log.Println("Successfully started VPN interface", interfaceName)
}

func DisableServer(interfaceName string, wgMainDir string) {
	util.ShellCommand("sudo", "wg-quick", "down", path.Join(wgMainDir, interfaceName, fmt.Sprintf("%s.conf", interfaceName)))
	log.Println("Successfully shut down interface", interfaceName)
}
