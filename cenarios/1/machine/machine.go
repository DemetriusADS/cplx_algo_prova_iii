package machine

import (
	"time"

	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/1/machine/ports/sensor"
)

const qtyMetrics = 2000

type Metric struct {
	Temperature sensor.SensorDTO
	Volume      sensor.SensorDTO
	Unstable    bool
}

type Machine struct {
	Metrics []*Metric
	Name    string
	isOn    bool

	VolumeSensor      sensor.Sensor
	TemperatureSensor sensor.Sensor
}

func NewMachine(name string, volumeSensor, temperatureSensor sensor.Sensor) *Machine {
	machine := Machine{
		Name:              name,
		isOn:              true,
		VolumeSensor:      volumeSensor,
		TemperatureSensor: temperatureSensor,
	}

	go func() {
		time.Sleep(5 * time.Second)
		machine.isOn = false
	}()

	return &machine
}

// Complexidade O(n)
// O codigo em questão possui uma complexidade O(n) pois o for é executado n vezes, sendo n a quantidade de métricas.
func (m *Machine) GenData() {
	for i := 0; i < qtyMetrics; i++ {
		// time.Sleep(1 * time.Second)
		temp := m.TemperatureSensor.Read()
		vol := m.VolumeSensor.Read()

		m.Metrics = append(m.Metrics, &Metric{
			Temperature: *temp,
			Volume:      *vol,
			Unstable:    true,
		})
	}
}

func (m *Machine) Read() []*Metric {
	return m.Metrics
}

func (m *Machine) IsOn() bool {
	return m.isOn
}
