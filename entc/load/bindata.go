// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package load

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x91\x41\x6b\xdc\x30\x10\x85\xcf\x9a\x5f\x31\x15\x5b\x22\x81\xa3\xd0\x6b\x61\x4f\xcd\x1e\x52\x68\x52\xd8\x40\x0f\xdb\x25\xc8\xf6\x78\x23\x6a\xcb\xae\xa4\x94\x06\xa1\xff\x5e\x24\xd9\x0b\x3d\xd9\xf2\x7b\xfa\xe6\x3d\x4f\x8c\xd8\xd3\x60\x2c\x21\x9f\xb4\xb1\x1c\x53\x82\xbb\x3b\xfc\x32\xf7\x84\x17\xb2\xe4\x74\xa0\x1e\xdb\x77\xbc\x21\x1b\xba\xeb\xa7\x1b\x85\xf7\x4f\xf8\xf8\xf4\x8c\x87\xfb\x87\x67\x05\x8b\xee\x7e\xe9\x0b\x61\x66\x00\x98\x69\x99\x5d\x40\x01\x8c\xcf\x9e\x03\xe3\xed\x7b\xa0\xfc\x12\x23\x06\x9a\x96\x51\x07\x42\x5e\x5d\xbe\x8c\x2c\xd2\xe2\x8c\x0d\x03\xf2\x8f\xbf\x39\xaa\xef\x2b\x31\x25\x90\x00\x31\xe2\xae\xd5\x9e\xf0\xf3\x1e\xcb\x73\xd3\xf3\xdd\x3f\xda\xa1\xef\x5e\x69\xd2\x1e\xf7\x78\x3a\x93\x0d\xea\xc1\x06\x72\x83\xee\x28\x16\xb4\xd3\xf6\x42\xb8\x7b\x69\x70\x67\xf5\x54\x30\xea\x51\x4f\xe4\x33\x9f\xb1\x18\x6f\x57\x7e\x4a\x2a\x1f\xae\x51\x7c\x4c\x7c\xbd\x93\x52\x53\x58\x64\x7b\xbc\x4d\x09\x12\xc0\xf0\x66\xbb\xd2\x59\x48\x8c\xc0\x72\x90\xd1\x58\xf2\x78\x3a\x9f\xce\xb9\x34\xb0\x61\x76\xf8\xd2\xac\xf9\xf2\xdc\x1a\x65\xcb\x1b\x81\xb1\xb6\x41\x72\x2e\x6b\xdf\xb4\xf3\xaf\x7a\x3c\x16\x51\x54\x8f\x04\xc6\xcc\x50\x1c\x1f\xf6\x68\xcd\x58\xee\xb0\x41\x9b\x51\x90\x73\x59\xce\x15\xea\xdc\x3d\xea\x65\x21\xdb\x8b\x72\x6c\xb0\x95\x90\xd5\xd9\xab\x63\xe8\xe7\xb7\xa0\x7e\x38\x13\x48\x94\x7d\xa8\xaf\xb3\xb1\x9b\xb1\xc6\x15\xfc\xa7\xe5\x52\xca\x6b\xb7\x6d\x4a\x1e\x3f\xbb\x52\xb2\xb2\xc8\xb9\xca\x3a\x06\x67\xec\x25\x7b\xd4\x21\x7b\x84\x94\xc5\x73\xf8\x6b\x82\xf8\x54\x48\xff\x6d\xbd\x96\xaa\x4b\x5f\x7f\x66\x4a\xf0\x2f\x00\x00\xff\xff\xb5\xb1\x2f\xf6\x87\x02\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 647, mode: os.FileMode(420), modTime: time.Unix(1566121519, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xdf\x6f\xdb\xb6\x13\x7f\xb6\xfe\x8a\xfb\x1a\x68\x20\x15\x86\xf3\xae\xc2\x0f\x5f\x6c\x19\x96\x0d\x4b\x87\x76\xdb\x4b\x51\xb4\x8c\x74\x72\xd8\xca\x94\x4b\xd2\x4e\xdb\xc0\xff\xfb\x70\x3f\x68\x49\x96\x9b\x01\x2d\x52\xa0\x88\x78\xbf\x78\xf7\xb9\x0f\x8f\xf4\xd6\x54\x1f\xcd\x1a\xa1\xed\x4c\x9d\x65\x76\xb3\xed\x7c\x84\x3c\x9b\xcd\xd1\x55\x5d\x6d\xdd\xfa\xf2\x43\xe8\xdc\x3c\x9b\xcd\x9b\xdb\xea\x12\x5d\x1c\x7c\x5e\x86\xea\x0e\x37\xe6\xb2\xb1\xd8\xd6\x2c\xdf\xb0\xda\x63\xd3\x62\x15\xe7\x59\x91\x65\x97\x97\xf0\x9a\xad\xc0\xe3\xd6\x63\x40\x17\x03\x18\x07\xe8\xe2\x52\x15\xf1\xce\x44\xb8\x37\x81\x73\xc0\x1a\x1a\xdf\x6d\xc0\x40\xd5\x6d\xb6\xad\xc5\x1a\x76\x01\x3d\x68\x9e\xcb\x2c\x7e\xd9\x62\x0a\x19\xa2\xdf\x55\x11\x1e\xb2\xd9\x8d\xd9\x20\x00\x90\xc4\xba\x35\x00\xbc\xa7\xb4\xcb\xb9\x33\x1b\x5c\x74\x1b\x1b\x71\xb3\x8d\x5f\xe6\xef\xb3\xd9\x55\xbd\xc6\x00\x00\x6f\xde\x3e\xa7\xcf\xa3\x25\x92\x7c\x6c\xfa\x0b\xd5\x15\xd8\x94\x3f\x93\x29\xd7\x7b\x62\x7b\xed\x6a\xfc\x8c\x81\x6c\xf9\x33\xd9\x5a\x91\x8f\x8c\x0f\x0c\x8b\x84\x9c\xa2\x22\xf2\xef\x00\x45\x1c\xa7\x98\x0c\x61\x79\x0c\x99\xbf\x28\x88\xfc\xe3\x0a\x97\x22\x50\x7b\xda\xe2\xc4\xde\x68\xc4\xb3\xf1\xa3\x59\x8f\xcd\x5f\xdb\xaf\x29\xfc\x73\xeb\x62\x72\x55\xf3\x60\xbf\x9e\x84\xff\xe9\xce\xf8\x80\x6c\xf7\x7c\x10\x5f\xed\x2b\xd1\x8e\x5d\xfe\x76\xf6\xd3\x4e\x36\xb9\xed\xba\xf6\x64\x8b\x1d\x6b\xc7\x1e\x37\xb6\x6d\xcd\x6d\x8b\xe7\x3d\x9c\x6a\xc7\x3e\x2f\xb7\xd1\x76\xce\xb4\xe7\x7d\x3a\xd5\x8e\x7d\x7e\xc6\xc6\xec\xda\xf8\x8d\xcc\x6a\xd1\x8e\x5d\xfe\x31\xad\xd6\x02\xd6\x45\xf4\x8d\xa9\xf0\xe1\x90\x5c\xf6\xa4\x9d\x38\xd8\xda\xc4\xce\x07\x18\xe0\x3b\x70\x50\xed\x19\x2e\xf2\x49\x98\x52\x91\xc5\xdf\xc1\x44\xf6\x3b\x43\x44\x6d\xe3\x7f\x33\x70\x6c\xf8\x08\xf5\x4e\x0c\x4f\x39\xf7\x0a\x1b\xd9\x7c\x6c\xe7\xb1\x79\x37\xdd\xfd\x15\x36\x4a\xcf\xd1\x60\xf0\xd8\x7c\x83\x65\xda\xc9\x47\xe8\x75\xed\xf6\xe8\x03\x9e\x9a\x5a\x11\x9f\x6e\xff\x69\x67\x3d\xd6\x27\xb6\x5e\xc5\x67\xba\x26\x83\x66\xda\x36\x91\x7f\x47\xdf\xc4\xb1\x6f\x9c\x56\x7a\x64\xec\x23\x95\x0a\x64\x93\xf9\x4b\x53\xf5\xec\x50\x7d\xf3\x76\xdc\x92\x33\x33\x55\x8a\xbc\xc1\x7b\x8e\x5d\x79\x34\x11\xb9\x46\x2d\x88\x62\x4b\x55\x72\x0f\xf5\xc7\x64\x99\x35\x3b\x57\x25\xd7\x1c\x8f\x64\x2e\xb4\xb7\x0f\xd9\xcc\x21\x94\x2b\xb8\xa0\xe5\x43\x36\x63\x8e\x96\x5c\x24\x2e\xe9\x3b\x2f\x16\xd9\x8c\x09\x99\xa4\xf4\xad\x52\xb3\x16\x21\x49\xcd\x5a\x84\x4a\xb5\x92\x84\xfa\x2d\x0a\x01\xb1\x64\xeb\xeb\x20\x2b\xd1\x28\x3b\x4a\xd1\xe8\x2a\x45\x93\xae\x97\xac\x4a\x2b\xd6\x1d\xb2\x99\x6d\x80\xb3\xc7\xe5\xff\x43\xe8\xaa\xbc\x78\x01\x08\xff\x5b\x81\xb3\x2d\x95\x36\x73\x9c\x02\xac\x7a\x04\x0a\xf6\xf3\x18\x77\xde\x81\x43\xc5\xf6\x0f\xe3\xc3\x9d\x69\xf5\x36\xe5\x1b\x9f\xce\x3b\x0e\x6f\xe7\x23\xa8\xf4\xd5\x81\x81\xdf\x5e\xbf\xbc\x21\x67\xe6\x57\x65\x1c\xdc\x22\xd4\x48\xae\xb5\x98\x50\x00\x75\xee\x6e\x3f\x60\x15\xf5\x8f\x76\x65\xb4\x69\x1e\xd2\xde\x44\x5b\xdd\xa9\x80\xfc\x16\xde\xbc\xbd\xfd\x12\x71\x01\xe8\x3d\xfd\xef\x7c\x41\xa5\x05\x6e\x9a\xf8\x3e\x08\xde\xd6\xd5\xd6\x63\x15\x73\x7d\x76\x70\xa3\x5e\x36\x1a\xb9\x28\xb4\x9d\x87\x6c\xa6\x1c\xe3\x90\xe5\x0a\x82\x69\x50\xd8\x98\x6c\x05\x59\xef\x87\x58\x26\xcc\x6c\xbb\x80\x66\x13\x97\x57\x94\x4b\x93\xcf\x35\xf1\x67\x9f\x4a\x78\xb6\x9f\x2f\x20\xf0\x3e\x1c\x5c\xc0\x6e\x3a\x0f\xef\x16\xd0\xd0\x56\xde\x38\xe2\xaa\x50\x9f\xa2\x06\x16\x5f\xf0\xf6\xb4\x1e\xf0\x0f\xa0\x19\x30\x70\x40\x41\x52\xf4\x24\x1c\xb0\x90\x15\x89\x87\x72\x6f\x94\x49\xce\x2b\xd5\xf4\x4c\x24\xcd\x98\x8b\xe9\x86\x2a\x45\xf9\xab\x09\x2a\x50\x75\xba\x28\x4b\xf5\x4d\x6b\x55\xa7\x3b\x31\xa9\xd3\xba\xcf\x49\x2f\x9f\x12\x5a\x74\x39\xe7\xa5\x92\xbc\x60\x9b\x43\x36\x23\xf4\xc3\x02\xba\x8f\x84\x4d\xb3\xcc\xe5\x15\x42\xef\x06\x5f\xbc\x20\x31\x23\x45\x2f\x05\x6e\x1f\x6b\xf2\x82\x65\x0d\x2f\x60\x05\x17\xa4\xee\xc3\x55\x93\x70\xfa\xac\x18\x85\xd4\xc7\x04\xd9\x55\xc9\xe0\x18\x38\xbd\x43\x56\x70\xa1\x76\x1a\x3e\x2c\x75\x96\xad\xc0\x6c\xb7\xe8\xea\x3c\x49\x16\x10\x1a\x21\x81\xbc\x2a\x87\x8c\xe3\xf7\xe7\x53\x12\x0e\x7b\xc2\xf1\xee\xc2\xb7\xa5\xbc\x7b\x07\xa9\x5e\x49\x6a\xfd\x88\x90\x28\xe9\xc1\x3a\xcc\x59\x1f\xb7\x4f\x99\xb5\xad\x3f\xf7\x79\x6b\x0e\x9a\x79\x7a\x5a\x0f\x72\xbf\x4e\x49\x5e\xf0\x17\x37\x91\xaa\x20\xfa\xd9\xfa\x33\xd7\xa6\xd4\x93\x8e\x94\x2c\xd6\xd3\x3e\x3e\x0d\xa4\x18\x9f\x85\xc3\x68\x5a\xd2\xfd\xb4\xd4\xa1\x95\x87\x42\x47\x67\x3f\x3c\xe0\xde\x9b\x6d\xe0\xa9\x27\xc5\x26\x5a\x6c\x30\xde\x75\x35\xdc\xdb\x78\x07\x1e\xab\x6e\x8f\x1e\x62\x07\xe8\xc2\xce\x23\xb8\x0e\xb6\xc6\xd9\x8a\x1e\x6b\xb0\x91\xf0\xd6\xad\x75\x48\x4e\x66\xd3\x64\x42\x36\xe9\x1e\x3d\xfe\x66\x38\x9d\x95\x35\x36\xe8\x81\xc2\xe5\xbc\xa6\xae\xed\x19\x64\x49\x86\xae\x8d\xfd\xb0\x87\x33\xf2\x5f\x9d\x69\x5f\xaa\x48\x12\xd6\x4e\xee\xf9\x80\x34\xe9\x08\x38\xdb\xca\xb9\x38\xd0\xc9\x51\xec\x46\xee\x79\xb1\x60\xab\x1e\x40\xe1\xe4\x04\x3f\x11\xff\x28\x7c\xc3\x83\x36\x41\x4f\x4e\x86\x80\x47\x86\x4f\x88\x9d\x54\x73\x06\x3a\xd4\x13\xf9\x18\x72\x52\xc4\x04\xb8\x74\x24\x26\xd0\x25\xc5\x8f\x82\x37\x3e\xf1\x13\xf8\xec\xf1\xd7\xee\xf1\xbd\xf9\x84\x08\xa6\xa2\xce\x60\x68\x8f\xb3\xe1\x31\x14\x53\x35\x3d\x8e\x5c\xe8\xf1\xe5\x10\x61\xf8\x76\x28\x46\x2b\xca\x8d\x66\x54\x5c\xfe\x6e\x5d\x9d\x17\xb0\x5a\x1d\xf5\x7f\x46\xcf\xa9\xd3\xe5\x10\x97\x57\x2d\x6e\xf2\xd1\xe8\x88\xd9\x21\xfb\x37\x00\x00\xff\xff\xdb\x1a\x88\xcc\x61\x11\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 4449, mode: os.FileMode(420), modTime: time.Unix(1566124960, 0)}
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
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
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
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
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