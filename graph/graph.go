package graph

import (
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"proyecto/Sort"
)

type Graph struct {
	Algorithm Sort.ISortAlgorithm
	pipes []*Rect
	Rect  *Rect
	Color color.Color
	Title string
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

func (g *Graph) Init()  {
	g.Rect.Color = color.White
	g.Algorithm.Init()

	arr := g.Algorithm.GetArray()
	pw := g.Rect.W / float64 (len(*arr)) // width of pipe
	mx := findMax(*arr)

	if mx == 0 {
		mx = 1
	}

	for indx, val := range *arr {
		indxF := float64(indx)
		valF := float64(val)
		g.pipes = append(g.pipes, &Rect { X: indxF*pw + g.Rect.X, Y: g.Rect.Y, W: pw, H: valF*g.Rect.H / mx } )
		g.pipes[indx].Color = g.Color
	}

}

func (g Graph) Draw(ctx *canvas.Context) {

	// marco
	g.Rect.Draw(ctx)
	ctx.Stroke()

	//test
	itm, more := g.Algorithm.Sort()

	// pipes
	for _, r := range g.pipes {
		r.Draw(ctx)
		ctx.Fill()

		r.Color = g.Color
	}

	if more {
		g.pipes[itm.IndexFrom].X, g.pipes[itm.IndexTo].X = g.pipes[itm.IndexTo].X, g.pipes[itm.IndexFrom].X
		g.pipes[itm.IndexFrom], g.pipes[itm.IndexTo] = g.pipes[itm.IndexTo], g.pipes[itm.IndexFrom]
		g.pipes[itm.IndexFrom].Color, g.pipes[itm.IndexTo].Color = colornames.Red, colornames.Red
	}


	ctx.Push()
	x := g.Rect.X + g.Rect.W/2
	y := g.Rect.Y + g.Rect.H - 10

	ctx.ScaleAbout(-1, 1, x, y)
	ctx.RotateAbout(math.Pi, x, y)

	ctx.SetColor(colornames.White)
	ctx.DrawStringAnchored(g.Title, x, y, 0.5, 0.5)
	ctx.Stroke()
	ctx.Pop()
	
}