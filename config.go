package hydra

type Config struct {
	IPFSAddr string
	IPFSPort string
	Topics   []string
}

func DefaultConfig() *Config {
	return &Config{
		IPFSAddr: "localhost",
		IPFSPort: "5001",
		Topics:   []string{},
	}
}
