// Code generated by go-bindata.
// sources:
// ui/index.html
// ui/style.css
// ui/ui.js
// DO NOT EDIT!

package ui

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _indexHtml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x4d\x6f\xdb\x30\x0c\xbd\xf7\x57\x70\x46\x0f\x0d\x10\x7f\xa4\x58\x87\x21\xb0\x73\x19\xb0\x5b\x81\x02\xcd\xce\x85\x26\x31\xb1\x36\x49\x36\x24\x25\x4d\x10\xf4\xbf\x8f\x96\x6c\x47\xe9\x69\xa7\x80\x7c\x8f\x8f\x7c\xa4\x9c\xfa\x8b\xe8\xb8\x3f\xf7\x08\xad\xd7\x6a\x73\x57\x4f\x3f\xc8\xc4\xe6\x0e\xa0\xd6\xe8\x19\xf0\x96\x59\x87\xbe\xc9\x0e\x7e\x97\x7f\xcf\xae\x80\x61\x1a\x9b\xec\x28\xf1\xbd\xef\xac\xcf\x80\x77\xc6\xa3\x21\xe2\xbb\x14\xbe\x6d\x04\x1e\x25\xc7\x3c\x04\x4b\x2d\x8d\xd4\x07\x9d\x3b\xce\x14\x36\xab\x25\x85\x5e\x32\x35\xc5\x51\xd5\x4b\xaf\x70\xf3\x22\x4f\x12\x79\x67\xb1\x2e\x63\x62\x80\x94\x34\x7f\xc1\xa2\x6a\x32\xe7\xcf\x0a\x5d\x8b\x48\x1d\x5b\x8b\xbb\x26\x2b\xdf\x4a\xe6\x68\x42\x57\x06\xac\xe0\xce\x45\x3d\xc7\xad\xec\x3d\x08\xdc\xa1\x05\x67\x79\x4a\x3d\xc8\xe2\x0f\xd1\xea\x32\x92\xc8\x75\x19\x6d\xd7\xbf\x3b\x71\x0e\xe5\xed\x2a\x9d\x85\xa2\x21\xd9\x9f\x72\xcd\x78\x2b\x0d\xe6\x4a\x3a\x3f\xe4\x00\x2e\x97\x1c\x2c\x33\x7b\x84\x7b\x02\x97\x70\x8f\x47\x5a\x84\x83\x75\x03\x05\x7c\x7c\x04\x4e\x52\x09\x52\x34\xd9\xe5\x12\xc8\x04\x67\x51\x84\x28\x42\x1e\x81\x2b\x9a\xb0\xc9\xba\x23\xda\x61\xb5\x33\x38\x28\x4c\x20\xd5\xe5\x4c\x08\x8b\x83\xd3\xab\x50\x5d\xf6\x09\x3b\x11\xeb\x6d\xb7\x8f\xe4\x19\x1e\x09\x61\x63\xe3\xc5\xd6\xe4\x03\x1e\xa4\x11\x78\x9a\x1d\x3c\x08\xe4\xf0\xa0\xd0\x4c\x99\xc5\x62\x51\xbc\x7a\xe6\xb1\x78\x19\x45\x83\x83\xba\x24\xb5\xa4\xf9\xa7\x70\x9e\xdc\x51\xe9\x21\x0e\xfd\xdf\x9d\x6e\x9d\xdd\x48\xa7\x26\x63\x55\x62\x31\xb9\x4a\xc0\x86\x73\x4c\xdd\xc6\xa3\x8c\xc3\x79\xa9\x91\x26\xd3\x7d\x38\xcb\x1c\xbd\x69\xa9\xe8\xc8\x63\x51\xb1\x9d\xf2\xe9\xcd\x82\x82\xeb\x99\x99\xc6\x98\xab\xe3\x65\x3e\x95\x16\xbf\xb6\x3f\x8a\x9f\x9d\xd5\xcc\x43\xf6\x58\x55\xdf\xf2\x6a\x95\x57\x8f\xb0\x7a\x5a\x57\x5f\xd7\xd5\x13\x3c\xbf\x6e\xb3\x60\x78\xd0\x4c\x9b\x5c\xc5\x9e\x69\xe9\x6c\x8f\x37\x1e\xd2\xcb\x0f\xbe\xd1\x88\x2b\x9e\xac\x8c\x88\xf3\x33\xbc\xbe\xdd\x99\x9d\xc2\xe3\xfb\xae\xcb\xf8\x41\xd0\x17\x10\xfe\x1d\xfe\x05\x00\x00\xff\xff\x97\xa6\x98\xc4\x35\x04\x00\x00"

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _styleCss = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\xd1\x6e\x9b\x30\x14\x7d\xcf\x57\x58\x9d\x26\xb5\x52\x5c\xd1\xb4\xe9\x5a\xfa\xb2\x69\x2f\x7b\xda\x3f\x18\xfb\x02\x77\x35\x36\xb2\x4d\xa0\x9b\xf2\xef\xb3\x81\x50\xa0\x90\xf6\x61\x0b\x0a\xc2\x36\x5c\x9f\x73\xee\xf1\xc9\x5d\x21\xb7\x24\xd1\xe2\x85\xfc\xd9\x10\xff\xab\x51\xb8\x3c\x26\x37\x51\xf4\xf9\xa9\x9d\xc8\x01\xb3\xdc\x8d\x67\x0a\x66\x32\x54\x31\x89\xba\x61\xc9\x84\x40\x95\x0d\xe3\x84\xf1\xe7\xcc\xe8\x4a\x89\x98\x7c\x82\xc7\x70\x75\x0b\xa9\x56\x8e\xa6\xac\x40\xf9\x12\x93\x6f\x06\x99\xdf\xf9\x07\xc8\x03\x38\xe4\x6c\x4b\x2c\x53\x96\x5a\x30\x98\x3e\x6d\x8e\x9b\x4d\x7e\xd3\x23\x72\xd0\x38\xca\x24\x66\x7e\x4b\x0e\xca\x81\x19\x95\xb3\xf8\x1b\x62\x72\x1f\x95\xee\x43\xd8\xba\x65\x9a\x68\xe7\x74\x11\x93\x5d\x54\x36\x4f\xcb\xb4\x27\x2c\xee\x38\x4b\xf7\x51\x0b\xab\x6c\x68\xc1\x78\x8e\x0a\xa8\x44\xeb\x7a\x8c\x02\x6d\x29\x99\xa7\x95\x4a\xe8\x2b\x86\x27\x2a\xd0\x00\x77\xa8\x03\x74\x2d\xab\x42\x75\x6b\x2d\x1b\x8a\x0e\x0a\x1b\x13\xeb\x0c\x38\x9e\x77\x2b\xbf\x2a\xeb\x30\x7d\xa1\xdc\x93\xf3\x5c\xa7\x8c\x0b\xd6\xd0\x1e\xe8\x43\xb4\x0e\xfd\xa4\x01\xab\x9c\x6e\x31\x5f\xeb\x03\x98\x03\x42\xbd\xd6\xe3\x44\x37\xd4\xe6\x4c\xe8\x3a\x26\xb7\x65\x33\xfc\x4d\x96\xb0\xcb\x68\xdb\x5e\xd7\xbb\xab\xb7\xca\xd4\xb9\x27\xb1\x2c\x6d\x80\x77\x56\x9b\x54\x86\xed\x8c\xae\x49\x6d\x58\xb9\x20\xcc\x98\xfa\x59\x5d\x86\x2e\xef\xc3\xa6\x13\xc2\x65\x4f\x79\xcd\x17\xfb\x93\x8a\x6f\x6d\xbe\xe4\xbb\x50\xdb\xb7\x9f\xfa\xcf\x0d\x58\xdb\x17\x9f\x18\xfb\xe2\xbb\xae\x0c\x82\x21\x3f\xa1\xbe\xd8\x92\x7e\xb4\x25\x85\x56\xda\x96\x8c\xc3\xc8\xbd\x75\xbf\x69\xa2\xa5\xe8\x8a\x97\x46\x67\xa3\xca\x7d\xa7\x76\xaf\xdd\x1e\x70\x0e\xc8\x97\x3d\x3e\xa2\x36\xbc\x70\x73\x9a\x49\xb4\x11\x60\xa8\x61\x02\x2b\x2f\xf4\x97\x61\x7e\x6c\x7a\xfe\x00\xf7\xfc\x71\x06\x4b\xe0\x61\x6e\xf9\x44\x6a\xfe\xfc\x2f\xb2\xe2\x74\xca\xc2\x42\x68\x60\x67\x90\x1c\x85\x00\xb5\x8e\x3b\xc0\xb3\x8e\xb9\x2a\x68\x16\x06\x70\xf0\xcd\x9a\x09\xf8\xf8\x7e\x9a\xbd\x9e\x96\xbe\xc0\x5b\xeb\x2c\xd8\xe6\x7f\x9d\x9c\x09\x0e\x6f\x1b\x35\xb6\xda\xd4\x36\xa3\x12\xe6\x44\xae\x6c\x42\x85\xaf\x05\x08\x64\xfe\x50\x49\xc2\x94\x20\x97\xf3\xfc\xb8\xea\x8b\x0e\x51\x3b\x8b\xd5\xbb\x21\x56\x8f\xed\xbd\xbd\x2d\xf8\x7f\x31\x53\x8e\xdd\xe9\x9f\x3b\x7a\xfa\xee\x49\xc6\xe3\x3b\x70\xf7\x1f\x81\x7b\x3b\xc0\x3d\xa3\xd4\x98\xc9\x2c\x16\x87\x64\x5a\x49\xed\xf1\xa7\xeb\xac\x76\x53\x56\x7f\x03\x00\x00\xff\xff\x84\x47\x3d\xe2\x61\x07\x00\x00"

func styleCssBytes() ([]byte, error) {
	return bindataRead(
		_styleCss,
		"style.css",
	)
}

func styleCss() (*asset, error) {
	bytes, err := styleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "style.css", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _uiJs = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x8f\xb1\x0a\xc2\x40\x10\x44\xfb\xfb\x8a\xb3\xba\x04\x4c\x7e\x20\x95\x88\x60\x91\x74\xfe\xc0\xb9\xb7\xc4\xe0\xe6\x4e\xb2\x1b\x08\x48\xfe\xdd\x8d\x06\x14\x2b\x1b\xb7\xdc\x99\x37\xc3\xb8\x91\xd1\xb2\x0c\x1d\x88\xab\x8c\x01\xf2\xcc\xb6\xf1\x70\xe9\x22\xd6\x1d\x8b\xc5\x49\x30\x06\xb6\xc7\x53\x53\x1f\x08\x7b\x8c\x62\xef\xc6\xea\x79\x11\xf5\x61\xd8\x7b\xa2\xb3\x87\x6b\x96\xaf\xc2\x72\x90\x22\x27\xc2\x92\x52\x9b\xb9\xdd\xd8\x2e\x1c\x86\x8d\xcb\xab\xa7\x65\x36\xf3\x57\xd9\x3f\x8b\x42\x82\x71\xf9\x97\x03\xb6\xba\x09\x87\x35\x3f\x73\xb7\xa9\xe8\x5f\xfd\x05\xa9\xe2\xb6\x9f\xdb\x35\xe2\x17\xf2\x0d\x29\xf0\x08\x00\x00\xff\xff\x36\xf4\x16\xc4\x50\x01\x00\x00"

func uiJsBytes() ([]byte, error) {
	return bindataRead(
		_uiJs,
		"ui.js",
	)
}

func uiJs() (*asset, error) {
	bytes, err := uiJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ui.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"index.html": indexHtml,
	"style.css": styleCss,
	"ui.js": uiJs,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"style.css": &bintree{styleCss, map[string]*bintree{}},
	"ui.js": &bintree{uiJs, map[string]*bintree{}},
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
