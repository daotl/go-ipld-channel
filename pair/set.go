// Modified from: https://github.com/ipfs/go-cid/blob/979bf3fb8572224c2b2fbfaf153f94b98734807c/set.go

package pair

import (
	"github.com/ipfs/go-cid"
)

// Set is a implementation of a set of ChannelCidChannelPairPair, that is, a structure
// to which holds a single copy of every CidChannelPairs that is added to it.
type Set struct {
	set map[CidChannelPair]struct{}
}

// NewSet initializes and returns a new Set.
func NewSet() *Set {
	return &Set{set: make(map[CidChannelPair]struct{})}
}

// NewSetFromPairs initializes and returns a new Set from a slice of CidChannelPairs.
func NewSetFromPairs(pairs []CidChannelPair) *Set {
	set := NewSet()
	for _, p := range pairs {
		set.Add(p)
	}
	return set
}

// Add puts a CidChannelPair in the Set.
func (s *Set) Add(c CidChannelPair) {
	s.set[c] = struct{}{}
}

// Has returns if the Set contains a given CidChannelPair.
func (s *Set) Has(c CidChannelPair) bool {
	_, ok := s.set[c]
	return ok
}

// Remove deletes a CidChannelPair from the Set.
func (s *Set) Remove(c CidChannelPair) {
	delete(s.set, c)
}

// Len returns how many elements the Set has.
func (s *Set) Len() int {
	return len(s.set)
}

// Keys returns the CidChannelPairs in the set.
func (s *Set) Keys() []CidChannelPair {
	out := make([]CidChannelPair, 0, len(s.set))
	for k := range s.set {
		out = append(out, k)
	}
	return out
}

// Cids returns all the unique Cids in the CidChannelPairs in the set.
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

// Visit adds a CidChannelPair to the set only if it is
// not in it already.
func (s *Set) Visit(c CidChannelPair) bool {
	if !s.Has(c) {
		s.Add(c)
		return true
	}

	return false
}

// ForEach allows to run a custom function on each
// CidChannelPair in the set.
func (s *Set) ForEach(f func(c CidChannelPair) error) error {
	for c := range s.set {
		err := f(c)
		if err != nil {
			return err
		}
	}
	return nil
}
