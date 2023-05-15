package main

import (
	cenario1 "github.com/DemetriusADS/cplx_algo_prova_ii/cmd/cenario_1"
	cenario2 "github.com/DemetriusADS/cplx_algo_prova_ii/cmd/cenario_2"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "monitor", DisableFlagsInUseLine: true}

	rootCmd.AddCommand(cenario1.Start)
	rootCmd.AddCommand(cenario2.Start)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
