package main

import "fmt"

func main() {
	consumption := 13
	volume := 45
	result :=toplivo(consumption,volume)
	fmt.Println(result,"km")
}
func toplivo(fuelCons,fuelVol int)  int{
	const km  = 100
	total := 0

	total = fuelVol*km/fuelCons
	return total
}
