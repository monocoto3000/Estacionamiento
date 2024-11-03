package main

import (
	"fyne.io/fyne/v2/app"
	"estacionamiento/services"
	"estacionamiento/simulators"
	"estacionamiento/ui"
)

func main() {
	a := app.New()

	parkingService := services.NewParkingService(a)
	simulator := simulators.NewParkingSimulator(parkingService)
	parkingUI := ui.NewParkingUI(parkingService)

	parkingUI.Setup()

	go simulator.IniciarSimulacion()

	parkingUI.Run()
}
