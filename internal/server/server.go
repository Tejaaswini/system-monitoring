package server

import (
    "fmt"
    "net/http"
    "sysmon/internal/collector"
)

func StartServer() {
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        cpuUsage := collector.CollectCPUUsage()
        memUsage := collector.CollectMemoryUsage()
        diskUsage := collector.CollectDiskUsage()
        netUsage := collector.CollectNetworkUsage()
        
        fmt.Fprintf(w, "CPU Usage: %.2f%%\n", cpuUsage)
        fmt.Fprintf(w, "Memory Usage: %.2f%%\n", memUsage)
        fmt.Fprintf(w, "Disk Usage: %.2f%%\n", diskUsage)
        fmt.Fprintf(w, "Network Usage: %.2f KBps\n", netUsage)
    })

    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
