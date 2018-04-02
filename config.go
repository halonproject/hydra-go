package hydra

// Config holds the needed values for either a producer or consumer to connect to
// an IPFS client and a initial set of topics that the producer or consumer are
// subscribed on creation.
type Config struct {
	IPFSAddr string
	IPFSPort string
	Topics   []string
}

// DefaultConfig returns a hydra config that points to localhost:5001 and has no
// topics configured to pull from.
func DefaultConfig() *Config {
	return &Config{
		IPFSAddr: "localhost",
		IPFSPort: "5001",
		Topics:   []string{},
	}
}
