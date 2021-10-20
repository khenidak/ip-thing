package generators

import (
	"testing"

	"github.com/khenidak/ip-thing/pkg/generators/types"
)

func TestIhlGen(t *testing.T) {

	resultChan := make(chan types.Result, 1)

	ihl, err := NewIhlGenerator(func(r types.Result, ihl uint8) {
		hdr := r.(*IPv4Header)
		hdr.Ihl = ihl
	})
	if err != nil {
		t.Fatalf("failed to create ihl gen:%v", err)
	}

	if err := ihl.Generate(NewIPv4Header(), resultChan); err != nil {
		t.Fatalf("failed to Generate with err:%v", err)
	}

	close(resultChan)
	if len(resultChan) != 1 {
		t.Fatalf("expected 1 results, got %v", len(resultChan))
	}

	for rr := range resultChan {
		hdr := rr.(*IPv4Header)
		if hdr.Ihl != 20 {
			t.Fatalf("expected ihl==20 got %v", hdr.Ihl)
		}
	}
}

func TestIhlGenChained(t *testing.T) {
	resultChan := make(chan types.Result, 10)

	firstGen, err := NewIhlGenerator(func(r types.Result, ihl uint8) {
		hdr := r.(*IPv4Header)
		hdr.Ihl = ihl
	})
	if err != nil {
		t.Fatalf("failed to create ihl gen:%v", err)
	}

	secondGen, err := NewOctetGenerator(0, 9, func(r types.Result, octet uint8) {
		hdr := r.(*IPv4Header)
		hdr.SrcIPOctet_0 = octet
	})
	if err != nil {
		t.Fatalf("failed to create octet gen:%v", err)
	}

	// chain them
	firstGen.Chain(secondGen)

	if err := firstGen.Generate(NewIPv4Header(), resultChan); err != nil {
		t.Fatalf("failed to Generate with err:%v", err)
	}

	close(resultChan)

	if len(resultChan) != 10 /*because ihl stamps the same value, we only get parent 10 */ {
		t.Fatalf("expected 10 results, got %v", len(resultChan))
	}

	// remove by key
	for rr := range resultChan {
		hdr := rr.(*IPv4Header)
		// we are not testing octet, only ihl
		if hdr.Ihl != 20 {
			t.Fatalf("expected ihl == 20 got %v", hdr.Ihl)
		}
	}
}
