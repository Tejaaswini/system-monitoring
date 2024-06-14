package collector

import (
    "fmt"
    "github.com/shirou/gopsutil/cpu"
)

func CollectCPUUsage() float64 {
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        fmt.Println("Error getting CPU usage:", err)
        return 0
    }
    return cpuPercent[0]
}
