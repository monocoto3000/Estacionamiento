package vehicle

import (
	"fmt"
	"time"
	// "fyne.io/fyne/v2/theme"
	"estacionamiento/config"
	"estacionamiento/models"
	"estacionamiento/services/door"
)

type EntradaService struct {
	estacionamiento *models.Estacionamiento
	doorService     *door.DoorService
}

func NewEntradaService(estacionamiento *models.Estacionamiento, doorService *door.DoorService) *EntradaService {
	return &EntradaService{
		estacionamiento: estacionamiento,
		doorService:     doorService,
	}
}

func (s *EntradaService) EntrarVehiculo(id int) int {
	s.estacionamiento.Mutex.Lock()
	defer s.estacionamiento.Mutex.Unlock()
	fmt.Print("LUGARES OCUPADOS: ", s.estacionamiento.Ocupados)

	if s.estacionamiento.Ocupados >= config.TOTAL_ESPACIOS || s.estacionamiento.ColaEntrada.Len() > 0 {
		vehiculo := models.VehiculoEspera{ID: id}
		elemento := s.estacionamiento.ColaEntrada.PushBack(vehiculo)
		fmt.Printf("--- Vehículo %d agregado a la cola de entrada (pos: %d)\n", id, s.estacionamiento.ColaEntrada.Len())
		for s.estacionamiento.Ocupados >= config.TOTAL_ESPACIOS || elemento != s.estacionamiento.ColaEntrada.Front() {
			s.estacionamiento.EsperaEntrada.Wait()
		}
		s.estacionamiento.ColaEntrada.Remove(elemento)
	}

	for s.estacionamiento.VehiculosEnPuerta > 0 && s.estacionamiento.EstadoPuerta == models.SALIENDO {
		s.doorService.MostrarEstadoPuerta(id, "Esperando para entrar")
		s.estacionamiento.EsperaEntrada.Wait()
	}

	s.estacionamiento.VehiculosEnPuerta++
	s.estacionamiento.EstadoPuerta = models.ENTRANDO
	s.doorService.ActualizarPuerta()
	s.doorService.MostrarEstadoPuerta(id, "// Entrando //")

	time.Sleep(time.Duration((config.TIEMPO_PUERTA_ENTRADA) * time.Millisecond))

	espacioAsignado := s.asignarEspacio(id)

	s.estacionamiento.VehiculosEnPuerta--
	if s.estacionamiento.VehiculosEnPuerta == 0 {
		s.estacionamiento.EstadoPuerta = models.LIBRE
		s.doorService.ActualizarPuerta()
		s.estacionamiento.EsperaSalida.Broadcast()
	}
	s.doorService.MostrarEstadoPuerta(id, "--- Terminó de entrar ---")

	return espacioAsignado
}

func (s *EntradaService) asignarEspacio(id int) int {
	espacioAsignado := -1
	for i := 0; i < config.TOTAL_ESPACIOS; i++ {
		if !s.estacionamiento.Espacios[i] {
			s.estacionamiento.Espacios[i] = true
			espacioAsignado = i
			s.estacionamiento.Ocupados++
			// s.estacionamiento.EspaciosCanvas[i].FillColor = theme.PrimaryColor()
			// s.estacionamiento.EspaciosCanvas[i].Refresh()
			break
		}
	}
	fmt.Printf("🚙 Vehículo %d entrando al cajón %d\n", id, espacioAsignado)
	return espacioAsignado
}