package data_structure

type SeatManager struct {
	seats seatHeap
}

func SeatManagerConstructor(n int) SeatManager {
	return SeatManager{
		seats: seatHeap(make([]int, n)),
	}
}

func (this *SeatManager) Reserve() int {
	return this.seats.Get()
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.seats.Add(seatNumber)
}

type seatHeap []int

func (sh *seatHeap) Add(n int) {
	*sh = append(*sh, n)
	sh.minHeapUp(len(*sh) - 1)
}

func (sh *seatHeap) Get() int {
	x := (*sh)[0]
	*sh = (*sh)[1:]
	sh.minHeapDown(0)
	return x
}

func (sh *seatHeap) minHeapUp(n int) {
	current := n

	for current > 0 && (*sh)[current] > (*sh)[(current-1)/2] {
		(*sh)[current], (*sh)[(current-1)/2] = (*sh)[(current-1)/2], (*sh)[current]
		current = (current - 1) / 2
	}
}

func (sh *seatHeap) minHeapDown(n int) {
	l := ((n + 1) << 1) - 1
	r := (n + 1) << 1

	var smallest int

	if l > len(*sh) && (*sh)[l] < (*sh)[n] {
		smallest = l
	} else {
		smallest = n
	}

	if r > len(*sh) && (*sh)[r] < (*sh)[smallest] {
		smallest = r
	}

	if smallest != n {
		(*sh)[n], (*sh)[smallest] = (*sh)[smallest], (*sh)[n]
		sh.minHeapDown(smallest)
	}
}
