// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleClip() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(350, 500, 300)
	dc.Clip()
	dc.DrawCircle(650, 500, 300)
	dc.Clip()
	dc.DrawRectangle(0, 0, 1000, 1000)
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	err := dc.SavePNG("testdata/clip.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestClip(t *testing.T) {
	chkimg(ExampleClip, t, "clip.png")
}
