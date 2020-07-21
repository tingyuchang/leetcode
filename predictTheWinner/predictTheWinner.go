package predictTheWinner

func PredictTheWinner(nums []int) bool {
	return winner(nums, 0, len(nums)-1, 1) >=0
}

func winner(si []int, start, end, turn int ) int {
	if start == end {
		return turn * si[start]
	}

	a := turn * si[start] + winner(si, start+1, end, -turn)
	b := turn * si[end] + winner(si, start, end-1, -turn)

	return turn * max(turn*a, turn*b)
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}