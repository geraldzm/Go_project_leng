package main

import (
	"proyecto/Sort"
	"proyecto/funciones"
	"proyecto/graph"

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
		//time.Sleep(time.Millisecond * 50)
	}
}

func main() {

	// ---------- bubble sort
	arreglo_base := funciones.Funcion1(100, 17)

	rsb := make([]int, 100)
	copy(rsb, arreglo_base)

	graphs[0] = &graph.Graph{
		Rect: &graph.Rect{
			X: 20, Y: 1, W: 400, H: 250,
		},
		Color:     &colornames.Green,
		Algorithm: &Sort.BubbleSort{Array: &rsb},
		Title:     "Bubble sort",
	}

	rsq := make([]int, 100)
	copy(rsq, arreglo_base)

	// ---------- quick sort
	graphs[1] = &graph.Graph{
		Rect: &graph.Rect{
			X: 430, Y: 1, W: 400, H: 250,
		},
		Color:     &colornames.Greenyellow,
		Algorithm: &Sort.QuickSort{Array: &rsq},
		Title:     "Quick Sort",
	}

	rsi := make([]int, 100)
	copy(rsi, arreglo_base)

	// ---------- insertion sort
	graphs[2] = &graph.Graph{
		Rect: &graph.Rect{
			X: 20, Y: 280, W: 400, H: 250,
		},
		Color:     &colornames.Firebrick,
		Algorithm: &Sort.InsertionSort{Array: &rsi},
		Title:     "Insertion Sort",
	}

	rso := make([]int, 100)
	copy(rso, arreglo_base)

	// ---------- selection sort
	graphs[3] = &graph.Graph{
		Rect: &graph.Rect{
			X: 430, Y: 280, W: 400, H: 250,
		},
		Color:     &colornames.Coral,
		Algorithm: &Sort.SelectionSort{Array: &rso},
		Title:     "Selection Sort",
	}

	rsoo := make([]int, 100)
	copy(rsoo, arreglo_base)

	// ---------- heap sort
	graphs[4] = &graph.Graph{
		Rect: &graph.Rect{
			X: 225, Y: 540, W: 400, H: 250,
		},
		Color:     &colornames.Blue,
		Algorithm: &Sort.HeapSort{Array: &rsoo},
		Title:     "Heap Sort",
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
