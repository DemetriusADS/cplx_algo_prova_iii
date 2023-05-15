package monitor

import (
	"errors"
	"fmt"
	"time"

	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/1/machine"
)

type Monitor struct {
	machines []*machine.Machine
}

func NewMonitor(machines []*machine.Machine) *Monitor {
	return &Monitor{
		machines: machines,
	}
}

// Complexidade O(n^4)
// O código em questão possui tal complexidade pois o método Start() possui um loop for para executar o polling a todo instante. (Como em Go nao existe a estrutura while ou do while, é utilizando esse for, sem condicoes). Na sequencia, é executado outro loop for para verificar se a máquina está ligada. Caso esteja, é executado outro loop for para verificar se a máquina está com as métricas instáveis. Caso esteja, é executado outro loop for para ajustar a temperatura da máquina.
// Caso a maquina esteja desligada, é executado outro loop for para verificar se todas as máquinas estão desligadas. Caso estejam, o monitoramento é finalizado.
func (m *Monitor) Start() {
	fmt.Printf("Iniciando monitoramento\n")
	machinesOff := []string{}
	for {
		for _, machine := range m.machines {
			if continueLoop, err := m.checkMachineStatus(machine, &machinesOff); err != nil {
				fmt.Printf("ERRO: %s. Finalizando monitoramento\n", err.Error())
				return
			} else if continueLoop {
				continue
			}

			metrics := machine.Read()
			for _, metric := range metrics {
				now := time.Now().Format("2006-01-02 15:04:05")
				if !metric.Unstable {
					continue
				}
				fmt.Printf("A MAQUINA: %s, POSSUI METRICAS INSTÁVEIS. CALIBRANDO...\n", machine.Name)
				fmt.Printf("TEMPERATURA ATUAL: %f\n", metric.Temperature.Value)
				if metric.Volume.Value > 0 {
					metric.Temperature.Value = metric.Volume.Value * 2.5
				} else {
					metric.Temperature.Value = 0
				}
				metric.Temperature.Time = now
				metric.Volume.Time = now
				fmt.Printf("TEMPERATURA AJUSTADA: %f\n", metric.Temperature.Value)
				metric.Unstable = false
			}
		}
	}
}

// Complexidade O(n)
// O código em questão possui tal complexidade pois o método checkMachineStatus() verifica se a máquina está ligada. Caso esteja, é executado um loop for para verificar se essa máquina já está registrada como desligada. Caso esteja, o loop é finalizado. Caso não esteja, a máquina é registrada e o loop encerrado.
// Mas, primeiramente, para evitar processamento desnecessário e otimizar o código, é verificado se todas as máquinas estão desligadas. Caso estejam, o monitoramento é finalizado.
func (m *Monitor) checkMachineStatus(machine *machine.Machine, machinesOff *[]string) (bool, error) {
	if !machine.IsOn() {
		if len(*machinesOff) == len(m.machines) {
			return true, errors.New("todas as maquinas estão desligadas")
		}
		for _, name := range *machinesOff {
			if name == machine.Name {
				return true, nil
			}
		}
		*machinesOff = append(*machinesOff, machine.Name)
		return true, nil
	}
	return false, nil
}

/**
* TODO List:
* 1 - Inicializar 10 maquinas
* 2 - Gerar um menu interativo para o usuario
* 3 - Menu deve conter: Listar maquinas, Listas as métricas de cada maquina monitorada, ordenacao crescente dos dados das maquinas
* por fim, um algoritmo O(N^2) ou N^3.
**/
