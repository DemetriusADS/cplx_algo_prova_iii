package machine

import (
	"fmt"
	"time"

	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/1/machine/ports/sensor"
)

const qtyMetrics = 40

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
		time.Sleep(50 * time.Second)
		machine.isOn = false
	}()

	return &machine
}

func (m *Machine) GenData() {
	for i := 0; i < qtyMetrics; i++ {
		time.Sleep(1 * time.Second)
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

func (m *Machine) FixTemperature() {
	fmt.Printf("CALIBRANDO A MAQUINA %s PARA A TEMPERATURA IDEAL\n", m.Name)
	for _, metric := range m.Read() {
		now := time.Now().Format("2006-01-02 15:04:05")
		if !metric.Unstable {
			continue
		}
		newTemp := metric.Volume.Value * 2.5
		if newTemp < 100 {
			newTemp = 100
		}
		fmt.Printf("TEMPERATURA ATUAL: %f\n", metric.Temperature.Value)
		metric.Temperature.Value = newTemp
		metric.Temperature.Time = now
		metric.Volume.Time = now
		fmt.Printf("TEMPERATURA AJUSTADA: %f\n", metric.Temperature.Value)
		metric.Unstable = false
	}
}

func (m *Machine) IsOn() bool {
	return m.isOn
}

// Aqui escolhi utilizar o algoritmo de ordenação Bubble Sort, pois ele é simples e fácil de implementar.
// O Bubble Sort é um algoritmo de ordenação simples que percorre o array várias vezes, comparando elementos adjacentes e os trocando de posição se estiverem na ordem errada.
// Sua complexidade é O(n²) por conter dois "for" aninhados.
