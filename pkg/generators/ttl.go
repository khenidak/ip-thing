package generators

import (
	"github.com/khenidak/ip-thing/pkg/generators/types"
)

// ttl is same size as octet so wrap octet generator
func NewDefaultTTLGenerator(resultfn func(r types.Result, ttl uint8)) (types.Generator, error) {
	return NewTTLGenerator(1, 255, resultfn)
}

func NewTTLGenerator(min, max uint8, resultfn func(r types.Result, ttl uint8)) (types.Generator, error) {
	return NewOctetGenerator(min, max, resultfn)
}
