// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package bundle_test

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/joshdk/tfbundle/bundle"
	"github.com/joshdk/tfbundle/cmd"
)

const (
	artifactContents = `this is a test`
	shimContents     = `output "filename" {
  value = "${path.module}/artifact/test.txt"
}

output "source_code_hash" {
  value = "${base64sha256(file("${path.module}/artifact/test.txt"))}"
}
`
)

func TestBundle(t *testing.T) {

	tests := []struct {
		name   string
		action func(t *testing.T, path string, module string)
	}{
		{
			name: "bundle cli",
			action: func(t *testing.T, path string, module string) {
				app := cmd.Cmd()

				args := []string{"tfbundle", path, module}

				status := app.Run(args)

				assert.Equal(t, 0, status)
			},
		},
		{
			name: "bundle file",
			action: func(t *testing.T, path string, module string) {
				moduleContents, err := bundle.File(path)

				require.Nil(t, err)

				err = ioutil.WriteFile(module, moduleContents, 0644)

				require.Nil(t, err)
			},
		},
		{
			name: "bundle reader",
			action: func(t *testing.T, path string, module string) {
				reader, err := os.Open(path)

				require.Nil(t, err)

				moduleContents, err := bundle.Reader(reader, path)

				require.Nil(t, err)

				err = ioutil.WriteFile(module, moduleContents, 0644)

				require.Nil(t, err)
			},
		},
		{
			name: "bundle string",
			action: func(t *testing.T, path string, module string) {
				content, err := ioutil.ReadFile(path)

				require.Nil(t, err)

				moduleContents, err := bundle.String(string(content), path)

				require.Nil(t, err)

				err = ioutil.WriteFile(module, moduleContents, 0644)

				require.Nil(t, err)
			},
		},
		{
			name: "bundle bytes",
			action: func(t *testing.T, path string, module string) {
				content, err := ioutil.ReadFile(path)

				require.Nil(t, err)

				moduleContents, err := bundle.Bytes(content, path)

				require.Nil(t, err)

				err = ioutil.WriteFile(module, moduleContents, 0644)

				require.Nil(t, err)
			},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)

		t.Run(name, func(t *testing.T) {

			tempDir, err := ioutil.TempDir("", "")
			defer func() {
				if err := os.RemoveAll(tempDir); err != nil {
					panic(err.Error())
				}
			}()

			require.Nil(t, err)

			artifactPath := path.Join(tempDir, "test.txt")

			modulePath := path.Join(tempDir, "module.tgz")

			err = ioutil.WriteFile(artifactPath, []byte(artifactContents), 0600)

			require.Nil(t, err)

			test.action(t, artifactPath, modulePath)

			checkArchive(t, modulePath)
		})

	}

}

func checkArchive(t *testing.T, archivePath string) {

	archiveContents, err := ioutil.ReadFile(archivePath)

	require.Nil(t, err)

	// Open the tar archive for reading.
	r := bytes.NewReader(archiveContents)

	gr, err := gzip.NewReader(r)

	require.Nil(t, err)

	tr := tar.NewReader(gr)

	type entry struct {
		name    string
		flag    byte
		mode    int64
		content []byte
	}

	actual := []entry{}

	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}

		require.Nil(t, err)

		contents, err := ioutil.ReadAll(tr)

		require.Nil(t, err)

		actual = append(actual, entry{
			name:    hdr.Name,
			flag:    hdr.Typeflag,
			mode:    hdr.Mode,
			content: contents,
		})
	}

	expected := []entry{
		{
			name:    "main.tf",
			flag:    tar.TypeReg,
			mode:    0644,
			content: []byte(shimContents),
		},
		{
			name:    "artifact",
			flag:    tar.TypeDir,
			mode:    0755,
			content: []byte{},
		},
		{
			name:    "artifact/test.txt",
			flag:    tar.TypeReg,
			mode:    0644,
			content: []byte(artifactContents),
		},
	}

	assert.Equal(t, expected, actual)
}
