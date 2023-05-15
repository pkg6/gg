// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
	"github.com/go-fonts/liberation/liberationsansregular"
)

func ExampleText() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	err := dc.LoadFontFaceFromBytes(liberationsansregular.TTF, 96)
	if err != nil {
		log.Fatalf("could not load font: %+v", err)
	}
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)

	err = dc.SavePNG("testdata/text.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestText(t *testing.T) {
	chkimg(ExampleText, t, "text.png")
}
