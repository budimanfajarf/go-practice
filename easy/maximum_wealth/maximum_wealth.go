package maximum_wealth

func MaximumWealth(accounts [][]int) int {
	maxWealth := 0

	for _, balances := range accounts {
		wealth := 0

		for _, balance := range balances {
			wealth += balance
		}

		if wealth >= maxWealth {
			maxWealth = wealth
		}
	}

	return maxWealth
}
