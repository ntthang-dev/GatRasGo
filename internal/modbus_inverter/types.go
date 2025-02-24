// Data structures
package modbus_inverter

import "time"

// DeviceConfig định nghĩa cấu hình cho 1 inverter
type DeviceConfig struct {
    IP          string        `yaml:"ip"`
    Port        int           `yaml:"port"`
    SlaveID     byte          `yaml:"slave_id"`
    Timeout     time.Duration `yaml:"timeout"`
}

// RegisterMap ánh xạ địa chỉ Modbus
type RegisterMap struct {
    PowerOutput uint16 `yaml:"power_output"` // 40001
    SetPoint    uint16 `yaml:"set_point"`    // 40013
}

// ResponseData chứa dữ liệu đọc từ inverter
type ResponseData struct {
    Timestamp time.Time
    Power     float32
    Voltage   float32
    Error     error
}