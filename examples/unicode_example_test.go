// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
	"github.com/go-fonts/xolonium/xoloniumregular"
)

func ExampleUnicode() {
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
	chkimg(ExampleUnicode, t, "unicode.png")
}
