// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"math"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleConcat() {
	im1, err := gg.LoadPNG("testdata/baboon.png")
	if err != nil {
		panic(err)
	}

	im2, err := gg.LoadPNG("testdata/gopher.png")
	if err != nil {
		panic(err)
	}

	s1 := im1.Bounds().Size()
	s2 := im2.Bounds().Size()

	width := int(math.Max(float64(s1.X), float64(s2.X)))
	height := s1.Y + s2.Y

	dc := gg.NewContext(width, height)
	dc.DrawImage(im1, 0, 0)
	dc.DrawImage(im2, 0, s1.Y)

	err = dc.SavePNG("testdata/concat.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestConcat(t *testing.T) {
	chkimg(ExampleConcat, t, "concat.png")
}
