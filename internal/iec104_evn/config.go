// Package connfig  IEC 60870-5-104 sever send data to EVN

package iec104_evn

type EVNConfig struct {
	ListenPort  int    `yaml:"listen_port"`  // Port server (default 2404)
	LinkAddress uint16 `yaml:"link_address"` // Link Adress (EVN)
	ASDUAddress uint16 `yaml:"asdu_address"` // Adress ASDU
}

func DefaultConfig() *EVNConfig {
	return &EVNConfig{
		ListenPort: 2404,
	}
}
