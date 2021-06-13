// Modified from: https://github.com/ipfs/go-cid/blob/979bf3fb8572224c2b2fbfaf153f94b98734807c/set.go

package pair

import (
	"github.com/ipfs/go-cid"
)

// Set is a implementation of a set of ChannelChannelCidPairPair, that is, a structure
// to which holds a single copy of every ChannelCidPairs that is added to it.
type Set struct {
	set map[ChannelCidPair]struct{}
}

// NewSet initializes and returns a new Set.
func NewSet() *Set {
	return &Set{set: make(map[ChannelCidPair]struct{})}
}

// NewSetFromPairs initializes and returns a new Set from a slice of ChannelCidPairs.
func NewSetFromPairs(pairs []ChannelCidPair) *Set {
	set := NewSet()
	for _, p := range pairs {
		set.Add(p)
	}
	return set
}

// Add puts a ChannelCidPair in the Set.
func (s *Set) Add(c ChannelCidPair) {
	s.set[c] = struct{}{}
}

// Has returns if the Set contains a given ChannelCidPair.
func (s *Set) Has(c ChannelCidPair) bool {
	_, ok := s.set[c]
	return ok
}

// Remove deletes a ChannelCidPair from the Set.
func (s *Set) Remove(c ChannelCidPair) {
	delete(s.set, c)
}

// Len returns how many elements the Set has.
func (s *Set) Len() int {
	return len(s.set)
}

// Keys returns the ChannelCidPairs in the set.
func (s *Set) Keys() []ChannelCidPair {
	out := make([]ChannelCidPair, 0, len(s.set))
	for k := range s.set {
		out = append(out, k)
	}
	return out
}

// Cids returns all the unique Cids in the ChannelCidPairs in the set.
func (s *Set) Cids() []cid.Cid {
	cidSet := make(map[cid.Cid]struct{})
	for k := range s.set {
		cidSet[k.Cid] = struct{}{}
	}

	out := make([]cid.Cid, 0, len(cidSet))
	for k := range cidSet {
		out = append(out, k)
	}
	return out
}

// Visit adds a ChannelCidPair to the set only if it is
// not in it already.
func (s *Set) Visit(c ChannelCidPair) bool {
	if !s.Has(c) {
		s.Add(c)
		return true
	}

	return false
}

// ForEach allows to run a custom function on each
// ChannelCidPair in the set.
func (s *Set) ForEach(f func(c ChannelCidPair) error) error {
	for c := range s.set {
		err := f(c)
		if err != nil {
			return err
		}
	}
	return nil
}
