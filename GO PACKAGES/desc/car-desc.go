package desc

import "example.com/car/shared"

func GetCarDescription(value int) {
	switch value {
	case 1:
		shared.PrintText("Toyota is Cap")
	case 2:
		shared.PrintText("Benz is kinda Cap")
	case 3:
		shared.PrintText("BMW is the real deal my nigga")
	default:
		shared.PrintText("Pussy ass nigga, check the choices")
	}
}
