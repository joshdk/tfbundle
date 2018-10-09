// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package shim_test

import (
	"fmt"

	"github.com/joshdk/tfbundle/shim"
)

func ExampleRender() {
	body, err := shim.Render(nil, "hello.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(body)
}
