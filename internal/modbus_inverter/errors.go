// Custom errors
package modbus_inverter

import "fmt"

// Custom error types
var (
	ErrConnectionFailed = fmt.Errorf("modbus connection failed")
	ErrTimeout          = fmt.Errorf("modbus request timeout")
	ErrInvalidRegister  = fmt.Errorf("invalid register address")
	ErrCRC              = fmt.Errorf("crc check failed")
)

// WrappedError bao bọc lỗi gốc với context
type WrappedError struct {
	Message string
	Err     error
}

func (e *WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

// NewWrappedError tạo lỗi bao bọc
func NewWrappedError(msg string, err error) *WrappedError {
	return &WrappedError{Message: msg, Err: err}
}
