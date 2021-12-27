package main

type ParkingLot struct {
	LotSize int
}

func ParkCar(lot *ParkingLot, carSize int) {
	lot.LotSize -= carSize
}

func (lot *ParkingLot) ParkCar(carSize int) {
	lot.LotSize -= carSize
}

func main() {
	lot := &ParkingLot{100}
	ParkCar(lot, 10)

	lot.ParkCar(10)
}
