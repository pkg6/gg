// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"codeberg.org/go-fonts/xolonium/xoloniumregular"
	"git.sr.ht/~sbinet/gg"
)

func Example_unicode() {
	const S = 4096 * 2
	const T = 16 * 2
	const F = 28
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	err := dc.LoadFontFaceFromBytes(xoloniumregular.TTF, F)
	if err != nil {
		log.Fatalf("could not load font: %+v", err)
	}
	for r := 0; r < 256; r++ {
		for c := 0; c < 256; c++ {
			i := r*256 + c
			x := float64(c*T) + T/2
			y := float64(r*T) + T/2
			r := rune(i)
			dc.DrawStringAnchored(string(r), x, y, 0.5, 0.5)
		}
	}

	err = dc.SavePNG("testdata/unicode.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestUnicode(t *testing.T) {
	chkimg(Example_unicode, t, "unicode.png")
}
