package main

import (
    "sysmon/internal/collector"
    "sysmon/internal/alert"
    "sysmon/internal/server"
    "time"
)

func main() {
    cpuThreshold := 80.0
    memThreshold := 80.0

    // Start the server in a separate goroutine
    go server.StartServer()

    for {
        cpuUsage := collector.CollectCPUUsage()
        memUsage := collector.CollectMemoryUsage()

        if cpuUsage > cpuThreshold {
            alert.SendAlert("High CPU usage!")
        }

        if memUsage > memThreshold {
            alert.SendAlert("High memory usage!")
        }

        // Sleep for 5 seconds
        time.Sleep(5 * time.Second)
    }
}
