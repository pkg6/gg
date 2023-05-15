// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleMask() {
	im, err := gg.LoadImage("testdata/baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(512, 512)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	dc.Clip()
	dc.DrawImage(im, 0, 0)

	err = dc.SavePNG("testdata/mask.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestMask(t *testing.T) {
	chkimg(ExampleMask, t, "mask.png")
}
