package services

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"estacionamiento/config"
	"estacionamiento/models"
)

type ParkingService struct {
	estacionamiento *models.Estacionamiento
}

func NewParkingService(app fyne.App) *ParkingService {
	e := &models.Estacionamiento{
		Espacios:        make([]bool, config.TOTAL_ESPACIOS),
		ColaEntrada:     list.New(),
		ColaSalida:      list.New(),
		EstadoPuerta: 	 models.LIBRE,
		Puerta:          canvas.NewRectangle(theme.BackgroundColor()), 
	}
	e.EsperaEntrada = sync.NewCond(&e.Mutex)
	e.EsperaSalida = sync.NewCond(&e.Mutex)
	return &ParkingService{estacionamiento: e}
}

func (s *ParkingService) GetEstacionamiento() *models.Estacionamiento {
	return s.estacionamiento
}

func (s *ParkingService) MostrarEstadoPuerta(id int, accion string) {
	estado := "🟢 LIBRE"
	if s.estacionamiento.VehiculosEnPuerta > 0 {
		if s.estacionamiento.EstadoPuerta == models.ENTRANDO {
			estado = "🔵 EN USO (Entrando)"
		} else {
			estado = "🔴 EN USO (Saliendo)"
		}
	}
	fmt.Printf("Puerta: %s | Vehículo %d: %s | Cola Entrada: %d | Cola Salida: %d\n", estado, id, accion, s.estacionamiento.ColaEntrada.Len(), s.estacionamiento.ColaSalida.Len())
}

func (s *ParkingService) ActualizarPuerta() {
	switch s.estacionamiento.EstadoPuerta {
	case models.ENTRANDO:
		s.estacionamiento.Puerta.FillColor = theme.PrimaryColor()
	case models.SALIENDO:
		s.estacionamiento.Puerta.FillColor = theme.ErrorColor()
	default:
		s.estacionamiento.Puerta.FillColor = theme.SuccessColor()
	}
	s.estacionamiento.Puerta.Refresh()
}

func (s *ParkingService) EntrarVehiculo(id int) int {
	s.estacionamiento.Mutex.Lock()
	defer s.estacionamiento.Mutex.Unlock()

	if s.estacionamiento.Ocupados >= config.TOTAL_ESPACIOS || s.estacionamiento.ColaEntrada.Len() > 0 {
		vehiculo := models.VehiculoEspera{ID: id, Timestamp: time.Now()}
		elemento := s.estacionamiento.ColaEntrada.PushBack(vehiculo)
		fmt.Printf("🚗 Vehículo %d agregado a la cola de entrada (posición: %d)\n", id, s.estacionamiento.ColaEntrada.Len())
		for s.estacionamiento.Ocupados >= config.TOTAL_ESPACIOS || elemento != s.estacionamiento.ColaEntrada.Front() {
			s.estacionamiento.EsperaEntrada.Wait()
		}
		s.estacionamiento.ColaEntrada.Remove(elemento)
	}

	for s.estacionamiento.VehiculosEnPuerta > 0 && s.estacionamiento.EstadoPuerta == models.SALIENDO {
		s.MostrarEstadoPuerta(id, "🕐 Esperando para entrar")
		s.estacionamiento.EsperaEntrada.Wait()
	}

	s.estacionamiento.VehiculosEnPuerta++
	s.estacionamiento.EstadoPuerta = models.ENTRANDO
	s.ActualizarPuerta()
	s.MostrarEstadoPuerta(id, "🚪 Entrando")

	time.Sleep(time.Duration((config.TIEMPO_PUERTA_ENTRADA) * time.Millisecond))

	espacioAsignado := -1
	for i := 0; i < config.TOTAL_ESPACIOS; i++ {
		if !s.estacionamiento.Espacios[i] {
			s.estacionamiento.Espacios[i] = true
			espacioAsignado = i
			s.estacionamiento.Ocupados++
			s.estacionamiento.EspaciosCanvas[i].FillColor = theme.PrimaryColor()
			s.estacionamiento.EspaciosCanvas[i].Refresh()
			break
		}
	}

	fmt.Printf("🚙 Vehículo %d entrando al cajón %d\n", id, espacioAsignado)

	s.estacionamiento.VehiculosEnPuerta--
	if s.estacionamiento.VehiculosEnPuerta == 0 {
		s.estacionamiento.EstadoPuerta = models.LIBRE
		s.ActualizarPuerta()
		s.estacionamiento.EsperaSalida.Broadcast()
	}
	s.MostrarEstadoPuerta(id, "✅ Terminó de entrar")

	return espacioAsignado
}

func (s *ParkingService) SalirVehiculo(id int, espacio int) {
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
		s.MostrarEstadoPuerta(id, "🕐 Esperando para salir")
		s.estacionamiento.EsperaSalida.Wait()
	}

	s.estacionamiento.VehiculosEnPuerta++
	s.estacionamiento.EstadoPuerta = models.SALIENDO
	s.ActualizarPuerta()
	s.MostrarEstadoPuerta(id, "🚪 Saliendo")

	time.Sleep(time.Duration(config.TIEMPO_PUERTA_SALIDA) * time.Millisecond)

	s.estacionamiento.Espacios[espacio] = false
	s.estacionamiento.Ocupados--
	s.estacionamiento.EspaciosCanvas[espacio].FillColor = theme.BackgroundColor()
	s.estacionamiento.EspaciosCanvas[espacio].Refresh()
	fmt.Printf("🚙 Vehículo %d saliendo del cajón %d\n", id, espacio)

	s.estacionamiento.VehiculosEnPuerta--
	if s.estacionamiento.VehiculosEnPuerta == 0 {
		s.estacionamiento.EstadoPuerta = models.LIBRE
		s.ActualizarPuerta()
		if s.estacionamiento.ColaEntrada.Len() > 0 {
			s.estacionamiento.EsperaEntrada.Broadcast()
		} else {
			s.estacionamiento.EsperaSalida.Broadcast()
		}
	}
	s.MostrarEstadoPuerta(id, "✅ Terminó de salir")

	if s.estacionamiento.ColaEntrada.Len() > 0 {
		s.estacionamiento.EsperaEntrada.Broadcast()
	}
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