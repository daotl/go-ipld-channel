package channel

// Channel represents an channel within which IPLD blocks can be store and
// exchanged, possibly with access control.
type Channel string

const (
	// PublicChannel represents the default public channel without access control.
	PublicChannel Channel = ""
)
