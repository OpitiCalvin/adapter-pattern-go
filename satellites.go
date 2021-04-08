package main

import "fmt"

type satellite interface {
	insertSatelliteIntoStarlinkPort(f9r *falcon9Rocket)
}

type starlinkSatellite struct {
	name string
}

type oco2Satellite struct {
	name string
}

func newStarlinkSatellite() *starlinkSatellite {
	return &starlinkSatellite{
		name: "Starlink Satellite",
	}
}

func newOco2Satellite() *oco2Satellite {
	return &oco2Satellite{
		name: "OCO2 Satellite",
	}
}

func (sls *starlinkSatellite) insertSatelliteIntoStarlinkPort(f9r *falcon9Rocket) {
	f9r.payload = append(f9r.payload, sls)
	fmt.Println("Starlink satellite is attached to Falcon 9 Rocket")
}

func (oco2s *oco2Satellite) insertSatelliteIntoStarlinkPort(f9r *falcon9Rocket) {
	f9r.payload = append(f9r.payload, oco2s)
	fmt.Println("OCO2 satellite is attached to Falcon 9 Rocket")
}
