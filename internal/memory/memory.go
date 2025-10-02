package memory

import (
    "fmt"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/JuanRZ24/aegis/internal/monitor"
    "github.com/JuanRZ24/aegis/internal/ui"
)

type MemoryMonitor struct{}

func (m MemoryMonitor) Name() string {
    return "Memory"
}

func (m MemoryMonitor) Collect() monitor.Data {
    vm, _ := mem.VirtualMemory()
    metrics := make(map[string]string)

    metrics["Total"] = fmt.Sprintf("%d MB", vm.Total/1024/1024)
    metrics["Used"] = fmt.Sprintf("%d MB", vm.Used/1024/1024)
    metrics["UsedPercent"] = fmt.Sprintf("%.2f%%", vm.UsedPercent)

    return monitor.Data{Metrics: metrics}
}

func (m MemoryMonitor) Format(data monitor.Data) string {
    percent := 0.0
    fmt.Sscanf(data.Metrics["UsedPercent"], "%f%%", &percent)
    bar := ui.ProgressBar(percent, 20)
	coloredBar := ui.Colorize( bar,percent)
	coloredPercent := ui.Colorize(data.Metrics["UsedPercent"],percent)
    result := "[Memory]\n"
    result += fmt.Sprintf("  Used: %s (%s) / %s %s\n",
        data.Metrics["Used"],
        coloredPercent,
        data.Metrics["Total"],
        coloredBar,
    )
    return result
}
