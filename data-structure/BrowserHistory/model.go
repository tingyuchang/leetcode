package BrowserHistory

type BrowserHistory struct {
	history []string
	current int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{[]string{homepage}, 0}
}

func (this *BrowserHistory) Visit(url string) {
	if this.current == len(this.history)-1 {
		this.history = append(this.history, url)
		this.current++
	} else {
		this.current++
		this.history[this.current] = url
		this.history = this.history[:this.current+1]
	}

}

func (this *BrowserHistory) Back(steps int) string {
	if this.current-steps <= 0 {
		this.current = 0
	} else {
		this.current -= steps
	}

	return this.history[this.current]
}

func (this *BrowserHistory) Forward(steps int) string {
	if (steps + this.current) < len(this.history)-1 {
		this.current += steps
	} else {
		this.current = len(this.history) - 1
	}
	return this.history[this.current]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
