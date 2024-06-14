package collector

import (
    "fmt"
    "github.com/shirou/gopsutil/net"
)

func CollectNetworkUsage() float64 {
    ioCounters, err := net.IOCounters(false)
    if err != nil {
        fmt.Println("Error getting network usage:", err)
        return 0
    }
    if len(ioCounters) > 0 {
        return float64(ioCounters[0].BytesSent+ioCounters[0].BytesRecv) / 1024 // in KB
    }
    return 0
}
