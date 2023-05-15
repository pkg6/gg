// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"math/rand"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleLines() {
	const W = 1024
	const H = 1024

	var (
		rnd = rand.New(rand.NewSource(1234))
		dc  = gg.NewContext(W, H)
	)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	for i := 0; i < 1000; i++ {
		x1 := rnd.Float64() * W
		y1 := rnd.Float64() * H
		x2 := rnd.Float64() * W
		y2 := rnd.Float64() * H
		r := rnd.Float64()
		g := rnd.Float64()
		b := rnd.Float64()
		a := rnd.Float64()*0.5 + 0.5
		w := rnd.Float64()*4 + 1
		dc.SetRGBA(r, g, b, a)
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}

	err := dc.SavePNG("testdata/lines.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestLines(t *testing.T) {
	chkimg(ExampleLines, t, "lines.png")
}
