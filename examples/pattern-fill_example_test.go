// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExamplePatternFill() {
	im, err := gg.LoadPNG("testdata/baboon.png")
	if err != nil {
		panic(err)
	}
	pattern := gg.NewSurfacePattern(im, gg.RepeatBoth)
	dc := gg.NewContext(600, 600)
	dc.MoveTo(20, 20)
	dc.LineTo(590, 20)
	dc.LineTo(590, 590)
	dc.LineTo(20, 590)
	dc.ClosePath()
	dc.SetFillStyle(pattern)
	dc.Fill()

	err = dc.SavePNG("testdata/pattern-fill.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestPatternFill(t *testing.T) {
	chkimg(ExamplePatternFill, t, "pattern-fill.png")
}
