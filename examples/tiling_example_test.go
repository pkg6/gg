// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleTiling() {
	const (
		NX = 4
		NY = 3
	)
	im, err := gg.LoadPNG("testdata/gopher.png")
	if err != nil {
		panic(err)
	}
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	dc := gg.NewContext(w*NX, h*NY)
	for y := 0; y < NY; y++ {
		for x := 0; x < NX; x++ {
			dc.DrawImage(im, x*w, y*h)
		}
	}

	err = dc.SavePNG("testdata/tiling.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestTiling(t *testing.T) {
	chkimg(ExampleTiling, t, "tiling.png")
}
