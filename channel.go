package channel

// Channel represents an IPLD channel within which blocks can be store and
// exchanged, possibly with access control.
type Channel string

const (
	// PublicChannel represents the default public IPLD channel without access control.
	PublicChannel Channel = ""
)

func (c Channel) toString() string {
	return string(c)
}
