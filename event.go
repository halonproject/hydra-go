package hydra

// Event defines either an error or message passed from IPFS pubsub.
type Event interface {
	String() string
}
