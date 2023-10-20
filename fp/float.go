package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	b := float32(balance)
	if b < 0 {
		return 3.213
	} else if b >= 1000 && b < 5000 {
		return 1.621
	} else if b >= 5000 {
		return 2.475
	} else {
		return 0.5
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {

	i := 1
	var tem = balance + Interest(balance)
	if tem >= targetBalance {
		return 0
	}
	for true {
		if tem >= targetBalance {
			break
		}
		tem = tem + Interest(tem)
		i++
	}
	return i
}
