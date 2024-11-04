package door

import (
	"fmt"
	"fyne.io/fyne/v2/theme"
	"estacionamiento/models"
)

type DoorService struct {
	estacionamiento *models.Estacionamiento
}

func NewDoorService(estacionamiento *models.Estacionamiento) *DoorService {
	return &DoorService{
		estacionamiento: estacionamiento,
	}
}

func (s *DoorService) MostrarEstadoPuerta(id int, accion string) {
	estado := "🟢 LIBRE"
	if s.estacionamiento.VehiculosEnPuerta > 0 {
		if s.estacionamiento.EstadoPuerta == models.ENTRANDO {
			estado = "🔵 EN USO (Entrando)"
		} else {
			estado = "🔴 EN USO (Saliendo)"
		}
	}
	fmt.Printf("Puerta: %s | Vehículo %d: %s | Cola Entrada: %d | Cola Salida: %d\n", 
		estado, id, accion, s.estacionamiento.ColaEntrada.Len(), s.estacionamiento.ColaSalida.Len())
}

func (s *DoorService) ActualizarPuerta() {
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