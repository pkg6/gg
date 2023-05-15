// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleCrisp() {
	const W = 1000
	const H = 1000
	const Minor = 10
	const Major = 100

	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// minor grid
	for x := Minor; x < W; x += Minor {
		fx := float64(x) + 0.5
		dc.DrawLine(fx, 0, fx, H)
	}
	for y := Minor; y < H; y += Minor {
		fy := float64(y) + 0.5
		dc.DrawLine(0, fy, W, fy)
	}
	dc.SetLineWidth(1)
	dc.SetRGBA(0, 0, 0, 0.25)
	dc.Stroke()

	// major grid
	for x := Major; x < W; x += Major {
		fx := float64(x) + 0.5
		dc.DrawLine(fx, 0, fx, H)
	}
	for y := Major; y < H; y += Major {
		fy := float64(y) + 0.5
		dc.DrawLine(0, fy, W, fy)
	}
	dc.SetLineWidth(1)
	dc.SetRGBA(0, 0, 0, 0.5)
	dc.Stroke()

	err := dc.SavePNG("testdata/crisp.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestCrisp(t *testing.T) {
	chkimg(ExampleCrisp, t, "crisp.png")
}
