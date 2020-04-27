package prototype

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrototype(t *testing.T) {
	h1 := make(Header, 0)

	h1.Set("key1", "value1")
	h1.Set("key2", "value2")
	h1.Set("key3", "value3")
	h1.Set("key4", "value4")

	h2 := h1.Clone()

	diff := cmp.Diff(h1, h2)
	if diff != "" {
		t.Errorf(diff)
	}

	h1.Set("key2", "value2-changed")
	h1k := h1.Get("key2")
	if h1k != "value2-changed" {
		t.Errorf("h1 set key2 = \"value2-changed\" failed, get \"%s\".", h1k)
	}

	h2k := h2.Get("key2")
	if h2k != "value2" {
		t.Errorf("h2 get key2 \"%s\".", h2k)
	}

}
