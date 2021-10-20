package generators

import (
	"github.com/khenidak/ip-thing/pkg/generators/types"
)

type IPv4Header struct {
	//
	// generally we break it down to smallest unit
	// allowing generators to partition generation process
	// anyway they like
	//
	// the generators written for this are all using little endian

	// every field must be copied manually in Clone()
	// source IP
	SrcIPOctet_0 uint8
	SrcIPOctet_1 uint8
	SrcIPOctet_2 uint8
	SrcIPOctet_3 uint8

	// destination IP
	DstIPOctet_0 uint8
	DstIPOctet_1 uint8
	DstIPOctet_2 uint8
	DstIPOctet_3 uint8

	Ihl     uint8 // only 4bits used, we have a static generator that can be turned into dynamic calc as needed
	Version uint8 // only 4bits, no generation needed for this. Always fixed value (4)

	TTL uint8
}

func NewIPv4Header() types.Result {
	return &IPv4Header{
		Version: uint8(4),
	}
}
func (iph *IPv4Header) Clone() types.Result {

	// this is a lot of copy, but we need a unique header
	// for each result
	return &IPv4Header{
		SrcIPOctet_0: iph.SrcIPOctet_0,
		SrcIPOctet_1: iph.SrcIPOctet_1,
		SrcIPOctet_2: iph.SrcIPOctet_2,
		SrcIPOctet_3: iph.SrcIPOctet_3,

		DstIPOctet_0: iph.DstIPOctet_0,
		DstIPOctet_1: iph.DstIPOctet_1,
		DstIPOctet_2: iph.DstIPOctet_2,
		DstIPOctet_3: iph.DstIPOctet_3,

		Ihl:     iph.Ihl,
		Version: iph.Version,
		TTL:     iph.TTL,
	}
}
