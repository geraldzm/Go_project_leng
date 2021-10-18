package funciones

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Codigo tomado de: https://github.com/gonum/plot/wiki/Example-plots
// Este programa toma un array de enteros y los grafica como grafico de barras
func Funcion2(array []int) error {
	group := plotter.Values{}

	group = append(group, intArr2Float(array)...)

	p := plot.New()

	p.Title.Text = "Grafico de barras"
	p.Y.Label.Text = ""

	w := vg.Points(20)

	bars, err := plotter.NewBarChart(group, w)

	if err != nil {
		panic(err)
	}

	bars.LineStyle.Width = vg.Length(0)
	bars.Color = plotutil.Color(0)
	bars.Offset = -w

	p.Add(bars)

	p.Legend.Add("")
	p.Legend.Top = true

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}

	return nil
}

func intArr2Float(array []int) []float64 {
	floatArr := make([]float64, len(array))

	for i := 0; i < len(array); i++ {
		floatArr[i] = float64(array[i])
	}

	return floatArr
}
