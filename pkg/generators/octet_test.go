package generators

import (
	"testing"

	"github.com/khenidak/ip-thing/pkg/generators/types"
)

func TestOctetGen(t *testing.T) {

	resultChan := make(chan types.Result, 50)

	octetGen, err := NewOctetGenerator(0, 49, func(r types.Result, octet uint8) {
		hdr := r.(*IPv4Header)
		hdr.SrcIPOctet_0 = octet
	})
	if err != nil {
		t.Fatalf("failed to create octet gen:%v", err)
	}

	if err := octetGen.Generate(NewIPv4Header(), resultChan); err != nil {
		t.Fatalf("failed to Generate with err:%v", err)
	}

	close(resultChan)
	if len(resultChan) != 50 {
		t.Fatalf("expected 50 results, got %v", len(resultChan))
	}

	all := make([]bool, 50)
	for rr := range resultChan {
		hdr := rr.(*IPv4Header)
		all[hdr.SrcIPOctet_0] = true
	}

	for idx, octetSet := range all {
		if !octetSet {
			t.Fatalf("expected %v: in the result but was not there", idx)
		}
	}
}

type pesudoKey struct {
	first  uint8
	second uint8
}

func TestOctetGenChained(t *testing.T) {
	resultChan := make(chan types.Result, 10*10)

	firstGen, err := NewOctetGenerator(0, 9, func(r types.Result, octet uint8) {
		hdr := r.(*IPv4Header)
		hdr.SrcIPOctet_0 = octet
	})
	if err != nil {
		t.Fatalf("failed to create octet gen:%v", err)
	}

	secondGen, err := NewOctetGenerator(0, 9, func(r types.Result, octet uint8) {
		hdr := r.(*IPv4Header)
		hdr.SrcIPOctet_1 = octet
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

	if len(resultChan) != 100 {
		t.Fatalf("expected 100 results, got %v", len(resultChan))
	}

	// create expected results
	all := make(map[pesudoKey]struct{})
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			key := pesudoKey{
				first:  uint8(i),
				second: uint8(j),
			}

			all[key] = struct{}{}
		}

	}

	// remove by key
	for rr := range resultChan {
		hdr := rr.(*IPv4Header)
		key := pesudoKey{
			first:  uint8(hdr.SrcIPOctet_0),
			second: uint8(hdr.SrcIPOctet_1),
		}

		delete(all, key)

	}

	if len(all) != 0 {
		t.Fatalf("the values %+v was not in the result, and was expected", all)
	}
}
