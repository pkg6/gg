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

func ExampleStar() {
	makePoints := func(n int, x, y, r float64) []gg.Point {
		result := make([]gg.Point, n)
		for i := 0; i < n; i++ {
			a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
			result[i] = gg.Point{x + r*math.Cos(a), y + r*math.Sin(a)}
		}
		return result
	}

	var (
		n      = 5
		points = makePoints(n, 512, 512, 400)
		dc     = gg.NewContext(1024, 1024)
	)
	dc.SetHexColor("fff")
	dc.Clear()
	for i := 0; i < n+1; i++ {
		index := (i * 2) % n
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	dc.SetRGBA(0, 0.5, 0, 1)
	dc.SetFillRule(gg.FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetRGBA(0, 1, 0, 0.5)
	dc.SetLineWidth(16)
	dc.Stroke()

	err := dc.SavePNG("testdata/star.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestStar(t *testing.T) {
	chkimg(ExampleStar, t, "star.png")
}
