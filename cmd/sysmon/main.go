package main

import (
    "sysmon/internal/collector"
    "sysmon/internal/alert"
    "sysmon/internal/config"
    "sysmon/internal/server"
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
        diskUsage := collector.CollectDiskUsage()
        netUsage := collector.CollectNetworkUsage()

        if cpuUsage > cfg.Thresholds.CPU {
            alert.SendAlert("High CPU usage!")
        }

        if memUsage > cfg.Thresholds.Memory {
            alert.SendAlert("High memory usage!")
        }

        if diskUsage > cfg.Thresholds.Disk {
            alert.SendAlert("High disk usage!")
        }

        if netUsage > cfg.Thresholds.Network {
            alert.SendAlert("High network usage!")
        }

        // Sleep for 5 seconds
        time.Sleep(5 * time.Second)
    }
}
