// module types.go in package iec104_evn
// định ngĩa các dữ liệu đặc trưng của IEC104
// chuẩn hoá lênh nhận từ EVN, cấu trúc dữ liệu Telemetry, cấu hình nâng cao cho sever
package iec104_evn

import "github.com/robinson/gos7"

// EVNCommand định nghĩa lệnh từ EVN
type EVNCommand struct {
    Type    iec104.ASDUType
    Address uint16   // IOA - Information Object Address
    Value   float32  // Giá trị điều khiển
    Qualifier byte   // Qualifier cho lệnh
}

// TelemetryData dữ liệu gửi lên EVN
type TelemetryData struct {
    Timestamp   int64     // Unix timestamp
    Power       float32   // P-out (kW)
    Voltage     float32   // Ua (V)
    IsValid     bool      // Data validity flag
}

// ServerConfig cấu hình mở rộng
type ServerConfig struct {
    MaxConnections int    `yaml:"max_connections"`
    TLSEnabled     bool   `yaml:"tls_enabled"`
}