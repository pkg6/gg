// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"math/rand"
	"testing"

	"git.sr.ht/~sbinet/gg"
	"github.com/go-fonts/liberation/liberationsansbold"
	"github.com/go-fonts/liberation/liberationsansregular"
)

func ExampleScatter() {
	const S = 1024
	const P = 64
	dc := gg.NewContext(S, S)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	points := CreatePoints(1000)
	dc.Translate(P, P)
	dc.Scale(S-P*2, S-P*2)
	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		dc.MoveTo(x, 0)
		dc.LineTo(x, 1)
		dc.MoveTo(0, x)
		dc.LineTo(1, x)
	}
	dc.SetRGBA(0, 0, 0, 0.25)
	dc.SetLineWidth(1)
	dc.Stroke()
	// draw axes
	dc.MoveTo(0, 0)
	dc.LineTo(1, 0)
	dc.MoveTo(0, 0)
	dc.LineTo(0, 1)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(4)
	dc.Stroke()
	// draw points
	dc.SetRGBA(0, 0, 1, 0.5)
	for _, p := range points {
		dc.DrawCircle(p.X, p.Y, 3.0/S)
		dc.Fill()
	}
	// draw text
	dc.Identity()
	dc.SetRGB(0, 0, 0)

	err := dc.LoadFontFaceFromBytes(liberationsansbold.TTF, 24)
	if err != nil {
		log.Fatalf("could not load font: %+v", err)
	}
	dc.DrawStringAnchored("Chart Title", S/2, P/2, 0.5, 0.5)

	err = dc.LoadFontFaceFromBytes(liberationsansregular.TTF, 18)
	if err != nil {
		log.Fatalf("could not load font: %+v", err)
	}
	dc.DrawStringAnchored("X Axis Title", S/2, S-P/2, 0.5, 0.5)

	err = dc.SavePNG("testdata/scatter.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func CreatePoints(n int) []gg.Point {
	var (
		rnd = rand.New(rand.NewSource(1234))
		pts = make([]gg.Point, n)
	)
	for i := 0; i < n; i++ {
		x := 0.5 + rnd.NormFloat64()*0.1
		y := x + rnd.NormFloat64()*0.1
		pts[i] = gg.Point{x, y}
	}
	return pts
}

func TestScatter(t *testing.T) {
	chkimg(ExampleScatter, t, "scatter.png")
}
