package minify

import (
	"bytes"
	"io"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/xml"
)

func MinifyCustomXml(reader io.Reader) (io.Reader, error) {
	xml := Minifier{}
	xml.KeepWhitespace = false

	var buf bytes.Buffer
	if err := xml.Minify(&buf, reader); err != nil {
		return nil, err
	}
	return &buf, nil
}

func MinifyXml(reader io.Reader) (io.Reader, error) {
	m := minify.New()
	m.AddFunc("text/html", xml.Minify)
	m.Add("text/html", &xml.Minifier{KeepWhitespace: false})
	var buf bytes.Buffer
	if err := m.Minify("text/html", &buf, reader); err != nil {
		return nil, err
	}
	return &buf, nil
}

func MinifyHtml(reader io.Reader) (io.Reader, error) {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.Add("text/html", &html.Minifier{
		KeepEndTags:      true,
		KeepDocumentTags: true,
	})
	var buf bytes.Buffer
	if err := m.Minify("text/html", &buf, reader); err != nil {
		return nil, err
	}
	return &buf, nil
}
