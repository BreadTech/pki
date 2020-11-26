// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// 0001_init.up.sql
// bindata.go
package migrations

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

var __0001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x51\xcd\x6a\x02\x31\x10\x3e\x27\x4f\x31\xcc\x49\x61\xdf\xa0\xa7\xb8\x3b\x4a\xa8\x26\x36\x66\x41\x4f\x25\x74\x43\x1b\x74\x53\x71\xa3\xe0\xdb\x97\x88\xed\xae\x2d\x6d\x3d\xe6\x9b\xc9\xf7\x33\x5f\x69\x48\x58\x02\x2b\x26\x73\x02\x39\x05\xa5\x2d\xd0\x5a\xae\xec\x0a\xd0\xc7\x14\x52\xf0\x1d\xc2\x88\x33\x0c\x0d\x42\x5d\xcb\x0a\x96\x46\x2e\x84\xd9\xc0\x23\x6d\x0a\xce\x30\xba\xd6\x23\x58\x5a\x5b\xa8\x95\x7c\xaa\x29\x83\xdb\x10\x1b\x04\xa9\x2c\xcd\xc8\x64\xa0\x71\x09\x61\x32\xd7\x13\x3e\x7e\xe0\xfc\x2f\xd5\x83\xdf\xb9\x14\xde\x63\xf7\x16\xf6\xdf\xa4\x33\xd1\xc5\xd4\xf9\x79\x08\xf5\x0e\x0a\xce\x06\xee\x60\x94\x7f\x16\x30\xf8\x33\x2e\x38\x9b\x6a\x43\x72\xa6\xae\x1b\x83\x19\x18\x9a\x92\x21\x55\xd2\x6d\x78\xbc\x0c\x39\x63\x5a\x41\x45\x73\xb2\x04\xa5\x58\x95\xa2\x22\xd0\x0a\xea\x65\x95\xb3\x28\x0d\xa2\xb4\x52\xab\x7f\x03\xfa\x93\x8f\xe9\xbe\x64\xc3\x5b\xdc\x0c\x5a\x9f\x5c\xe3\x92\xeb\x41\xe0\x0c\xdd\x4b\x5e\xfe\x3a\x05\xa6\xd0\xfa\x2e\xb9\x76\x8f\x90\x3d\x5a\xb9\xb8\x94\xd3\x85\xd7\xe8\x0f\x3d\x59\x7e\xbb\x74\x3c\xf8\x3b\x1b\xfa\x14\xff\x19\x61\xeb\xcf\xbd\xfa\xc9\xed\x8e\x57\xca\x5f\x7a\xc9\xeb\xe3\x2c\xf7\x11\x00\x00\xff\xff\x9c\x91\xe9\xba\x86\x02\x00\x00")

func _0001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_initUpSql,
		"0001_init.up.sql",
	)
}

func _0001_initUpSql() (*asset, error) {
	bytes, err := _0001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_init.up.sql", size: 646, mode: os.FileMode(420), modTime: time.Unix(1590828240, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1590828243, 0)}
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
	"0001_init.up.sql": _0001_initUpSql,
	"bindata.go":       bindataGo,
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
	"0001_init.up.sql": &bintree{_0001_initUpSql, map[string]*bintree{}},
	"bindata.go":       &bintree{bindataGo, map[string]*bintree{}},
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
