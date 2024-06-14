package server

import (
    "encoding/json"
    "fmt"
    "html/template"
    "net/http"
    "strconv"
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
    http.HandleFunc("/system", serveSystemInfo)
    http.HandleFunc("/memory", serveMemoryInfo)
    http.HandleFunc("/swap", serveSwapInfo)
    http.HandleFunc("/disks", serveDiskInfo)
    http.HandleFunc("/network", serveNetworkInfo)
    http.HandleFunc("/processes", serveProcesses)
    http.HandleFunc("/processor-usage-historical", serveProcessorUsageHistorical)
    http.HandleFunc("/memory-historical", serveMemoryHistorical)
    http.HandleFunc("/services", serveServices)
    http.HandleFunc("/custom-metric-names", serveCustomMetricNames)
    http.HandleFunc("/custom", serveCustomMetricValues)

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

func serveSystemInfo(w http.ResponseWriter, r *http.Request) {
    info, err := collector.CollectCPUInfo()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(info)
}

func serveMemoryInfo(w http.ResponseWriter, r *http.Request) {
    info, err := collector.CollectMemoryInfo()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(info)
}

func serveSwapInfo(w http.ResponseWriter, r *http.Request) {
    info, err := collector.CollectSwapInfo()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(info)
}

func serveDiskInfo(w http.ResponseWriter, r *http.Request) {
    info, err := collector.CollectDiskUsage()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(info)
}

func serveNetworkInfo(w http.ResponseWriter, r *http.Request) {
    info, err := collector.CollectNetworkUsage()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(info)
}

func serveProcesses(w http.ResponseWriter, r *http.Request) {
    limitStr := r.URL.Query().Get("limit")
    limit, err := strconv.Atoi(limitStr)
    if err != nil {
        limit = 10
    }
    processes, err := collector.CollectTopProcesses(limit)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(processes)
}

func serveProcessorUsageHistorical(w http.ResponseWriter, r *http.Request) {
    load, err := collector.CollectCPULoad()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(load)
}

func serveMemoryHistorical(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(metrics.Memory)
}

func serveServices(w http.ResponseWriter, r *http.Request) {
    services, err := collector.CollectServices()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(services)
}

func serveCustomMetricNames(w http.ResponseWriter, r *http.Request) {
    names, err := collector.CollectCustomMetricNames()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(names)
}

func serveCustomMetricValues(w http.ResponseWriter, r *http.Request) {
    metricName := r.URL.Query().Get("custom-metric")
    values, err := collector.CollectCustomMetricValues(metricName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(values)
}

func CollectMetrics() {
    for {
        cpuUsage := collector.CollectCPUUsage()
        memUsage := collector.CollectMemoryUsage()
        diskUsage, diskErr := collector.CollectDiskUsage()
        netUsage, netErr := collector.CollectNetworkUsage()

        if diskErr != nil {
            fmt.Println("Error collecting disk usage:", diskErr)
        } else {
            for _, usage := range diskUsage {
                metrics.Disk = append(metrics.Disk, Metric{Time: time.Now(), Value: usage.UsedPercent})
            }
        }

        if netErr != nil {
            fmt.Println("Error collecting network usage:", netErr)
        } else {
            for _, usage := range netUsage {
                metrics.Network = append(metrics.Network, Metric{Time: time.Now(), Value: float64(usage.BytesSent + usage.BytesRecv)})
            }
        }

        metrics.CPU = append(metrics.CPU, Metric{Time: time.Now(), Value: cpuUsage})
        metrics.Memory = append(metrics.Memory, Metric{Time: time.Now(), Value: memUsage})

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
