package main

import (
	"test-mem-hazmat/gopacking"
	"test-mem-hazmat/shippingoptions"
)

func main() {
	shp := shippingoptions.ShippingOptions{
		Gopacking: gopacking.Gopacking{},
	}

	shp.CallCandidates(2500000, 1)
	shp.CallCandidates(2500000, 2)

	shp.SendAttributes(100000, 1)
	shp.SendAttributes(100000, 2)

}
