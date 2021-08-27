package main

import (
	"io"
	"io/ioutil"
	"path"
	"strings"
)

type protobufImportModuleProvider struct {
	dir string
}

func (pi *protobufImportModuleProvider) Provide(module string) (io.Reader, error) {
	modulePath := path.Join(pi.dir, module)
	if strings.HasPrefix(module, "google/protobuf/") {
		modulePath = path.Join(pi.dir, "_protobuf", "src", module)
	}

	raw, err := ioutil.ReadFile(modulePath)
	if err != nil {
		return nil, err
	}

	r := strings.NewReader(string(raw[:]))
	return r, nil
}
