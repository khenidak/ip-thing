package generators

import (
	"fmt"

	"github.com/khenidak/ip-thing/pkg/generators/types"
)

type octetGenerator struct {
	*types.CommonGenerator
	min      uint8
	max      uint8
	resultfn func(types.Result, uint8)
}

func NewOctetGenerator(min, max uint8, resultfn func(types.Result, uint8)) (types.Generator, error) {
	if min >= max {
		return nil, fmt.Errorf("min %v must be less than max %v ", min, max)
	}

	octetG := &octetGenerator{}

	octetG.CommonGenerator = &types.CommonGenerator{}
	octetG.min = min
	octetG.max = max
	octetG.resultfn = resultfn

	return octetG, nil
}

func (o *octetGenerator) Generate(r types.Result, results chan<- types.Result) error {
	if err := o.IsValid(); err != nil {
		return err
	}

	count := int(o.min)
	for {
		if count > int(o.max) {
			break
		}

		// write result
		o.resultfn(r, uint8(count))

		if child := o.ChainedChild(); child != nil {

			// if there is a child set, call it
			if err := child.Generate(r.Clone(), results); err != nil {
				return err
			}
		} else {
			// no child, complete this result
			results <- r.Clone()
		}

		count = count + 1
	}

	return nil
}
