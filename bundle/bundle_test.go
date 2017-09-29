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
	"io/ioutil" //"os"
	"os"
	"path" //"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/joshdk/tfbundle/bundle"
	"github.com/joshdk/tfbundle/cmd"
)

const (
	artifactContents = `this is a test`
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

	fmt.Println(archiveContents)

	require.Nil(t, err)

	// Open the tar archive for reading.
	r := bytes.NewReader(archiveContents)

	gr, err := gzip.NewReader(r)

	require.Nil(t, err)

	tr := tar.NewReader(gr)

	type entry struct {
		flag byte
		mode int64
	}

	actual := map[string]entry{}

	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}

		require.Nil(t, err)

		fmt.Printf("Contents of %s:\n", hdr.Name)

		_, found := actual[hdr.Name]
		require.False(t, found, "duplicate entries in tarball")

		actual[hdr.Name] = entry{
			flag: hdr.Typeflag,
			mode: hdr.Mode,
		}
	}

	expected := map[string]entry{
		"main.tf": {
			flag: tar.TypeReg,
			mode: 0644,
		},
		"artifact": {
			flag: tar.TypeDir,
			mode: 0755,
		},
		"artifact/test.txt": {
			flag: tar.TypeReg,
			mode: 0644,
		},
	}

	assert.Equal(t, expected, actual)

}
