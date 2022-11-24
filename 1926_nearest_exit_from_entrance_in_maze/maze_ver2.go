const (
    initial int8 = iota
    up
    down
    left
    right
)

func nearestExit(maze [][]byte, entrance []int) int {
    fmt.Println("START maze is", maze, " entrance is ", entrance)
    if len(maze) == 0 {
        return -1
    }

    // find exits
    exits := make(map[int]map[int]struct{})
    for i := 0; i < len(maze); i++ {
        for j := 0; j < len(maze[0]); j++ {
            if i == 0 || i == len(maze) - 1 || j == 0 || j == len(maze[0]) - 1 {
                if _, ok := exits[i]; !ok {
                    exits[i] = make(map[int]struct{})
                }
                exits[i][j] = struct{}{}
            }
        }
    }
    
    directionsMatrix := make([][]int8, len(maze))
    for i := range directionsMatrix {
        directionsMatrix[i] = make([]int8, len(maze[0]))
    }

    directionsMatrix[entrance[0]][entrance[1]] = initial

    // map[x]map[y]stepCount
    nodes := make(map[int]map[int]int)
    nodes[entrance[0]] = map[int]int{entrance[1]: 0}
    return bfs(nodes, maze, directionsMatrix, 0, exits)
}

func bfs(nodes map[int]map[int]int, maze [][]byte, directionsMatrix [][]int8, step int, exits map[int]map[int]struct{}) int {
    step++
    fmt.Println(nodes)
    fmt.Printf("step %d: ", step)
    old := make(map[int]int)
    for x, ySteps := range nodes {
        for y := range ySteps {
            old[x] = y
        }
    }

    for x, ySteps := range nodes {
        //fmt.Printf("%d %d, ", node[0], node[1])
        for y, _ := range ySteps {

            step = move(x, y, x, y + 1, step, down, nodes, directionsMatrix, maze, exits)
            if step != -1 {
                return step
            }
            step = move(x, y, x, y - 1, step, up, nodes, directionsMatrix, maze, exits)
            if step != -1 {
                return step
            }
            step = move(x, y, x - 1, y, step, left, nodes, directionsMatrix, maze, exits)
            if step != -1 {
                return step
            }
            step = move(x, y, x + 1, y, step, right, nodes, directionsMatrix, maze, exits)
            if step != -1 {
                return step
            }
        }
    }
    
    for x, y := range old {
        ySteps := nodes[x]
        delete(ySteps, y)
        if len(ySteps) == 0 {
            delete(nodes, x)
        }
    }

    if len(nodes) == 0 {
        return -1
    }

    return bfs(nodes, maze, directionsMatrix, step, exits)
}

func move(x1, y1, x2, y2, step int, 
    direction int8, 
    nodes map[int]map[int]int, 
    directionsMatrix [][]int8, 
    maze [][]byte, 
    exits map[int]map[int]struct{}) int {
        fmt.Println("direction is ", direction)
        fmt.Println("old is ", x1, y1)
        fmt.Println("new is ", x2, y2)
        
        fmt.Println("exits ", exits)
    if ((direction == left && x2 > 0) || (direction == right && x2 < len(maze)) || (direction == up && y2 > 0) || (direction == down && y2 < len(maze[0]))) && 
        directionsMatrix[x1][y1] != opposite(direction) && 
            maze[x2][y2] == '.' {
        if _, ok := exits[x2][y2]; ok {
            return step
        }
        directionsMatrix[y2][x2] = direction
        // if we come again to the same node we rewrite number of steps with min
        if stepsToNode, ok := nodes[x2][y2]; ok {
            if step > stepsToNode {
                nodes[x2][y2] = stepsToNode
            } else {
                nodes[x2][y2] = step
            }
        } else { // first time here
            nodes[x1] = map[int]int{y2: step}
        }
    }

    return -1
}

func opposite(in int8) int8 {
    switch in {
        case up:
            return down
        case down:
            return up
        case left: 
            return right
        case right:
            return left
    }

    return initial
}