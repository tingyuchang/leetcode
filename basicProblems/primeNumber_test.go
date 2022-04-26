package basicProblems

import "testing"

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name        string
		args        args
		wantIsPrime bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsPrime := isPrime(tt.args.n); gotIsPrime != tt.wantIsPrime {
				t.Errorf("isPrime() = %v, want %v", gotIsPrime, tt.wantIsPrime)
			}
		})
	}
}
