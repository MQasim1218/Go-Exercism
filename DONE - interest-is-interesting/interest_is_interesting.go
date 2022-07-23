package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) (ir float32) {
	if balance < float64(0.000) {
		ir = 3.213
	} else if balance < float64(1000) {
		ir = 0.5
	} else if balance < float64(5000) {
		ir = 1.621
	} else {
		ir = 2.475
	}
	return
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return (balance * float64(InterestRate(balance))) / 100.0
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	for ; balance <= targetBalance; years++ {
		balance = AnnualBalanceUpdate(balance)
	}
	return years
}
