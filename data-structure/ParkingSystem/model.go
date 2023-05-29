package ParkingSystem

const (
	BIG = iota
	MEDIUM
	SMALL
)

type ParkingSystem struct {
	space []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	space := make([]int, 3)
	space[BIG] = big
	space[MEDIUM] = medium
	space[SMALL] = small

	return ParkingSystem{space: space}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case BIG - 1:
		if this.space[BIG] <= 0 {
			return false
		} else {
			this.space[BIG] -= 1
		}
	case MEDIUM - 1:
		if this.space[MEDIUM] <= 0 {
			return false
		} else {
			this.space[MEDIUM] -= 1
		}
	case SMALL - 1:
		if this.space[SMALL] <= 0 {
			return false
		} else {
			this.space[SMALL] -= 1
		}
	}
	return true
}
