package cpu

import (
    "fmt"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/JuanRZ24/aegis/internal/monitor"
    "github.com/JuanRZ24/aegis/internal/ui"
)

type CPUMonitor struct{}

func (c CPUMonitor) Name() string {
    return "CPU"
}

func (c CPUMonitor) Collect() monitor.Data {
    usage, _ := cpu.Percent(0, true) // uso por núcleo
    metrics := make(map[string]string)

    for i, u := range usage {
        key := fmt.Sprintf("Core %d", i)
        metrics[key] = fmt.Sprintf("%.2f%%", u)
    }

    return monitor.Data{Metrics: metrics}
}

func (c CPUMonitor) Format(data monitor.Data) string {
    result := "[CPU]\n"
    for k, v := range data.Metrics {
        percent := 0.0
        fmt.Sscanf(v, "%f%%", &percent)
        bar := ui.ProgressBar(percent, 20)
		coloredBar := ui.Colorize(bar,percent)
		coloredVal := ui.Colorize(v, percent)
        result += fmt.Sprintf("  %s:\t %s\t %s\n", k, coloredVal, coloredBar)
    }
    return result
}
