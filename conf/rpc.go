package conf

import "github.com/AstraProtocol/astra-go-sdk/config"

func GetEvmRpcConfig() config.Config {
	return config.Config{configuration.EvmRpc.ChainID,
		configuration.EvmRpc.EndPoint,
		configuration.EvmRpc.PrefixAddress,
		configuration.EvmRpc.TokenSymbol,
	}
}
