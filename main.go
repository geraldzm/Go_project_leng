package main

import (
	"github.com/h8gi/canvas"
	"golang.org/x/image/colornames"
	"proyecto/Sort"
	"proyecto/funciones"
	"proyecto/graph"
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
		//time.Sleep(time.Millisecond * 50)
	}
}

func main() {

	// ---------- bubble sort

	rsb := funciones.Funcion1(100, 17)

	graphs[0] = &graph.Graph{
		Rect: &graph.Rect{
			X: 20, Y: 1, W: 400, H: 250,
		},
		Color:     &colornames.Green,
		Algorithm: &Sort.BubbleSort{Array: &rsb },
		Title:     "Bubble sort",
	}

	rsq := funciones.Funcion1(100, 17)

	// ---------- quick sort
	graphs[1] = &graph.Graph{
		Rect: &graph.Rect{
			X: 430, Y: 1, W: 400, H: 250,
		},
		Color:     &colornames.Greenyellow,
		Algorithm: &Sort.QuickSort{Array: &rsq},
		Title:     "Quick sort",
	}

	rsi := funciones.Funcion1(100, 17)

	// ---------- insertion sort
	graphs[2] = &graph.Graph{
		Rect: &graph.Rect{
			X: 20, Y: 280, W: 400, H: 250,
		},
		Color:     &colornames.Firebrick,
		Algorithm: &Sort.InsertionSort{Array: &rsi},
		Title:     "Insertion sort",
	}

	rso := funciones.Funcion1(100, 17)


	// ---------- other
	graphs[3] = &graph.Graph{
		Rect: &graph.Rect{
			X: 430, Y: 280, W: 400, H: 250,
		},
		Color:     &colornames.Coral,
		Algorithm: &Sort.QuickSort{Array: &rso},
		Title:     "Quick sort",
	}

	rsoo := funciones.Funcion1(100, 17)

	// ---------- other
	graphs[4] = &graph.Graph{
		Rect: &graph.Rect{
			X: 225, Y: 540, W: 400, H: 250,
		},
		Color:     &colornames.Blue,
		Algorithm: &Sort.QuickSort{Array: &rsoo},
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
