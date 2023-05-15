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
	Metrics  []*Metric
	Name     string
	isOn     bool
	mChannel chan string
	mQuit    chan bool

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
		if machine.mQuit != nil {
			machine.mQuit <- true
		}
	}()

	return &machine
}

func (m *Machine) RegisterChannel(c chan string, q chan bool) {
	m.mChannel = c
	m.mQuit = q
}

// Complexidade O(n)
// Justificativa: A função GenData é executada em O(n), pois há um loop que executa n vezes, sendo n a quantidade de métricas.
func (m *Machine) GenData() {
	for i := 0; i < qtyMetrics; i++ {
		// time.Sleep(1 * time.Second)
		temp := m.TemperatureSensor.Read()
		vol := m.VolumeSensor.Read()
		metricGen := &Metric{
			Temperature: *temp,
			Volume:      *vol,
			Unstable:    true,
		}
		m.Metrics = append(m.Metrics, metricGen)
		go func(metric *Metric) {
			time.Sleep(1 * time.Second)
			m.FixTemperature(metric)
		}(metricGen)
		m.mChannel <- fmt.Sprintf("%s\n Temperatura: %f\n Volume: %f\n Leitura Calibrada: %t\n", m.Name, metricGen.Temperature.Value, metricGen.Volume.Value, !metricGen.Unstable)
	}
}

func (m *Machine) Read() []*Metric {
	return m.Metrics
}

// Complexidade O(1)
// Justificativa: A função FixTemperature é executada em uma goroutine, portanto, não é considerada na complexidade.
// Mas, analisando o código em si, a função FixTemperature é executada em O(1), pois não há nenhum loop ou estrutura de repetição.
func (m *Machine) FixTemperature(metric *Metric) {
	now := time.Now().Format("2006-01-02 15:04:05")
	if !metric.Unstable {
		return
	}
	fmt.Printf("TEMPERATURA ATUAL DA %s: %f\n", m.Name, metric.Temperature.Value)
	if metric.Volume.Value > 0 {
		metric.Temperature.Value = metric.Volume.Value * 2.5
	} else {
		metric.Temperature.Value = 0
	}
	metric.Temperature.Time = now
	metric.Volume.Time = now
	fmt.Printf("TEMPERATURA AJUSTADA DA %s:: %f\n", m.Name, metric.Temperature.Value)
	metric.Unstable = false
}

func (m *Machine) IsOn() bool {
	return m.isOn
}
