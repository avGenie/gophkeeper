package config

const (
	grpcConnection = ":9090"
	logLevel       = "debug"
)

type Config struct {
	GRPCConnection string
	LogLevel       string
	LogPath        string
}

func NewConfig() Config {
	return Config{
		GRPCConnection: grpcConnection,
		LogLevel:       logLevel,
		LogPath:        "client.log",
	}
}
