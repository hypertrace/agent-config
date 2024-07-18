package main

import (
	"html/template"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Loaders struct {
	MainType  string
	Header    string
	EnvPrefix string
}

func copyTemplateFiles(srcDir, outDir string, settings Loaders) error {
	return filepath.Walk(srcDir, func(fpath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		content, err := os.ReadFile(fpath)
		if err != nil {
			return err
		}

		fileRelativePath := fpath[len(srcDir):]
		tmpl := template.Must(template.New(fileRelativePath).Parse(string(content)))

		outputFilepath := path.Join(outDir, fileRelativePath)
		if strings.HasSuffix(outputFilepath, ".tmpl") {
			outputFilepath = outputFilepath[:len(outputFilepath)-5]
		}

		err = os.MkdirAll(path.Dir(outputFilepath), 0755)
		if err != nil {
			return err
		}

		f, err := os.Create(outputFilepath)
		if err != nil {
			return err
		}
		defer f.Close()

		return tmpl.Execute(f, settings)
	})
}
