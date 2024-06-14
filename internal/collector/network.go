package collector

import (
    "github.com/shirou/gopsutil/net"
)

func CollectNetworkUsage() (map[string]net.IOCountersStat, error) {
    counters, err := net.IOCounters(true)
    if err != nil {
        return nil, err
    }

    usage := make(map[string]net.IOCountersStat)
    for _, counter := range counters {
        usage[counter.Name] = counter
    }
    return usage, nil
}
