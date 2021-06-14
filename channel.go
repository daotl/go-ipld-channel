package channel

// Channel represents an IPLD channel within which blocks can be store and
// exchanged, possibly with access control.
type Channel string

const (
	// PublicChannel represents the default public IPLD channel without access control.
	PublicChannel Channel = ""
)

func (c Channel) String() string {
	return string(c)
}

// StringsToChannels converts a slice of strings to a slice of Channels.
func StringToChannels(strs []string) []Channel {
	chns := make([]Channel, 0, len(strs))
	for i, s := range strs {
		chns[i] = Channel(s)
	}
	return chns
}

// ChannelsToStrings converts a slice of channels to a slice of strings.
func ChannelsToStrings(chns []Channel) []string {
	strs := make([]string, 0, len(chns))
	for i, chn := range chns {
		strs[i] = chn.String()
	}
	return strs
}
