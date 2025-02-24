package iec104_evn

import (
	"encoding/binary"
	"fmt"
	"net"
)

// ASDU Type IDs theo chuẩn IEC 104
const (
	C_SC_NA_1 = 45 // Single command (T45)
	C_SE_NC_1 = 50 // Setpoint command (T50)
	M_ME_NC_1 = 13 // Measured value (T13)
)

type EVNServer struct {
	listener net.Listener
}

func NewServer(port int) (*EVNServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	return &EVNServer{listener: listener}, nil
}

func (s *EVNServer) Start() {
	for {
		conn, _ := s.listener.Accept()
		go s.handleConnection(conn)
	}
}

// Xử lý kết nối IEC 104
func (s *EVNServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		// Parse ASDU frame
		if n >= 6 {
			typeID := buf[0]
			switch typeID {
			case C_SC_NA_1:
				// Xử lý lệnh T45
				ioAddress := binary.BigEndian.Uint16(buf[1:3])
				value := buf[5] & 0x01 // Bit đầu tiên là giá trị
				s.handleControlCommand(ioAddress, value)
			}
		}
	}
}
