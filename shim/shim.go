// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

// Package shim implements functions for rendering the Terraform shim.
package shim

import (
	"bytes"
	"text/template"
)

const (
	body = `output "filename" {
  value = "${path.module}/artifact/{{ .Artifact }}"
}

output "source_code_hash" {
  value = "${base64sha256(file("${path.module}/artifact/{{ .Artifact }}"))}"
}
`
)

var (
	tpl = template.Must(template.New("main.tf").Parse(body))
)

//Render templates the given file name over the Terraform shim.
func Render(file string) ([]byte, error) {

	var buf bytes.Buffer
	var ctx = map[string]string{"Artifact": file}

	if err := tpl.Execute(&buf, ctx); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
