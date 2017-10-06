// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetNames(t *testing.T) {
	expected := []string{
		"main.tf",
	}

	actual := AssetNames()

	assert.Equal(t, expected, actual)
}

func TestAsset(t *testing.T) {
	actual := MustAsset("main.tf")

	assert.NotEmpty(t, actual)
}
