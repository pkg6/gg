// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build example

package main

import (
	"flag"
	"log"

	"git.sr.ht/~sbinet/gg"
	"github.com/go-fonts/liberation/liberationsansregular"
)

func main() {
	var (
		fname = flag.String("o", "text.png", "output file")
	)
	flag.Parse()

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

	err = dc.SavePNG(*fname)
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}
