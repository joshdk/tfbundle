// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package shim_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/joshdk/tfbundle/shim"
)

func TestRaw(t *testing.T) {
	actual := shim.Raw()

	assert.NotEmpty(t, actual)
}

func TestRender(t *testing.T) {
	actual, err := shim.Render("lambda.zip")

	require.Nil(t, err)

	assert.NotEmpty(t, actual)
}
