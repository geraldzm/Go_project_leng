package graph

import (
	"image/color"
	"math"
	"proyecto/Sort"
	"strconv"

	"github.com/fogleman/gg"
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

type Graph struct {
	Algorithm         Sort.ISortAlgorithm
	pipes             []*Rect
	Rect              *Rect
	Color             color.Color
	Title             string
	finishInformation string
}

func findMax(a []int) float64 {
	max := a[0]

	for _, value := range a {
		if value > max {
			max = value
		}
	}

	return float64(max)
}

func (g *Graph) Init() {
	g.Rect.Color = color.White
	g.Algorithm.Init()

	arr := g.Algorithm.GetArray()
	pw := g.Rect.W / float64(len(*arr)) // width of pipe
	mx := findMax(*arr)

	if mx == 0 {
		mx = 1
	}

	for indx, val := range *arr {
		indxF := float64(indx)
		valF := float64(val)
		g.pipes = append(g.pipes, &Rect{X: indxF*pw + g.Rect.X, Y: g.Rect.Y, W: pw, H: valF * g.Rect.H / mx})
		g.pipes[indx].Color = g.Color
	}

}

func drawText(ctx *canvas.Context, txt *string, x float64, y float64) {

	ctx.Push()

	ctx.ScaleAbout(-1, 1, x, y)
	ctx.RotateAbout(math.Pi, x, y)

	ctx.SetColor(colornames.White)
	ctx.DrawStringWrapped(*txt, x, y, 0.5, 0.5, 150, 1, gg.AlignCenter)
	ctx.Stroke()
	ctx.Pop()

}

func (g *Graph) Draw(ctx *canvas.Context) {

	// marco
	g.Rect.Draw(ctx)
	ctx.Stroke()

	// get change
	itm, more := g.Algorithm.Sort()

	// pipes
	for _, r := range g.pipes {
		r.Draw(ctx)
		ctx.Fill()

		r.Color = g.Color
	}

	if more {
		if !itm.Finished {
			g.pipes[itm.IndexFrom].X, g.pipes[itm.IndexTo].X = g.pipes[itm.IndexTo].X, g.pipes[itm.IndexFrom].X
			g.pipes[itm.IndexFrom], g.pipes[itm.IndexTo] = g.pipes[itm.IndexTo], g.pipes[itm.IndexFrom]
			g.pipes[itm.IndexFrom].Color, g.pipes[itm.IndexTo].Color = colornames.Red, colornames.Red
		} else {
			g.finishInformation = "Start: " + itm.TimeStart + " End: " + itm.TimeEnd + " Total time: " + itm.TotalTime + " Comp: " + strconv.Itoa(itm.TotalComp) + " Swaps: " + strconv.Itoa(itm.TotalSwaps) + " Iter: " + strconv.Itoa(itm.TotalIter)
		}
	}

	x := g.Rect.X + g.Rect.W/2
	y := g.Rect.Y + g.Rect.H

	// Title
	drawText(ctx, &g.Title, x, y-10)

	// Info if exists
	drawText(ctx, &g.finishInformation, x, y-50)

}
