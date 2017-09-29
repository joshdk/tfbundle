// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package cmd

import (
	"testing"

	"github.com/palantir/pkg/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommand(t *testing.T) {
	app := Cmd()

	called := false

	app.Action = func(ctx cli.Context) error {
		called = true
		return nil
	}

	status := app.Run([]string{"tfbundle", "test.txt", "module.tgz"})

	require.Zero(t, status)

	assert.True(t, called)

}
