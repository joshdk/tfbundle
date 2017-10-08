// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

//go:generate go-bindata -o data.go -pkg shim -nocompress -nometadata main.tf

// Package shim implements functions for rendering the Terraform shim.
package shim

import (
	"bytes"
	"text/template"
)

func Raw() []byte {
	return MustAsset("main.tf")
}

//Render templates the given file name over the Terraform shim.
func Render(file string) ([]byte, error) {

	var buf bytes.Buffer
	var ctx = map[string]string{"Artifact": file}

	tpl, err := template.New("main.tf").Parse(string(Raw()))
	if err != nil {
		return nil, err
	}

	if err := tpl.Execute(&buf, ctx); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
