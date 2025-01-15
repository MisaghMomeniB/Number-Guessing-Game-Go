// Enhanced Go System Monitoring Tool
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// Importing necessary packages for system monitoring
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// SystemStats holds the system usage statistics
// The fields are structured to be serialized into JSON

type SystemStats struct {
	CPUUsage      float64 `json:"cpu_usage"`
	MemoryUsage   string  `json:"memory_usage"`
	MemoryPercent float64 `json:"memory_percent"`
	DiskUsage     string  `json:"disk_usage"`
	Uptime        string  `json:"uptime"`
}

// formatBytes converts bytes into a human-readable format (e.g., KB, MB, GB)
func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	exp := int64(0)
	for n := bytes / unit; n >= unit; n /= unit {
		exp++
	}
	pre := "KMGTPE"[exp : exp+1] // Selects the appropriate size unit
	return fmt.Sprintf("%.2f %sB", float64(bytes)/float64(uint64(unit)<<exp*10), pre)
}

// getSystemStats collects CPU, memory, and disk usage statistics
func getSystemStats() (SystemStats, error) {
	// Fetch CPU usage percentage
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return SystemStats{}, fmt.Errorf("CPU stats error: %v", err)
	}

	// Fetch memory usage statistics
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return SystemStats{}, fmt.Errorf("Memory stats error: %v", err)
	}

	// Fetch disk usage statistics for the root directory
	diskStats, err := disk.Usage("/")
	if err != nil {
		return SystemStats{}, fmt.Errorf("Disk stats error: %v", err)
	}

	// Fetch system uptime
	hostInfo, err := host.Info()
	if err != nil {
		return SystemStats{}, fmt.Errorf("Host stats error: %v", err)
	}

	uptime := time.Duration(hostInfo.Uptime) * time.Second

	// Return collected statistics in a structured format
	return SystemStats{
		CPUUsage:      cpuPercent[0],
		MemoryUsage:   formatBytes(memStats.Used),
		MemoryPercent: memStats.UsedPercent,
		DiskUsage:     formatBytes(diskStats.Used),
		Uptime:        uptime.String(),
	}, nil
}

// statsHandler handles HTTP requests and responds with system statistics in JSON format
func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := getSystemStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Encode the system stats into JSON and send as response
	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// main initializes the HTTP server and starts listening for incoming requests
func main() {
	http.HandleFunc("/stats", statsHandler)
	fmt.Println("Server is running on http://localhost:8080/stats")
	http.ListenAndServe(":8080", nil)
}
