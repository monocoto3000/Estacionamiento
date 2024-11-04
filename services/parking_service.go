package services

import (
	"fyne.io/fyne/v2"
	"container/list"
	"sync"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"estacionamiento/config"
	"estacionamiento/models"
	"estacionamiento/services/door"
    "estacionamiento/services/vehicle"
)

type ParkingService struct {
	estacionamiento *models.Estacionamiento
	entradaService  *vehicle.EntradaService
	salidaService   *vehicle.SalidaService
	doorService     *door.DoorService
}

func NewParkingService(app fyne.App) *ParkingService {
	e := &models.Estacionamiento{
		Espacios:         make([]bool, config.TOTAL_ESPACIOS),
		ColaEntrada:      list.New(),
		ColaSalida:       list.New(),
		EstadoPuerta:     models.LIBRE,
		Puerta:           canvas.NewRectangle(theme.BackgroundColor()),
	}
	e.EsperaEntrada = sync.NewCond(&e.Mutex)
	e.EsperaSalida = sync.NewCond(&e.Mutex)

	doorService := door.NewDoorService(e)
	
	service := &ParkingService{
		estacionamiento: e,
		doorService:     doorService,
	}
	
	service.entradaService = vehicle.NewEntradaService(e, doorService)
	service.salidaService = vehicle.NewSalidaService(e, doorService)
	
	return service
}

func (s *ParkingService) EntrarVehiculo(id int) int {
	return s.entradaService.EntrarVehiculo(id)
}

func (s *ParkingService) SalirVehiculo(id int, espacio int) {
	s.salidaService.SalirVehiculo(id, espacio)
}

// Interfaz D:

func (s *ParkingService) GetEstacionamiento() *models.Estacionamiento {
	return s.estacionamiento
}

func (s *ParkingService) GetColaEntrada() int {
	return s.estacionamiento.ColaEntrada.Len()
}

func (s *ParkingService) GetColaSalida() int {
	return s.estacionamiento.ColaSalida.Len()
}

func (s *ParkingService) EspacioOcupado(i int) bool {
	return s.estacionamiento.Espacios[i]
}