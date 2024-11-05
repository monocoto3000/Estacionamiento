package vehicle

import (
    "fmt"
    "sync"
    "estacionamiento/models"
)

type FreeSpaceService struct {
    estacionamiento *models.Estacionamiento
    mu              sync.Mutex
}

func NewFreeSpaceService(estacionamiento *models.Estacionamiento) *FreeSpaceService {
    return &FreeSpaceService{
        estacionamiento: estacionamiento,
    }
}

func (s *FreeSpaceService) LiberarEspacio(id int, espacio int) {
    s.mu.Lock()
    defer s.mu.Unlock()

    s.estacionamiento.Mutex.Lock()
    defer s.estacionamiento.Mutex.Unlock()

    s.estacionamiento.Espacios[espacio] = false
    s.estacionamiento.Ocupados--
    fmt.Printf("🚙 !!!!!!!!!!!! Vehículo %d liberó el espacio %d\n", id, espacio)

    if s.estacionamiento.ColaEntrada.Len() > 0 {
        s.estacionamiento.EsperaEntrada.Broadcast()
    }
}
