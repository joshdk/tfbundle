// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package bundle

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/joshdk/tfbundle/shim"
)

func File(path string) ([]byte, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return Reader(reader, path)
}

func Reader(content io.Reader, name string) ([]byte, error) {
	body, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, err
	}

	return Bytes(body, name)
}

func String(content string, name string) ([]byte, error) {
	return Bytes([]byte(content), name)
}

func Bytes(content []byte, name string) ([]byte, error) {

	var buf bytes.Buffer

	gzipWriter := gzip.NewWriter(&buf)

	tarWriter := tar.NewWriter(gzipWriter)

	base := path.Base(name)

	rendered, err := shim.Render(base)
	if err != nil {
		return nil, err
	}

	if err := writeFile("main.tf", rendered, tarWriter); err != nil {
		return nil, err
	}

	if err := writeDir("artifact", tarWriter); err != nil {
		return nil, err
	}

	if err := writeFile("artifact/"+base, content, tarWriter); err != nil {
		return nil, err
	}

	if err := tarWriter.Close(); err != nil {
		return nil, err
	}

	if err := gzipWriter.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func writeFile(name string, contents []byte, tarWriter *tar.Writer) error {

	var now = time.Now()

	header := tar.Header{
		Name:       name,
		Size:       int64(len(contents)),
		Mode:       0644,
		Typeflag:   tar.TypeReg,
		AccessTime: now,
		ChangeTime: now,
		ModTime:    now,
	}

	if err := tarWriter.WriteHeader(&header); err != nil {
		return err
	}

	if _, err := tarWriter.Write(contents); err != nil {
		return err
	}

	return nil
}

func writeDir(name string, tarWriter *tar.Writer) error {

	var now = time.Now()

	header := tar.Header{
		Name:       name,
		Mode:       0755,
		Typeflag:   tar.TypeDir,
		AccessTime: now,
		ChangeTime: now,
		ModTime:    now,
	}

	if err := tarWriter.WriteHeader(&header); err != nil {
		return err
	}

	return nil
}
