package scripts

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
	"wg-vpn-onboarder/wgv/models"
	"wg-vpn-onboarder/wgv/util"
)

func OnboardNewClient() {
	// fmt.Print("== Please enter the name of the interface ")

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
	// fmt.Print("== Please enter the number associated with the interface you would like to use: ")
	interfaceChoiceIdx := util.GetInput[int](
		"Please enter the number associated with the interface you would like to use",
		1,
		func(res int) bool {
			return res >= 1 && res <= len(folders)
		},
	)

	// var interfaceChoiceIdx int
	// _, err := fmt.Scan(&interfaceChoiceIdx)
	// if err != nil {
	// 	log.Fatal("Error while reading user input. Please re-run to try again.")
	// }

	interfaceChoiceIdx-- // Convert back to 0-indexed

	interfaceName := folders[interfaceChoiceIdx].Name()
	log.Println("Generating configuration for new client on interface", interfaceName)

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
		Creds:   clientCreds,
		ID:      clientID,
		Address: fmt.Sprintf("%s.%d", network.AddressPrefix, clientID+1), // Server has ID of 1, clients start from 2

		PeerPublicKey:  network.Server.Creds.PublicKey,
		PeerEndpoint:   network.Server.Endpoint,
		PeerListenPort: network.Server.ListenPort,
		NetworkAddress: network.Address,
	}

	network.Clients = append(network.Clients, newClient)

	// Write network configuration data to JSON file
	networkJSON, _ := json.Marshal(network)
	networkJSONWriter, _ := os.OpenFile(fmt.Sprintf("%s/%s.json", interfaceDir, interfaceName), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	fmt.Fprint(networkJSONWriter, string(networkJSON))

	// Format and write network configuration file
	networkConfigWriter, _ := os.OpenFile(interfaceConfigPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	networkConfigTemplate, _ := os.ReadFile("templates/server_config.conf")
	networkTmpl, _ := template.New("networkConfigParser").Parse(string(networkConfigTemplate))
	networkTmpl.Execute(networkConfigWriter, &network)

	// Format and write new client configuration file
	clientConfigPath := path.Join(interfaceDir, fmt.Sprintf("client%d.conf", clientID))
	clientConfigWriter, _ := os.OpenFile(clientConfigPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	clientConfigTemplate, _ := os.ReadFile("templates/client_peer.conf")
	clientTmpl, _ := template.New("clientConfigParser").Parse(string(clientConfigTemplate))
	clientTmpl.Execute(clientConfigWriter, &newClient)

	log.Printf("Successfully generated client configuration 'client%d'!\n", clientID)
}
