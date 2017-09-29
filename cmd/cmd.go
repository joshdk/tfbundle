// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/palantir/pkg/cli"
	"github.com/palantir/pkg/cli/flag"

	"github.com/joshdk/tfbundle/bundle"
)

var (
	artifactParam = flag.StringParam{
		Name:  "artifact",
		Usage: "Path to artifact for bundling",
	}

	moduleParam = flag.StringParam{
		Name:  "module",
		Usage: "Path to bundled module",
	}
)

func Cmd() *cli.App {

	app := cli.NewApp()

	app.Name = "tfbundle"
	app.Description = "Bundle a single artifact as a Terraform module"

	app.Flags = []flag.Flag{
		artifactParam,
		moduleParam,
	}

	app.ErrorHandler = func(ctx cli.Context, err error) int {
		fmt.Printf("tfbundle: %s\n", err.Error())
		return 1
	}

	app.Action = func(ctx cli.Context) error {

		artifact := ctx.String(artifactParam.Name)

		module := ctx.String(moduleParam.Name)

		artifactContents, err := ioutil.ReadFile(artifact)
		if err != nil {
			return err
		}

		moduleContents, err := bundle.Bytes(artifactContents, artifact)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(module, moduleContents, 0644); err != nil {
			return err
		}

		return nil
	}

	return app
}
