// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build example

package main

import (
	"flag"
	"image/color"
	"log"

	"git.sr.ht/~sbinet/gg"
)

func main() {
	var (
		fname = flag.String("o", "gradient-radial.png", "output file")
	)
	flag.Parse()

	dc := gg.NewContext(400, 200)

	grad := gg.NewRadialGradient(100, 100, 10, 100, 120, 80)
	grad.AddColorStop(0, color.RGBA{0, 255, 0, 255})
	grad.AddColorStop(1, color.RGBA{0, 0, 255, 255})

	dc.SetFillStyle(grad)
	dc.DrawRectangle(0, 0, 200, 200)
	dc.Fill()

	dc.SetColor(color.White)
	dc.DrawCircle(100, 100, 10)
	dc.Stroke()
	dc.DrawCircle(100, 120, 80)
	dc.Stroke()

	err := dc.SavePNG(*fname)
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}
