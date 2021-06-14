package pair

import (
	"github.com/ipfs/go-cid"

	channel "github.com/daotl/go-ipld-channel"
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

func CidsToPairs(cids []cid.Cid, chn channel.Channel) []ChannelCidPair {
	keys := make([]ChannelCidPair, 0, len(cids))
	for _, c := range cids {
		keys = append(keys, ChannelCidPair{chn, c})
	}
	return keys
}

func PairsToCids(keys []ChannelCidPair) []cid.Cid {
	cids := make([]cid.Cid, 0, len(keys))
	for _, key := range keys {
		cids = append(cids, key.Cid)
	}
	return cids
}
