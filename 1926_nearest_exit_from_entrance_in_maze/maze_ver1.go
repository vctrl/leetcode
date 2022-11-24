const (
    Up int = iota
    Down
    Left
    Right
)


func nearestExit(maze [][]byte, entrance []int) int {
    stepCountMatrix := make([][]int, 0, len(maze))
    
    //exits := make([][]int, 0, 5)

    exits := make(map[int]map[int]struct{})
    for i := 0; i < len(maze); i++ {
        stepCountRow := make([]int, len(maze[0]), len(maze[0]))
        stepCountMatrix = append(stepCountMatrix, stepCountRow)
        for j := 0; j < len(maze[0]); j++ {
            stepCountMatrix[i][j] = math.MaxInt32
            if i == 0 || i == len(maze)-1 ||  j == 0 ||  j == len(maze[0]) - 1 && maze[i][j] == '.' {
                if entrance[0] == i && entrance[1] == j {
                    continue
                }

                //exits = append(exits, []int{i, j})
                _, ok := exits[i]
                if !ok {
                    exits[i] = make(map[int]struct{})
                }
                exits[i][j] = struct{}{}
            }
        }
    }

    step := 0
    r, c := entrance[0], entrance[1]
    stepCountMatrix[r][c] = step
    //(stepCountMatrix)

    if len(maze) == 0 {
        return 0
    }

    h, w := len(maze), len(maze[0])
    //fmt.Printf("h = %d, w = %d, r = %d, c = %d\n", h, w, r, c)
    res := -1
    if c - 1 >= 0 && maze[r][c - 1] != '+' {
        result := nearestExitVariant(maze, []int{r, c - 1},exits, step, Left, stepCountMatrix)
        if result != -1 {
            if res == -1 {
                res = result
            } else if result < res {
                res = result
            }
        }
    }
    if c + 1 < w && maze[r][c + 1] != '+' {
         result := nearestExitVariant(maze, []int{r, c + 1},exits, step, Right, stepCountMatrix)
        if result != -1 {
            if res == -1 {
                res = result
            } else if result < res {
                res = result
            }
        }
    }
    if r - 1 >= 0 && maze[r - 1][c] != '+' {
        result := nearestExitVariant(maze, []int{r - 1, c},exits, step, Up, stepCountMatrix)
        if result != -1 {
            if res == -1 {
                res = result
            } else if result < res {
                res = result
            }
        }
    }
    if r + 1 < h && maze[r + 1][c] != '+' {
        result := nearestExitVariant(maze, []int{r + 1, c}, exits, step, Down, stepCountMatrix)
        if result != -1 {
            if res == -1 {
                res = result
            } else if result < res {
                res = result
            }
        }
    }

    return res
}

func nearestExitVariant(maze [][]byte, entrance []int, exits map[int]map[int]struct{}, step, direction int, stepCountMatrix [][]int) int {
    //printMaze(maze, entrance)
    //fmt.Println()
    step++
    
    r, c := entrance[0], entrance[1]
    if _, ok := exits[r][c]; ok {
        return step
    }

    if step > stepCountMatrix[r][c] {
        return -1
    }

    stepCountMatrix[r][c] = step

    //printMatrix(stepCountMatrix)
    if len(maze) == 0 {
        return 0
    }

    h, w := len(maze), len(maze[0])
    //fmt.Printf("h = %d, w = %d, r = %d, c = %d\n", h, w, r, c)
    //fmt.Println()
    res := -1
    if direction != Right && c - 1 >= 0 && maze[r][c - 1] != '+' {
        result := nearestExitVariant(maze, []int{r, c - 1},exits, step, Left, stepCountMatrix)
        if result != -1 {
            if res == -1 {
                res = result
            } else if result < res {
                res = result
            }
        }
    }
    if direction != Left && c + 1 < w && maze[r][c + 1] != '+'  {
        result := nearestExitVariant(maze, []int{r, c + 1},exits, step, Right, stepCountMatrix)
        //("result is %d\n", result)
        if result != -1 {
            if res == -1 {
                res = result
            } else if result < res {
                res = result
            }
        }
    }
    if direction != Down && r - 1 >= 0 && maze[r - 1][c] != '+' {
        result := nearestExitVariant(maze, []int{r - 1, c},exits, step, Up, stepCountMatrix)
        if result != -1 {
            if res == -1 {
                res = result
            } else if result < res {
                res = result
            }
        }
    }
    if direction != Up && r + 1 < h && maze[r + 1][c] != '+' {
        result := nearestExitVariant(maze, []int{r + 1, c}, exits, step, Down, stepCountMatrix)
        if result != -1 {
            if res == -1 {
                res = result
            } else if result < res {
                res = result
            }
        }
    }

    return res
}

func printMaze(maze [][]byte, entrance []int) {
    for i := 0; i < len(maze); i++ {
        for j := 0; j < len(maze[0]); j++ {
            if i == entrance[0] && j == entrance[1] {
                fmt.Printf("%c ", 'X')
                continue
            }
            fmt.Printf("%c ", maze[i][j])
        }
        fmt.Println()
    }
}

func printMatrix(m [][]int) {
    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m); j++ {
            fmt.Printf("%d ", m[i][j])
        }
        fmt.Println()
    }
    fmt.Println()
}