package hashing

import (
	"testing"
)

func TestBasicFlow(t *testing.T) {
	//hashFunc := func(val []byte) string {
	//	hash := sha1.Sum(val)
	//	hex.EncodeToString(hash[:])
	//}

	hashFunc := func(s string) string {
		return s
	}
	ring := NewHashRing(hashFunc)

	// a b c d e f g
	ring.addServer("c")
	ring.addServer("e")
	ring.addServer("g")
	ring.addServer("a")
	ring.addServer("b")

	ring.store("aaa")
	ring.ListServers()

}
