package Underground

import "fmt"

type Station struct {
	totalTimeSpent int
	totalPeople    int
	endStations    map[string]*Station
}

func (s *Station) calculateAVGTime() float64 {
	return float64(s.totalTimeSpent) / float64(s.totalPeople)
}

type Passenger struct {
	startStation string
	startTime    int
}

type UndergroundSystem struct {
	// startStation -> endStation
	stations   map[string]*Station
	passengers map[int]*Passenger
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		stations:   map[string]*Station{},
		passengers: map[int]*Passenger{},
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	if _, ok := this.passengers[id]; !ok {

		this.passengers[id] = &Passenger{
			startStation: stationName,
			startTime:    t,
		}
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if _, ok := this.passengers[id]; ok {
		passenger := this.passengers[id]
		timeSpent := t - passenger.startTime
		station, ok2 := this.stations[passenger.startStation]

		if !ok2 {
			endStation := &Station{
				totalTimeSpent: timeSpent,
				totalPeople:    1,
				endStations:    nil,
			}
			this.stations[passenger.startStation] = &Station{
				endStations: map[string]*Station{stationName: endStation},
			}
		} else {
			fmt.Println(station.endStations, "id: ", id, "end: ", stationName)
			endStation, ok3 := station.endStations[stationName]

			if !ok3 {
				endStation = &Station{
					totalTimeSpent: timeSpent,
					totalPeople:    1,
					endStations:    nil,
				}
				station.endStations[stationName] = endStation
			} else {
				endStation.totalPeople = endStation.totalPeople + 1
				endStation.totalTimeSpent = endStation.totalTimeSpent + timeSpent
			}
		}

		delete(this.passengers, id)
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	fmt.Println(this.stations)
	return this.stations[startStation].endStations[endStation].calculateAVGTime()
}
