package graph

import (
	"github.com/h8gi/canvas"
	"image/color"
)

type Rect struct {
	X, Y, W, H float64
	Color color.Color
}

func (r Rect) Draw(ctx *canvas.Context) {

	ctx.SetColor(r.Color)
	ctx.DrawRectangle(r.X, r.Y, r.W, r.H)

}
