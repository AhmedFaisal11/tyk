// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// schema/3.0.3.json
package schema

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

var __303Json = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\xeb\x6f\xdb\xb6\xf6\xdf\xfd\x57\x10\x5e\x3f\x74\x58\x1c\xa7\xed\xef\x07\xdc\xf5\xcb\xd0\xa6\x19\x66\xac\x5d\x02\xdf\xdc\xa1\xb8\x89\x7b\xc1\x48\xc7\x31\x57\x49\x54\x29\xaa\xb1\xd7\xe4\x7f\xbf\xd0\x5b\x94\x48\x3d\x29\x3f\x7a\x5d\x60\x58\x4b\x91\x87\xe7\x49\x1e\x9e\x73\x48\x7f\x1b\x21\x34\x26\xe6\xf8\x35\x1a\xaf\x38\x77\xbd\xd7\xd3\xa9\xe7\x82\x71\x4a\x5d\x70\xb0\x4b\xbc\x53\xca\xee\xa7\x14\x7b\xd3\x57\xa7\x67\x53\xcf\x58\x81\x8d\xa7\x2f\xcf\x5e\xbe\x98\x9c\xfd\x3c\x79\xf9\x8f\xf1\x49\x30\xfc\x59\xd4\x9e\xc0\x78\x3d\x9d\xfe\xe5\x51\x67\x12\xb5\x86\x00\x4c\x86\x97\x7c\x72\xf6\x7f\x31\x84\x1f\xa2\x71\x26\x78\x06\x23\x2e\x27\xd4\x09\xc6\x5e\xaf\x00\xe5\x9a\x10\x5d\xa2\x4b\x17\x9c\x37\x57\x33\xf4\xf5\xd5\xe9\xd9\xe9\x1a\x99\xd4\xf0\x6d\x70\xb8\x77\x82\xb0\x87\x4c\x58\x12\x07\x4c\x74\xb7\x41\xd5\xa8\x87\xa3\x5f\x45\x73\xf2\x8d\x0b\xc1\x64\xf4\xee\x2f\x30\x78\xd4\xc6\xe0\x8b\x4f\x18\x04\x4c\xb8\x19\x21\x84\xd0\x38\x06\x11\x7e\x0e\xf8\xe3\x2c\x69\xf2\x77\x17\xf3\x95\x37\x1e\x21\xb4\x08\xc7\xba\x8c\xba\xc0\x38\x01\x6f\xfc\x1a\x7d\x13\x47\x27\x0d\xb9\x69\x3d\xce\x88\x73\x1f\xc3\x8a\xa0\x71\x60\x21\xf9\x9f\x5e\xdd\xde\x9e\x9e\xdd\xde\x9e\xde\xde\x9a\xcf\x27\xa7\x3f\xfd\xf8\xcb\xb3\x71\xd8\xed\x29\x8f\x45\x0e\xe6\x33\x06\xcb\x60\xe0\x0f\xd3\x90\x13\x24\x60\x9a\x37\x9d\x05\xbd\x84\x71\xb0\x0e\xa6\xc0\xd6\x3b\x6a\x78\xf5\xe3\x2f\xb2\xde\x21\xaf\x71\x28\x1e\x01\xa0\x07\xec\x2b\x30\x4f\x46\x1f\x66\x0c\x6f\x32\xf2\x08\x07\x3b\xdf\x4f\x39\xeb\x3f\x43\x90\xe3\xb8\xdb\x53\x61\x3a\xc3\x67\x84\x6f\xf4\xce\x17\xc1\x9c\x47\xa2\x0f\x08\x95\x4f\xce\xf1\xbd\x56\x42\xaf\xf1\x7d\x3a\x51\x3a\xda\x77\xc8\x17\x1f\x66\x31\x0c\xce\x7c\x10\x70\x88\x54\xae\x56\x72\x57\x89\x66\x66\x23\x0d\x6a\xbb\xd4\x09\x2c\xa6\x7e\xf8\x79\xd6\x77\x94\xf0\xe1\x29\xd2\xf1\x48\x49\xaf\x24\xaa\xfe\x69\x3d\x49\xff\x91\x0d\xc0\xa6\x19\x02\xc5\x96\x30\x66\x89\x2d\x0f\x62\xcb\x4f\xe7\xcd\x60\xcd\x61\x09\x0c\x1c\x03\x64\xfc\xce\xd9\x6b\xd8\x5e\xb2\xd9\x8c\xb2\xf8\x9f\x8b\xa2\x91\x49\xf0\x8f\x68\xb8\xbd\x0d\xc6\x3d\x13\x5a\xd5\x36\x1b\x7e\x5b\x52\x66\x63\x1e\x7c\xf5\x19\x99\xb0\x14\xf3\xb4\xcf\x93\x54\x9b\x66\x05\x13\x6e\x47\x1d\x27\xdc\x82\x1c\x1e\xe3\xc0\x0a\x53\xeb\x14\x28\x56\x90\x1a\x41\xa8\xa6\x33\x23\x21\x37\x93\xb8\x54\xb7\x1e\xce\x81\xd9\xde\xe5\x32\xb0\x72\x62\xd4\x21\xd0\x92\xd1\xb9\x69\x0c\xea\x70\x6c\xf0\x22\x7c\x95\xbe\x47\x9d\xa5\x90\x2c\x62\x80\xe3\x95\x30\x95\x43\x7a\x1f\x77\x96\x42\x4a\x64\xd4\x94\x6b\xa5\xc5\xa1\x4e\x7b\x73\x16\x28\x07\x50\x65\x8d\x82\x76\x9e\x97\xd8\xa7\x54\x50\x95\x82\x39\xd8\xee\xa4\x5f\x3e\xb3\x06\xd3\x0a\xb0\x31\xe9\x08\x3d\x1a\xba\x17\xc2\x79\x5f\xd2\xc8\x96\xab\x47\x28\x9a\xc6\x2b\xc5\xde\x08\x72\x87\x2c\x8f\x9d\x92\xce\x1c\x0f\x78\xd1\x98\xe1\xf5\x8c\x1b\x62\x61\xfe\x8a\x19\xc1\x77\x56\x01\x19\x35\xa5\x55\x0c\xcc\x8f\xaf\x71\xf4\xfe\x8c\xa7\x1d\xe7\x86\x3c\xed\x91\xcc\x53\xfc\x3a\xcb\xde\x84\x25\xf6\x2d\xde\x5c\xfe\xe0\xf8\xb6\x4a\x08\xa2\xb3\x89\xe4\x0e\x67\xa5\xcc\x05\xf6\x9e\x94\xd1\xdc\x92\xe6\xed\x74\x73\x93\xb9\xc2\x6d\xf7\xb7\xe8\xfc\xda\xca\x5a\xaa\x49\x8c\xc8\xbc\xc1\x93\xbf\xdf\x4c\xfe\x7d\x36\xf9\x39\x3c\xff\x4d\xfe\xb3\xf8\xa9\xe8\x8e\x86\x1d\xa9\x03\x97\x4b\x41\xd1\x92\x3f\xc5\xae\xa8\xc2\x04\xa3\xd3\x7a\x69\xc0\xd3\x49\x1f\xa8\xf3\xf2\xda\x9d\x02\x2e\xb4\x2c\x46\xaa\xaf\x72\x25\x65\xe0\xb9\xd4\xf1\xda\x2d\x52\x7b\xc7\xf6\x2a\x06\xf5\xe4\x7c\xc4\x1e\xfd\x8c\x77\x31\xc3\x36\x70\xf1\xac\x8f\x8e\x9c\xcf\x4e\xdc\x31\x7f\xf4\xb3\x1e\xd6\xd8\x76\x5b\xee\xcb\xff\x3b\x8c\xbf\x88\xb8\x33\xc4\x52\xf3\xc5\x07\x8f\xbf\xa5\x66\x99\x7d\x47\xde\xc7\x80\x13\x16\x6d\xf4\xf3\x7f\x05\xd8\x3c\x2e\x37\x0a\xc0\xbf\x85\xcc\xd1\xcf\xf4\x24\xc0\x1a\x3a\x06\x47\xb5\xaf\x8e\x18\x47\x4c\xd2\x2f\x04\x8b\x38\x9f\x8f\xac\x57\xc4\xd7\x9c\xcf\xfa\x19\x6e\x60\xcb\xba\xc3\xc6\x91\xe9\x8a\xf0\x68\xcc\x1e\x4d\x8c\x1f\x15\x70\xda\xe6\x91\x3e\xc9\x4e\x76\x3e\xf5\x75\x0e\x9b\xdb\xbe\xc5\x89\x6b\x45\x72\x96\x8e\x76\x7c\xfb\x0e\x98\xa8\x64\x36\x71\x88\x1d\x46\x03\xce\x84\x76\x58\x1b\x96\xef\x91\xaf\xf0\x21\xed\x90\xe6\x8a\x4a\x33\xe3\x75\xdc\xa5\x6a\x5a\x85\xeb\x99\x4c\x53\x0d\xe3\x8e\x52\x0b\xb0\x23\xe2\x9e\x05\x14\x32\x41\x94\x70\x4b\xd1\xef\x81\x5b\x35\x8c\xee\xb8\xe1\xf5\x7b\x70\xee\xf9\x4a\x05\x99\x38\x1c\xee\x2b\x24\xa6\xa2\xb8\x1f\x54\x05\x1d\xf2\xd9\xb2\xa4\x72\x87\x38\x2c\x83\x7b\x58\x2b\xb4\x19\xaf\x67\x92\xb0\x53\x6f\xd6\xf4\x02\xda\x8a\x33\x62\xa2\x55\xbb\xe6\x28\xb7\x85\xde\x2c\xea\x0f\xb9\x15\x9f\x72\x31\xcd\xe1\x02\x92\x45\x4c\x13\xb1\xbc\x10\x3e\x28\x72\xe3\x45\x8c\x35\x84\x4f\x3b\x21\xa4\xd6\x87\x78\xf6\xc6\x26\x18\x53\x20\x7a\x03\x12\xdc\x91\x5c\x47\x91\x5c\x05\x90\x74\x8b\x41\x52\xef\x06\xc9\xc4\xb4\x90\x12\xe7\xd0\x52\xcc\x58\xee\xce\x94\xbc\x9e\xe6\x31\xc9\x82\xff\xd2\x10\x92\xc2\x21\x7a\xaa\x25\x09\x5b\x96\x7a\x9b\x6e\xae\xf0\x2a\xaf\xae\xec\x7a\xb5\x0a\xcf\x96\x9c\xb9\xc6\xf0\x94\x1e\xa2\xe8\xc7\x2d\xe4\x7e\x5b\x8e\x41\x09\x65\x47\x06\x29\x18\x84\x9d\xcd\x91\x41\x55\x0c\x92\x11\x7c\xf0\xeb\x86\xe2\xd0\xa0\x25\x87\xfa\xfd\xaa\x42\x03\xfa\xf7\x5d\x33\x6a\x80\x15\x1d\x4a\xb5\x5a\x29\x5c\x33\xa5\xaf\xd3\x33\xe7\x9f\xfa\xfa\x1d\x52\xbe\xe5\x6c\xb1\xe0\x15\xf8\x96\x55\x48\x9b\x4b\x59\xd1\xda\xb7\x36\x49\x40\xb1\x4d\x1c\xcc\x29\x6b\x56\x94\xf5\x4e\x18\xa2\xf0\x72\xb1\x79\xe9\x58\x1b\xed\xe8\x3e\x30\xc2\x61\x10\xc8\x71\x5a\x48\x29\x01\x45\xad\x6f\x05\x9f\x2a\xea\x7d\xcb\xe2\x77\x19\x18\x98\xab\xcf\x05\x9d\xe9\x5a\xdb\xa5\xea\x17\x39\xba\x1f\x3f\xbc\xdf\x8f\x7a\x82\x77\x0a\x95\x6c\x59\x27\x12\xef\x1e\x9b\x3f\x5a\x55\x67\x09\xa3\xba\xc4\xa5\xb0\xeb\x06\x1f\xb5\x6f\x58\x8d\xea\x4f\x46\xf9\xff\x27\xfc\x0c\x04\xbb\x83\xc2\xc3\x60\x9c\xe7\xe2\x1e\x45\xa9\x72\xb8\x2e\x83\x25\x59\x77\xc1\x08\x73\xce\xc8\x9d\xcf\xf5\xaf\xa2\x0f\x0c\xbb\xae\x2e\xe3\xdd\xa1\xe9\xa5\x05\x0f\x3d\xaa\xb3\xb2\x1d\xb4\xb1\xd1\xf5\xdc\x76\x3b\xa4\x36\xb7\xe4\x24\xca\xf3\x8a\x7b\xe0\x24\x1a\xd4\xe1\xe0\x28\x5d\x15\xdd\xb5\x89\x1f\xc0\x24\xf8\x3a\x00\x5e\x8b\x59\xeb\x6c\xdd\x96\x44\x29\x4b\x95\xed\x4a\x90\x3b\x5c\x21\x32\x49\xf6\xae\xf5\x3b\xb0\x43\xc9\x48\x26\x96\x76\x1e\x64\xfb\xc2\xa3\x2d\x29\xb7\xa2\xea\x67\x0f\x16\x2a\x70\x0c\x6a\x0e\xe2\x50\x29\x38\x91\xcc\xb7\x0f\x96\x97\x75\x8a\x23\xa8\x99\x80\x1b\x9d\x3f\x42\xa1\x7e\xbc\x9c\x5f\x24\x9a\x57\xc2\x63\x21\x18\xf7\x45\x49\x83\xdb\x9b\xb6\x6f\xdb\x98\x29\x8f\x67\x83\x16\xdc\x5b\x7e\xfd\xf1\xed\xcf\x52\xaf\x43\xbe\x47\x11\xfb\x17\x3d\xe4\xd5\x93\xe9\x75\xe9\xac\xee\x71\x89\xe1\x4e\xc4\xd8\xb2\xe8\xc3\x85\xed\xf2\x4d\xa5\x32\x74\x86\xef\xf1\x4d\x6d\x4d\x43\x83\xfc\x94\x47\x8a\x4b\xb2\x2a\x9c\x55\xea\x2a\x2a\xbe\x6b\x51\xb3\x96\x4a\x35\xa7\xe6\x10\xde\x4a\xd6\x2f\x87\xef\xd3\x0b\xd8\x0b\xcf\xba\x98\x6f\x15\xe0\x8a\x49\xd7\x62\x72\xfd\xc5\xd1\xb7\xa9\x01\x78\x48\xbe\xfb\x80\x1e\xc4\x49\x1b\x48\x91\x61\x7e\xbc\x9c\x9f\xc7\xe6\x51\xe3\x89\x5c\x15\x2f\xe6\x2b\xf7\xb5\xda\x1b\xe8\xd3\x66\xe1\xcf\x60\xc6\x19\x07\x5b\xbe\x12\x6a\xdc\xb3\xd3\x89\x7a\xec\xda\x31\x09\xad\xb7\xeb\xdd\x79\x67\xe5\x97\x2d\x84\xa1\xcd\x73\xb8\x4d\x5e\xb9\x40\xca\xe3\x44\xfd\xbd\x9b\x2d\x66\x93\x95\x97\x5c\xb6\xbb\x36\x09\xb4\x56\x15\x05\xb5\x5e\xba\x9e\xdf\x03\x7f\x74\x7d\xfe\xe8\x52\x8f\x3f\x9a\x60\x01\x87\x47\x1a\xea\x90\xf7\xb8\x02\x6c\x3e\xba\x98\x1b\xab\x47\xce\xb0\x01\x3f\x96\x9e\x89\x90\x93\x78\xe9\x02\xab\x48\xa5\x68\xb4\xd4\x6c\xa6\xce\xe1\xd0\xec\x82\x5d\xb2\xcc\xd5\xda\x76\xe1\x5d\x14\xd4\x4d\x39\xdb\x5e\x57\xdd\xdd\xca\x30\x68\x62\x8d\x26\x22\x9c\x29\xfd\xd7\x2a\xdc\x8e\xcb\x45\xf8\xa7\xdb\x72\x51\x38\x26\x26\x37\x9b\x74\xfa\xfb\xca\x1b\x53\x5b\x76\xfa\x95\xf7\x68\xab\xef\x95\xca\xdd\xa9\x6e\xf7\x27\xb6\xe4\x4d\xab\x2e\x2f\xec\x41\xa8\x70\xc0\x90\x81\xe4\xa1\x2c\xa4\xd7\x7d\x51\x3e\x9a\xa5\xa6\x77\x97\x2e\x55\x6b\x57\x40\x7f\x8e\xb2\xd7\x6b\x03\x8a\xd7\x19\x7a\x2e\x46\xb2\xdb\xe2\xdb\x58\x89\x5a\x4b\xe3\xe6\xc5\xe4\xff\x17\xcf\x7f\x79\x7d\x7b\x6b\x7e\x7b\xf9\xf4\xf8\xf1\x63\xd9\xf9\x3a\x50\x56\xb4\xf2\x04\x95\x91\x91\x36\xaf\x9a\x94\x2d\xb7\x81\x5e\xd6\xae\xd6\x4a\x33\x96\x97\x7a\xd6\xbe\x0b\x22\x20\x7d\x8d\xef\xbb\xbb\xb4\xdb\x79\xed\x68\xef\x9c\xca\x1d\xae\x78\x72\xd4\x3a\x0b\xb0\xd5\xe3\x49\x3d\xe5\xf0\x3d\x3d\x5a\x25\x89\x86\xe5\x64\x50\x78\x72\x35\xee\x82\xb0\x63\xa2\x24\x38\x8a\x30\x03\x64\xfb\xdc\xc7\x96\xb5\x41\xe9\x85\xbb\x4c\x5e\xc5\xbb\x1f\x72\x01\xe6\xc2\xb1\x27\x92\xc6\x9c\x53\xb9\x90\x5a\x7f\x29\x16\xa7\xa6\x22\xea\x1a\x12\x11\xc7\xb5\x15\x34\x9c\x20\xcc\x91\x05\xd8\xe3\x88\x3a\x80\x88\x87\x52\xd4\xdb\x53\x17\x27\x06\x04\xe2\x8c\x62\xdc\x70\x51\x12\x64\x79\xc7\x12\xf4\x4e\x31\x5b\x36\x5f\xdd\x86\xd2\x10\x5a\x09\xd3\x72\xe2\x46\xe4\x31\xb5\x01\x65\xe6\x17\x32\xd8\xa1\x1c\x85\xc9\x17\x30\x11\x59\xa6\xbc\x27\x1e\x72\x19\x78\x01\x74\xd1\xf7\x2f\x05\x77\x8b\xf8\x4a\x05\x50\x4f\x4c\xc2\xa0\x30\xa3\x55\xfa\xb2\x28\xb4\x3c\xb5\xda\xed\x3b\x63\x93\x64\xb4\xf6\x05\x1f\x31\x4b\xb6\x2f\x58\x81\xea\x21\x97\xdd\xe2\xe3\xb5\x46\x48\x6e\x96\xa3\x7c\x4b\x16\x5a\x4f\x22\x24\x3b\xa8\x77\x25\x9d\x36\xc8\x63\xfe\x7d\x27\xf9\x77\x85\xd3\x78\x4c\x95\x77\x80\x74\x4c\x95\x1f\x53\xe5\xdf\x53\xaa\xbc\xe2\xd8\x2b\x2e\xf7\xc9\x1e\x74\x00\x49\xf6\x56\x70\xd2\x6d\xf4\x3d\x35\x14\xc7\x60\xc5\xbe\x9b\x0e\x50\x1f\x29\xd2\xbe\xc8\x4a\x3a\x37\xf5\xe0\x95\x90\x88\x83\x5c\xcc\x57\xa2\xcd\xa8\x3d\xf4\xf4\x4b\x7e\x85\x12\xc6\x2a\x6f\x9e\xca\xf6\xf9\xb0\x55\x5a\xc3\x85\x62\xa5\x5d\x15\xd5\x7f\x51\xb5\xec\x4a\x37\xb0\x9a\x49\x6c\xcc\x19\x59\x8f\x8b\x86\x8a\xd0\xd8\xc2\x77\x60\xc9\x3e\x94\xeb\xcb\x50\x91\x0f\x08\x55\xd7\x99\xc9\x90\x57\x78\x24\xd5\xf8\x0b\x89\xa3\x18\x11\xe5\x66\x22\x8d\x86\x37\x56\x94\x2f\x3e\xb0\x42\x58\x5c\xa7\xb4\x23\xf0\x43\x8b\x7b\x49\x99\x2d\x95\xa9\x8b\x0d\x78\x07\x16\xb1\x09\xcf\x1d\xbd\xf3\xb4\x12\xb7\xba\x83\x09\xe0\x5e\x46\x5b\x50\x1b\xd5\x08\x31\x1a\x48\x64\xd1\xfd\x9e\xe1\x64\xb6\x92\xde\xd2\xd1\x2e\x34\x5d\x06\xa7\x8d\xaf\x06\xa5\x9f\x09\x0c\xc7\xd7\x18\xfe\x56\x8c\x61\x10\x5d\x95\x6e\x76\x73\x69\x2e\xbb\x65\x0c\xb6\x10\x24\x1a\x3c\x0e\xbb\x17\xae\xb6\x14\x33\xbd\x27\xd8\xed\xba\x8c\x82\x5e\x14\x9e\xb1\xcc\xa9\x46\x8d\x5b\x23\x67\xe1\x9b\xab\xd9\xef\xb0\x51\xbd\x8d\xd9\xce\xa1\xfb\xed\xfa\xfa\x4a\x0f\xa4\xcb\x37\x3e\x5f\xbd\xd4\x04\xcb\x05\x67\x66\x9e\x53\xc7\x01\x83\x2b\x41\x4a\x6d\x50\xca\x9c\xce\xc6\x18\xf6\xce\xa9\x64\x13\x57\x5f\x59\xbd\xa5\xe7\x6d\x2a\x97\xfc\x0e\x9b\xda\x88\xf4\x70\xb1\xaa\x06\x38\x4a\x76\xe5\xd4\xff\x29\x34\x96\x37\x02\x39\x39\x07\xf6\xc3\x07\x12\xb3\xea\xac\x83\x5e\x34\xfe\xa4\xa0\x49\x8d\xf5\xce\x2b\xce\x5f\xcd\xb5\xdc\x3c\x77\x80\x19\xb0\x5f\x7b\x3c\x5d\xd2\xef\x07\xac\xb4\x18\xcc\x8a\x73\x57\xa1\x5f\xdb\x55\x90\x8e\x67\xd9\xb7\xa1\x10\x9a\xbb\x62\x52\x71\xd7\xf1\x2e\x4f\x7d\xd0\xe5\xd3\xcd\xdb\xbb\xc5\xcd\x05\x2c\x6e\xde\xe0\xc5\xcd\x9c\x45\x7f\x9f\xb3\xc5\x33\x8d\x1e\xe7\x1f\xd4\x41\x32\xea\x64\x89\x84\xaa\x24\x82\xa8\xa7\xc2\x47\x75\x79\x62\x37\x16\x2a\x73\x1c\x35\xbc\xed\xc6\xdd\xaa\x84\x47\x9d\x2f\x2a\xdd\x8e\xb5\xed\x83\x4b\x8b\x3e\xb4\x29\x5c\xd6\x62\xc9\x14\x07\x24\xd5\xef\x15\x11\x72\xcd\x2a\xc6\x03\x2e\xfd\x9a\x27\xe6\xa0\xb7\x9d\x2a\xbf\x49\x9b\xec\x69\x7e\x92\x7f\xb5\xa9\x18\xd1\xa4\x06\xf9\xf9\xeb\xb5\xa1\x84\xae\xd6\xa2\x93\x03\xd6\x95\x4c\xf5\x7b\xa4\x44\x89\xed\x5a\xc4\x20\x0d\x7f\xfe\x71\x16\xf7\x4e\xe7\x56\xd5\xd6\x7b\xde\x03\x65\xa5\x83\x9f\x2a\x26\x1d\xf5\xae\x01\x6a\x58\x04\x1c\x7e\xce\xc0\x04\x87\x13\x6c\x35\x5c\x21\xce\x8b\xc3\xd4\x33\x04\xcb\x13\x65\xe4\xef\x30\x74\x7d\x2e\x49\x17\x2a\x8e\x73\xc5\x61\x32\x42\x76\xa8\x28\x65\xa1\x75\x5e\x4a\x04\x0e\x05\xb6\x28\x24\x32\xa9\xdb\xe6\x32\x4c\x09\xd6\x50\x76\xcd\x60\xc9\xc0\x5b\x0d\x39\x45\x4c\xbb\xf6\x10\x4c\xab\x97\xc1\x76\xa0\x5a\x65\xd3\xed\xbe\x4b\xd1\xcf\xd0\x57\xa5\x52\x18\x47\x55\x3a\x38\x55\x92\xaf\xd3\x47\x75\x3a\xaa\x53\x27\x75\xaa\xd8\x94\x07\xd9\xfd\x74\xe8\xdb\xd6\x76\xc4\xa3\x62\x1f\xae\x62\x87\x8f\xd5\xf5\x70\xf8\x7b\x5e\x9b\x4d\x87\xcf\x6b\x9f\x29\xe8\x2e\xdb\x6e\x3f\xa1\xd9\x40\xbe\xea\x3c\x95\xec\x0a\xab\xc6\x50\xa8\x57\xfc\x55\x6c\xa4\x3e\x4f\x14\x6f\xe7\xed\x2a\xcc\x59\x2a\xea\x2f\x84\xff\xd2\x3b\xf4\x68\x66\x86\xf7\x08\xb2\x86\x39\x2c\xeb\x6e\x44\x54\x5c\x12\xc8\xab\xe8\x89\xf4\x43\xa0\x7c\xe9\x07\xf9\x55\x88\xf4\x36\x6b\x03\x5b\xa9\xbf\xb2\xd5\xec\x61\x91\xfe\x12\x12\x48\xb8\x28\x3f\x90\xd7\xda\xdc\xe3\xec\xf0\x75\x7d\xe4\x46\xaa\xb7\xc7\xf7\x57\x5b\xdf\x55\xd6\xf4\x46\x98\xa4\x0e\xa8\xba\x06\xa8\xb2\xfe\x47\x51\xfb\xa3\x7a\xed\x72\xaf\x8a\xa5\x3b\xec\x95\xa3\xe0\xbf\xa7\xd1\x7f\x03\x00\x00\xff\xff\x47\x33\xeb\x7e\xa8\x8b\x00\x00"

func _303JsonBytes() ([]byte, error) {
	return bindataRead(
		__303Json,
		"3.0.3.json",
	)
}

func _303Json() (*asset, error) {
	bytes, err := _303JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3.0.3.json", size: 35752, mode: os.FileMode(420), modTime: time.Unix(1648637554, 0)}
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
	"3.0.3.json": _303Json,
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
	"3.0.3.json": {_303Json, map[string]*bintree{}},
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
