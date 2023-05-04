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
	"github.com/go-fonts/anton/antonregular"
)

const (
	W = 1024
	H = 512
)

func main() {
	var (
		fname = flag.String("o", "gradient-text.png", "output file")
	)
	flag.Parse()

	dc := gg.NewContext(W, H)

	// draw text
	dc.SetRGB(0, 0, 0)

	err := dc.LoadFontFaceFromBytes(antonregular.TTF, 128)
	if err != nil {
		log.Fatalf("could not load bold font: %+v", err)
	}
	dc.DrawStringAnchored("Gradient Text", W/2, H/2, 0.5, 0.5)

	// get the context as an alpha mask
	mask := dc.AsMask()

	// clear the context
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// set a gradient
	g := gg.NewLinearGradient(0, 0, W, H)
	g.AddColorStop(0, color.RGBA{255, 0, 0, 255})
	g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	dc.SetFillStyle(g)

	// using the mask, fill the context with the gradient
	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, W, H)
	dc.Fill()

	err = dc.SavePNG(*fname)
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}
