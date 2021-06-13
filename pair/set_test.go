// Modified from: https://github.com/ipfs/go-cid/blob/979bf3fb8572224c2b2fbfaf153f94b98734807c/set_test.go

package pair

import (
	"crypto/rand"
	"errors"
	"testing"

	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
)

func makeRandomChannelCidPair(t *testing.T) ChannelCidPair {
	p := make([]byte, 256)
	_, err := rand.Read(p)
	if err != nil {
		t.Fatal(err)
	}

	h, err := mh.Sum(p, mh.SHA3, 4)
	if err != nil {
		t.Fatal(err)
	}

	cid := cid.NewCidV1(7, h)

	return PublicCidPair(cid)
}

func TestSet(t *testing.T) {
	pair := makeRandomChannelCidPair(t)
	pair2 := makeRandomChannelCidPair(t)
	s := NewSet()

	s.Add(pair)

	if !s.Has(pair) {
		t.Error("should have the ChannelCidPair")
	}

	if s.Len() != 1 {
		t.Error("should report 1 element")
	}

	keys := s.Keys()

	if len(keys) != 1 || !keys[0].Equals(pair) {
		t.Error("key should correspond to ChannelCidPair")
	}

	if s.Visit(pair) {
		t.Error("visit should return false")
	}

	var foreach []ChannelCidPair
	foreachF := func(p ChannelCidPair) error {
		foreach = append(foreach, p)
		return nil
	}

	if err := s.ForEach(foreachF); err != nil {
		t.Error(err)
	}

	if len(foreach) != 1 {
		t.Error("ForEach should have visited 1 element")
	}

	foreachErr := func(p ChannelCidPair) error {
		return errors.New("test")
	}

	if err := s.ForEach(foreachErr); err == nil {
		t.Error("Should have returned an error")
	}

	if !s.Visit(pair2) {
		t.Error("should have visited a new ChannelCidPair")
	}

	if s.Len() != 2 {
		t.Error("len should be 2 now")
	}

	s.Remove(pair2)

	if s.Len() != 1 {
		t.Error("len should be 1 now")
	}
}
