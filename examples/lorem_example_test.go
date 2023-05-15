// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

var lines = []string{
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
	"tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,",
	"quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo",
	"consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse",
	"cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat",
	"non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
}

func ExampleLorem() {
	const (
		W = 800
		H = 400
	)

	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	const h = 24
	for i, line := range lines {
		y := H/2 - h*len(lines)/2 + i*h
		dc.DrawStringAnchored(line, 400, float64(y), 0.5, 0.5)
	}

	err := dc.SavePNG("testdata/lorem.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestLorem(t *testing.T) {
	chkimg(ExampleLorem, t, "lorem.png")
}
