package models

import (
	"container/list"
	"sync"
	"time"

	"fyne.io/fyne/v2/canvas"
)

type Direccion int

const (
	ENTRANDO Direccion = iota
	SALIENDO
	LIBRE
)

type VehiculoEspera struct {
	ID        int
	Timestamp time.Time
}

type Estacionamiento struct {
	Espacios          []bool
	Ocupados          int
	EstadoPuerta   Direccion
	VehiculosEnPuerta int
	Mutex             sync.Mutex
	EsperaEntrada     *sync.Cond
	EsperaSalida      *sync.Cond
	ColaEntrada       *list.List
	ColaSalida        *list.List
	Puerta            *canvas.Rectangle
	EspaciosCanvas    []*canvas.Rectangle
}