package scripts

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/ji-mmyliu/wg-vpn-onboarder/models"
	"github.com/ji-mmyliu/wg-vpn-onboarder/templates"
	"github.com/ji-mmyliu/wg-vpn-onboarder/util"
)

func OnboardNewClient() {
	interfaceName, wgMainDir := ChooseExistingInterface()
	log.Println("Generating configuration for new client on interface", interfaceName)

	nickname := util.GetInput[string](
		"Please enter a nickname for this client (no spaces please)",
		"",
		func(s string) bool {
			return strings.Contains(s, " ")
		},
	)

	interfaceDir := path.Join(wgMainDir, interfaceName)
	interfaceConfigPath := path.Join(interfaceDir, fmt.Sprintf("%s.conf", interfaceName))
	interfaceJSONpath := path.Join(interfaceDir, fmt.Sprintf("%s.json", interfaceName))

	interfaceConfigData, err := os.ReadFile(interfaceJSONpath)
	if err != nil {
		log.Fatal("error occurred when reading server config JSON data:", err.Error())
	}
	var network models.Network
	if err := json.Unmarshal(interfaceConfigData, &network); err != nil {
		log.Fatal("error occurred when unmarshalling server config:", err.Error())
	}

	numExistingClients := len(network.Clients)
	clientCreds := GenerateKeyPair(fmt.Sprintf("%s/client%d", interfaceDir, numExistingClients+1))

	clientID := uint(numExistingClients) + 1

	newClient := models.Client{
		Creds:    clientCreds,
		ID:       clientID,
		Address:  fmt.Sprintf("%s.%d", network.AddressPrefix, clientID+1), // Server has ID of 1, clients start from 2
		Nickname: nickname,

		PeerPublicKey:  network.Server.Creds.PublicKey,
		PeerEndpoint:   network.Server.Endpoint,
		PeerListenPort: network.Server.ListenPort,
		NetworkAddress: network.Address,
		DnsServer:      network.DnsServer,
	}

	network.Clients = append(network.Clients, newClient)

	// Write network configuration data to JSON file
	networkJSON, _ := json.Marshal(network)
	networkJSONWriter, _ := os.OpenFile(fmt.Sprintf("%s/%s.json", interfaceDir, interfaceName), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	networkJSONWriter.Truncate(0)
	networkJSONWriter.Seek(0, 0)
	fmt.Fprint(networkJSONWriter, string(networkJSON))

	// Format and write network configuration file
	networkConfigWriter, _ := os.OpenFile(interfaceConfigPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	networkConfigTemplate, _ := templates.Asset("templates/server_config.conf")
	networkTmpl, _ := template.New("networkConfigParser").Parse(string(networkConfigTemplate))
	networkTmpl.Execute(networkConfigWriter, &network)

	// Format and write new client configuration file
	clientConfigPath := path.Join(interfaceDir, fmt.Sprintf("client%d.conf", clientID))
	clientConfigWriter, _ := os.OpenFile(clientConfigPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	clientConfigTemplate, _ := templates.Asset("templates/client_peer.conf")
	clientTmpl, _ := template.New("clientConfigParser").Parse(string(clientConfigTemplate))
	clientTmpl.Execute(clientConfigWriter, &newClient)

	log.Printf("Successfully generated client configuration 'client%d'!\n", clientID)
	log.Printf("Please make sure to restart the running server using '%s server restart' and export the client config using '%s client connect'", os.Args[0], os.Args[0])
}

func GetClientConfig() {
	interfaceName, wgMainDir := ChooseExistingInterface()
	clientConfigFilePath := ChooseExistingClient(path.Join(wgMainDir, interfaceName), interfaceName)

	options := []string{
		"QR Code (for the WireGuard mobile app)",
		"Paste configuration file content to output",
		"Paste configuration file location",
	}
	for idx, option := range options {
		fmt.Printf("(%d) %s\n", idx+1, option)
	}
	methodNum := util.GetInput[int](
		"Please enter a selection from above",
		1,
		func(res int) bool {
			return res >= 1 && res <= len(options)
		},
	)

	switch methodNum {
	case 1:
		GenerateConfigQRCode(clientConfigFilePath)
		break
	case 2:
		PasteConfigFileData(clientConfigFilePath)
		break
	case 3:
		configFileAbsPath, err := filepath.Abs(clientConfigFilePath)
		if err != nil {
			panic(fmt.Sprint("error occurred when fetching full path of config file:", err.Error()))
		}
		fmt.Printf("\nAbsolute path of client config file:\n%s\n", configFileAbsPath)
	}
}

func GenerateConfigQRCode(configFilePath string) {
	util.ShellCommandWithInputToConsole(configFilePath, "qrencode", "-t", "ansiutf8")
}

func PasteConfigFileData(configFilePath string) {
	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Sprint("error occurred when fetching config file text data:", err.Error()))
	}
	fmt.Printf("\n%s\n", string(configData))
}
