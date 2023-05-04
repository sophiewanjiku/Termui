package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	startYear       = 2021
	endYear         = 2030
	startPopulation = 408286745334
	startCars       = 2376908
)

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	defer termui.Close()

	//display current population
	populationWidget := widgets.NewParagraph()
	populationWidget.Title = "population"
	populationWidget.Text = fmt.Sprintf("%d", startPopulation)

	//display current number of cars
	carsWidget := widgets.NewParagraph()
	carsWidget.Title = "Number of cars"
	carsWidget.Text = fmt.Sprintf("%d", startCars)

	//grid layout
	grid := termui.NewGrid()
	grid.Set(
		termui.NewRow(1.0/2, populationWidget),
		termui.NewRow(1.0/2, carsWidget),
	)
	termui.Render(grid)

	//growth rate and increase in the number of cars
	rand.Seed(time.Now().UnixNano())
	population := startPopulation
	cars := startCars

	for year := startYear; year <= endYear; year++ {
		growthRate := 0.01 + rand.Float64()/100

		newPopulation := int(float64(population) * growthRate)

		//calculate the number of new cars added to the population
		carRate := 0.05 + rand.Float64()*0.05
		newCars := int(float64(newPopulation) * carRate)

		//update the population and number of cars
		population += newPopulation
		cars += newCars

		//update the widgets with the new values
		populationWidget.Text = fmt.Sprintf("%d", population)
		carsWidget.Text = fmt.Sprintf("%d", cars)

		//render the widgets
		termui.Render(grid)

		//
		time.Sleep(time.Second / 4)
	}

}
