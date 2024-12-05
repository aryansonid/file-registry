package config

// ReadRpcUrlConfig reads rpc url from config.json
func (v *ViperConfig) ReadRpcUrlConfig() string {
	return v.GetString("RPC_URL")
}

// ReadContractAddressConfig reads the contract address
func (v *ViperConfig) ReadContractAddressConfig() string {
	return v.GetString("CONTRACT_ADDRESS")
}

// ReadPrivateKey reads the private key
func (v *ViperConfig) ReadPrivateKey() string {
	return v.GetString("PRIVATE_KEY")
}
