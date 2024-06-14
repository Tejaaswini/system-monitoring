package collector

import (
	"errors"
	"time"
)

// CustomMetric represents a custom metric.
type CustomMetric struct {
	Name  string    `json:"name"`
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

// In-memory storage for custom metrics (this can be replaced by a more persistent storage solution)
var customMetrics = make(map[string][]CustomMetric)

// RegisterCustomMetric registers a new custom metric.
func RegisterCustomMetric(name string, value float64) {
	metric := CustomMetric{
		Name:  name,
		Time:  time.Now(),
		Value: value,
	}
	customMetrics[name] = append(customMetrics[name], metric)

	// Keep only the last 100 metrics for each custom metric
	if len(customMetrics[name]) > 100 {
		customMetrics[name] = customMetrics[name][1:]
	}
}

// CollectCustomMetricNames returns the names of all registered custom metrics.
func CollectCustomMetricNames() ([]string, error) {
	var names []string
	for name := range customMetrics {
		names = append(names, name)
	}
	return names, nil
}

// CollectCustomMetricValues returns the values of a specific custom metric.
func CollectCustomMetricValues(name string) ([]CustomMetric, error) {
	metrics, exists := customMetrics[name]
	if !exists {
		return nil, errors.New("custom metric not found")
	}
	return metrics, nil
}
