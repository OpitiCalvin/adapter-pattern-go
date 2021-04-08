package main

import "fmt"

type oco2SatelliteAdapter struct {
	oco2Satellite *oco2Satellite
}

func (oco2sa *oco2SatelliteAdapter) insertSatelliteIntoStarlinkPort(f9r *falcon9Rocket) {
	fmt.Println("Satellite adapter converts Starlink port to OCO2 port.")
	oco2sa.oco2Satellite.insertSatelliteIntoStarlinkPort(f9r)
}
