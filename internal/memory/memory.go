package memory

import(
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"aegis/internal/monitor"
)

type MemoryMonitor struct{}

func (m MemoryMonitor) Name() string{
	return "Memory"
}

func (m MemoryMonitor) Collect() monitor.Data{
	v,_ := mem.VirtualMemory()
	metrics := map[string]string{
		"Total": fmt.Sprintf("%v MB", v.Total/1024/1024),
        "Used":  fmt.Sprintf("%v MB (%.2f%%)", v.Used/1024/1024, v.UsedPercent),
        "Free":  fmt.Sprintf("%v MB", v.Free/1024/1024),
	}

	return monitor.Data{Metrics: metrics}
}


func (m MemoryMonitor) Format(data monitor.Data) string {
    result := "[Memory]\n"
    for k, v := range data.Metrics {
        result += fmt.Sprintf("  %s: %s\n", k, v)
    }
    return result
}