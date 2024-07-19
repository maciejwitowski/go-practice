package urlshortener

import (
	"testing"
)

func TestBasicFlow(t *testing.T) {
	shortener := UrlShortener{
		store: NewStore(),
	}

	input := LongUrl("https://edition.cnn.com/business/live-news/global-outage-intl-hnk/index.html")
	short := shortener.Shorten(input)
	url, _ := shortener.GetUrl(short)

	if url != input {
		t.Errorf("Expected the same url")
	}
}
