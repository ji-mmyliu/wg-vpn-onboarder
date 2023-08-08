package templates

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticTemplates(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should have command summary template",
			args: args{
				fileName: "help.txt",
			},
		},
		{
			name: "Should have version template",
			args: args{
				fileName: "version.txt",
			},
		},
		{
			name: "Should have server_config template",
			args: args{
				fileName: "server_config.conf",
			},
		},
		{
			name: "Should have client_peer template",
			args: args{
				fileName: "client_peer.conf",
			},
		},
		{
			name: "Should have server_peer template",
			args: args{
				fileName: "server_peer.conf",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedData, err := os.ReadFile(tt.args.fileName)
			assert.Nil(t, err)
			templateData, err := Asset(path.Join("wg-config-templates", tt.args.fileName))
			assert.Nil(t, err)

			assert.Equal(t, expectedData, templateData)
		})
	}
}
