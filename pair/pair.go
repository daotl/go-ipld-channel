package pair

import (
	"github.com/ipfs/go-cid"

	"github.com/daotl/go-ipld-channel"
)

// ChannelCidPair is a pair of channel and cid.
type ChannelCidPair struct {
	Chn channel.Channel
	Cid cid.Cid
}

func PublicCidPair(c cid.Cid) ChannelCidPair {
	return ChannelCidPair{
		Chn: channel.PublicChannel,
		Cid: c,
	}
}

// Defined returns true if a ChannelCidPair is defined
// Calling any other methods on an undefined ChannelCidPair will result in
// undefined behavior.
func (k ChannelCidPair) Defined() bool {
	return k.Cid.Defined()
}

// Equals checks that two ChannelCidPairs are the same.
func (k ChannelCidPair) Equals(o ChannelCidPair) bool {
	return k == o
}
