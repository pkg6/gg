// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
	"github.com/go-fonts/anton/antonregular"
)

func ExampleMeme() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	err := dc.LoadFontFaceFromBytes(antonregular.TTF, 96)
	if err != nil {
		log.Fatalf("could not load font: %+v", err)
	}
	dc.SetRGB(0, 0, 0)
	s := "ONE DOES NOT SIMPLY"
	n := 6 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := S/2 + float64(dx)
			y := S/2 + float64(dy)
			dc.DrawStringAnchored(s, x, y, 0.5, 0.5)
		}
	}
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(s, S/2, S/2, 0.5, 0.5)

	err = dc.SavePNG("testdata/meme.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestMeme(t *testing.T) {
	chkimg(ExampleMeme, t, "meme.png")
}
