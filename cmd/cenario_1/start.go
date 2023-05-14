package cenario1

import (
	"fmt"

	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/1/machine"
	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/1/machine/sensors/temperature"
	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/1/machine/sensors/volume"
	"github.com/DemetriusADS/cplx_algo_prova_ii/cenarios/1/monitor"
	"github.com/spf13/cobra"
)

const qtyMachines = 10

var StartCenario1 = &cobra.Command{
	Use:                   "cenario-1",
	Short:                 "Executa o cenario 1",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Executando o cenario 1\n")
		var machines []*machine.Machine
		for i := 0; i < qtyMachines; i++ {
			machine := machine.NewMachine(
				fmt.Sprintf("Maquina %d", i),
				volume.NewVolumeSensor(),
				temperature.NewTemperatureSensor(),
			)
			machines = append(machines, machine)
		}

		for i := 0; i < qtyMachines; i++ {
			go machines[i].GenData()
		}
		monitor.NewMonitor(machines).Start()

	},
}
