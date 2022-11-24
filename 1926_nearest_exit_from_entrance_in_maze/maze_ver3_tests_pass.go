func nearestExit(maze [][]byte, entrance []int) int {
    exits := make(map[int]map[int]struct{})

    if len(maze) == 0 {
        return -1
    }
    for i := 0; i < len(maze); i++ {
        for j := 0; j < len(maze[0]); j++ {
            if (i == 0 || i == len(maze) - 1 || j == 0 || j == len(maze[0]) - 1) && 
                (i != entrance[0] || j != entrance[1]) && maze[i][j] == '.' { // entrance cannot be exit
                if _, ok := exits[i]; !ok {
                    exits[i] = make(map[int]struct{})
                }

                exits[i][j] = struct{}{}
            }
        }
    }

    nodes := [][]int{entrance}
    y, x := entrance[0], entrance[1]
    maze[y][x] = 'x'
    step := 0

    for len(nodes) > 0 {
        step++
        count := len(nodes)
        for _, node := range nodes {
            y, x := node[0], node[1]
            if y - 1 >= 0 && maze[y - 1][x] == '.' {
                if isExit(y - 1, x, exits) {
                    return step
                }
                maze[y - 1][x] = 'x'
                nodes = append(nodes, []int{y - 1, x})
            }
            if y + 1 < len(maze) && maze[y + 1][x] == '.' {
                if isExit(y + 1, x, exits) {
                    return step
                }
                maze[y + 1][x] = 'x'
                nodes = append(nodes, []int{y + 1, x})
            }
            if x - 1 >= 0 && maze[y][x - 1] == '.' {
                if isExit(y, x - 1, exits) {
                    return step
                }
                maze[y][x - 1] = 'x'
                nodes = append(nodes, []int{y, x - 1})
            }
            if x + 1 < len(maze[0]) && maze[y][x + 1] == '.' {
                if isExit(y, x + 1, exits) {
                    return step
                }
                maze[y][x + 1] = 'x'
                nodes = append(nodes, []int{y, x + 1})
            }
        }

        nodes = nodes[count:]
    }
    

    return -1
}

func isExit(y, x int, exits map[int]map[int]struct{}) bool {
    if exitY, ok := exits[y]; ok {
        if _, ok2 := exitY[x]; ok2 {
            return true
        }
    }

    return false
}