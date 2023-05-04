package greedy

func PredictPartyVictory(senate string) string {
	return predictPartyVictory(senate)
}

func predictPartyVictory(senate string) string {
	n := len(senate)

	queueR := make([]int, 0)
	queueD := make([]int, 0)

	for i := range senate {
		if senate[i] == 'R' {
			queueR = append(queueR, i)
		} else {
			queueD = append(queueD, i)
		}
	}

	for len(queueR) != 0 && len(queueD) != 0 {
		r := queueR[0]
		queueR = queueR[1:]
		d := queueD[0]
		queueD = queueD[1:]

		if r < d {
			queueR = append(queueR, r+n)
		} else {
			queueD = append(queueD, d+n)
		}
	}

	if len(queueR) == 0 {
		return "Dire"
	}

	return "Radiant"
}
