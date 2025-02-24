## Test modbus

``` bash
go test ./internal/modbus_inverter -v
```

## Mock IEC 104 Sever

``` bash
mockServer := NewMockEVNServer()
```