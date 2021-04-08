package main

import "fmt"

type falcon9Rocket struct {
	payload []interface{}
}

func (f9r *falcon9Rocket) insertSatelliteIntoStarlinkPort(s satellite) {
	fmt.Println("Attaching satellite to Falcon 9 Rocket")
	s.insertSatelliteIntoStarlinkPort(f9r)
}

func (f9r *falcon9Rocket) getPayload() {
	fmt.Printf("Falcon 9 payload: %s\n", f9r.payload)
}
