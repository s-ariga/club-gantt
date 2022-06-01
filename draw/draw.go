/**
 * Seiichi Ariga <seiichi.ariga@gmail.com>
 */

package draw

import (
	"github.com/s-ariga/club-gantt/csv"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

const (
	WIDTH = 1920
	HEIGHT = 1080
)

// SVGで描画する。canvasを作成して、各funcに渡していく
func DrawEventsGantt(events []csv.Event, filename string) error {

	c := canvas.New(WIDTH, HEIGHT)
	ctx := canvas.NewContext(c)

	ctx.SetFillColor(canvas.White)

	renderers.Write(filename, c, canvas.DPMM(5.0))

	return nil
}
