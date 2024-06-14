package collector

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

func CollectDiskUsage() float64 {
	diskStat, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Error getting disk usage:", err)
		return 0
	}
	return diskStat.UsedPercent
}
