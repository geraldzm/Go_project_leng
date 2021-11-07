package main

import (
	"proyecto/Sort"
	"proyecto/graph"
	"time"

	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
)

var graphs [5]*graph.Graph

func setUp(ctx *canvas.Context) {
	ctx.SetColor(colornames.Black)
	ctx.Clear()
	ctx.SetColor(colornames.White)
	ctx.SetLineWidth(2)
}

func draw(ctx *canvas.Context) {
	ctx.SetColor(colornames.Black)
	ctx.Clear()

	for _, v := range graphs {
		v.Draw(ctx)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// ---------- bubble sort
	graphs[0] = &graph.Graph{
		Rect: &graph.Rect{
			X: 20, Y: 1, W: 400, H: 250,
		},
		Color:     &colornames.Green,
		Algorithm: &Sort.BubbleSort{Array: &[]int{32, 4, 7, -5, 8, 3, 5, 23, 45, 2, 4, 38, 6, 40, 17}},
		Title:     "Bubble sort",
	}

	// ---------- quick sort
	graphs[1] = &graph.Graph{
		Rect: &graph.Rect{
			X: 430, Y: 1, W: 400, H: 250,
		},
		Color:     &colornames.Greenyellow,
		Algorithm: &Sort.BubbleSort{Array: &[]int{8, 3, 3, 5, 23, 13, 4, 14, 5, 8, 2, 4, 23, 6, 35, 17}},
		Title:     "Quick sort",
	}

	// ----------
	graphs[2] = &graph.Graph{
		Rect: &graph.Rect{
			X: 20, Y: 280, W: 400, H: 250,
		},
		Color:     &colornames.Firebrick,
		Algorithm: &Sort.BubbleSort{Array: &[]int{8, 3, 3, 5, 23, 9, 5, 9, 5, 18, 13, 4, 23, 6, 35, 17}},
		Title:     "Quick sort",
	}
	// ---------- other
	graphs[3] = &graph.Graph{
		Rect: &graph.Rect{
			X: 430, Y: 280, W: 400, H: 250,
		},
		Color:     &colornames.Coral,
		Algorithm: &Sort.BubbleSort{Array: &[]int{8, 3, 3, 5, 23, 13, 4, 7, 5, 8, 2, 4, 23, 6, 35, 17}},
		Title:     "Quick sort",
	}

	// ---------- other
	graphs[4] = &graph.Graph{
		Rect: &graph.Rect{
			X: 225, Y: 540, W: 400, H: 250,
		},
		Color:     &colornames.Blue,
		Algorithm: &Sort.BubbleSort{Array: &[]int{8, 3, 3, 35, 17, 23, 9, 5, 32, 4, 8, 10, 5, 8, 31, 17}},
		Title:     "Quick sort",
	}

	for _, v := range graphs {
		v.Init()
	}

	c := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     850,
		Height:    810,
		FrameRate: 15,
		Title:     "Graphs",
	})

	c.Setup(setUp)
	c.Draw(draw)
}
