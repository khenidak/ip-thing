package generators

import (
	"github.com/khenidak/ip-thing/pkg/generators/types"
)

type ihlGenerator struct {
	*types.CommonGenerator
	resultfn func(types.Result, uint8)
}

func NewIhlGenerator(resultfn func(types.Result, uint8 /* ilh is 4 bits */)) (types.Generator, error) {
	ihlG := &ihlGenerator{}
	ihlG.CommonGenerator = &types.CommonGenerator{}
	ihlG.resultfn = resultfn

	return ihlG, nil
}

func (ihl *ihlGenerator) Generate(r types.Result, results chan<- types.Result) error {
	if err := ihl.IsValid(); err != nil {
		return err
	}

	// if you want dynamic option based len
	// then create another generator, chain it as the last
	// one and do dynamic calc
	// this generator assumes no options were set on header
	const val = uint8(20)
	ihl.resultfn(r, val)

	if child := ihl.ChainedChild(); child != nil {
		// if there is a child set, call it
		if err := child.Generate(r.Clone(), results); err != nil {
			return err
		}
	} else {
		// no child, complete this result
		results <- r // no need to clone it since it is only one
	}

	return nil
}
