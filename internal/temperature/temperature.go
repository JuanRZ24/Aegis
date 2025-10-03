package temperature

import (
    "fmt"

    "github.com/shirou/gopsutil/v3/host"
    "github.com/JuanRZ24/aegis/internal/monitor"
    "github.com/JuanRZ24/aegis/internal/ui"
)

type TempMonitor struct{}

func (t TempMonitor) Name() string {
    return "Temperature"
}

func (t TempMonitor) Collect() monitor.Data {
    temps, _ := host.SensorsTemperatures()
    metrics := make(map[string]string)

    for _, temp := range temps {
        key := temp.SensorKey
        val := fmt.Sprintf("%.1f°C", temp.Temperature)
        metrics[key] = val
    }

    return monitor.Data{Metrics: metrics}
}

func (t TempMonitor) Format(data monitor.Data) string {
    result := "[Temperature]\n"
    if len(data.Metrics) == 0 {
        result += "  No sensors detected (maybe not supported on this OS)\n"
        return result
    }

    for k, v := range data.Metrics {
        percent := 0.0
        fmt.Sscanf(v, "%f°C", &percent)
        bar := ui.ProgressBar(percent, 20)
        colored := ui.Colorize(v, percent)
        result += fmt.Sprintf("  %s: %s %s\n", k, colored, bar)
    }
    return result
}

