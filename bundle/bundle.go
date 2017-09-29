// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package bundle

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"path"

	"github.com/joshdk/tfbundle/shim"
)

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

	header := tar.Header{
		Name:     name,
		Size:     int64(len(contents)),
		Mode:     0644,
		Typeflag: tar.TypeReg,
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

	header := tar.Header{
		Name:     name,
		Mode:     0755,
		Typeflag: tar.TypeDir,
	}

	if err := tarWriter.WriteHeader(&header); err != nil {
		return err
	}

	return nil
}
