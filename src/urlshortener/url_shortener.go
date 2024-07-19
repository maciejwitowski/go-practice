package urlshortener

import (
	"encoding/base64"
	"encoding/binary"
)

type UrlShortener struct {
	store *Store
}

type ShortUrl string

type LongUrl string

func (us *UrlShortener) Shorten(url LongUrl) ShortUrl {
	short, exists := us.store.getShort(url)

	if exists {
		return short
	}

	short = us.generateShort()
	us.store.put(url, short)
	return short
}

func (us *UrlShortener) GetUrl(short ShortUrl) (LongUrl, bool) {
	return us.store.getLong(short)
}

func (us *UrlShortener) generateShort() ShortUrl {
	id := us.store.nextID()
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(id))
	return ShortUrl(base64.StdEncoding.EncodeToString(buf))
}

// Store Could be DB
type Store struct {
	idCounter    int64
	shortsLookup map[LongUrl]ShortUrl
	longsLookup  map[ShortUrl]LongUrl
}

func NewStore() *Store {
	return &Store{
		idCounter:    0,
		shortsLookup: make(map[LongUrl]ShortUrl),
		longsLookup:  make(map[ShortUrl]LongUrl),
	}
}

// Could use a bloom filter
func (s *Store) getShort(long LongUrl) (ShortUrl, bool) {
	short, exists := s.shortsLookup[long]
	return short, exists
}

func (s *Store) getLong(short ShortUrl) (LongUrl, bool) {
	long, exists := s.longsLookup[short]
	return long, exists
}

func (s *Store) nextID() int64 {
	return s.idCounter + 1
}

func (s *Store) put(longUrl LongUrl, shortUrl ShortUrl) {
	s.shortsLookup[longUrl] = shortUrl
	s.longsLookup[shortUrl] = longUrl
	s.idCounter = s.idCounter + 1
}
