package task_2

import "erew/internal/memory"

/*
 dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - 1], dp[i - 1][j - 1}) + array[i][j]
 We suppose turtle can choose where it starts in first line (from 0 to m)
*/

func GetTurtle(u *memory.Unique, workerId int, table [][]*memory.Erew[int]) {
	var (
		n = len(table)
		m = len(table[0])
	)

	for i := 1; i < n; i++ {
		var aboveLeft, above, aboveRight int
		if workerId >= 1 {
			aboveLeft = table[i-1][workerId-1].Read()
		} else {
			u.Skip(1)
		}

		above = table[i-1][workerId].Read()

		if workerId < m-1 {
			aboveRight = table[i-1][workerId+1].Read()
		} else {
			u.Skip(1)
		}

		// TODO(kasalakhov) think about negative numbers (initialize above variables as -INF)
		value := table[i][workerId].Read()
		result := max(aboveLeft, above, aboveRight) + value
		table[i][workerId].Write(result)
	}
}
