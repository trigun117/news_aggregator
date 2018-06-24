package fetch

import (
	"testing"
)

func TestFetchFreshNews(t *testing.T) {
	if err := FreshNews(); err != nil {
		t.Fail()
	}
}
