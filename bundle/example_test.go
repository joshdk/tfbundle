// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package bundle_test

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/joshdk/tfbundle/bundle"
)

func ExampleFile() {
	body, err := bundle.File("hello.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := ioutil.WriteFile("module.tgz", body, 0400); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func ExampleReader() {
	content := strings.NewReader("Hello, world!")

	body, err := bundle.Reader(content, "hello.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := ioutil.WriteFile("module.tgz", body, 0400); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func ExampleString() {
	content := "Hello, world!"

	body, err := bundle.String(content, "hello.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := ioutil.WriteFile("module.tgz", body, 0400); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func ExampleBytes() {
	content := []byte("Hello, world!")

	body, err := bundle.Bytes(content, "hello.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := ioutil.WriteFile("module.tgz", body, 0400); err != nil {
		fmt.Println(err.Error())
		return
	}
}
