package collector

import (
    "github.com/shirou/gopsutil/disk"
)

func CollectDiskUsage() (map[string]*disk.UsageStat, error) {
    partitions, err := disk.Partitions(false)
    if err != nil {
        return nil, err
    }

    usage := make(map[string]*disk.UsageStat)
    for _, partition := range partitions {
        usageStat, err := disk.Usage(partition.Mountpoint)
        if err != nil {
            continue
        }
        usage[partition.Mountpoint] = usageStat
    }
    return usage, nil
}
