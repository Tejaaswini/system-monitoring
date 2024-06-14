package collector

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
)

func CollectMemoryUsage() float64 {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting memory usage:", err)
		return 0
	}
	return vmStat.UsedPercent
}
