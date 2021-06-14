package pair

import (
	"github.com/ipfs/go-cid"

	channel "github.com/daotl/go-ipld-channel"
)

// ChannelCidPair is a pair of IPLD channel and cid.
type ChannelCidPair struct {
	Chn channel.Channel
	Cid cid.Cid
}

// Defined returns true if a ChannelCidPair is defined
// Calling any other methods on an undefined ChannelCidPair will result in
// undefined behavior.
func (p ChannelCidPair) Defined() bool {
	return p.Cid.Defined()
}

// Equals checks that two ChannelCidPairs are the same.
func (p ChannelCidPair) Equals(o ChannelCidPair) bool {
	return p == o
}

// PublicCidPair returns a pair consisting of the public IPLD channel and the specified cid.
func PublicCidPair(c cid.Cid) ChannelCidPair {
	return ChannelCidPair{
		Chn: channel.PublicChannel,
		Cid: c,
	}
}

func CidsToPairs(cids []cid.Cid, chn channel.Channel) []ChannelCidPair {
	pairs := make([]ChannelCidPair, 0, len(cids))
	for _, c := range cids {
		pairs = append(pairs, ChannelCidPair{chn, c})
	}
	return pairs
}

func PairsToCids(pairs []ChannelCidPair) []cid.Cid {
	cids := make([]cid.Cid, 0, len(pairs))
	for _, key := range pairs {
		cids = append(cids, key.Cid)
	}
	return cids
}
