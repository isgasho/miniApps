package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(src io.Reader, dest string) error {
	var size int64
	buff := bytes.NewBuffer([]byte{})
	_, err := io.Copy(buff, src)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(buff.Bytes())
	size = int64(reader.Len())
	r, err := zip.NewReader(reader, size)
	if err != nil {
		return err
	}
	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: invalid filepath", fpath)
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}
		outfile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(outfile, rc)
		outfile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func fileOut(data io.Reader, filename string) (io.Reader, error) {
	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2)
	if _, err := io.Copy(w, data); err != nil {
		return nil, err
	}
	myData, err := ioutil.ReadAll(&buf1)
	if err != nil {
		return nil, err
	}
	ioutil.WriteFile(filename, myData, 0766)
	return &buf2, err
}
