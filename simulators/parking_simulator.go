package simulators

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"estacionamiento/config"
	"estacionamiento/services"
)

type ParkingSimulator struct {
	service *services.ParkingService
	wg      sync.WaitGroup
}

func NewParkingSimulator(service *services.ParkingService) *ParkingSimulator {
	return &ParkingSimulator{
		service: service,
	}
}

func (s *ParkingSimulator) simularVehiculo(id int, r *rand.Rand) {
	defer s.wg.Done()

	espacio := s.service.EntrarVehiculo(id)
	if espacio != -1 {
		tiempoEstacionado := config.MIN_TIEMPO_ESTAC + r.Float64()*(config.MAX_TIEMPO_ESTAC-config.MIN_TIEMPO_ESTAC)
		fmt.Printf("🕐 Vehículo %d estacionado por %.2f segundos\n", id, tiempoEstacionado)
		time.Sleep(time.Duration(tiempoEstacionado * float64(time.Second)))
		s.service.SalirVehiculo(id, espacio)
	}
}

func (s *ParkingSimulator) IniciarSimulacion() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("🅿️  Iniciando simulación del estacionamiento")

	lambda := 0.9
	for i := 0; i < config.TOTAL_VEHICULOS; i++ {
		s.wg.Add(1)
		tiempoEspera := r.ExpFloat64() / lambda
		time.Sleep(time.Duration(tiempoEspera * float64(time.Second)))
		go s.simularVehiculo(i, rand.New(rand.NewSource(time.Now().UnixNano())))
	}

	s.wg.Wait()
	fmt.Println("✨ Simulación completada")
}