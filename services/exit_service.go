package services

import (
	"fmt"
	"time"
	"fyne.io/fyne/v2/theme"
	"estacionamiento/config"
	"estacionamiento/models"
)

type SalidaService struct {
	estacionamiento *models.Estacionamiento
	doorService    *DoorService
}

func NewSalidaService(estacionamiento *models.Estacionamiento, doorService *DoorService) *SalidaService {
	return &SalidaService{
		estacionamiento: estacionamiento,
		doorService:     doorService,
	}
}

func (s *SalidaService) SalirVehiculo(id int, espacio int) {
	s.estacionamiento.Mutex.Lock()
	defer s.estacionamiento.Mutex.Unlock()

	if s.estacionamiento.ColaSalida.Len() > 0 || (s.estacionamiento.VehiculosEnPuerta > 0 && s.estacionamiento.EstadoPuerta == models.ENTRANDO) {
		vehiculo := models.VehiculoEspera{ID: id, Timestamp: time.Now()}
		elemento := s.estacionamiento.ColaSalida.PushBack(vehiculo)
		fmt.Printf("🚗 Vehículo %d agregado a la cola de salida (posición: %d)\n", id, s.estacionamiento.ColaSalida.Len())
		for elemento != s.estacionamiento.ColaSalida.Front() || (s.estacionamiento.VehiculosEnPuerta > 0 && s.estacionamiento.EstadoPuerta == models.ENTRANDO) {
			s.estacionamiento.EsperaSalida.Wait()
		}
		s.estacionamiento.ColaSalida.Remove(elemento)
	}

	for s.estacionamiento.VehiculosEnPuerta > 0 && s.estacionamiento.EstadoPuerta == models.ENTRANDO {
		s.doorService.MostrarEstadoPuerta(id, "🕐 Esperando para salir")
		s.estacionamiento.EsperaSalida.Wait()
	}

	s.estacionamiento.VehiculosEnPuerta++
	s.estacionamiento.EstadoPuerta = models.SALIENDO
	s.doorService.ActualizarPuerta()
	s.doorService.MostrarEstadoPuerta(id, "🚪 Saliendo")

	time.Sleep(time.Duration(config.TIEMPO_PUERTA_SALIDA) * time.Millisecond)

	s.liberarEspacio(id, espacio)

	s.estacionamiento.VehiculosEnPuerta--
	if s.estacionamiento.VehiculosEnPuerta == 0 {
		s.estacionamiento.EstadoPuerta = models.LIBRE
		s.doorService.ActualizarPuerta()
		if s.estacionamiento.ColaEntrada.Len() > 0 {
			s.estacionamiento.EsperaEntrada.Broadcast()
		} else {
			s.estacionamiento.EsperaSalida.Broadcast()
		}
	}
	s.doorService.MostrarEstadoPuerta(id, "✅ Terminó de salir")

	if s.estacionamiento.ColaEntrada.Len() > 0 {
		s.estacionamiento.EsperaEntrada.Broadcast()
	}
}

func (s *SalidaService) liberarEspacio(id int, espacio int) {
	s.estacionamiento.Espacios[espacio] = false
	s.estacionamiento.Ocupados--
	s.estacionamiento.EspaciosCanvas[espacio].FillColor = theme.BackgroundColor()
	s.estacionamiento.EspaciosCanvas[espacio].Refresh()
	fmt.Printf("🚙 Vehículo %d saliendo del cajón %d\n", id, espacio)
}