// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x57\xcd\x6e\xe3\x36\x10\xbe\xeb\x29\x26\x32\xf6\xb6\x58\xf4\xec\x1b\x23\xcb\x8e\xba\xb2\x64\x48\x72\xb6\xe9\x45\xa0\xa5\xb1\xcd\x46\x22\x05\x92\xda\x34\x7b\xeb\x7b\xf5\x9d\xfa\x0a\x05\x29\xda\x96\x1d\xa7\xf1\xa2\xb9\xc4\x02\x38\x3f\xfc\x66\xe6\x9b\x8f\x93\x4a\xb4\xad\xe0\x5e\x42\x96\x61\x19\xfe\x16\xe5\x45\x3e\x05\x3f\xa1\x2d\x02\x6d\x24\xd2\xfa\x15\xf0\x4f\xa6\xb4\xf2\xbd\x68\x55\x26\x69\x71\x3a\xb4\x6a\x90\x2a\x84\x2d\x6b\x1a\x60\x1c\xf4\x1e\xa1\x11\x15\x6d\x20\x5a\x81\x18\xbe\xd5\xab\xd2\xd8\x82\x42\xad\x19\xdf\x41\x47\x77\xe8\x7b\xde\xa4\x6a\x7a\xa5\x51\x7a\x41\xbc\xce\x8b\x30\x2b\x67\x61\x1c\x16\x61\x39\x27\x51\x1c\xce\xa6\xe0\x57\x94\x03\x17\x1a\x6a\x6c\x50\x23\xb8\xe3\x26\x48\xd5\x4b\x89\x5c\x83\xd2\x54\xa3\x7f\x74\x10\xe5\x36\xb5\x6c\x9d\x24\x51\xb2\x98\x82\x5f\xec\x47\x66\xca\x3a\x93\x3d\xe7\x8c\xef\xde\x18\xc5\x69\x40\xe2\x29\xf8\x51\xdb\x09\xa9\x8f\x56\x15\xe5\xc6\x6a\x83\xd0\x77\x3b\x49\x6b\xac\x6d\xe6\x12\x6b\xe4\x9a\xd1\xc6\x3b\x4b\xba\xcc\xc2\x3c\x5d\x67\x41\x38\x05\x7f\x4e\x59\x83\x35\x68\xe1\xf2\xbf\x83\x62\x8f\x12\x4d\x1e\x94\x03\x55\x4a\x54\x8c\x6a\xac\x61\x2f\x94\x86\x9e\xd7\x28\x41\xef\x99\x82\x67\x7c\xf5\xdf\x71\x5b\xfe\x9e\x26\x3f\xe5\xfb\x87\xe0\x78\xc5\xf7\x9c\xac\xe3\xa2\x0c\xb2\x70\x16\x26\x45\x44\xe2\x32\x20\x89\x45\x61\x08\x3b\x05\x7f\x86\x5b\xda\x37\x1a\x4e\x37\x1d\x41\x31\x04\xad\xfd\xa1\x5d\x82\x87\x30\xf8\x7a\xaa\x9a\xc5\xfc\x64\xc5\x4d\x0f\x9d\x4c\x6d\x3b\xd8\xb6\x52\xf6\x77\xaf\x50\xda\x33\xbe\xe7\x79\x13\x03\x85\xf7\x90\xe6\x45\x49\xe2\x2c\x24\xb3\xa7\x53\x9b\x3d\x18\x94\x2e\x7b\xd1\xa1\x64\x2d\x8e\x09\x5c\x05\x67\x40\xd9\xe0\xe3\x5c\x8c\x40\x7a\x61\x7a\x6f\x93\x71\x45\xbf\xe6\xb7\xbc\x7f\x2a\x57\x59\xfa\x6b\x18\x14\xff\x2b\x44\x27\xc5\x1f\x58\x69\xdf\xcb\x9f\xf2\x22\x5c\x96\x6e\x9a\xe6\xe9\x3a\x99\x7d\x30\x4c\x5b\x26\x95\xfe\xe7\xef\xbf\x7c\x2f\x49\x8d\x1d\x79\x24\x51\x4c\xee\x63\x53\xae\x44\x40\xd4\x01\xfd\x4e\x59\x43\x37\x0d\xfa\x5e\x94\x0f\x1d\x6d\xef\x30\x9a\x25\x36\xb4\xf7\xe0\xd4\x24\xec\x0f\x78\x47\xcb\x55\x9a\x15\x65\x98\x65\x69\x56\x26\xeb\x38\x1e\x9c\xd6\x54\x53\x73\x4d\x67\xf6\x42\x15\x6c\x45\xcf\xeb\x3b\x70\x99\x56\x7b\xac\x9e\x6d\x9e\xee\xc8\x96\x35\x78\x77\xee\xd4\xb8\x2b\x1f\x49\xbc\x36\x99\x86\x6d\xa7\x5f\x07\xbf\x82\x43\xc3\x38\xc2\x27\x75\x7e\xfe\x5b\x96\x26\x8b\x72\x9e\x66\x4b\x62\x52\x8f\x78\x25\xa4\xc4\x4a\xc3\x10\x40\xc8\x96\xea\x77\x8d\x47\x4d\x3d\x06\x36\x18\x75\xa4\xd0\xc3\x25\xde\xf5\xe1\x4a\x9e\xac\x97\x53\xf0\x09\x68\xa1\x69\x03\x62\x0b\x9f\x14\x48\xf1\xa2\xcc\x4f\x7b\x01\x2a\x11\xe8\x86\x9b\x84\x9a\xcf\xa0\x9e\x59\x77\xd5\x4f\x4e\x1e\xcf\x87\x56\xd1\xef\xae\x5b\x3e\xa9\xcf\x20\x91\x2a\xc1\xa7\x26\x09\xf7\x37\x31\x53\xe1\xad\x48\x9e\x7f\x4b\xb3\x99\xbd\xc7\x92\x14\xc1\x83\x9b\xae\x8e\x2a\xf5\x22\x64\x6d\x7a\x8d\x1d\xc0\xf1\xbd\x34\x8b\x16\x51\xe2\xee\x3d\x3e\x2f\x24\xdb\x31\x4e\x9b\xf7\x0c\xd7\xf9\x89\x06\x49\x50\x44\x36\xdb\xe2\x30\x9c\x8e\x37\x91\x9b\xc6\x1a\x15\x5e\x70\x4d\x2b\x6d\x4b\x4f\xeb\x96\x71\xa6\xb4\xa4\x5a\xc8\x3b\xe7\x70\x8c\x7e\x22\x40\xf5\xd5\xde\x3a\xb4\x1d\x4c\x66\xcb\x28\x79\xcb\x3a\x26\x68\xed\x98\xc7\x3a\x1d\x52\x78\xc3\x3c\x77\xe7\x49\x67\x61\x4c\x8a\x70\x36\x1a\xd0\xb5\x31\xdb\x53\x93\xfa\x78\x0c\xdd\xf4\xd9\x14\xe2\x19\x59\x1d\x33\x58\xaf\x66\xe4\x98\x41\x53\xd3\xee\x32\x30\xd6\x6c\x88\xfb\x18\x66\xd1\xfc\xa9\x0c\xd2\xd9\x68\x53\x3d\xa2\x64\x5b\x56\x51\xcd\x04\x87\x4a\xd4\x08\x28\xa5\x90\xbe\x17\x2e\x49\x14\x97\xb3\x28\x77\x73\xba\xa4\xac\x39\xec\x41\x65\xfb\xa7\x66\xea\x46\x60\x0f\xde\xc6\xe5\x0d\x5b\xe3\xb0\xa5\xba\xda\xc3\xd6\xf6\xd7\x40\x10\x86\x97\x8f\xfd\x93\x9b\xaf\x63\xae\x06\x9a\xff\x20\xe5\x43\x8f\x5c\x3a\xb1\xcc\x30\x05\xff\x45\x0a\xbe\x3b\xd1\x36\x08\x39\x32\xf1\x26\x12\x77\x4c\xf0\x03\x81\x66\xe1\x22\x4a\x93\x5b\x97\x23\x0c\xc6\x1f\x51\xa8\xd9\x69\x26\x94\xf9\x7f\x08\x64\xf6\xe2\xcd\x61\xec\x52\xfc\x88\xa7\x1b\xca\xcf\x35\xc2\xc0\x89\xc1\x00\xda\x0e\xf5\x78\x63\x5c\xa1\xc3\x4a\xf0\x2d\xdb\xf5\xd2\xf6\x84\x2d\x4a\xb4\x24\x8b\xf0\x7d\x57\xac\xa5\x3b\xbc\xc9\x91\xe7\x4d\x44\x87\x5c\x69\x5a\x3d\x7b\x8b\xb0\x38\xc0\x7c\x28\x51\x22\x0e\x48\x5a\x9a\x33\x58\xb9\xc6\x5f\x62\xbb\x41\x79\x9c\x1d\x32\x9b\x8d\x67\x65\x83\xc8\x81\xd6\x4e\xe3\x38\x93\xa3\x5a\x70\xc3\xf5\xbe\x54\x70\x06\xd7\x74\xc2\xc1\xf6\x81\xe4\xa5\xc3\xd4\x2c\x3b\x67\x30\xc2\xff\x88\x67\x70\x65\xe4\xbd\x09\x17\x35\x7a\x89\x19\xbd\x83\x3e\x70\x5a\xaf\x2c\x48\xfe\xd5\x90\x75\x5d\x83\x39\x64\xda\xd2\xc9\x46\xfb\x79\x28\xb5\x53\x7f\x9f\xbb\x01\xe5\x17\xca\x34\x30\x0d\xb5\xe0\xf8\xc5\x04\xd8\xd0\xea\xb9\xef\x48\x55\x89\x9e\x6b\x6f\x45\x32\xb2\x2c\xc3\xe5\xaa\x78\xb2\x7b\x48\xf5\xdb\x2d\xab\x98\x51\x9e\x1d\x95\xb4\x45\x8d\x52\x99\x75\x5c\x94\xf9\x7a\x65\x28\xdf\xc0\xc9\x55\xdf\x99\x4d\x65\x1a\xef\xb5\x33\xf2\xf4\x5c\x20\x9d\x91\xc5\x30\xb4\x47\xc5\x71\x4f\x82\xaf\xeb\x55\x49\x82\x20\x5d\x27\x3f\xa3\x3d\xce\x12\xbf\x59\x84\x78\x13\xd3\xe7\x17\x62\xf3\x86\x68\xc6\xea\x27\x82\xb8\xaa\xde\xdb\x1c\x3d\x77\xc7\x79\x14\x87\x03\x7d\xbb\x79\x70\x9d\xef\xca\xa6\x8f\x97\xb2\x92\x02\x36\xb8\x15\x12\x41\xbd\x30\x5d\xed\xcd\x2b\x62\x74\x80\x0e\xd7\x3e\x9b\xd7\x21\xca\xdb\x37\xc0\x06\x8d\xb1\x31\xc4\x1a\xfa\xce\x0e\xd4\xc8\x2c\x0b\xf3\x22\xcd\xc2\xb7\x76\x12\x95\x16\x92\xf1\xdd\x30\x82\x87\xf1\xc8\x50\x89\x5e\x56\xf8\x16\xc3\xd1\x35\x3f\xbc\xdc\x49\xda\x5d\x57\x9e\xa7\x99\x39\xea\xcc\x03\xf4\x1b\x6c\x84\xd9\x25\x5a\x9c\x13\x52\x61\xe4\xbe\xe8\x50\xba\xa5\x74\x9c\xa7\x0e\xa5\x91\x50\x6e\xa2\xcc\xae\x1b\x56\xe1\x71\xa1\x0e\xd5\xc8\x5f\x79\xb5\x97\x82\xb3\x1f\x06\x27\x85\x72\x58\x57\xbf\xb8\xc5\x19\xa7\x8b\x28\xb9\xb4\x59\x8f\xf5\x82\x91\xe8\x77\xee\xf4\x69\x01\x16\xa7\xe7\x60\x27\xc5\x9e\x6d\x98\x56\x60\xce\xb8\x18\x5b\x29\x5a\x68\xc4\x6e\x67\xaa\xc4\xf8\x97\x5b\xe4\x86\x37\xa9\x98\xf2\x82\x28\xb7\x2c\x70\x49\x0d\x46\xfc\x32\x05\x9a\xaa\xe7\x4b\x1a\x30\xa6\xdf\xdb\xc0\x12\xac\xf7\xb8\x2c\x83\x34\x99\x47\x8b\xd3\x9b\x23\x18\x53\xef\x9b\xc7\xc7\xc9\xe0\xf2\xe1\x5a\x5c\xd2\xf6\x7b\xd5\xaa\xb1\x6b\xc4\x6b\x6b\x09\xa5\xa1\xfc\xd6\xaa\x79\xde\xbf\x01\x00\x00\xff\xff\x2b\xc2\x01\x01\xb0\x0f\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 4016, mode: os.FileMode(420), modTime: time.Unix(1607505446, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x57\x5d\x57\xda\x5c\x16\xbe\x3f\xbf\x82\x05\xab\x77\xb3\x66\xcd\x75\xef\x8e\x21\x62\xa6\x21\x61\x25\xc1\x8e\x73\x93\x65\x91\xe9\x38\x55\x70\x89\xce\xc5\x5c\x15\x3f\xc1\x82\x38\xad\xfa\xfa\x41\xab\x58\xac\xbe\x5a\x40\x6b\x2b\x08\xa8\x7f\x26\xe7\x24\x5c\xf9\x17\xde\x75\xce\x49\x42\xc0\xd2\xd7\x4b\xe3\x7e\xf6\xde\x67\xef\x67\x3f\x7b\x13\x88\x25\xa7\xa7\x93\x09\x20\xc1\x30\xaf\xf3\xff\x10\x54\x4d\x7d\xee\xf3\xa3\x8d\xbc\x79\x72\x81\xea\xdf\x50\x65\x07\x15\x4f\xfd\x40\x88\xe8\x92\xac\x75\x0d\xac\x5a\x1d\x15\x4f\xcd\xab\x96\xd9\x3a\xb0\xaa\x77\xe6\x6d\xb5\x53\xfa\xd1\xf9\x78\x84\x4a\xe7\x68\x65\x97\x7d\x17\x22\x7e\x00\x02\xb1\xa9\xf9\xd4\x5c\x7c\x16\x70\x62\x54\xd5\x78\x45\x0f\xf2\x22\xaf\xf1\xfa\x30\x14\x44\x3e\xf8\xdc\xe7\xc7\xbf\x1d\xe2\xab\x2d\x94\x39\xec\xec\x96\xd1\xed\x07\x94\xcd\x9b\x6b\xd7\xf8\x6d\xda\xdc\x5b\xea\xec\xaf\x98\x77\x65\xbf\x0b\x15\x54\x9a\x84\x12\x95\x24\x41\x0a\x3d\xf7\xf9\x99\x81\xd1\xc8\xa3\xe2\xa9\x75\xbf\x61\x95\x72\x46\xa3\xf2\xd0\x4e\x3f\x82\x88\x32\x07\x45\xf2\xae\x5a\x1b\x2d\x1f\x33\x98\x1d\x38\xbf\x6a\x36\x4f\x68\xa2\xb3\xf1\x89\x78\x62\x6e\x72\x7c\x0a\xf4\xe4\xa8\x2b\xbc\x2a\x47\x15\x8e\x27\x78\x96\x66\xf9\xd2\xfa\x7e\xfc\xd0\x4e\x5b\xb5\x63\xf3\x64\xa7\xf3\xfe\xd8\x68\xbc\xc3\xc5\x2c\x5a\xbe\xb2\xd2\x9b\x46\xa3\x85\x8b\x4d\xff\x00\x27\xfa\x3f\x65\xe9\xa9\x9e\x50\xa1\x66\x6e\x9e\xa2\x1c\x75\x36\x0c\xa3\xa2\xa6\x73\x0a\x1f\xe4\x25\x4d\x80\xa2\xce\x41\x89\xbe\x8d\xc5\x21\xd5\x68\xed\x58\xd5\x32\x5a\xad\xe0\x7c\xd5\x68\xe4\xad\xc5\x5b\x16\x84\x16\x84\xf6\x97\x1b\xe1\xb9\x17\xdd\xd2\xb3\x88\xac\xd7\x0c\x60\x34\xd6\xcd\xcd\x53\x9c\xa9\x93\x8f\xfb\x0d\xb4\x91\x23\x95\xf9\x77\x32\x35\x07\x46\x64\x55\xd3\xa1\xa8\xf0\x30\x38\xd6\xe5\x01\x7b\xac\x87\x28\xf6\xab\xa9\xb5\x1b\xe8\xf1\x63\x5d\x9c\xd9\x2a\xb0\xc7\x3a\xcd\x7e\xec\x40\x1f\x1a\xd3\x23\x8a\xfc\x77\x9e\xd3\x9e\xea\xab\x74\x63\xee\x57\xe9\xbb\xd5\x31\x55\xe3\xc3\xba\xcd\xdf\x61\x39\x2a\x05\x6d\xfa\x2e\x67\x18\x59\x71\xf1\x2b\x2e\x36\x85\x08\x2b\x93\x4c\x4c\xe1\x28\x14\x44\x38\x24\x92\xaa\x0a\x11\x9f\xf5\x63\x09\x37\x37\x48\x85\xae\xaf\xfc\x40\x50\x19\x9d\x68\x8a\x5d\x02\x53\x66\x31\x5f\x7e\x56\x2c\x21\x1c\x91\x15\x4d\xe7\x15\x45\x56\x74\x29\x2a\x12\x02\xe2\xe2\x19\xce\xde\xa1\xcc\x05\x2a\xd4\x18\xc4\xdc\x5b\xc2\x5b\x17\x38\x5f\xa5\x8f\xa9\xe3\xcf\x6f\xf1\xc1\xb1\xed\x6d\x7b\xd5\x68\x5d\xd3\xbc\xbc\x0e\x89\x2b\x7d\x14\x8a\x51\x92\xde\xb3\x94\xcf\x2a\xe5\x70\x31\x6b\xfe\xde\x64\x7e\x7a\x8d\x5f\x2a\xb2\x14\xd2\x87\x65\x25\x0c\x35\xd7\xdc\x3c\xaf\xa1\xc2\x67\x7c\xd8\x46\xed\x82\xd1\xc8\xe3\xca\x67\xb3\xd4\x87\xf3\x10\xcd\x5b\x38\x3b\x5c\xf6\x8e\x0c\x5d\xe6\x02\xd5\x56\x3a\xef\x8f\x7b\x91\x76\xd3\xa4\x68\x98\x74\x6b\xf9\xd2\x67\x63\x68\x72\x2c\x1a\x6a\x34\x1e\xda\x39\xab\x7e\x65\xdd\xaf\xfe\x14\xac\xc2\x51\xde\x65\x17\x71\x60\xdc\x7f\x24\x0c\xb3\xbb\x9e\x43\xeb\x07\x68\xff\xf0\xa1\xbd\xf7\x2c\xe5\x07\x00\x04\xe6\x53\xf1\x59\x10\x81\xaa\xfa\x52\x56\x82\x34\xe1\x30\xd4\xb8\x11\x3a\xf1\x2b\xe6\x61\xba\xb3\xb9\x6b\xd5\x6a\x7e\x20\x2b\x42\x48\x90\xec\x37\xb9\x26\xeb\x07\xbd\x56\x51\xb5\xab\x1c\x90\xd3\x04\x9a\x0b\x9b\x0b\x5c\x3c\x43\x1b\x64\x2a\x59\xb7\xac\xf4\x26\xd1\xba\x6a\xc9\xdc\x58\x41\xff\xdf\xa1\xad\xa2\x68\x6f\xcd\x08\x41\x2a\x65\x86\xa7\x16\x30\x18\x16\xa4\x41\xf3\xeb\x1b\x9f\x98\x9e\x4c\xf8\x98\x39\x9b\x4a\xeb\xe8\xdc\x33\xc9\xde\xec\x14\x5e\x84\x1a\x1f\xf4\x8c\x87\x9d\xe6\xb7\x92\xab\x22\x6c\x18\xfc\x40\x0c\xc2\x88\x1b\x34\x1a\x09\x42\x1a\x94\x7c\xed\x09\x66\xdc\x57\xf1\xe6\x0d\x8d\x34\xca\x2b\xc2\xf0\x98\xce\xc9\x41\x8f\x5e\x77\xce\x72\x56\x2d\xed\xa9\x16\x1f\x86\x82\xa8\x07\x05\xd5\x1e\x98\xce\x42\xd5\x68\x5d\xb3\xa5\x60\x1d\x9d\x9b\x5f\xd2\x83\xca\xe5\x60\xbd\xcd\x60\x68\x94\xbb\xe9\x2c\xe7\xdd\x21\xb7\xe5\xcb\x6d\xb0\x4a\xfe\xea\xca\x98\xa3\x58\xae\x86\xb1\x76\x3a\x02\xd6\x8b\xa5\xe3\xe8\x45\xe1\xcc\x76\x6f\xfb\x01\x08\xcc\xc6\x5f\x4f\x26\x13\x8e\x16\x29\x7c\x48\x90\xa5\x27\xad\x03\x94\x6b\xa2\x83\x03\xaf\x16\x79\x44\x1c\x04\xfe\x97\x4c\xc4\x1d\xaf\x64\x11\x3c\xcd\xa7\xe3\xa1\x47\xe2\x16\x4f\xcd\xdb\x6f\x56\xb5\x84\x32\xef\x7b\xd7\x1d\x13\x1a\x6b\xbd\x8e\x0a\xdb\x6c\x4a\x99\xb4\x7a\xf5\xa5\xb3\x9c\x37\x6f\x99\x40\x0a\x61\x18\xe2\x07\x01\xb7\x8a\x68\xb1\x30\x00\x08\x02\xc9\x99\x78\x22\x35\x37\x1e\x7b\x03\x42\xbc\xe6\x54\xc9\xa9\x6f\x57\x23\x68\x49\x88\xfd\xcc\x6c\xf2\x3f\xf1\xd8\x5c\x38\x3e\xfd\x2a\x3e\xeb\xf2\x18\x06\x3d\x4d\xa4\x2f\xc4\xf5\x16\x5a\x3b\xf4\x20\xdc\xf5\x67\xf3\x7c\xd0\xec\x30\xaa\x3f\xda\x7d\x0e\x6a\x04\xaa\xba\x5d\x29\x02\xa1\xc6\xde\xe5\xf3\xd0\x4e\x3f\xc2\x82\x40\x22\x39\x11\x07\x12\x99\x01\x67\xfb\xd9\x07\x88\xae\x41\xf5\x05\xd5\xb8\x6b\xa3\xb5\x6d\xad\x2d\x98\x0b\x37\x78\xeb\xa2\xb3\x5a\xc0\x1f\xf2\xc6\x6d\x91\x08\x5d\xf1\x14\x67\x4f\xac\x52\xee\x2f\x3e\xab\x56\x37\x2b\x59\x74\xb7\x8c\xaa\x8b\x46\xeb\x2b\xfb\x8c\xaa\x39\x5c\xdb\xfa\x2b\x09\xf3\x6a\x3c\xf6\x66\x7e\x06\xc6\x62\xc9\xf9\xc4\x1c\x88\x40\x05\x86\x75\x3e\x1c\xd1\xc6\x48\x84\xc2\x02\xde\xba\x70\x16\x10\x79\xb8\x1a\x8d\x10\xbd\xa4\x1a\x99\xc7\x9b\x35\x9c\x23\x07\x93\x79\xd9\x42\x9f\xde\xf9\x41\xdf\x9e\xc7\x87\xa5\xce\x59\xce\x33\x4d\x36\x03\x87\x20\xf7\x22\x1a\xd1\x21\xc7\xc9\x51\xe9\xa9\xdb\x16\x95\x57\x8d\xd6\xad\xf5\xfd\x0b\x2a\xd4\x07\xec\x5c\x10\x98\x99\x1a\x4f\xf4\x1d\x40\x7f\xe2\xd6\xcb\xe7\xc7\x6e\x3d\x87\xe4\x10\xad\x13\xb0\x73\x1f\x16\x44\x9e\x49\xa1\x43\x5f\xb6\xd9\x6d\xff\x24\x55\xb6\x45\xd1\x4a\x1e\x65\x56\x71\xfe\xc8\x9b\x7f\xcf\xe0\x30\x8f\xee\x55\xc9\xba\xc7\xac\x7f\x72\x55\x2a\xbc\xaa\xc9\x0a\xdf\x67\x8e\xd3\x47\xa8\x9c\x77\xcc\x5d\xfe\x2a\xf1\x54\x72\x7e\x36\x16\x7f\x5c\x12\xcf\x33\x7e\x91\xbc\xb7\x69\x7d\x87\x51\x97\xcf\x3d\x67\xd0\xe5\x27\xa3\xb9\xde\xc7\x6a\xeb\x7e\x9f\xec\xe1\x4a\x99\xd1\xd3\x4e\x71\x6a\x62\x7c\x86\xed\x06\x77\xb1\xb0\x4a\xa2\x8d\x1c\xae\x1c\xa3\xcc\x05\x21\x96\xbd\x20\x9a\x7f\xb3\xf7\x88\x28\x87\x04\xa9\x1f\xe1\xae\x11\x76\x0c\xd2\x00\xd4\xba\xbb\x1c\xd8\x2f\x03\xf3\x4b\x1a\x57\x8e\xc8\xbf\x18\xc4\xdc\x6d\x75\x76\x57\x7c\x03\x16\x2a\x08\xc4\x26\x53\x80\x13\x54\x3a\x70\xfd\x53\x48\x32\x75\xce\x4f\x9c\x3d\xc7\x85\x82\xd1\xa8\x98\x7b\x4b\x46\xab\x85\xd6\x4a\x04\xfd\xdf\x69\x2e\x99\xf8\xd7\xe4\x6b\x30\x1a\xd6\x39\x59\x1a\x16\x42\xdd\xfb\x95\xe9\x99\xe7\x7e\xed\xda\xf4\xff\x50\x71\x4d\xbb\xf5\xf5\x70\xf6\xd7\x55\x06\x00\xfc\x11\x00\x00\xff\xff\x87\x78\x0f\x9c\x69\x0d\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 3433, mode: os.FileMode(420), modTime: time.Unix(1607505446, 0)}
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
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
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
