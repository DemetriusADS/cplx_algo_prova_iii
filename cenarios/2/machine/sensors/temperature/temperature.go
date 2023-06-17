package temperature

import (
	"math/rand"
	"time"

	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/2/machine/ports/sensor"
)

type Temparature struct {
	sensor.SensorDTO
}

func NewTemperatureSensor() *Temparature {
	return &Temparature{}
}

func (t *Temparature) Read() *sensor.SensorDTO {
	t.Time = time.Now().Format("2006-01-02 15:04:05")
	t.Value = rand.Float64() * 100
	return &t.SensorDTO
}
