package design

type ParkingSystem struct {
	empty []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{empty: []int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.empty[carType-1] > 0 {
		this.empty[carType-1]--
		return true
	}

	// Else, return False
	return false
}
