package generators

import (
	"testing"

	"github.com/khenidak/ip-thing/pkg/generators/types"
)

func TestTTLGen(t *testing.T) {

	resultChan := make(chan types.Result, 255)

	ttlGen, err := NewDefaultTTLGenerator(func(r types.Result, ttl uint8) {
		hdr := r.(*IPv4Header)
		hdr.TTL = ttl
	})
	if err != nil {
		t.Fatalf("failed to create TTL gen:%v", err)
	}

	if err := ttlGen.Generate(NewIPv4Header(), resultChan); err != nil {
		t.Fatalf("failed to Generate with err:%v", err)
	}

	close(resultChan)
	if len(resultChan) != 255 {
		t.Fatalf("expected 50 results, got %v", len(resultChan))
	}

	all := make([]bool, 255)
	for rr := range resultChan {
		hdr := rr.(*IPv4Header)
		if hdr.TTL == 0 {
			// should never be zero
			t.Fatalf("TTL can not be zero")
		}
		all[hdr.TTL-1] = true // align 1..255 values to 255 element slice
	}

	for idx, octetSet := range all {
		if !octetSet {
			t.Fatalf("expected %v: in the result but was not there", idx+1)
		}
	}
}

// since it wraps octet, we don't need to test chaining, since it is done elsewhere
