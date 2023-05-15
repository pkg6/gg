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

func ExampleGofont() {
	font, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face, err := opentype.NewFace(font, &opentype.FaceOptions{
		Size: 48,
		DPI:  72,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer face.Close()

	dc := gg.NewContext(1024, 1024)
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Hello, world!", 512, 512, 0.5, 0.5)

	err = dc.SavePNG("testdata/gofont.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestGofont(t *testing.T) {
	chkimg(ExampleGofont, t, "gofont.png")
}
