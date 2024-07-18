package hashing

import (
	"container/list"
	"fmt"
	"strings"
)

type HashRing struct {
	hashFunc func(string) string
	Servers  *list.List
}

func NewHashRing(hashFunc func(string) string) *HashRing {
	return &HashRing{
		hashFunc: hashFunc,
		Servers:  list.New(),
	}
}

type Server struct {
	ID   string
	keys []string
	hash string
}

func (s Server) String() string {
	return fmt.Sprintf("Server[%s][%s]", s.ID, strings.Join(s.keys, ", "))
}

type Key struct {
	hash string
}

func (hr *HashRing) store(key string) {
	inserted := false

	for e := hr.Servers.Front(); e != nil; e = e.Next() {
		server := e.Value.(*Server)
		if key < server.hash {
			server.keys = append(server.keys, key)
			inserted = true
			break
		}
	}

	// Server with lower hash wasn't found. Add as the last one
	if !inserted {
		server := hr.Servers.Back().Value.(Server)
		server.keys = append(server.keys, key)
	}
}

func (hr *HashRing) findServerByKey(key string) *Server {
	for e := hr.Servers.Front(); e != nil; e = e.Next() {
		server := e.Value.(*Server)
		if key < server.hash {
			for _, keyInServer := range server.keys {
				if key == keyInServer {
					return server
				}
			}
		}
	}

	return nil
}

func (hr *HashRing) addServer(id string) { // ip net.IPAddr) {
	server := Server{
		ID:   id,
		keys: make([]string, 0),
		hash: hr.hashFunc(id),
	}

	if hr.Servers.Len() == 0 {
		hr.Servers.PushFront(&server)
		return
	}

	inserted := false

	for e := hr.Servers.Front(); e != nil; e = e.Next() {
		if server.hash < e.Value.(*Server).hash {
			hr.Servers.InsertBefore(&server, e)
			inserted = true
			break
		}
	}

	// Server with lower hash wasn't found. Add as the last one
	if !inserted {
		hr.Servers.InsertAfter(&server, hr.Servers.Back())
	}
}

func (hr *HashRing) ListServers() {
	for e := hr.Servers.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*Server))
	}
}
