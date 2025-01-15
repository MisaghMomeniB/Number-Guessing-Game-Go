// Enhanced Go System Monitoring Tool
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
	CPUUsage      float64 `json:"cpu_usage"`
	MemoryUsage   string  `json:"memory_usage"`
	MemoryPercent float64 `json:"memory_percent"`
	DiskUsage     string  `json:"disk_usage"`
	Uptime        string  `json:"uptime"`
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	exp := int64(0)
	for n := bytes / unit; n >= unit; n /= unit {
		exp++
	}
	pre := "KMGTPE"[exp : exp+1]
	return fmt.Sprintf("%.2f %sB", float64(bytes)/float64(uint64(unit)<<exp*10), pre)
}

func getSystemStats() (SystemStats, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return SystemStats{}, fmt.Errorf("CPU stats error: %v", err)
	}

	memStats, err := mem.VirtualMemory()
	if err != nil {
		return SystemStats{}, fmt.Errorf("Memory stats error: %v", err)
	}

	diskStats, err := disk.Usage("/")
	if err != nil {
		return SystemStats{}, fmt.Errorf("Disk stats error: %v", err)
	}

	hostInfo, err := host.Info()
	if err != nil {
		return SystemStats{}, fmt.Errorf("Host stats error: %v", err)
	}

	uptime := time.Duration(hostInfo.Uptime) * time.Second

	return SystemStats{
		CPUUsage:      cpuPercent[0],
		MemoryUsage:   formatBytes(memStats.Used),
		MemoryPercent: memStats.UsedPercent,
		DiskUsage:     formatBytes(diskStats.Used),
		Uptime:        uptime.String(),
	}, nil
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats, err := getSystemStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/stats", statsHandler)
	fmt.Println("Server is running on http://localhost:8080/stats")
	http.ListenAndServe(":8080", nil)
}
