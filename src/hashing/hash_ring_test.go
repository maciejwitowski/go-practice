package hashing

import (
	"testing"
)

func TestBasicFlow(t *testing.T) {
	// Commented as for simplicity I used a stubbed hash func which just returns key as is without hashing
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

	ring.store("a1")
	ring.store("a2")
	ring.store("b1")
	ring.store("b2")

	server := ring.findServerByKey("b1")
	if server.ID != "c" {
		t.Errorf("Expected c to contain b1")
	}
}
