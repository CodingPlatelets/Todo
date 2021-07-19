package config

var serverConfig map[string]interface{}

func init() {
	serverConfig = make(map[string]interface{})

	serverConfig["host"] = "localhost"
	serverConfig["port"] = "5001"

	serverConfig["mode"] = "debug"
}

func GetServerConfig() map[string]interface{} {
	return serverConfig
}
