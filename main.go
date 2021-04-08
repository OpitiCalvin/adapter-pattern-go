package main

func main() {
	// falcon9Rocket is our client, it expects a Starlink Satellite
	falcon9Rocket := &falcon9Rocket{}
	falcon9Rocket.getPayload()

	// instantiate satellites
	starlinkSatellite := newStarlinkSatellite()
	oco2Satellite := newOco2Satellite()

	// insert the expected Starlink satellite
	falcon9Rocket.insertSatelliteIntoStarlinkPort(starlinkSatellite)

	// allow a falcon9Rocket to accept an oco2Satellite via adapter
	oco2SatelliteAdapter := &oco2SatelliteAdapter{
		oco2Satellite: oco2Satellite,
	}

	// use the adapter to insert an OCO2 satellite
	falcon9Rocket.insertSatelliteIntoStarlinkPort(oco2SatelliteAdapter)
	falcon9Rocket.getPayload()
}
