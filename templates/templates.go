// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/client_peer.conf
// templates/help.txt
// templates/server_config.conf
// templates/server_peer.conf
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

var _templatesClient_peerConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xce\xcf\x6a\xc3\x30\x0c\x06\xf0\xbb\x9e\xc2\x2f\xd0\x0c\xb2\xed\x32\xe8\x61\xac\x3b\x94\x8e\x62\xe8\xb1\xf4\x90\xd5\x5f\xc0\xcc\xd8\x41\xf6\x12\x86\xd0\xbb\x8f\xfc\x21\xe9\x51\xbf\x4f\x42\xdf\xf5\x18\x0b\xb8\x6d\xee\xb8\x91\x65\xdf\x37\x05\x27\xfc\x99\xbd\x11\xa9\x3e\x18\x2e\x57\x9b\xaa\xd2\xbb\x73\x8c\x9c\xe7\x7c\x19\x54\x9f\x9e\x6b\x12\xf1\xad\xa9\x0e\x31\x5f\xc0\x3d\x58\xf5\x70\xbe\xcc\x6b\x0f\x26\x82\xe8\x54\x89\xae\x16\xe0\x1b\xd9\xdf\xef\xe0\xef\xeb\xc3\x11\x57\x52\xa5\xcf\xe8\xba\xe4\x63\xd9\xd2\x55\x76\xaa\x6f\x22\x3b\x33\xe9\x97\xcf\x05\xd1\x26\x2e\x63\xc5\x10\xd2\x00\x77\xb4\x4b\xcb\x33\xca\x90\xf8\x67\x2b\x5b\xbf\x90\x05\xe7\xe9\xa6\x9c\x80\xae\x09\xbe\x87\xd9\x9b\xfa\xf5\x3f\x00\x00\xff\xff\x57\x7b\x39\x30\x0e\x01\x00\x00")

func templatesClient_peerConfBytes() ([]byte, error) {
	return bindataRead(
		_templatesClient_peerConf,
		"templates/client_peer.conf",
	)
}

func templatesClient_peerConf() (*asset, error) {
	bytes, err := templatesClient_peerConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/client_peer.conf", size: 270, mode: os.FileMode(420), modTime: time.Unix(1689613889, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
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

var _templatesServer_configConf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xcf\x4a\xc4\x30\x10\x87\xef\x79\x8a\x39\x78\x6d\xc4\x3f\x78\x10\x7a\x28\x55\x61\x51\x34\xac\x8a\x87\x65\x91\x6c\x32\x5d\x22\x31\x2d\x93\xd8\xb2\x84\x79\x77\x59\x6d\x5d\x7b\xd8\x5b\x32\xbf\x6f\x86\x99\x6f\xb5\x08\x09\xa9\xd1\x06\xd7\x42\x91\xeb\x75\xc2\x7b\xdc\x41\x09\x39\xcb\x67\xa4\x1e\x49\xd6\x84\x36\xca\x43\xc8\x2c\x2a\x6b\x09\x63\xfc\xc5\xc6\x0f\xf3\xe9\xc5\xb9\x78\x70\x31\x61\x50\x2d\xa5\xd9\x8c\x43\x99\x59\x28\xc2\xd7\x0e\x4a\x88\xbb\x68\x92\x87\x62\x80\x80\x49\xba\xae\xbf\x94\xae\x7b\x6f\x5a\x1a\x34\xd9\xf2\xec\x38\x77\x25\x4d\x1b\x1a\xa9\xbd\x97\x23\xed\xc2\xf6\x5f\x83\xeb\x92\xde\x78\x8c\x50\x54\x70\xf7\xb4\x7c\xab\x96\x37\x50\xb8\xfd\x3a\x7f\xd7\x3e\xea\x4f\x64\x86\xe2\x03\xaa\xba\xbe\x55\x2f\x22\x67\xd2\x61\x8b\x70\x62\xbc\xc3\x90\xe0\xba\x04\x59\xff\x3c\x23\xb3\x58\x29\x44\x5a\x0b\xf5\xb5\xf1\xce\x4c\x82\x46\x74\x12\x34\x65\x7b\x3f\xde\xb7\x03\xda\x85\x8a\x33\x70\x66\x2a\x67\x0c\x96\xf9\x3b\x00\x00\xff\xff\xcf\xbe\x1a\x35\x81\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/server_config.conf", size: 385, mode: os.FileMode(420), modTime: time.Unix(1689630847, 0)}
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

var _templatesVersionTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4f\x2f\x53\x28\x4b\x2d\x2a\xce\xcc\xcf\x53\x28\x33\xd4\x33\xd0\x33\x04\x04\x00\x00\xff\xff\xdb\x85\x49\x6d\x12\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/version.txt", size: 18, mode: os.FileMode(420), modTime: time.Unix(1689637116, 0)}
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
	"templates/client_peer.conf":   templatesClient_peerConf,
	"templates/help.txt":           templatesHelpTxt,
	"templates/server_config.conf": templatesServer_configConf,
	"templates/server_peer.conf":   templatesServer_peerConf,
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
		"client_peer.conf":   &bintree{templatesClient_peerConf, map[string]*bintree{}},
		"help.txt":           &bintree{templatesHelpTxt, map[string]*bintree{}},
		"server_config.conf": &bintree{templatesServer_configConf, map[string]*bintree{}},
		"server_peer.conf":   &bintree{templatesServer_peerConf, map[string]*bintree{}},
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
