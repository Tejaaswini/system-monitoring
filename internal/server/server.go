package server

import (
    "encoding/json"
    "fmt"
    "html/template"
    "net/http"
    "sysmon/internal/collector"
    "time"
)

type Metrics struct {
    CPU     []Metric `json:"cpu"`
    Memory  []Metric `json:"memory"`
    Disk    []Metric `json:"disk"`
    Network []Metric `json:"network"`
}

type Metric struct {
    Time  time.Time `json:"time"`
    Value float64   `json:"value"`
}

var metrics Metrics

func StartServer() {
    http.HandleFunc("/", serveIndex)
    http.HandleFunc("/metrics", serveMetrics)

    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("internal/server/templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

func serveMetrics(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(metrics)
}

func CollectMetrics() {
    for {
        metrics.CPU = append(metrics.CPU, Metric{Time: time.Now(), Value: collector.CollectCPUUsage()})
        metrics.Memory = append(metrics.Memory, Metric{Time: time.Now(), Value: collector.CollectMemoryUsage()})
        metrics.Disk = append(metrics.Disk, Metric{Time: time.Now(), Value: collector.CollectDiskUsage()})
        metrics.Network = append(metrics.Network, Metric{Time: time.Now(), Value: collector.CollectNetworkUsage()})

        // Keep only the last 100 metrics
        if len(metrics.CPU) > 100 {
            metrics.CPU = metrics.CPU[1:]
        }
        if len(metrics.Memory) > 100 {
            metrics.Memory = metrics.Memory[1:]
        }
        if len(metrics.Disk) > 100 {
            metrics.Disk = metrics.Disk[1:]
        }
        if len(metrics.Network) > 100 {
            metrics.Network = metrics.Network[1:]
        }

        time.Sleep(5 * time.Second)
    }
}
