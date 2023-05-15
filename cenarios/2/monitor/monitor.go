package monitor

import (
	"fmt"

	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/2/machine"
)

type Monitor struct {
	machines  []*machine.Machine
	mChannel  chan string
	mqChannel chan bool
}

func NewMonitor(machines []*machine.Machine, mChannel chan string, mqChannel chan bool) *Monitor {
	return &Monitor{
		machines:  machines,
		mChannel:  mChannel,
		mqChannel: mqChannel,
	}
}

//Complexidade O(1) ou O(n^2)
// Mais uma vez, o for sem condicao de parada, torna a complexidade linear para o primeiro case.
// Para o segundo case, a complexidade é O(n^2), pois o for é executado n vezes, sendo n a quantidade de máquinas e em seguida o for é executado n vezes, sendo n a quantidade de métricas.

func (m *Monitor) Start() {
	fmt.Printf("Iniciando monitoramento\n")
	machinesOff := 0
	for {
		select {
		case msg := <-m.mChannel:
			fmt.Print(msg)
		case <-m.mqChannel:
			machinesOff++
			if machinesOff == len(m.machines) {
				fmt.Printf("Todas as maquinas estão desligadas\n")
				fmt.Printf("------- Realizando leitura para o desligamento do monitoramento -------\n")
				for _, machine := range m.machines {
					fmt.Printf("Lendo metricas da %s\n", machine.Name)
					metrics := machine.Read()
					for _, metric := range metrics {
						fmt.Printf("Temperatura: %f\n Volume: %f\n Leitura Estável: %t\n", metric.Temperature.Value, metric.Volume.Value, !metric.Unstable)
					}
				}
				return
			}
		}
	}
}
