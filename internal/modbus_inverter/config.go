// Package connection configuration for Modbus Inverter

package modbus_inverter

import "time"

type InverterConfig struct {
	IP      string        `yaml:"ip"`      // IP inverter (e.g., "192.168.1.10")
	Port    int           `yaml:"port"`    // Port (default 502)
	Timeout time.Duration `yaml:"timeout"` // Timeout (e.g., 10 * time.Second)
	Retries int           `yaml:"retries"` // Count retry when disconneted
}

// Default config for inverter
func DefaultConfig() *InverterConfig {
	return &InverterConfig{
		Port:    502,
		Timeout: 10 * time.Second,
		Retries: 3,
	}
}
