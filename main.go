package main

import (
	"proyecto/Sort"
	"proyecto/funciones"
	"proyecto/graph"

	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

	// La semilla deberá ser un número primo entre 11 y 101.
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a seed: ")
	str1, _ := reader.ReadString('\n')

	// remove newline
	str1 = strings.Replace(str1, "\n", "", -1)

	// convert string variable to int variable
	seed, e := strconv.Atoi(str1)
	if e != nil {
		fmt.Println("conversion error:", str1)
		return
	}

	arreglo_base := funciones.Funcion1(100, seed)

	// ---------- bubble sort
	rsb := make([]int, 100)
	copy(rsb, arreglo_base)

	graphs[0] = &graph.Graph{
		Rect: &graph.Rect{
			X: 20, Y: 20, W: 400, H: 250,
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
			X: 430, Y: 120, W: 400, H: 250,
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
			X: 20, Y: 540, W: 400, H: 250,
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
			X: 430, Y: 385, W: 400, H: 250,
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
