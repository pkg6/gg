// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//go:build example

package main

import (
	"flag"
	"log"

	"git.sr.ht/~sbinet/gg"
)

func main() {
	var (
		fname = flag.String("o", "mask.png", "output file")
	)
	flag.Parse()

	im, err := gg.LoadImage("testdata/baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(512, 512)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	dc.Clip()
	dc.DrawImage(im, 0, 0)

	err = dc.SavePNG(*fname)
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}
