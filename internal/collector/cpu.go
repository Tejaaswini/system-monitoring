package collector

import (
    "github.com/shirou/gopsutil/cpu"
)

func CollectCPUUsage() float64 {
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        return 0
    }
    return cpuPercent[0]
}

func CollectCPUInfo() ([]cpu.InfoStat, error) {
    return cpu.Info()
}

func CollectCPULoad() ([]cpu.TimesStat, error) {
    return cpu.Times(false) // Collect CPU times for the overall system (false for per-CPU).
}
