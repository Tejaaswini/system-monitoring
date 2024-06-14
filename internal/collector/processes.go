package collector

import (
    "sort"

    "github.com/shirou/gopsutil/process"
)

type ProcessUsage struct {
    PID   int32   `json:"pid"`
    Name  string  `json:"name"`
    CPU   float64 `json:"cpu"`
    Memory float32 `json:"memory"`
}

func CollectTopProcesses(limit int) ([]ProcessUsage, error) {
    processes, err := process.Processes()
    if err != nil {
        return nil, err
    }

    var usage []ProcessUsage
    for _, p := range processes {
        cpuPercent, err := p.CPUPercent()
        if err != nil {
            continue
        }

        memPercent, err := p.MemoryPercent()
        if err != nil {
            continue
        }

        name, err := p.Name()
        if err != nil {
            continue
        }

        usage = append(usage, ProcessUsage{
            PID:    p.Pid,
            Name:   name,
            CPU:    cpuPercent,
            Memory: memPercent,
        })
    }

    sort.Slice(usage, func(i, j int) bool {
        return usage[i].CPU > usage[j].CPU
    })

    if len(usage) > limit {
        usage = usage[:limit]
    }

    return usage, nil
}
