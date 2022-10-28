package shippingoptions

import (
	"test-mem-hazmat/externalteam"
	"test-mem-hazmat/gopacking"
)

type ShippingOptions struct {
	Gopacking    gopacking.Gopacking
	ExternalTeam externalteam.ExternalTeam
}

func (s *ShippingOptions) CallCandidates(times, proposta int) {
	s.Gopacking.Candidates(times, proposta)
}

func (s *ShippingOptions) SendAttributes(times, proposta int) {
	s.ExternalTeam.UnpackProposta(times, proposta)

}
