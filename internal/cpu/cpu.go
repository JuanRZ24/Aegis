package CPU

import(
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"aegis/internal/monitor"
)

Type CPUMonitor struct {}


func (c CPUMonitor) Name() string{
	return "CPU"
}

func (c CPUMonitor) Collect() monitor.Data {
	usage, _ := cpu.Percent(0,true)//uso por nucleo
	metrics := make(map[string]string)

	for i, u := range usage {
		key := fmt.Sprintf("Core %d", i)
		metrics[key] = fmt.Sprintf("%.2f%%",u)
	}

	return monitor.Data{Metrics:metrics}
}

func (c CPUMonitor) Format(data monitor.Data) string{
	result := "[CPU]\n"
	for k,v := range data.Metrics{
		result += fmt.Sprintf(" %s: %s\n",k,v)
	}
	return result
}