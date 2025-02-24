// Core Modbus client

package modbus_inverter

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"

	"github.com/goburrow/modbus"
)

type InverterClient struct {
	handler *modbus.TCPClientHandler
	client  modbus.Client
	config  *InverterConfig
}

// Init connec to inverter
func NewClient(config *InverterConfig) (*InverterClient, error) {
	handler := modbus.NewTCPClientHandler(fmt.Sprintf("%s:%d", config.IP, config.Port))
	handler.Timeout = config.Timeout
	handler.SlaveId = 1 // ID Modbus devices (inverter model)

	if err := handler.Connect(); err != nil {
		return nil, fmt.Errorf("modbus connection failed: %v", err)
	}

	return &InverterClient{
		handler: handler,
		client:  modbus.NewClient(handler),
		config:  config,
	}, nil
}

// Read power P-out from Holding Register 40001 (EVN)
func (c *InverterClient) ReadPowerOutput() (float32, error) {
	data, err := c.client.ReadHoldingRegisters(40000, 2) // Read 2 registers (32-bit float)
	if err != nil {
		return 0, fmt.Errorf("read failed: %v", err)
	}
	return bytesToFloat32(data), nil
}

// Write SetPoint P-out (%) to Holding Register 40013 (EVN)
func (c *InverterClient) WriteSetPoint(value float32) error {
	bytes := float32ToBytes(value)
	_, err := c.client.WriteMultipleRegisters(40012, 2, bytes) // Write 2 registers
	return err
}

// Transfer float32 <-> bytes
func float32ToBytes(f float32) []byte {
	bits := math.Float32bits(f)
	return []byte{byte(bits >> 24), byte(bits >> 16), byte(bits >> 8), byte(bits)}
}

func bytesToFloat32(b []byte) float32 {
	bits := binary.BigEndian.Uint32(b)
	return math.Float32frombits(bits)
}

// Close connection
func (c *InverterClient) Close() {
	c.handler.Close()
}

// Error handling
func (c *InverterClient) ReadPowerOutput() (float32, error) {
	if c.handler == nil {
		return 0, NewWrappedError("read failed", ErrConnectionFailed)
	}

	data, err := c.client.ReadHoldingRegisters(40000, 2)
	if err != nil {
		// Phân loại lỗi
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return 0, NewWrappedError("read timeout", ErrTimeout)
		}
		return 0, NewWrappedError("read failed", err)
	}

	if len(data) != 4 {
		return 0, NewWrappedError("invalid data length", ErrInvalidRegister)
	}

	return bytesToFloat32(data), nil
}
