// Code generated by go-bindata.
// sources:
// templates/cf_lb.tf
// templates/network.tf
// templates/network_security_group.tf
// templates/output.tf
// templates/resource_group.tf
// templates/storage.tf
// templates/storage.tf.orig
// templates/template
// templates/tls.tf
// templates/vars.tf
// DO NOT EDIT!

package azure

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

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xcb\x6e\xeb\x36\x13\x5e\xff\x7a\x8a\x01\xff\x6e\xe5\x24\x4d\x70\xd0\x8d\x37\x6d\x17\x67\x59\x20\x28\xba\x28\x0a\x62\x24\x8d\x2d\x22\x34\xc9\x43\x8e\x9c\xb8\x07\x7a\xf7\x82\xba\xd8\x96\x44\xb9\x49\xea\x9d\x87\xdf\xdc\xbe\xb9\x68\x8e\xe8\x15\x16\x9a\x40\x84\x53\x60\x3a\xc8\xca\x1e\x50\x19\x01\xdf\x33\x00\x3e\x39\x82\x2d\x88\xc0\x5e\x99\xbd\xc8\x00\x2a\xda\x61\xa3\x39\x0a\xe9\x0d\x0f\x4e\xd3\xa6\xb4\x07\x91\xb5\x59\x76\xb1\xe4\x76\x6f\xb2\x24\xcf\xb2\xc0\x40\x5f\x9e\xd2\xb6\x16\x1a\x0e\x43\x78\xb5\xbe\x5a\x85\xff\x1f\x7e\xf1\x84\x4c\x80\x80\xce\x69\x55\x22\x2b\x6b\x60\x8f\x4c\xaf\x78\x02\x65\x80\x6b\x82\x57\x2a\x64\x20\x7f\x24\x1f\xc0\x53\xb0\x8d\x2f\x09\xf6\xde\x36\x2e\x3b\xff\x15\xf8\x77\xe3\xc9\x1f\xe4\x51\x79\x6e\x50\x4b\x43\xfc\x6a\xfd\x8b\x00\x51\xee\xf2\xe3\x90\xbd\xc1\x03\xc1\xec\xb7\x05\xf1\xc3\xf7\x23\xfa\x0d\x99\xa3\x54\x55\x9b\x77\x78\xe2\xc8\xcd\x68\x5e\x76\xde\x64\xa7\xde\xe1\x47\x6f\x53\xc0\xa6\xb0\xa1\xde\x44\x54\x1b\xb5\xb1\xaa\x3c\x85\x20\x83\xc3\x92\xce\xde\xfe\x1c\xdc\x0d\x01\xca\x52\x55\xbe\x15\x7f\x65\x00\xda\x0e\xf9\xa7\xa2\xf3\xb4\x57\xd6\xb4\x1d\x6d\xcb\xac\x43\x53\xc4\x90\x41\x84\xa6\x78\x58\xcf\x35\x99\x6c\xaf\xfb\xb0\x96\xef\x07\x12\x9e\x71\x9f\xe0\x6b\x86\xd8\x74\x5c\x2f\x19\x73\x9e\x76\xea\x6d\x1e\xb3\x32\x4c\xde\xa0\x1e\x28\x4b\x33\xe1\x9a\x42\xab\x52\x2a\x27\x40\xe8\xe2\x06\x15\xb7\x28\xd1\x45\xae\x9c\x58\x29\xc9\x8d\xda\xac\x30\x38\x51\x78\x17\x93\xe7\x2c\xe4\x48\x09\xea\x73\x2c\x5b\x10\xd5\xc9\xe0\x41\x95\xb3\x21\x32\xa9\x29\x4a\x70\x74\x85\x92\x03\x4a\x80\x38\x4f\xcc\xbb\x27\x05\x9d\xcb\x47\xfd\xff\x3c\x2d\x2b\xdd\xff\x07\x05\x86\xdf\x9f\x45\x06\x19\x40\x78\x69\xba\xe8\x16\xf1\x6d\x41\x3c\x33\x9a\x0a\x7d\x25\x9f\x0f\xa8\xb5\xe8\x50\xac\xc8\xa7\x51\xfd\x7b\x89\x0e\x4b\xc5\xa7\xf3\xfb\x8f\x19\x40\x9b\xc5\x02\x78\x5b\xd0\xb5\xaf\x2d\x88\xdf\xa2\xec\xfe\xa1\x57\x75\xde\xb2\x2d\xad\x8e\x0f\x5f\x99\xdd\x20\x45\xae\xa3\xe4\x4e\xdb\xbd\x32\xbd\xac\xb6\xa1\xdb\xae\x9d\x68\xd3\x33\x38\xd9\xcc\x6d\x8f\xeb\xda\xfb\x88\xd1\xe2\x97\xfb\x21\xfc\x03\xd9\x86\x2f\x82\xc6\xd4\x84\x9a\xeb\x93\xe4\xda\x53\xa8\xad\xae\x60\x0b\x8f\x5d\xd0\x91\x9f\xa1\x18\xb1\x6f\x4a\x6b\x76\x6a\xdf\xf8\x9e\xd3\x04\x69\xa9\xbe\x1f\xf4\x73\xe5\xf2\x89\x7e\x1f\x60\xbf\x28\xa4\xaa\x16\xbd\x9c\x1e\x6b\x55\xb5\x77\xbd\x4e\xb8\xbb\x60\x7b\xc9\x26\x2e\xaa\x4b\xf1\xfb\xf0\x77\xde\x1a\x26\x53\x49\x67\x3d\xaf\xc4\x3c\x62\x22\xa4\x66\x76\x61\x20\x3e\x6a\x5c\x60\x4f\x4f\x8f\x63\x25\x3f\x6a\x54\xdb\xfd\x9a\xcd\xa7\xc7\x65\xa8\x9f\xa5\x7a\x34\xb0\xc2\xf5\x72\x01\xa8\x6a\xca\xf9\x19\xb1\xd1\x45\xa4\x5a\x8c\x09\x17\x58\xbe\xc4\xc8\xce\xcb\xd4\x5a\x3d\xeb\xe4\x45\x34\x83\x4e\x3e\xe8\xe4\x51\xe7\x52\x97\xd1\x62\xe4\x5b\x06\x62\x56\x66\x1f\xd2\x83\x98\xda\x74\x37\x96\x7e\x5e\x50\x5e\x73\xe0\x61\x1c\xad\x7d\x51\xd4\xdd\x18\x95\xc4\xdd\x4e\x99\x38\x9b\x5b\x10\xbf\xaa\x10\xaf\x8a\x2a\x51\x97\x2b\x97\x3f\xdd\x4f\x27\x73\x1e\xd1\x65\x4c\x3d\x7d\x6b\x28\xb0\x1c\x07\x6c\x84\x3c\x8c\x06\x0a\x92\xf3\xc4\xa6\x0b\xa0\x63\x3a\x04\xdd\x5d\x45\x6a\x17\xb7\xe9\x7c\x5d\x84\xa0\xf3\xf8\xda\xbb\xac\x90\xf1\x42\x7d\xbc\x8e\xfa\x53\xaa\x1d\x17\x47\x7f\x2a\x4d\x21\xa3\xf4\x6a\x46\xba\x1a\x68\x15\x98\x0c\xf9\x9b\x35\xf8\x74\x4d\xa2\x0b\x1d\x78\xe8\xc4\xd5\x4e\x9f\x30\xf4\x89\x16\x9f\x0c\xe6\x82\xee\x2b\xc3\x6b\x23\x9f\xac\x73\x42\xff\xeb\x45\x67\x56\xb1\x15\xaf\xb3\xe2\xf5\xc4\x8f\x3d\xe3\x6d\x13\xfb\x5f\xfa\x46\xd3\xbf\xf3\xff\x6e\xd2\xfd\xb7\xb1\x53\xa2\x61\xd9\x1d\xcb\x29\x6b\x3f\x63\x88\xdf\xfd\xf8\x6f\xd2\x0a\x72\xb9\x71\x3e\x5e\xeb\xd4\xea\xb8\xbe\x02\xdf\xb9\x35\x56\x36\xc6\xc7\xee\xc1\xeb\xd5\xd0\xc6\x3b\xc7\x36\xec\x1a\x06\x81\xce\x8d\x57\x4b\x67\x31\x1e\x2c\xff\x3b\xa2\x6e\x66\xb6\x13\x57\xce\x78\x74\x8f\x1f\x9e\xf6\x9f\x00\x00\x00\xff\xff\x2e\x9d\xc3\x44\x2c\x0d\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 3372, mode: os.FileMode(480), modTime: time.Unix(1512068353, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetworkTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x5d\x4e\x86\x30\x10\x45\xdf\xbb\x8a\xc9\xc4\x57\xd8\x81\x2b\x31\xa6\x29\xed\x88\x8d\xd0\x92\xe9\x8f\x46\xd2\xbd\x9b\x12\x30\xa1\x42\xfc\xfa\x7c\x66\x7a\xcf\x1d\xa6\xe0\x13\x6b\x02\x54\xdf\x89\x89\x67\x99\x2d\xc7\xa4\x26\xe9\x28\x7e\x7a\xfe\x40\xc0\xc1\x87\x77\x84\x55\x00\x38\x35\x13\x34\xef\x19\xf0\x69\xcd\x8a\x7b\x72\x59\x5a\x53\xba\x8a\x77\xd9\xa1\x00\x50\xc6\x30\x85\x20\xc3\xa2\x34\xfd\xf2\x2f\xfb\xc0\xfe\x83\xd4\xd6\x70\xc1\x57\x01\x30\x79\xad\xa2\xf5\xee\x72\x3f\xd3\x68\xbd\x2b\x75\xef\x91\x5a\x8e\xec\xd3\x22\xb7\x58\x1b\x77\x48\x9c\x81\xbe\x46\xea\x2b\x55\x50\x14\x21\xfe\x4a\x87\x34\x38\x8a\xff\xba\xde\xc8\x86\x93\xec\xc2\xf4\x66\xbf\xda\x81\xb3\xec\x8d\xc3\xc3\x12\x00\xcd\x99\x2e\x3a\x68\x88\xa6\x84\x9f\x00\x00\x00\xff\xff\x2b\x4f\xd2\x88\xf9\x01\x00\x00")

func templatesNetworkTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetworkTf,
		"templates/network.tf",
	)
}

func templatesNetworkTf() (*asset, error) {
	bytes, err := templatesNetworkTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network.tf", size: 505, mode: os.FileMode(480), modTime: time.Unix(1512068137, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetwork_security_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\xcd\x6e\x9b\x40\x10\x80\xef\x3c\xc5\x68\xd5\x53\x24\x5b\x29\x81\xc8\x97\x1c\x7a\xec\xbd\x77\xb4\xde\x1d\xf0\xaa\x78\x07\xcd\x2e\x4e\xdb\x88\x77\xaf\x16\x6a\x27\x38\x38\xc2\xa8\x55\x85\x05\xe7\x6f\xfe\x76\x3e\x8d\x60\x74\x54\xb3\x42\x10\xf2\x57\xcd\xc8\xfb\xcc\xa2\x7f\x26\xfe\x9e\x39\x54\x35\x1b\xff\x33\x2b\x98\xea\x4a\x80\xd8\x92\xdb\x09\x78\x89\x00\xac\xdc\x23\x9c\x7d\x4f\x20\x3e\xbd\x1c\x24\xaf\xd1\x1e\x32\xa3\x9b\x55\x8b\x47\x00\x25\x29\xe9\x0d\xd9\x41\x98\xb1\x30\x64\x9b\xc0\x1d\x3b\xe9\xea\x65\x6d\x8d\x96\x3b\x36\xd6\x07\xd6\x21\xff\x3a\x50\x8d\x88\x22\x00\x2f\x0b\xd7\x36\x07\x80\xf6\x60\x98\xec\x1e\xad\x7f\xd7\x56\xa8\xd4\x44\x4d\x14\x5d\x31\xb8\xca\xaf\x18\x5b\xe5\xf3\x1e\x9a\xeb\x12\x05\x08\xf7\xd1\xae\x2f\x0e\xef\xba\x95\x57\x6c\x28\x24\x1b\x8e\x89\xef\xef\x23\x00\x6d\x18\xd5\xf9\x13\xbd\xe6\xfd\x6a\xb7\x54\x5b\x1d\xb2\x49\xa5\xd0\xb9\x8b\x1d\x7c\x29\x4b\x7a\xee\xaa\x92\x27\x45\xe5\x05\xee\x9b\xaa\x02\xf5\xe7\x39\x2b\x62\x9f\xb1\xb4\x05\xf6\xa9\xbb\xc0\x68\x74\xde\xd8\x76\x81\xef\xc0\x27\x10\x71\xfc\x26\x91\xd4\x9a\xd1\xb9\xac\x62\xcc\xcd\x8f\x0f\x12\x9d\x83\x47\x66\x48\x81\xde\x03\x8f\x50\x01\x60\x58\xde\x01\xa1\x86\xc1\x5e\xb6\xab\x44\x09\x81\x2b\x59\xa0\xf5\x13\x7c\x79\x13\x3c\x42\x9b\xcf\xf3\xd6\xe6\x71\xf3\xb8\x59\xc4\xe9\x8b\xd3\xad\x93\x78\xaa\x3b\xa7\xf8\x11\xfa\xc4\xf3\xd6\x27\x4e\xd3\x34\x5d\xfc\x39\xf9\xa3\xad\x9b\x60\x4d\x88\x1a\xe1\xca\xc3\xff\x70\xe5\xee\x2f\x99\x92\x3e\x2c\x9a\x9c\x34\x51\x8c\x7a\x57\x6f\x27\xa8\x72\x8c\x1c\xa1\x4b\x32\xef\xd3\xb2\xd9\x24\xc9\xa2\xcc\xab\x32\xf9\x6a\xe7\x7d\xf5\x0f\xcf\xcb\xcc\xff\x64\x92\xe4\x06\x2f\x8c\xca\xa7\xca\x52\x52\x31\xe5\xbc\x74\x81\xb7\xff\xe3\x92\xdc\xbc\x2e\xbf\x03\x00\x00\xff\xff\x77\x27\x82\x78\x45\x11\x00\x00")

func templatesNetwork_security_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetwork_security_groupTf,
		"templates/network_security_group.tf",
	)
}

func templatesNetwork_security_groupTf() (*asset, error) {
	bytes, err := templatesNetwork_security_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network_security_group.tf", size: 4421, mode: os.FileMode(480), modTime: time.Unix(1512068137, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutputTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xdd\x6e\xa3\x30\x10\x85\xef\x79\x0a\x0b\xed\x35\x95\x22\x71\x13\x69\x9f\xc5\x32\x66\xda\xcc\xd6\xd8\xd6\xfc\xb0\xed\x56\x79\xf7\x15\x4b\x89\xe2\x26\x40\xb6\xb9\x0c\xdf\x7c\xe7\x68\xc6\x49\x25\xab\x98\xba\x4b\x7c\xb2\x11\xe4\x77\xa2\x57\x1b\xdd\x00\xb5\xf9\xa8\x8c\x19\x5d\x50\x30\x3f\x4d\xfd\xe3\xc3\xfd\x51\x02\x1a\xec\x88\x24\xea\xc2\x02\x37\xd3\x64\x33\x4d\x9c\xeb\xea\x5c\x55\x85\x90\xb5\x8b\x20\xdb\xbe\x99\xd9\xd4\x10\x70\x52\xf2\x60\x5f\x28\x69\xde\xd6\x95\xec\x76\x3b\x49\xe4\x5e\xc0\x3a\xef\x93\xc6\xbd\x9a\x25\xbc\x29\xee\xe1\xd9\x69\x10\xcb\xe0\x95\x50\xde\xe7\x2e\xab\xea\x65\xed\x25\xbe\x96\x00\x6f\x02\x14\x5d\xb0\xb8\x6e\xcc\xda\x05\xf4\x16\x3f\x25\x98\xad\xeb\x7b\x02\xe6\x52\xd5\x23\x81\x97\x44\xcb\xd7\x2f\xbe\x93\x48\xe6\xe3\xd3\xd3\x23\xde\xe3\xa1\x6d\xdb\xf6\x76\x15\xe3\xc0\x36\x13\x8e\x4e\xc0\xbe\xc2\xfb\x75\xc2\xf4\xfb\xd7\x5a\x42\xc1\x34\xcb\x60\x73\xf5\xa7\xcd\x30\x9c\xeb\xca\x18\x86\xc8\x28\x38\x4e\x0d\x85\x14\xee\x27\xce\x3d\xff\x3f\xf0\x32\x67\x53\x86\xc8\x7c\xba\xc9\x7c\x76\x81\x8b\xd0\x5f\x3a\xe4\x2e\xbd\x59\xa5\xf0\x8d\x7b\x1c\x0f\x87\x62\x69\xcb\x5b\xf0\xd8\xd3\x8d\x6e\x74\xd4\x5c\x03\xe5\x35\x31\x7e\x3e\x8c\xd5\xd9\x82\x28\x87\x59\x3b\xf6\x84\x59\x30\x45\x8b\xfd\xdd\xb5\x4d\x8a\x2f\xdc\xee\x49\x04\xa2\x8b\xb2\x65\xbc\x10\xbb\x2e\x1f\x10\xb6\x5d\x17\xe2\x51\x17\x83\x27\x90\x3d\xdf\x4c\xad\x39\xff\x06\x00\x00\xff\xff\xf8\xc9\xc8\x2d\x43\x05\x00\x00")

func templatesOutputTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutputTf,
		"templates/output.tf",
	)
}

func templatesOutputTf() (*asset, error) {
	bytes, err := templatesOutputTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output.tf", size: 1347, mode: os.FileMode(480), modTime: time.Unix(1512068061, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResource_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x8a\xc4\x20\x10\x45\xf7\x9e\xe2\x23\xb3\x9d\xdc\x60\xce\x22\x15\x53\x64\x84\x44\x43\xa9\x59\x4c\xf0\xee\x83\x42\xa7\x3b\x24\xdd\x0d\x5d\x4b\xf9\xf5\xeb\xf9\x84\x63\xc8\x62\x19\x9a\xfe\xb2\xb0\xcc\xe6\xf6\x62\x46\x09\x79\xd1\xd0\x7d\x88\xbf\x1a\x9b\x02\x3c\xcd\x8c\x3a\x3f\xd0\x5f\xdb\x4a\xd2\xb1\x5f\x8d\x1b\xca\x77\xcb\x28\x60\x0a\x96\x92\x0b\xfe\x9e\x10\x1e\x5d\xf0\x45\x2b\x05\x24\x1a\x63\x2b\x02\xd8\xaf\x4e\x82\x9f\xd9\xa7\x53\x5b\x2d\x2a\xaa\x28\x75\x86\x5b\x72\x3f\x39\x6b\xdc\x13\xae\xab\x79\xcf\xfa\x72\x6b\xe7\x07\x8e\x66\xcc\xf1\x6a\x5b\xb8\x76\xd8\xd5\x8b\x5d\x8d\xb7\x9a\xfd\x0f\x86\x86\x41\x38\x46\x43\xd3\xa3\xb7\x98\x28\x39\xfb\x89\xb0\xff\x00\x00\x00\xff\xff\x58\xec\xbb\xb4\xcd\x01\x00\x00")

func templatesResource_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResource_groupTf,
		"templates/resource_group.tf",
	)
}

func templatesResource_groupTf() (*asset, error) {
	bytes, err := templatesResource_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resource_group.tf", size: 461, mode: os.FileMode(480), modTime: time.Unix(1512068137, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorageTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\xcd\x6e\x83\x30\x0c\xc7\xef\x79\x0a\x2b\xda\x99\x37\xd8\x79\xf7\xf5\x01\x22\x13\x2c\x16\x29\xd8\xc8\x31\x48\x5b\xc5\xbb\x4f\x64\x2d\x2b\x48\x55\x7b\x2c\x57\x7e\xfe\x7f\x45\xa9\xc8\xa4\x91\xc0\xe3\xcf\xa4\xa4\x43\x28\x26\x8a\x3d\x05\x8c\x51\x26\x36\x0f\xbe\x95\xf2\xe5\xe1\xec\x00\x18\x07\x82\xc3\xf7\x0e\xfe\xed\x3c\xa3\x36\x25\x0d\x63\xa6\x40\x3c\x87\xd4\x2d\xde\x01\x5c\xc5\x43\xaf\x32\x8d\xa1\x5e\x57\xfc\xea\xb5\x07\x9a\xd5\xa8\x59\xa9\xc5\x3b\x07\x90\x25\xa2\x25\xe1\xbd\x8d\x52\x9f\x84\xab\xfe\x25\x62\xb0\x44\x7a\xcc\x74\x32\xe4\x0e\xb5\xbb\xe5\x94\xc6\x9c\xfe\x34\x83\x7d\x8f\x35\xcc\xc7\xe7\xa9\x9a\x19\xf6\xa5\x76\x04\x20\x9e\x93\x0a\x0f\xc4\xf6\x6f\x7b\x53\x6b\x71\x8b\x73\xf7\x87\x8b\xc2\x86\x89\x49\x1f\x4e\x57\x83\x56\xe4\xce\x58\xf0\xf4\x5c\x00\x87\x77\xbb\x08\xec\xee\x0f\xc8\x41\x60\xcb\xbd\xfe\xa7\x52\xb6\x89\x46\x4d\x33\x1a\xf9\xe7\x6b\x17\xa3\x21\x52\xce\x0f\xaa\x6f\xd8\x4b\xd7\x6f\xb3\xb4\x6b\xf7\xdf\x00\x00\x00\xff\xff\x0f\x2e\x7d\x9b\x2b\x03\x00\x00")

func templatesStorageTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorageTf,
		"templates/storage.tf",
	)
}

func templatesStorageTf() (*asset, error) {
	bytes, err := templatesStorageTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage.tf", size: 811, mode: os.FileMode(480), modTime: time.Unix(1512068297, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorageTfOrig = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\xcd\x6e\x83\x30\x0c\xc7\xef\x79\x0a\x2b\xda\x99\x37\xd8\x79\xf7\xf5\x01\x22\x13\x2c\x16\x29\xd8\xc8\x31\x48\x5b\xc5\xbb\x4f\x64\x2d\x2b\x48\x55\x7b\x2c\x57\x7e\xfe\x7f\x45\xa9\xc8\xa4\x91\xc0\xe3\xcf\xa4\xa4\x43\x28\x26\x8a\x3d\x05\x8c\x51\x26\x36\x0f\xbe\x95\xf2\xe5\xe1\xec\x00\x18\x07\x82\xc3\xf7\x0e\xfe\xed\x3c\xa3\x36\x25\x0d\x63\xa6\x40\x3c\x87\xd4\x2d\xde\x01\x5c\xc5\x43\xaf\x32\x8d\xa1\x5e\x57\xfc\xea\xb5\x07\x9a\xd5\xa8\x59\xa9\xc5\x3b\x07\x90\x25\xa2\x25\xe1\xbd\x8d\x52\x9f\x84\xab\xfe\x25\x62\xb0\x44\x7a\xcc\x74\x32\xe4\x0e\xb5\xbb\xe5\x94\xc6\x9c\xfe\x34\x83\x7d\x8f\x35\xcc\xc7\xe7\xa9\x9a\x19\xf6\xa5\x76\x04\x20\x9e\x93\x0a\x0f\xc4\xf6\x6f\x7b\x53\x6b\x71\x8b\x73\xf7\x87\x8b\xc2\x86\x89\x49\x1f\x4e\x57\x83\x56\xe4\xce\x58\xf0\xf4\x5c\x00\x87\x77\xbb\x08\xec\xee\x0f\xc8\x41\x60\xcb\xbd\xfe\xa7\x52\xb6\x89\x46\x4d\x33\x1a\xf9\xe7\x6b\x17\xa3\x21\x52\xce\x0f\xaa\x6f\xd8\x4b\xd7\x6f\xb3\xb4\x6b\xf7\xdf\x00\x00\x00\xff\xff\x0f\x2e\x7d\x9b\x2b\x03\x00\x00")

func templatesStorageTfOrigBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorageTfOrig,
		"templates/storage.tf.orig",
	)
}

func templatesStorageTfOrig() (*asset, error) {
	bytes, err := templatesStorageTfOrigBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage.tf.orig", size: 811, mode: os.FileMode(480), modTime: time.Unix(1512068247, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4d\x6f\xe3\x36\x10\x3d\xd7\xbf\x82\x10\x7a\x5a\x20\x4e\xe2\xb5\x83\x74\x01\x1f\xf6\x54\xf4\xba\xe9\xad\x28\x08\x8a\x1a\xdb\xec\xd2\xa4\xc0\x0f\x25\x69\xe0\xff\x5e\x50\x94\x64\x89\xa2\x2c\x39\x4d\x8b\x75\x10\xe7\x26\xce\x3c\xce\x3c\x3e\x0e\x27\x64\x41\x14\x23\x29\x07\x94\x80\x28\x30\xcb\x12\xf4\x72\x98\xcd\x8e\x5f\x15\x6c\x99\x14\x09\x7a\x99\x21\x64\x9e\x73\x40\x6b\x94\x68\xa3\x98\xd8\x26\xb3\x8e\xa1\x66\xfb\x9c\x03\x8e\xa3\x68\x9b\x6a\xaa\x58\x6e\x98\x14\x91\x61\x03\x82\x08\x13\x19\xa0\x9c\xc1\xa9\x01\x0d\x54\x81\x09\x07\x05\x98\x47\xa9\xbe\x63\xca\x32\xe5\x23\xcf\x60\x43\x2c\x37\x2e\xf8\xdb\x9b\x79\xf9\x77\x7d\x7b\x17\x64\xc0\x84\x01\x25\x08\x1f\xf1\x5b\x2c\x4b\xbf\x5c\xc9\x82\x65\xa0\x50\x42\xfe\xb6\x0a\xd4\xde\x7b\x04\x99\x22\xe7\xfa\xf3\x4b\x41\xd4\x3c\x18\x39\x24\x8e\xd1\x3a\x71\x54\xfd\x1a\xeb\x66\xa4\xb4\x6b\x78\xe8\xd9\x35\x23\x6d\x3b\x4f\x4b\xcc\xce\x8f\x1c\x5c\x06\x0a\xb4\xb4\x8a\x42\x93\x00\xae\xbf\xe0\xad\x92\x36\x4f\x50\x92\x4a\xbd\xf3\x69\x09\xb2\x87\xee\xc4\x7e\xa1\x0f\x57\xa5\xcd\x0c\x21\x2e\x29\x71\xa9\x1d\x2d\xbc\x74\x0e\xc9\xcc\x25\x4a\xb6\xba\x04\x42\x08\x44\xc1\x94\x14\x7b\x10\xa6\x87\xe6\x80\x0e\x8e\xdd\x7e\x70\xb9\x4d\x39\xa3\x98\x0d\xc4\x15\xfb\x8d\xc7\x7a\xd2\xab\x89\x1f\xa1\x2e\x33\xb8\x3b\x6b\xe9\x10\xe7\x70\xee\x66\x9c\x3b\xf3\x12\xa6\xc9\x01\x93\x2c\x53\xa0\x35\x26\xbc\xcd\x9b\x36\xc4\x30\xfa\x1a\xc2\xfa\x7c\x15\x4c\x19\x4b\x38\xae\x36\xc3\x28\x6b\x51\xb2\xae\x0a\xe1\x66\xa8\xa3\xd5\x39\xa1\xd0\xd8\xff\x51\x39\xb4\xb7\xdb\x21\xf9\x73\x80\xdf\xa9\xb4\x4e\x65\x33\x2a\x12\x6d\x53\xe1\xea\xc1\x98\x42\xe2\xc9\xea\x4e\xb2\xb9\x82\x0d\x7b\x0a\x1d\xba\xc9\x0e\x49\x63\xba\x24\x82\x65\x8a\x70\x10\x58\x04\x24\x44\x38\x30\x52\x91\x2d\x60\x42\xa9\xb4\x62\x9c\x8c\x63\x81\x6a\x57\xf0\x7f\xbd\x3e\xa1\x0c\x62\xeb\x5f\x85\x88\x0d\x03\xe5\x0c\x1e\x0c\x11\x19\x51\x59\x7b\x4c\x41\xce\x99\xc7\xc1\xf5\x01\xf4\xeb\xb7\x87\x37\xaa\x2a\x35\x59\x54\x0a\x43\x98\x00\x35\xa1\xba\xac\x2b\x93\xc1\xba\x30\x7d\xf5\x83\xb5\x8a\xa9\x27\x30\x09\x00\x9a\xb8\xdd\xb8\x53\x6d\x4d\x51\xae\x58\x41\x0c\x0c\xed\x93\x48\xda\xda\xc0\x9e\x02\xe7\x23\xa9\x37\x66\x3f\x74\xfa\x29\x97\x69\x7c\x7b\xd4\x1b\x4d\x03\xb5\x8a\x99\xe7\x53\x87\xdd\x58\xc5\x18\x3a\x4b\xde\xba\xd6\xbd\x89\xd4\x07\x13\xa7\x9b\x33\xd2\xa6\x9b\xcb\x4e\x5a\x59\x0e\x4e\xec\x13\x1a\x88\x5e\xee\xda\xaf\x78\xae\x98\x74\x58\x51\x97\xc5\xcd\x8d\xeb\x1b\x99\x02\x1a\xef\x30\xd6\x28\xf9\x4d\xa4\xd2\x8a\xba\xca\x81\xd6\x43\xd3\x7f\xe5\x5c\x3e\xfa\x29\xa5\x91\x54\xf2\xb8\xd9\xef\x34\x2f\x77\x93\x67\x32\x97\xca\x60\x45\xc4\xb6\xdb\xa2\x7c\x4a\xca\x7e\x56\x1b\x26\x7c\x31\x0d\xec\xd6\x28\x59\x2c\x5a\x30\xb1\x63\x30\x06\x13\xd8\xd5\x26\xa7\x3a\xa6\xe9\x05\x22\xae\xd9\x88\x8e\xe2\x86\xe3\x0d\xc3\x90\x3e\xca\x7e\x80\x6c\xc1\x9d\x9f\xe7\xca\xa4\xe5\x3b\xae\x96\xdb\x8b\x55\xcb\xdd\xfd\xdd\xfd\x87\x5e\xda\x7a\xf1\xeb\x28\xd5\x2b\x25\xd3\xb8\x8f\xab\x66\x71\xb1\xaa\x59\xac\x56\xab\xd5\x87\x6c\x2a\xd9\x64\x42\x9f\x2f\x16\xe7\x34\x2e\x91\xcf\xff\xbb\x44\x3e\xbd\x89\x40\x56\x9f\x3f\xd4\x51\xa9\x83\x2a\xc8\x76\x36\x3d\x5f\x21\xb5\xe3\xb8\x4a\x96\x17\x5b\x48\xee\xef\x97\xcb\x0f\xa5\xd4\x4a\xd9\x5c\xed\x8c\xc9\xff\xb3\x62\x72\xb9\x5d\xca\x72\xf9\xee\xea\x09\xdd\xbc\x56\x23\x5c\x6e\x5f\x51\x4c\xbc\xdf\x7b\x6e\x4a\x96\xef\x5c\x25\xd2\x9a\xdc\x1a\xdf\xa5\x76\x6e\x1a\xbd\x1a\x0a\xc2\xed\x59\x97\x8e\x5d\x40\x7f\xe9\x7a\x1a\xcf\xdb\x9c\x84\x89\xd0\x36\x08\x77\xfa\x5e\xb8\x1b\x5d\xe4\x8e\x69\x38\xcc\x13\xb7\x4d\x21\x70\xf5\x36\xd4\xbb\x4f\x19\x80\x9e\x56\xf2\xeb\x19\xe0\xa9\x7a\x8b\x62\xc3\x88\xcd\x6b\x82\x07\x39\xbe\x2a\x74\xa1\xea\x7f\x2b\xea\x51\x87\xf7\x53\x03\x57\x1e\x1b\x5f\xae\xaf\xa7\xc0\x7e\xa9\x9a\xf7\x90\x89\x62\xef\xf6\x40\x79\xdf\x88\xbf\xc3\x73\x2f\x60\xc3\x3b\xe3\xf3\xda\x69\xde\xfa\x88\x73\xd8\xfb\x3b\x41\x10\x9a\x19\x56\x38\x5f\xa3\x2c\xc4\x67\xf3\x31\x9e\x37\x59\xe3\x83\x65\x0e\x42\xeb\x5d\x6f\xbe\x0d\xe1\xba\x33\xe1\x5f\x76\x9f\xa7\xf2\x09\x5b\xc5\xbb\xb4\x4d\xa3\x6b\xd1\xe1\xaa\xff\x2c\xd9\x82\x8b\xbc\x2d\xb4\x5c\x23\x4f\x93\x81\x6f\xc7\x22\x78\x19\x08\x28\x49\x8e\x44\x7a\x28\xc2\xb7\xae\xae\xef\xf6\x0e\xee\xdb\xc3\xd7\xb2\x88\x69\x82\x53\x66\x34\x5a\xa3\xe5\xcd\x2f\x77\xb3\xc3\xec\x9f\x00\x00\x00\xff\xff\x83\x56\x05\x44\x2a\x1e\x00\x00")

func templatesTemplateBytes() ([]byte, error) {
	return bindataRead(
		_templatesTemplate,
		"templates/template",
	)
}

func templatesTemplate() (*asset, error) {
	bytes, err := templatesTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/template", size: 7722, mode: os.FileMode(480), modTime: time.Unix(1510606149, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTlsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0a\x02\x31\x0c\x05\xd0\x7d\x4e\xf1\xc9\x09\x5c\x88\xe0\xa2\x0b\xaf\xa0\x07\x08\xad\x04\x5b\x6c\xa9\x24\xb1\x30\x0c\x73\xf7\x79\xa6\x3e\xff\xf6\x56\x70\x74\x97\x9f\xb5\x95\x43\xe5\xab\x1b\x83\xcb\xf4\x2a\x6b\x38\x63\x27\x20\xf7\xcf\xb4\x16\x75\x20\x81\x9f\xaf\x07\x13\x60\x9e\xa5\xb4\x70\x20\xe1\x7a\xb9\xdf\xe8\xa0\x33\x00\x00\xff\xff\x52\x4d\xac\xad\x51\x00\x00\x00")

func templatesTlsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesTlsTf,
		"templates/tls.tf",
	)
}

func templatesTlsTf() (*asset, error) {
	bytes, err := templatesTlsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tls.tf", size: 81, mode: os.FileMode(480), modTime: time.Unix(1512068061, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xc1\x4e\xc4\x20\x10\x86\xef\x7d\x8a\x09\xf1\xbc\xba\xc6\x78\xf3\x59\x1a\xb6\x1d\x37\x13\xe9\x40\x86\x29\x46\x9b\xbe\xbb\xa1\x44\x54\x42\x5c\xb8\xf1\x7d\x3f\x99\xf9\x93\x15\xb2\x17\x87\x60\x90\xd3\x48\xb3\x81\x6d\x1f\x86\x9f\x57\xc1\x2b\x79\x36\xb0\x0d\x00\xfa\x11\x10\x5e\xc0\x44\x15\xe2\xab\x19\xfe\x88\x91\x96\xe0\x70\xec\xff\x12\xd7\x4b\x9c\x84\x82\x92\xe7\x0e\x56\x64\xcb\xda\x01\x93\x23\xfc\x0f\x44\x9c\x04\xb5\x85\x8c\xfa\xee\xe5\x6d\x9c\x68\x96\x32\xf9\x8c\xaf\x76\x75\x9a\x87\x3f\x3f\x9c\x8e\x7b\x7f\x7e\x6e\x36\x20\x56\x14\xb6\xee\x46\xee\xf1\xe9\xc8\x05\xf1\x89\x66\x14\x30\xf6\x73\x15\x94\xa5\x24\x9a\x4d\x73\xf2\x6e\x4b\x56\x4e\x0d\xd8\x4d\x2e\xf4\x7b\x6f\x28\xa7\xca\x15\x1c\x5a\x6d\xa1\xd5\x2a\xf8\xad\x95\x4e\x3a\x5a\x01\x7b\x9e\xfe\x2b\x00\x00\xff\xff\xa6\xc7\xe4\x1f\xf6\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 502, mode: os.FileMode(480), modTime: time.Unix(1512068137, 0)}
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
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/network.tf": templatesNetworkTf,
	"templates/network_security_group.tf": templatesNetwork_security_groupTf,
	"templates/output.tf": templatesOutputTf,
	"templates/resource_group.tf": templatesResource_groupTf,
	"templates/storage.tf": templatesStorageTf,
	"templates/storage.tf.orig": templatesStorageTfOrig,
	"templates/template": templatesTemplate,
	"templates/tls.tf": templatesTlsTf,
	"templates/vars.tf": templatesVarsTf,
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
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"network.tf": &bintree{templatesNetworkTf, map[string]*bintree{}},
		"network_security_group.tf": &bintree{templatesNetwork_security_groupTf, map[string]*bintree{}},
		"output.tf": &bintree{templatesOutputTf, map[string]*bintree{}},
		"resource_group.tf": &bintree{templatesResource_groupTf, map[string]*bintree{}},
		"storage.tf": &bintree{templatesStorageTf, map[string]*bintree{}},
		"storage.tf.orig": &bintree{templatesStorageTfOrig, map[string]*bintree{}},
		"template": &bintree{templatesTemplate, map[string]*bintree{}},
		"tls.tf": &bintree{templatesTlsTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

