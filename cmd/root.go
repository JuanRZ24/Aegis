package cmd

import (
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "time"

    "github.com/spf13/cobra"
    "github.com/JuanRZ24/aegis/internal/cpu"
    "github.com/JuanRZ24/aegis/internal/memory"
    "github.com/JuanRZ24/aegis/internal/disk"
    "github.com/JuanRZ24/aegis/internal/process"
    "github.com/JuanRZ24/aegis/internal/monitor"
    "github.com/JuanRZ24/aegis/internal/temperature"
)

var rootCmd = &cobra.Command{
    Use:   "aegis",
    Short: "Aegis - Server Monitoring CLI",
    Run: func(cmd *cobra.Command, args []string) {
        monitors := []monitor.ResourceMonitor{
            cpu.CPUMonitor{},
            memory.MemoryMonitor{},
            disk.DiskMonitor{},
            process.ProcessMonitor{Limit: 5},
            temperature.TempMonitor{},

        }

        for {
            clearScreen()

            fmt.Println("=== Aegis - Server Monitoring ===")
            for _, m := range monitors {
                data := m.Collect()
                fmt.Println(m.Format(data))
            }

            time.Sleep(2 * time.Second)
        }
    },
}

func clearScreen() {
    switch runtime.GOOS {
    case "windows":
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        _ = cmd.Run()
    default:
        fmt.Print("\033[H\033[2J")
    }
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
