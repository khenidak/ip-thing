package types

import ()

var NoMoreErr error

// each generator must implement this interface. An error of type NoMoreErr
// means the generator can not generate any more for that parent
// any other error is treated as failure in the process and must bubble up
type Generator interface {
	Generate(r Result, results chan<- Result) error
	Chain(g Generator)
}

type CommonGenerator struct {
	child Generator
	name  string
}

func (c *CommonGenerator) Chain(child Generator) {
	c.child = child
}

func (c *CommonGenerator) ChainedChild() Generator {
	return c.child
}

func (c *CommonGenerator) IsValid() error {
	// for when it is needed
	return nil
}
