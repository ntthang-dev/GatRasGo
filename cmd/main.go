package main

import (
	"time"

	"./internal/iec104_evn"
	"./internal/modbus_inverter"
)

func main() {
	// Cấu hình Modbus Inverter
	invConfig := modbus_inverter.DefaultConfig()
	invConfig.IP = "192.168.1.100"
	invClient, _ := modbus_inverter.NewClient(invConfig)
	defer invClient.Close()

	// Cấu hình IEC 104 Server
	evnConfig := iec104_evn.DefaultConfig()
	evnConfig.LinkAddress = 1
	evnServer, _ := iec104_evn.NewServer(evnConfig)
	defer evnServer.Stop()

	// Vòng lặp đọc dữ liệu từ inverter và gửi đến EVN
	for {
		pOut, _ := invClient.ReadPowerOutput()
		evnServer.SendPowerOutput(pOut) // Gửi T13
		time.Sleep(5 * time.Second)     // Độ phân giải 5 phút (đổi thành 300s)
	}
}
