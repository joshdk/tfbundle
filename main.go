// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joshdk/tfbundle/bundle"
)

var (
	// version is replaced with the (git describe) version string at build time.
	version = "development"
)

func main() {
	if err := mainCmd(); err != nil {
		fmt.Fprintf(os.Stderr, "tfbundle: %v\n", err)
		os.Exit(1)
	}
}

func mainCmd() error {
	var (
		// flagArtifact (-artifact) is the name of a file to be read from and
		// then bundled.
		flagArtifact = flag.String("artifact", "", "File name of artifact to bundle.")

		// flagModule (-module) is the name of a file to write the bundle to.
		flagModule = flag.String("module", "", "File name of Terraform module.")

		// versionFlag (-version) causes the version string to be printed and
		// then exits immediately.
		versionFlag = flag.Bool("version", false, fmt.Sprintf("Print the version (%s) and exit.", version))
	)
	flag.Parse()

	switch {
	case *versionFlag:
		// If the version flag (-version) was passed, print the version string
		// and exit immediately.
		fmt.Println(version)
		return nil
	case *flagArtifact == "":
		// Error if the artifact flag (-artifact) is empty.
		return fmt.Errorf("no artifact file name given")
	case *flagModule == "":
		// Error if the module flag (-module) is empty.
		return fmt.Errorf("no module file name given")
	}

	moduleContents, err := bundle.File(*flagArtifact)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(*flagModule, moduleContents, 0644); err != nil {
		return err
	}

	return nil
}
