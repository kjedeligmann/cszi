package main

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateLineItems(list []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(list); i++ {
		items = append(items, opts.LineData{Value: list[i]})
	}
	return items
}

func createHTML(list []float64) {
	xValues := make([]int, 0)

	for idx := range list {
		xValues = append(xValues, idx)
	}

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Логарифмічна функція правдоподібності для кожного ключа",
		}),
	)

	line.SetXAxis(xValues).AddSeries("", generateLineItems(list)).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Smooth: true,
			}),
		)

	f, _ := os.Create("line.html")
	line.Render(f)
}
