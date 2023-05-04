// Copyright ©2023 The gg Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"git.sr.ht/~sbinet/cmpimg"
)

func TestExamples(t *testing.T) {
	for _, tc := range []struct {
		name  string
		oname string
	}{
		{
			"beziers.go",
			"testdata/beziers.png",
		},
		{
			"circle.go",
			"testdata/circle.png",
		},
		{
			"clip.go",
			"testdata/clip.png",
		},
		{
			"concat.go",
			"testdata/concat.png",
		},
		{
			"crisp.go",
			"testdata/crisp.png",
		},
		{
			"cubic.go",
			"testdata/cubic.png",
		},
		{
			"ellipse.go",
			"testdata/ellipse.png",
		},
		{
			"gofont.go",
			"testdata/gofont.png",
		},
		{
			"gradient-conic.go",
			"testdata/gradient-conic.png",
		},
		{
			"gradient-linear.go",
			"testdata/gradient-linear.png",
		},
		{
			"gradient-radial.go",
			"testdata/gradient-radial.png",
		},
		{
			"gradient-text.go",
			"testdata/gradient-text.png",
		},
		{
			"invert-mask.go",
			"testdata/invert-mask.png",
		},
		{
			"lines.go",
			"testdata/lines.png",
		},
		{
			"linewidth.go",
			"testdata/linewidth.png",
		},
		{
			"lorem.go",
			"testdata/lorem.png",
		},
		{
			"mask.go",
			"testdata/mask.png",
		},
		{
			"meme.go",
			"testdata/meme.png",
		},
		{
			"openfill.go",
			"testdata/openfill.png",
		},
		{
			"pattern-fill.go",
			"testdata/pattern-fill.png",
		},
		{
			"quadratic.go",
			"testdata/quadratic.png",
		},
		{
			"rotated-image.go",
			"testdata/rotated-image.png",
		},
		{
			"rotated-text.go",
			"testdata/rotated-text.png",
		},
		{
			"scatter.go",
			"testdata/scatter.png",
		},
		{
			"sine.go",
			"testdata/sine.png",
		},
		{
			"spiral.go",
			"testdata/spiral.png",
		},
		{
			"star.go",
			"testdata/star.png",
		},
		{
			"stars.go",
			"testdata/stars.png",
		},
		{
			"text.go",
			"testdata/text.png",
		},
		{
			"tiling.go",
			"testdata/tiling.png",
		},
		{
			"unicode.go",
			"testdata/unicode.png",
		},
		{
			"wrap.go",
			"testdata/wrap.png",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fun := func() {
				cmd := exec.Command("go", "run", tc.name, "-o", tc.oname)
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("could not run %q:\n%s\nerr: %+v", tc.name, out, err)
				}
			}
			cmpimg.CheckPlot(fun, t, filepath.Base(tc.oname))

			_ = os.Remove(tc.oname)
		})
	}
}
