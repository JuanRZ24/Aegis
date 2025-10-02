package process

import (
    "fmt"
    "sort"

    "github.com/shirou/gopsutil/v3/process"
    "github.com/JuanRZ24/aegis/internal/monitor"
)

// ProcessMonitor con estado interno (limit = top N procesos)
type ProcessMonitor struct {
    Limit int
}

func (p ProcessMonitor) Name() string {
    return "Processes"
}

func (p ProcessMonitor) Collect() monitor.Data {
    procs, _ := process.Processes()
    metrics := make(map[string]string)

    type procInfo struct {
        pid  int32
        name string
        cpu  float64
        mem  float32
    }

    var infos []procInfo

    for _, proc := range procs {
        name, _ := proc.Name()
        cpuPercent, _ := proc.CPUPercent()
        memPercent, _ := proc.MemoryPercent()

        infos = append(infos, procInfo{
            pid:  proc.Pid,
            name: name,
            cpu:  cpuPercent,
            mem:  memPercent,
        })
    }

    // Ordenar por CPU descendente
    sort.Slice(infos, func(i, j int) bool {
        return infos[i].cpu > infos[j].cpu
    })

    // Limitar a los N procesos
    limit := p.Limit
    if limit > len(infos) {
        limit = len(infos)
    }

    for i := 0; i < limit; i++ {
        proc := infos[i]
        key := fmt.Sprintf("PID %d - %s", proc.pid, proc.name)
        val := fmt.Sprintf("CPU: %.2f%% | MEM: %.2f%%", proc.cpu, proc.mem)
        metrics[key] = val
    }

    return monitor.Data{Metrics: metrics}
}

func (p ProcessMonitor) Format(data monitor.Data) string {
    result := fmt.Sprintf("[Processes] (Top %d)\n", p.Limit)
    for k, v := range data.Metrics {
        result += fmt.Sprintf("  %s: %s\n", k, v)
    }
    return result
}
