func canCompleteCircuit(gas []int, cost []int) int {
	remains, debts, start := 0, 0, 0

	for i, g := range gas {
		remains += g - cost[i]
		if remains < 0 {
			start = i + 1
			debts += remains
			remains = 0
		}
	}

	if debts+remains < 0 {
		return -1
	}
	return start
}
