// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

func ExampleRotatedText() {
	const S = 400
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	font, err := opentype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(font, &opentype.FaceOptions{
		Size: 40,
		DPI:  72,
	})
	if err != nil {
		panic(err)
	}
	defer face.Close()

	dc.SetFontFace(face)
	text := "Hello, world!"
	w, h := dc.MeasureString(text)
	dc.Rotate(gg.Radians(10))
	dc.DrawRectangle(100, 180, w, h)
	dc.Stroke()
	dc.DrawStringAnchored(text, 100, 180, 0.0, 1)

	err = dc.SavePNG("testdata/rotated-text.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestRotatedText(t *testing.T) {
	chkimg(ExampleRotatedText, t, "rotated-text.png")
}
