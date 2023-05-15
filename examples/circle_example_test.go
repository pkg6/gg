// Copyright ©2022 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples_test

import (
	"log"
	"testing"

	"git.sr.ht/~sbinet/gg"
)

func ExampleCircle() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	err := dc.SavePNG("testdata/circle.png")
	if err != nil {
		log.Fatalf("could not save to file: %+v", err)
	}
}

func TestCircle(t *testing.T) {
	chkimg(ExampleCircle, t, "circle.png")
}
