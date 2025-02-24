package modbus_inverter

import "time"

// executeCommand thực thi Modbus command với retry logic
func (c *InverterClient) executeCommand(
	fn func() ([]byte, error),
) ([]byte, error) {
	var err error
	var result []byte

	for attempt := 0; attempt < c.config.Retries; attempt++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
		// Log retry attempt
		c.logger.Printf("Retry %d for command failed: %v", attempt+1, err)
		time.Sleep(time.Duration(attempt+1) * time.Second) // Exponential backoff
	}
	return nil, err
}

// ReadMultipleRegisters wrapper với retry
func (c *InverterClient) ReadMultipleRegisters(address, quantity uint16) ([]byte, error) {
	return c.executeCommand(func() ([]byte, error) {
		return c.client.ReadHoldingRegisters(address, quantity)
	})
}

// EXAMPLE
// Custom Command cho inverter Huawei
func (c *InverterClient) ReadHuaweiDiagnostics() ([]byte, error) {
	// Gửi command đặc biệt 0x03 đến register 40108
	return c.executeCommand(func() ([]byte, error) {
		return c.client.ReadHoldingRegisters(40107, 10) // Đọc 10 registers từ 40108
	})
}
