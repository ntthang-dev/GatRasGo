package iec104_evn

import (
	"github.com/simonvetter/iec104"
)

// handleControlCommand xử lý lệnh điều khiển từ EVN
func (s *EVNServer) handleControlCommand(asdu *iec104.ASDU) {
	cmd := EVNCommand{
		Type:      asdu.Type,
		Address:   asdu.Objects[0].IOA,
		Value:     asdu.Objects[0].Value.Float32(),
		Qualifier: asdu.Objects[0].Qualifier,
	}

	switch cmd.Type {
	case iec104.C_SC_NA_1: // Single command
		s.handleSingleCommand(cmd)
	case iec104.C_SE_NC_1: // Setpoint command
		s.handleSetpointCommand(cmd)
	default:
		s.logger.Printf("Unsupported command type: %v", cmd.Type)
	}
}

// handleSingleCommand xử lý lệnh đơn (ví dụ bật/tắt)
func (s *EVNServer) handleSingleCommand(cmd EVNCommand) {
	// Giả sử địa chỉ 11 là lệnh enable P-out
	if cmd.Address == 11 {
		s.EnableSystem(cmd.Value > 0)
		s.SendConfirmation(cmd.Address, true)
	}
}

// handleSetpointCommand xử lý lệnh setpoint
func (s *EVNServer) handleSetpointCommand(cmd EVNCommand) {
	// Địa chỉ 12 là SetPoint P-out
	if cmd.Address == 12 {
		err := s.modbusClient.WriteSetPoint(cmd.Value)
		if err == nil {
			s.SendConfirmation(cmd.Address, true)
		} else {
			s.SendConfirmation(cmd.Address, false)
		}
	}
}
