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

	// 上面已保障了从 start 能走到最后一个 gas(remains>0情况下)
	// 只要再保障 debts+remains>=0，则代表 start 能走到 start-1
	// 反证: 如果没有解，绝不会出现 debts+remains=0(假设有2个加油站)
	if debts+remains < 0 {
		return -1
	}
	return start
}
