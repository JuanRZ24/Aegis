package disk

import (
    "fmt"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/JuanRZ24/aegis/internal/monitor"
    "github.com/JuanRZ24/aegis/internal/ui"
)

type DiskMonitor struct{}

func (d DiskMonitor) Name() string {
    return "Disk"
}

func (d DiskMonitor) Collect() monitor.Data {
    usage, err := disk.Usage("/")
    if err != nil || usage == nil {
        return monitor.Data{Metrics: map[string]string{
            "Error": fmt.Sprintf("no se pudo leer disco en '/': %v", err),
        }}
    }

    metrics := make(map[string]string)
    metrics["Total"] = fmt.Sprintf("%.2f GB", float64(usage.Total)/(1024*1024*1024))
    metrics["Used"]  = fmt.Sprintf("%.2f GB", float64(usage.Used)/(1024*1024*1024))
    metrics["Free"]  = fmt.Sprintf("%.2f GB", float64(usage.Free)/(1024*1024*1024))
    metrics["UsedPercent"] = fmt.Sprintf("%.2f%%", usage.UsedPercent)

    return monitor.Data{Metrics: metrics}
}


func (d DiskMonitor) Format(data monitor.Data) string {
    if errMsg, ok := data.Metrics["Error"]; ok {
        return fmt.Sprintf("[Disk]\n  %s\n", errMsg)
    }

    percent := 0.0
    fmt.Sscanf(data.Metrics["UsedPercent"], "%f%%", &percent)
    bar := ui.ProgressBar(percent, 20)
    coloredBar := ui.Colorize(bar, percent)
    coloredVal := ui.Colorize(data.Metrics["UsedPercent"], percent)

    result := "[Disk]\n"
    result += fmt.Sprintf("  Used: %s (%s) / %s %s\n",
        data.Metrics["Used"],
        coloredVal,
        data.Metrics["Total"],
        coloredBar,
    )
    return result
}

