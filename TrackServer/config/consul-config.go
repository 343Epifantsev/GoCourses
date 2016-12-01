// consul-config
package config

import (
	consulapi "github.com/hashicorp/consul/api"
)

type TrackServerConfig struct {
	Port string
}

func RecieveConfig() (TrackServerConfig, error) {
	var result TrackServerConfig
	config := consulapi.DefaultConfig()
	config.Address = "138.197.17.115"
	consul, err := consulapi.NewClient(config)
	if err == nil {
		kv := consul.KV()
		port, _, err := kv.Get("trackserver/port", nil)
		if err == nil {
			result = TrackServerConfig{string(port.Value)}
		}
	}
	return result, err
}
