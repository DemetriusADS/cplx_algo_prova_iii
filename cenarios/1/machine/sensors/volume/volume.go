package volume

import (
	"math/rand"
	"time"

	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/2/machine/ports/sensor"
)

type Volume struct {
	sensor.SensorDTO
}

func NewVolumeSensor() *Volume {
	return &Volume{}
}

func (v *Volume) Read() *sensor.SensorDTO {
	v.Time = time.Now().Format("2006-01-02 15:04:05")

	if rand.Int()*10%5 == 0 {
		v.Value = rand.Float64() * 100
	}
	return &v.SensorDTO
}
