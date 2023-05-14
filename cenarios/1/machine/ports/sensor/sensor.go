package sensor

type SensorDTO struct {
	Value float64
	Time  string
}

type Sensor interface {
	Read() *SensorDTO
}
