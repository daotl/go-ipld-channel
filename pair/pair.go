package pair

import (
	"github.com/ipfs/go-cid"

	channel "github.com/daotl/go-ipld-channel"
)

// CidChannelPair is a pair of IPLD channel and cid.
type CidChannelPair struct {
	Cid     cid.Cid
	Channel channel.Channel
}

// Defined returns true if a CidChannelPair is defined
// Calling any other methods on an undefined CidChannelPair will result in
// undefined behavior.
func (p CidChannelPair) Defined() bool {
	return p.Cid.Defined()
}

// Equals checks that two CidChannelPairs are the same.
func (p CidChannelPair) Equals(o CidChannelPair) bool {
	return p == o
}

// PublicCidPair returns a pair consisting of the public IPLD channel and the specified cid.
func PublicCidPair(c cid.Cid) CidChannelPair {
	return CidChannelPair{
		Cid:     c,
		Channel: channel.PublicChannel,
	}
}

func CidsToPairs(cids []cid.Cid, chn channel.Channel) []CidChannelPair {
	pairs := make([]CidChannelPair, 0, len(cids))
	for _, c := range cids {
		pairs = append(pairs, CidChannelPair{c, chn})
	}
	return pairs
}

func PairsToCids(pairs []CidChannelPair) []cid.Cid {
	cids := make([]cid.Cid, 0, len(pairs))
	for _, key := range pairs {
		cids = append(cids, key.Cid)
	}
	return cids
}
