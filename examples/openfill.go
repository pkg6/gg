// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build example

package main

import (
	"flag"
	"log"
	"math"
	"math/rand"

	"git.sr.ht/~sbinet/gg"
)

func main() {
	var (
		fname = flag.String("o", "openfill.png", "output file")
	)
	flag.Parse()

	var (
		rnd = rand.New(rand.NewSource(1234))
		dc  = gg.NewContext(1000, 1000)
	)
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			x := float64(i)*100 + 50
			y := float64(j)*100 + 50
			a1 := rnd.Float64() * 2 * math.Pi
			a2 := a1 + rnd.Float64()*math.Pi + math.Pi/2
			dc.DrawArc(x, y, 40, a1, a2)
			// dc.ClosePath()
		}
	}
	dc.SetRGB(0, 0, 0)
	dc.FillPreserve()
	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(8)
	dc.StrokePreserve()
	dc.SetRGB(1, 0, 0)
	dc.SetLineWidth(4)
	dc.StrokePreserve()

	err := dc.SavePNG(*fname)
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}
