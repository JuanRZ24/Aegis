/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "aegis/internal/cpu"
    "aegis/internal/memory"
    "aegis/internal/monitor"
)




// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aegis",
	Short: "Aegis - Server Monitoring CLI",
	// has an action associated with it:
	 Run: func(cmd *cobra.Command, args []string) { 
		//lista de monitores que queremos ejecutar
		monitors := []monitor.ResourceMonitor{
			cpu.CPUMonitor{},
			memory.MemoryMonitor{},
		}

		for _, m := range monitors {
			data := m.Collect()
			fmt.Println(m.Format(data))
		}
	 },

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aegis.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


