package collector

import (
    "github.com/shirou/gopsutil/mem"
)

func CollectMemoryUsage() float64 {
    vmStat, err := mem.VirtualMemory()
    if err != nil {
        return 0
    }
    return vmStat.UsedPercent
}

func CollectMemoryInfo() (*mem.VirtualMemoryStat, error) {
    return mem.VirtualMemory()
}

func CollectSwapInfo() (*mem.SwapMemoryStat, error) {
    return mem.SwapMemory()
}
