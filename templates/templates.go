// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/help.txt
// templates/server_config.conf
// templates/server_peer.conf
// templates/templates.go
// templates/version.txt
package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesHelpTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xf4\x40\x0c\x84\xfb\x3c\xc5\xfc\xdd\x4f\x87\x8e\x8e\xf6\xa0\x41\x08\x4e\x3a\x89\xab\x4d\xe2\x24\x16\xbb\xeb\xc8\xeb\x5c\xc8\xdb\xa3\x64\x81\x0b\xd2\xb9\xdb\x9d\x6f\x3c\xe3\x13\x87\x5a\x23\xc3\x15\xde\x33\x4e\x62\xdc\x8d\x64\x0d\xde\x0e\x2f\x78\x4d\xef\x4a\xd6\xb0\xfd\xab\xf6\x3a\xcc\x26\x5d\xef\xf8\xbf\xbf\xc1\xee\x76\x77\x87\x27\x89\x71\xc6\xb3\x8c\x55\x75\x1c\x63\x24\x9b\xa1\x2d\x6a\x8d\x91\x52\x93\xef\xab\x6a\xea\xce\xc8\x6c\x67\x36\x24\x9e\xb0\xce\xde\x98\x9c\x33\x68\xfd\x5a\x42\x12\xfb\xa4\xf6\x01\x49\xce\xd6\x52\xcd\x5b\xdf\x38\x14\x1b\x8e\x4e\xe6\x8b\x6d\xb1\x14\x71\xcb\x35\x3a\xa5\xc2\xf5\xa3\xe7\xf2\xfc\xc3\xae\x70\x1d\x84\x93\x5f\xca\x1c\x8c\x07\xb2\xb5\xcd\xb7\x54\x6b\x6a\xa5\x1b\x8d\x5c\x34\xa1\x55\x03\x6d\xd3\x2e\x54\xe2\xda\x81\xc7\xcf\x41\x4b\xad\xeb\x7e\x09\x5c\x92\x7b\x0e\x3f\x97\xac\xf3\x20\x79\x08\x34\x67\x78\x2f\xb9\xa8\x91\x73\xa6\xae\x5c\x7f\x66\xcb\xcb\x82\x2b\x38\xff\x8a\xda\x62\x61\x25\x65\xa7\x10\xb8\xf9\x0a\x00\x00\xff\xff\x86\xf0\x74\x09\xcb\x01\x00\x00")

func templatesHelpTxtBytes() ([]byte, error) {
	return bindataRead(
		_templatesHelpTxt,
		"templates/help.txt",
	)
}

func templatesHelpTxt() (*asset, error) {
	bytes, err := templatesHelpTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/help.txt", size: 459, mode: os.FileMode(420), modTime: time.Unix(1689637104, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesServer_configConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\x4f\x4b\x33\x31\x10\xc6\xef\xf9\x14\x03\xef\x7b\xdd\x88\x7f\xf0\x20\xec\x61\xd9\x2a\x14\xa5\x86\xaa\x78\x28\x45\xd2\x64\xb6\x44\x63\x76\x99\xc4\x96\x12\xe6\xbb\x4b\xeb\xae\xed\x1e\xbc\x0d\xcf\xfc\xe6\x21\xf9\x2d\xa6\x21\x21\x35\xda\xe0\x52\x28\x72\x1b\x9d\xf0\x1e\x77\x50\x42\xce\xf2\x09\x69\x83\x24\x6b\x42\x1b\xe5\x71\xc9\x2c\x2a\x6b\x09\x63\x1c\x61\x7d\xc6\x7c\x76\x79\x21\x1e\x5c\x4c\x18\x54\x4b\x69\xc4\x1c\x63\x66\xa1\x08\x5f\x3a\x28\x21\xee\xa2\x49\x1e\x8a\x2d\x04\x4c\xd2\x75\x9b\x2b\xe9\xba\xb7\xa6\xa5\xad\x26\x5b\x9e\xff\xcd\x5d\x4b\xd3\x86\x46\x6a\xef\x65\x4f\xbb\xb0\x3e\x39\x70\x5d\xd2\x2b\x8f\x11\x8a\x0a\xee\x1e\xe7\xaf\xd5\x7c\x02\x85\xdb\x3f\xe7\xf7\xd3\x33\xfd\x89\xcc\x50\xbc\x43\x55\xd7\xb7\xea\x59\xe4\x4c\x3a\xac\x11\xfe\x1b\xef\x30\x24\xb8\x29\x41\xd6\x87\x31\x32\x8b\x7f\xf0\x33\x43\xce\x3d\x20\xa7\x13\x66\xc8\xd9\x35\xc3\x89\x9c\x39\xf3\x11\x0e\xbd\x47\xea\x34\xc3\x60\x99\xc5\x42\x21\xd2\x52\xa8\xaf\x95\x77\x66\x50\x3e\xe0\xbd\xf2\x61\xb7\x37\xee\x7d\xbb\x45\x3b\x55\x71\x04\x8e\xa4\xf7\xd5\xdf\x01\x00\x00\xff\xff\x87\xa7\x7c\x14\xd3\x01\x00\x00")

func templatesServer_configConfBytes() ([]byte, error) {
	return bindataRead(
		_templatesServer_configConf,
		"templates/server_config.conf",
	)
}

func templatesServer_configConf() (*asset, error) {
	bytes, err := templatesServer_configConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/server_config.conf", size: 467, mode: os.FileMode(420), modTime: time.Unix(1690161443, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesServer_peerConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\x49\xce\xc9\x4c\xcd\x2b\x51\xb0\xb2\x55\xd0\x73\x06\x33\x8b\x6b\x6b\xb9\xa2\x03\x52\x53\x8b\x62\xb9\x02\x4a\x93\x72\x32\x93\xbd\x53\x2b\x15\x6c\x15\xaa\xab\xa1\x4a\xf5\x9c\x8b\x52\x53\x8a\xf5\xe0\x72\xb5\xb5\x5c\x8e\x39\x39\xf9\xe5\xa9\x29\x9e\x01\xc5\x28\x0a\x1d\x53\x52\x8a\x52\x8b\x8b\x6b\x6b\xf5\x8d\x8d\xb8\xaa\xab\x53\xf3\x52\x6a\x6b\x01\x01\x00\x00\xff\xff\xad\x51\x8d\x87\x78\x00\x00\x00")

func templatesServer_peerConfBytes() ([]byte, error) {
	return bindataRead(
		_templatesServer_peerConf,
		"templates/server_peer.conf",
	)
}

func templatesServer_peerConf() (*asset, error) {
	bytes, err := templatesServer_peerConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/server_peer.conf", size: 120, mode: os.FileMode(420), modTime: time.Unix(1689553641, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTemplatesGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x5d\x6f\xdb\xca\x11\x7d\x26\x7f\xc5\x5e\x01\x37\x20\x0b\x57\xe2\xf7\x87\x00\x03\xc5\x4d\x52\x34\x0f\xcd\x05\x9a\xf4\xa9\x53\x04\x4b\x72\xd7\x21\x2a\x89\x2a\x49\xe7\x8e\x13\xf8\xbf\x17\x67\x67\xed\xca\x89\x93\x0b\x14\x7d\x90\x2d\x71\xb9\x33\x67\x66\xce\x9e\xb3\xbb\x9d\x7a\x39\x0d\x46\xdd\x98\x93\x99\xf5\x6a\x06\x65\xa7\x59\x9d\x75\xff\x2f\x7d\x63\xd4\x51\x8f\x27\xd5\xdd\xa9\x9b\xe9\x8f\xdd\x78\x1a\xf4\xaa\xd5\xab\x5f\xd5\xdb\x5f\xdf\xab\xd7\xaf\xde\xbc\xdf\xaa\xe8\x4f\x8f\xfb\xe2\x70\xb7\x53\xcb\x74\x3b\xf7\x66\xd9\xe3\xfb\x6a\x8e\xe7\x83\x5e\xcd\xb2\xfb\x68\x0e\xe7\xed\xca\xeb\xd3\xa7\x8b\x99\x3f\x99\xf9\x43\x3f\x9d\xec\x78\xb3\xc5\xbf\x67\xd7\xcf\xc6\xcc\xcf\xac\x3e\x7e\xdb\xde\x4c\x4f\x57\x3e\x99\x79\x19\xa7\x93\x4b\x78\x59\x47\x18\x8e\xc7\xf3\x34\xaf\x2a\x0a\x83\x4d\x77\xb7\x9a\x65\x13\x06\x9b\x7e\x3a\x9e\x67\xb3\x2c\xbb\x9b\xcf\xe3\x19\x0f\xec\x71\xc5\xbf\x71\x92\xbf\xbb\x71\xba\x5d\xc7\x03\x7e\x4c\x6e\xc3\x59\xaf\x1f\x77\x76\x3c\x18\x7c\xc1\x83\x65\x9d\xc7\xd3\x8d\x5b\x5b\xc7\xa3\xd9\x84\x71\x18\xda\xdb\x53\xaf\x7c\xc7\xfe\x66\xf4\x10\xb9\xd6\xfd\xe3\x9f\x48\x7b\xa5\x4e\xfa\x68\x94\x6c\x8b\x55\xf4\xf0\xd4\xcc\xf3\x34\xc7\xea\x4b\x18\xdc\x7c\x76\xbf\xd4\xfe\x5a\x01\xd5\xf6\xad\xf9\x0d\x41\xcc\x1c\x39\xd8\xf8\xfd\xcb\xad\xb5\x66\x76\x61\xe3\x38\x0c\x46\xeb\x36\xfc\x74\xad\x4e\xe3\x01\x21\x82\xd9\xac\xb7\xf3\x09\x3f\xaf\x94\x3d\xae\xdb\xd7\x88\x6e\xa3\x0d\x02\xa9\x9f\xff\xbd\x57\x3f\x7f\xda\x08\x12\x97\x2b\x0e\x83\xfb\x30\x0c\x3e\xe9\x59\x75\xb7\x56\x49\x1e\x49\x12\x06\x1f\x04\xce\xb5\x1a\xa7\xed\xcb\xe9\x7c\x17\xbd\xe8\x6e\xed\x95\xba\xf9\x1c\x87\x41\x7f\x78\xfd\x80\x74\xfb\xf2\x30\x2d\x26\x8a\xc3\xff\x17\x1e\x84\x91\xf8\xdf\x09\x64\xe6\x59\x70\xfb\x87\xdd\xad\xdd\xfe\x02\xe8\x51\x7c\x85\x37\xc2\xfb\x30\x5c\xef\xce\x46\xe9\x65\x31\x2b\x5a\x7e\xdb\xaf\x88\xe2\xea\xf3\xf3\x08\x83\xf1\x64\x27\xa5\xa6\x65\xfb\xe7\xf1\x60\xde\x9c\xec\xf4\xb8\xcf\x8f\xf0\xe1\xf9\x45\x04\x37\x43\xa5\xfc\x18\xc3\x60\x19\x3f\xbb\xdf\xe3\x69\xad\x8a\x30\x38\xe2\x4c\xa9\xc7\xa0\x7f\x9d\x06\xe3\x1e\xbe\x1f\x8f\x46\x81\x26\x5b\x7c\x43\x9e\xdd\x4e\xbd\x45\x2c\x5f\x02\x98\xe5\xda\x20\x1c\x8a\xec\xf8\x35\x88\xd8\xbd\x1f\xc5\x3e\x35\xc0\x3c\xee\xdd\xba\x9d\x12\xf5\x1d\x10\x5d\x46\x05\xc4\x1f\x44\xc5\xfb\x51\x2c\x05\x3c\x0d\xea\x36\x4a\x50\x14\xf2\x24\x28\x0a\xfd\x41\x50\xbc\x1f\xc5\x97\x6d\x78\x1a\xda\x6d\xff\x7e\xe8\xd1\xde\xb9\x6e\xfd\x38\x03\x5a\x19\xc5\xff\x6d\xeb\x37\x29\x2e\x7a\xfd\x66\x79\x35\xce\x4f\xd2\xfc\xf6\xd1\xac\x1f\xcd\xac\xb4\x1a\xc6\xd9\xf4\xeb\x34\xdf\xfd\x20\x9d\xdb\x1f\xc5\xaa\x9b\xa6\xc3\xb7\xa5\xbc\x98\x96\x2d\xea\x40\x8e\x9f\xae\x55\xf2\x30\x8a\xbb\xe5\x49\xca\x71\x51\xcb\xdd\xf2\x7b\xbd\x7b\x77\xb7\xc8\x3c\xcc\x6c\x75\x6f\xbe\xdc\x5f\xe4\xf3\xe4\xc6\x79\xfd\xf0\x28\x7d\x7f\x31\x87\xf3\x7b\x5e\xd5\xb5\x67\x76\xb4\x21\x4e\x2d\x71\xd3\x11\x27\x0d\x71\x92\x3c\xff\xb1\x96\xb8\xea\x89\x1b\x4b\xdc\xa5\xc4\x85\x21\xb6\x05\x71\x81\x75\x3c\x2f\x88\x6d\x47\x9c\xf7\xc4\x7d\x49\x6c\x7b\xe2\x61\x20\x2e\x10\xbb\x26\x6e\x8c\x7c\x6c\x45\xac\x13\xe2\x22\x95\x7c\x88\x93\x6b\xe2\xa6\x25\xd6\x1d\x71\x31\x10\x9b\x8c\x38\x2b\x88\xd3\x8a\xb8\xeb\x88\x4d\x47\xdc\x37\xf2\xbf\xec\xe5\xfb\xd0\x11\xeb\x9c\xb8\x2a\x88\x1b\xc4\xea\x88\x87\x8c\xb8\x6b\x65\xad\x1d\x88\x2b\x2b\x78\x4c\x4e\x9c\xe6\x82\xa3\xd4\xc4\x59\x4e\xdc\xe3\x59\x49\x3c\x20\x7f\x2e\x38\xaa\x8c\x78\x40\x2d\x83\xc4\x4d\x06\x59\x4f\x0c\x71\x66\x89\xeb\xc6\xe3\xb3\xc4\x85\x26\x1e\x80\x0f\xbd\x19\x04\x3b\x6a\x43\x2d\x7d\x4f\x9c\x55\xc4\xa5\x7f\xd7\x36\xc4\x9d\x95\x4f\x9f\x12\x1b\x43\x5c\x57\xc4\x75\x2d\x98\xb2\x5a\xea\xaf\x53\xe2\x1e\x31\x81\xb5\x27\x2e\x4b\xe2\xba\x24\x4e\x7b\xe2\x2a\x97\x9e\xb4\xa8\x3b\x25\xce\x80\x51\x0b\xd6\x36\x25\x2e\x33\xe2\x36\x97\x7c\xc0\x82\x35\x03\x2c\x46\xfa\x85\xd9\x55\x35\x71\x5e\xf9\x38\x46\xb0\x63\x1d\x35\xb6\x0d\x71\xdb\x4b\x2f\xaa\x46\x6a\x42\xaf\x8a\x8c\x38\xcd\x64\xb6\xba\x90\x1a\x13\xcc\xbf\xf5\x7b\x2b\xc9\xdd\xe3\x7d\xcc\x00\x3d\x6f\x88\x53\xcc\xaf\x93\x99\xa3\xb7\xa6\x12\x8e\x55\x83\xf0\x07\xeb\xae\x5e\x70\xa6\x94\xbe\xe9\x92\xb8\x47\xae\x52\x66\x6b\x6a\x62\x9b\x09\x97\x30\x2f\x8d\xbe\x25\xc4\xe9\x20\x7c\x43\xbd\x25\x7a\xaa\xa5\x47\xe8\x59\x52\x13\x77\xe0\x41\x29\x98\x3a\xcc\xbc\x20\xae\x7c\x4f\x90\xc3\xe1\x1a\x84\x4b\x79\x21\xfd\x44\xaf\x93\x5c\xf0\x0d\xe8\xb5\x91\x7d\xe0\xe2\xa0\x85\x63\x7d\x4d\xdc\x5b\xe1\x6e\x81\x7e\x0c\xc2\xc9\x1a\x1c\x69\x25\x5e\x9b\x11\xd7\x9d\x70\x26\xb7\xc4\x6d\x4d\xac\x7b\x62\x8b\x98\x09\x71\xdd\x0a\xef\x91\x17\xf3\x00\xa7\xc0\x2d\xf0\x56\x37\x32\xcb\x1a\xb5\x56\x52\x2f\x62\xd6\x38\x7b\x95\xf4\x0a\x9f\x06\xe7\xa3\x93\x3e\xe3\x5c\x36\x5a\x30\x82\xbf\x15\xb8\x51\x12\x57\xe8\x61\x4d\x9c\x26\xc4\x1d\xde\x43\x5e\xfd\xf4\x3c\xbb\xbd\x15\xb1\x05\xae\x42\x6a\x40\x7c\xcc\x57\xde\xdb\x3c\xdc\x53\xbe\xd6\x0f\xef\xa0\xcf\xdd\x4c\x1e\x7c\xf6\xe2\x66\x13\x06\xc1\x37\x0a\x74\x15\x06\xc1\xe6\xdb\xfb\xdf\xe6\x2a\x0c\x62\xc8\xd6\xb3\x69\x91\xf1\x0f\xce\xa8\x2f\x33\x3a\xa7\x7e\xbc\x0e\x7d\x07\xea\xef\x5d\x38\x1e\xef\x09\xce\xe9\xf7\xd7\x5f\x6b\xed\x17\xd8\xe6\x5e\x3d\x0b\xd9\xd9\xe6\x5e\x15\x65\x7b\xe5\x04\x7b\x7f\x69\x66\x51\x91\x25\xb1\x7b\x0e\x8b\xd9\x8b\x05\xfd\xfd\x34\x72\x94\x56\x4d\x5b\xe5\x75\x9a\x14\x57\x2a\x89\xef\xc3\x40\x23\xef\x0b\x57\xdf\x17\x57\xd4\x5e\xf9\xda\x00\x6a\xef\xfe\xde\x3f\x76\x58\x5f\x3d\x2f\xf1\xef\x2e\x2f\xcd\x2f\xa7\x93\xfd\x9f\xc4\x1e\x84\x68\x13\x11\x70\x90\x1d\xc2\x90\xa7\x42\x28\x88\x94\x13\xb6\x56\x0e\x72\xe2\x85\x07\xc4\x87\xe8\x37\x8d\x90\x16\xc4\x02\xe9\x4d\x2f\xc4\x1c\x5a\xe2\x4c\xcb\x1e\x1c\x44\x90\x4f\x6b\x7f\x08\x20\xae\xa5\x88\x38\xc4\xb7\xab\x88\x8b\x42\x84\x0f\x62\xd9\x82\x9c\xc0\x54\x89\x20\x19\x6f\x0e\x85\x37\x08\x1c\x16\x33\x10\xa7\x10\xb6\x5e\xc4\x1b\x87\x15\xe2\x81\x77\xb3\x54\xf0\x42\x38\xb5\xff\x8d\x0f\xc4\x07\x87\xc7\x24\x22\x64\xc0\x51\x67\x22\x12\x30\x11\xd4\x80\x98\x10\xeb\x32\x11\x41\x84\xf0\x41\x98\x70\x60\xaa\x96\xb8\xf1\xe2\x0c\x91\x29\xfc\x5e\x53\x7a\x81\x43\xce\x5e\xea\xc6\xba\xdb\x93\x7b\x51\xc7\x61\x1f\xa4\x9f\x75\xef\x4d\xa1\x15\x5c\xa9\x3f\xfc\x45\x4f\x9c\x36\x22\x46\xa8\xb5\xf2\x7d\x78\x30\x05\x08\x83\xf6\x66\x0a\x31\x84\x90\xa2\x86\xcc\x1b\x8d\xce\x44\xac\x11\x0f\x42\x81\xfa\x93\x42\xe2\xa2\xd7\x30\x18\x98\x0a\x04\xc5\xb4\xc4\x9d\x16\xd1\x74\x02\x54\x8a\xc8\x39\x33\xeb\xc4\x34\xc0\x0d\x27\xf6\x83\xe0\x81\x58\x62\x8e\x45\xe5\x4d\xc7\x8a\xf8\x80\x37\x98\x0f\x4c\x26\x87\x18\xb6\x22\xde\xd8\x37\x78\x01\xc3\x65\x22\x4d\x05\x17\xc4\x09\x78\x81\x13\xc2\xaf\x2d\xf1\x50\x4a\x5f\x92\x8c\xb8\x29\xc5\xd8\x21\xaa\xce\x18\x6a\xc9\x0d\x5e\xc2\xac\x60\xa0\x30\x5e\xcc\x08\xf3\x2f\x72\x11\xf5\xa1\x96\x7a\x60\x86\x65\x4b\x6c\x7c\xed\xce\x70\x7a\xc9\x6f\x8d\x9f\x19\xfa\x97\xc8\x2c\x21\x9a\x59\x2b\x62\x0f\x93\x83\x51\x83\xff\x79\x26\xe7\xe7\x81\xdf\xee\xf2\x90\x0b\x27\xf2\x41\xf8\xee\x84\xd7\xcf\xc6\x5d\x56\x5a\xe1\x19\x4c\x0c\x38\x9d\xd1\xb6\x62\x0a\xc8\x0f\xb3\xe8\xd0\x4f\x8f\x13\xe6\x80\x77\xab\xc4\xf3\xbe\x14\x5e\x61\xa6\xe8\x1d\x78\x0a\xb3\x40\x8f\xda\x52\xb8\x89\x5c\xa5\xef\x35\xb8\x8c\x78\xd6\x1b\x02\xea\xcf\x6b\x7f\xf1\x18\xfc\xb9\x81\xe9\x76\xfe\x92\x91\x0a\x27\x60\xd6\xce\xe4\x6b\xe9\x3d\xcc\x3c\x49\xe9\x3f\x01\x00\x00\xff\xff\x6b\x8a\x56\x98\x00\x10\x00\x00")

func templatesTemplatesGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesTemplatesGo,
		"templates/templates.go",
	)
}

func templatesTemplatesGo() (*asset, error) {
	bytes, err := templatesTemplatesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/templates.go", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1690161793, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVersionTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4f\x2f\x53\x28\x4b\x2d\x2a\xce\xcc\xcf\x53\x28\x33\xd4\x33\xd0\x33\x01\x04\x00\x00\xff\xff\x54\x71\x23\x1d\x12\x00\x00\x00")

func templatesVersionTxtBytes() ([]byte, error) {
	return bindataRead(
		_templatesVersionTxt,
		"templates/version.txt",
	)
}

func templatesVersionTxt() (*asset, error) {
	bytes, err := templatesVersionTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/version.txt", size: 18, mode: os.FileMode(420), modTime: time.Unix(1690161785, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/help.txt":           templatesHelpTxt,
	"templates/server_config.conf": templatesServer_configConf,
	"templates/server_peer.conf":   templatesServer_peerConf,
	"templates/templates.go":       templatesTemplatesGo,
	"templates/version.txt":        templatesVersionTxt,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"help.txt":           &bintree{templatesHelpTxt, map[string]*bintree{}},
		"server_config.conf": &bintree{templatesServer_configConf, map[string]*bintree{}},
		"server_peer.conf":   &bintree{templatesServer_peerConf, map[string]*bintree{}},
		"templates.go":       &bintree{templatesTemplatesGo, map[string]*bintree{}},
		"version.txt":        &bintree{templatesVersionTxt, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
