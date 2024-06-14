package main

import (
    "sysmon/internal/alert"
    "sysmon/internal/config"
    "sysmon/internal/server"
    "sysmon/internal/collector"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/net"
    "time"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        panic(err)
    }

    alert.InitConfig(cfg)

    // Start the server in a separate goroutine
    go server.StartServer()
    go server.CollectMetrics()

    for {
        cpuUsage := collector.CollectCPUUsage()
        memUsage := collector.CollectMemoryUsage()
        
        diskUsage, diskErr := collector.CollectDiskUsage()
        if diskErr != nil {
            // Handle disk collection error
            diskUsage = make(map[string]*disk.UsageStat)
        }

        netUsage, netErr := collector.CollectNetworkUsage()
        if netErr != nil {
            // Handle network collection error
            netUsage = make(map[string]net.IOCountersStat)
        }

        if cpuUsage > cfg.Thresholds.CPU {
            alert.SendAlert("High CPU usage!")
        }

        if memUsage > cfg.Thresholds.Memory {
            alert.SendAlert("High memory usage!")
        }

        for _, usage := range diskUsage {
            if usage.UsedPercent > cfg.Thresholds.Disk {
                alert.SendAlert("High disk usage on " + usage.Path)
            }
        }

        for _, usage := range netUsage {
            if float64(usage.BytesSent+usage.BytesRecv) > cfg.Thresholds.Network {
                alert.SendAlert("High network usage on " + usage.Name)
            }
        }

        // Sleep for 5 seconds
        time.Sleep(5 * time.Second)
    }
}
