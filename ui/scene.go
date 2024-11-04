package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"estacionamiento/services"
	"image/color"
	"strconv"
	"time"
)

type ParkingUI struct {
	window            fyne.Window
	service           *services.ParkingService
	colaEntradaLabel  *widget.Label
	colaSalidaLabel   *widget.Label
	carrosEntrada     []*canvas.Rectangle
	carrosSalida      []*canvas.Rectangle
	contenedorEntrada *fyne.Container
	contenedorSalida  *fyne.Container
}

func NewParkingUI(service *services.ParkingService) *ParkingUI {
	return &ParkingUI{
		service:       service,
		carrosEntrada: make([]*canvas.Rectangle, 0),
		carrosSalida:  make([]*canvas.Rectangle, 0),
	}
}

func (ui *ParkingUI) Setup() {
	a := app.New()
	ui.window = a.NewWindow("Simulación Estacionamiento :D")

	SetupView(ui)

	ui.window.Resize(fyne.NewSize(800, 500))
	go ui.actualizarUI()
}

func (ui *ParkingUI) actualizarCarrosEnCola(cantidad int, esEntrada bool) {
	var carros *[]*canvas.Rectangle
	var contenedor **fyne.Container

	colorEntrada := color.NRGBA{R: 0, G: 200, B: 0, A: 255}  
	colorSalida := color.NRGBA{R: 200, G: 0, B: 0, A: 255} 
	var colorCarro color.Color

	if esEntrada {
		carros = &ui.carrosEntrada
		contenedor = &ui.contenedorEntrada
		colorCarro = colorEntrada
	} else {
		carros = &ui.carrosSalida
		contenedor = &ui.contenedorSalida
		colorCarro = colorSalida
	}

	cantidadActual := len(*carros)
	if cantidad > cantidadActual {
		for i := cantidadActual; i < cantidad; i++ {
			carro := canvas.NewRectangle(colorCarro)
			carro.Resize(fyne.NewSize(30, 30))
			carro.SetMinSize(fyne.NewSize(30, 30))
			*carros = append(*carros, carro)
			(*contenedor).Add(carro)
		}
	} else if cantidad < cantidadActual {
		for i := cantidadActual - 1; i >= cantidad; i-- {
			(*contenedor).Remove((*carros)[i])
		}
		*carros = (*carros)[:cantidad]
	}
	
	(*contenedor).Refresh()
}

func (ui *ParkingUI) actualizarUI() {
	for {
		colaEntrada := ui.service.GetColaEntrada()
		colaSalida := ui.service.GetColaSalida()
		
		ui.colaEntradaLabel.SetText("Cola de Entrada: " + strconv.Itoa(colaEntrada))
		ui.colaSalidaLabel.SetText("Cola de Salida: " + strconv.Itoa(colaSalida))

		ui.actualizarCarrosEnCola(colaEntrada, true)
		ui.actualizarCarrosEnCola(colaSalida, false)

		for i := 0; i < len(ui.service.GetEstacionamiento().Espacios); i++ {
			if ui.service.EspacioOcupado(i) {
				ui.service.GetEstacionamiento().EspaciosCanvas[i].FillColor = theme.PrimaryColor()
			} else {
				ui.service.GetEstacionamiento().EspaciosCanvas[i].FillColor = theme.BackgroundColor()
			}
			ui.service.GetEstacionamiento().EspaciosCanvas[i].Refresh()
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func (ui *ParkingUI) Run() {
	ui.window.ShowAndRun()
}
