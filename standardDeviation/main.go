package main

import (
	"image/color"
	"math"

	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Create a new plot
	p := plot.New()
	p.Title.Text = "Normal Distribution with 1.7 Standard Deviations from the Mean"
	p.X.Label.Text = "Value"
	p.Y.Label.Text = "Probability Density"

	// Create a normal distribution with mean 0 and standard deviation 1
	normal := distuv.Normal{
		Mu:    0,
		Sigma: 1,
	}

	// Generate points for the normal distribution
	n := 1000
	pts := make(plotter.XYs, n)
	for i := range pts {
		x := -4 + 8*float64(i)/float64(n-1)
		pts[i].X = x
		pts[i].Y = normal.Prob(x)
	}

	// Plot the normal distribution
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	// Highlight the area within 1.7 standard deviations from the mean
	fillPts := make(plotter.XYs, 0)
	for i := range pts {
		if math.Abs(pts[i].X) <= 1.7 {
			fillPts = append(fillPts, pts[i])
		}
	}
	fill, err := plotter.NewPolygon(fillPts)
	if err != nil {
		panic(err)
	}
	fill.Color = color.NRGBA{R: 0, G: 0, B: 255, A: 128} // Set color with transparency
	fill.Color = color.NRGBA{R: 0, G: 0, B: 255, A: 128} // Set fill color with tra
	p.Add(fill)

	// Save the plot to a PNG file
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "normal_distribution.png"); err != nil {
		panic(err)
	}
}
