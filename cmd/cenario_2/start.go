package cenario1

import (
	"fmt"

	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/2/machine"
	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/2/machine/sensors/temperature"
	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/2/machine/sensors/volume"
	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/2/monitor"
	"github.com/spf13/cobra"
)

const qtyMachines = 10

var Start = &cobra.Command{
	Use:                   "cenario-2",
	Short:                 "Executa o cenario 2",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executando o cenario 2\n")
		var machines []*machine.Machine
		machineChannel := make(chan machine.MetricDTO)
		machineQuitChannel := make(chan bool)
		for i := 0; i < qtyMachines; i++ {
			machine := machine.NewMachine(
				fmt.Sprintf("Maquina %d", i),
				volume.NewVolumeSensor(),
				temperature.NewTemperatureSensor(),
			)
			machine.RegisterChannel(machineChannel, machineQuitChannel)
			machines = append(machines, machine)
		}
		for i := 0; i < qtyMachines; i++ {
			go machines[i].GenData()
		}
		monitor.NewMonitor(machines, machineChannel, machineQuitChannel).Start()

	},
}
