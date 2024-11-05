package vehicle

import (
	"fmt"
	"time"
	// "fyne.io/fyne/v2/theme"
	"estacionamiento/config"
	"estacionamiento/models"
	"estacionamiento/services/door"
)

type SalidaService struct {
	estacionamiento *models.Estacionamiento
	doorService    *door.DoorService
}

func NewSalidaService(estacionamiento *models.Estacionamiento, doorService *door.DoorService) *SalidaService {
	return &SalidaService{
		estacionamiento: estacionamiento,
		doorService:     doorService,
	}
}

func (s *SalidaService) SalirVehiculo(id int) {
    s.estacionamiento.Mutex.Lock()
    defer s.estacionamiento.Mutex.Unlock()
    fmt.Print("LUGARES OCUPADOS: ", s.estacionamiento.Ocupados)

    vehiculo := models.VehiculoEspera{ID: id}
    elemento := s.estacionamiento.ColaSalida.PushBack(vehiculo)
    fmt.Printf("Vehículo %d agregado a la cola de salida (pos: %d)\n", id, s.estacionamiento.ColaSalida.Len())

    for elemento != s.estacionamiento.ColaSalida.Front() || 
          (s.estacionamiento.VehiculosEnPuerta > 0 && s.estacionamiento.EstadoPuerta == models.ENTRANDO) {
        s.estacionamiento.EsperaSalida.Wait()
    }

    for s.estacionamiento.ColaEntrada.Len() > 0 && s.estacionamiento.Ocupados < config.TOTAL_ESPACIOS {
        s.estacionamiento.EsperaEntrada.Broadcast()
        s.estacionamiento.EsperaSalida.Wait()
    }

    s.estacionamiento.ColaSalida.Remove(elemento)

    for s.estacionamiento.VehiculosEnPuerta > 0 && s.estacionamiento.EstadoPuerta == models.ENTRANDO {
        s.doorService.MostrarEstadoPuerta(id, "--- Esperando para salir ---")
        s.estacionamiento.EsperaSalida.Wait()
    }

    s.estacionamiento.VehiculosEnPuerta++
    s.estacionamiento.EstadoPuerta = models.SALIENDO
    s.doorService.ActualizarPuerta()
    s.doorService.MostrarEstadoPuerta(id, "--- Saliendo ---")

    time.Sleep(time.Duration(config.TIEMPO_PUERTA_SALIDA) * time.Millisecond)

    s.estacionamiento.VehiculosEnPuerta--
    if s.estacionamiento.VehiculosEnPuerta == 0 {
        s.estacionamiento.EstadoPuerta = models.LIBRE
        s.doorService.ActualizarPuerta()
        if s.estacionamiento.ColaEntrada.Len() > 0 || s.estacionamiento.ColaSalida.Len() > 0 {
            s.estacionamiento.EsperaEntrada.Broadcast()
            s.estacionamiento.EsperaSalida.Broadcast()
        }
    }

    s.doorService.MostrarEstadoPuerta(id, "--- Terminó de salir ---")

    if s.estacionamiento.ColaEntrada.Len() > 0 {
        s.estacionamiento.EsperaEntrada.Broadcast()
    }
}
