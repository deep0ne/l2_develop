package util

type Config struct {
	ServerAddress string
}

func NewConfig(address string) Config {
	return Config{
		ServerAddress: address,
	}
}
