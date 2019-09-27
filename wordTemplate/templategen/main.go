package templategen

import (
	"bytes"
	"io"
	"path/filepath"
	"text/template"
)

func GenerateTemplate(templateDir string, extension string, firstTemplateName string, templateFns map[string]interface{}, data interface{}) (io.Reader, error) {

	globalFiles := filepath.Join(templateDir, extension)
	tmpl, err := template.New(firstTemplateName).Funcs(templateFns).ParseGlob(globalFiles)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}
