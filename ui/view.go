package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
	"estacionamiento/config"
)

func SetupView(ui *ParkingUI) {
	ui.colaEntradaLabel = widget.NewLabel("Cola de Entrada: 0")
	ui.colaSalidaLabel = widget.NewLabel("Cola de Salida: 0")

	ui.contenedorEntrada = container.NewGridWrap(fyne.NewSize(30, 30)) 
	ui.contenedorSalida = container.NewGridWrap(fyne.NewSize(30, 30))  

	ui.service.GetEstacionamiento().Puerta.SetMinSize(fyne.NewSize(50, 30))

	for i := 0; i < config.TOTAL_ESPACIOS; i++ {
		espacio := canvas.NewRectangle(theme.BackgroundColor())
		espacio.SetMinSize(fyne.NewSize(80, 100))
		ui.service.GetEstacionamiento().EspaciosCanvas = append(ui.service.GetEstacionamiento().EspaciosCanvas, espacio)
	}

	var canvasObjects []fyne.CanvasObject
	for _, espacio := range ui.service.GetEstacionamiento().EspaciosCanvas {
		canvasObjects = append(canvasObjects, espacio)
	}
	espaciosGrid := container.NewGridWithColumns(5, canvasObjects...)

	colaEntradaContainer := container.NewVBox(
		ui.colaEntradaLabel,
		ui.contenedorEntrada,
	)
	colaSalidaContainer := container.NewVBox(
		ui.colaSalidaLabel,
		ui.contenedorSalida,
	)

	colasContainer := container.NewHBox(
		colaEntradaContainer,
		widget.NewSeparator(),
		colaSalidaContainer,
	)

	puertaContainer := container.NewVBox(ui.service.GetEstacionamiento().Puerta)
	layout := container.NewVBox(
		colasContainer,
		puertaContainer,
		espaciosGrid,
	)

	ui.window.SetContent(layout)
}
