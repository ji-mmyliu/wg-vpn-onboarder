package scripts

import (
	"fmt"
	"os"
	"testing"
	"text/template"
)

func TestSetupWireguardServer(t *testing.T) {
	SetupWireguardServer()
}

func TestTemplating(t *testing.T) {
	f, _ := os.OpenFile("wg0.conf", os.O_WRONLY|os.O_CREATE, 0600)

	serverConfigData, _ := os.ReadFile("templates/server_config.conf")
	fmt.Println(string(serverConfigData))
	tmpl, _ := template.New("serverConfigParser").Parse(string(serverConfigData))
	tmpl.Execute(f, "")
}
