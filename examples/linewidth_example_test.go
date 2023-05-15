// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleLinewidth() {
	dc := gg.NewContext(1000, 1000)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	w := 0.1
	for i := 100; i <= 900; i += 20 {
		x := float64(i)
		dc.DrawLine(x+50, 0, x-50, 1000)
		dc.SetLineWidth(w)
		dc.Stroke()
		w += 0.1
	}

	err := dc.SavePNG("testdata/linewidth.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestLinewidth(t *testing.T) {
	chkimg(ExampleLinewidth, t, "linewidth.png")
}
