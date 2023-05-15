// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleEllipse() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}

	im, err := gg.LoadImage("testdata/gopher.png")
	if err != nil {
		panic(err)
	}
	dc.DrawImageAnchored(im, S/2, S/2, 0.5, 0.5)

	err = dc.SavePNG("testdata/ellipse.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestEllipse(t *testing.T) {
	chkimg(ExampleEllipse, t, "ellipse.png")
}
