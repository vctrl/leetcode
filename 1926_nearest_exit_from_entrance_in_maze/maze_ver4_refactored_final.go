func nearestExit(maze [][]byte, entrance []int) int {
	if len(maze) == 0 {
		return -1
	}

	nodes := [][]int{entrance}
	maze[entrance[0]][entrance[1]] = 'x'
	step := 0

	for len(nodes) > 0 {
		step++
		count := len(nodes)
		for _, node := range nodes {
			y, x := node[0], node[1]
			for _, next := range [][]int{{y - 1, x}, {y + 1, x}, {y, x - 1}, {y, x + 1}} {
				y2, x2 := next[0], next[1]
				if y2 < len(maze) && y2 >= 0 && x2 < len(maze[0]) && x2 >= 0 && maze[y2][x2] == '.' {
					if y2 == 0 || y2 == len(maze)-1 || x2 == 0 || x2 == len(maze[0])-1 {
						return step
					}
					maze[y2][x2] = 'x'
					nodes = append(nodes, []int{y2, x2})
				}
			}
		}

		nodes = nodes[count:]
	}

	return -1
}