// Code generated by go-bindata.
// sources:
// templates/home.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _templatesHomeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x94\x4b\x6f\xdb\x38\x10\x80\xcf\xf6\xaf\x98\x55\xb0\x10\xe0\xb5\x1e\x4e\xb2\x79\xc8\x92\x7b\x48\x9a\x53\x81\x16\x68\x80\xa0\x47\x5a\x1a\x59\x8c\x45\x51\x25\xc7\xb1\x5d\xc3\xff\xbd\x24\x25\x3b\x4e\x9a\x43\x03\xc4\x1a\xce\x70\x3e\xce\x8b\xdc\xed\xa0\xc0\x92\x37\x08\x5e\x25\x05\x7a\xb0\xdf\x0f\xd3\x7f\xee\xbf\xde\x3d\xfe\xf8\xf6\x19\x2a\x12\xf5\x6c\x98\x1e\x3e\xc8\x8a\xd9\x70\x90\x0a\x24\x06\x79\xc5\x94\x46\xca\xbc\x15\x95\xc1\x8d\x77\xd4\x37\x4c\x60\xe6\xbd\x70\x5c\xb7\x52\x91\x07\xb9\x6c\x08\x1b\xb3\x6f\xcd\x0b\xaa\xb2\x02\x5f\x78\x8e\x81\x5b\x8c\x81\x37\x9c\x38\xab\x03\x9d\xb3\x1a\xb3\x89\xa3\xd4\xbc\x59\x42\xa5\xb0\xcc\xfc\x8a\xa8\xd5\x49\x14\x95\x86\xa1\xc3\x85\x94\x8b\x1a\x59\xcb\x75\x98\x4b\x11\xe5\x5a\x7f\x2a\x99\xe0\xf5\x36\x7b\x92\x6a\xf9\xdf\x77\xd6\xe8\xe4\x32\x8e\xc7\xd7\x71\xec\x83\xc2\x3a\xf3\x35\x6d\x6b\xd4\x15\x22\xf9\x40\xdb\x16\x33\x9f\x70\x43\xd6\xd3\xb7\x27\x39\xf3\x6c\x08\xe6\xcf\x66\x08\x3b\x98\xcb\x4d\xa0\xf9\x2f\xde\x2c\x12\x23\xab\x02\x55\x60\x54\x53\xb0\x01\x58\x03\x26\x57\xe7\xe1\xff\xff\x4e\x61\xef\xbc\x46\x63\x18\x25\x73\x2c\xa5\x42\x2b\xb1\x92\x50\xbd\xa3\xf0\xa6\x42\xc5\xc9\x7a\x38\x97\xb9\x2c\xb6\xb0\x73\x22\x9c\x60\x61\x12\xde\x28\x14\xd3\xe1\x60\x30\x88\x46\x4e\xdf\xe5\x96\x80\x6f\xb3\x03\x9b\x9d\x3f\x06\x6d\x3e\x81\x36\xc4\x72\x3a\x8a\x1c\xa5\xe7\x9e\xf5\x75\x3e\xb2\x05\xdb\x74\x55\x4e\xe0\x26\x8e\xdb\xcd\xf4\xa8\x57\x0b\xde\x24\x10\x03\x5b\x91\x9c\xf6\x0c\xfb\x6b\xab\x2a\x38\xbd\x8d\xee\x10\x85\x90\x8d\xd4\x2d\xcb\xf1\xc0\x59\x57\x9c\x30\x70\xaa\x04\x5a\x65\x5a\xaa\x58\xfb\xf6\x10\x53\x3b\x22\x29\x12\x38\x3f\x39\xbf\xe0\xba\xad\x99\x21\x06\x42\x07\x65\x8d\x1f\x18\xd6\x38\x5f\x72\xfa\xd8\xf8\xaa\xec\x83\xb6\xe3\x32\x86\xb0\x60\x64\x7a\x10\x9a\x9c\x2a\xa9\x4e\x8a\xe0\xe2\x50\x7c\x51\x91\x09\xc3\x55\xb8\x4f\xce\x70\x02\x5d\x29\xe3\x6d\x6a\xf1\x86\xd8\x33\xb8\x58\x1c\x39\x15\xfe\x49\xe8\x8b\xfb\xaa\xda\x0f\xa3\xd1\x60\x10\xea\x55\x9e\xa3\xd6\xc6\xd7\xf4\x32\x97\xb5\x54\x09\x9c\x5d\xdc\xdf\xde\x5e\xc7\xb6\xbd\x7b\xf3\x1f\xa2\x52\x2e\xca\x93\x1d\x0f\x0f\x97\x93\x8b\xab\x6e\x87\x69\xed\x20\x8d\xfa\xf1\x4c\xa3\xee\xde\xa5\x76\x74\xec\xb4\xa6\x05\x7f\x01\x5e\x64\x5e\xdf\x72\xaf\x1b\xe1\xb4\x9a\xcc\xbe\x30\x4d\xf0\x84\xb8\xf4\x35\xdc\xb9\x6e\x6a\x78\x94\x90\xe6\xb2\xc0\x99\x30\x46\x54\x69\xe4\x16\x86\x3a\xe9\xfd\x5e\x71\xce\xc1\x33\x36\xa3\x72\x27\x75\x82\x95\x74\xae\x78\x4b\xa0\x55\x9e\x79\x87\x6b\x69\x41\xe1\xf3\xcf\x15\xaa\xad\xbb\x91\x9d\x18\x9c\x87\x93\xf0\x32\x14\xbc\x09\x9f\x1d\xac\x73\x9d\xbd\xa7\x44\xcf\x3a\x12\xe6\xc9\x69\xe8\xef\xf6\x76\xeb\x77\xfb\xd2\xa8\xab\x8a\x49\xc7\xbd\x51\xe6\x39\xc3\xa6\xb0\x8f\xd8\xef\x00\x00\x00\xff\xff\x8e\xb0\xf0\x74\xd9\x04\x00\x00")

func templatesHomeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHomeHtml,
		"templates/home.html",
	)
}

func templatesHomeHtml() (*asset, error) {
	bytes, err := templatesHomeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/home.html", size: 1241, mode: os.FileMode(436), modTime: time.Unix(1445905164, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"templates/home.html": templatesHomeHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"home.html": &bintree{templatesHomeHtml, map[string]*bintree{
		}},
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
