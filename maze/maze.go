package main
import (
	"fmt"
)

// 定义一个坐标结构体
type Point struct {
	i int
	j int
}

// 声明当前坐标四周的其他坐标
// 移动规则是 右上左下
var dirs = [4]Point{
	{0, 1}, {-1, 0}, {0, -1}, {1, 0},
}

// 下一个节点坐标
func (p Point) add(r Point) Point {
	return Point{p.i + r.i, p.j + r.j}
}

// 当前坐标的值， 验证是否越界，路是否可走
func (p Point) at(grid [][]int) (int, bool) {
	// 验证 x 轴是否越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	// 验证 y 轴是否越界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

// 迷宫探险
// maze [][]int 地图
// start Point 地图中起始点
// end Point 地图中终点
func walk(maze [][]int, start Point, end Point) [][]int {
	// 二位数组切片初始化如下
	// 初始化一个和 maze 一样的切片
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// 显示的时候，会显示将源切片和走过的路径全显示出来
	copy(steps, maze) // 目标切片， 源切片

	// 出口是一堵墙
	if steps[end.i][end.j] == 1 {
		fmt.Println("这是一堵墙，不能作为终点")
		return nil
	}
	
	// 开始探险
	q := []Point{start} // q = {0, 0}
	for len(q) > 0 {
		cur := q[0] // 获取起点的坐标
		q = q[1:] // 删除当前的位置

		for _, dir := range dirs {
			nextPoint := cur.add(dir) // 移动到下一个坐标
			// 是否越界
			val, ok := nextPoint.at(maze)
			if !ok || val == 1 {
				// 数组越界 或者 当前位置是一堵墙
				continue
			}

			// 是否走过
			val, ok = nextPoint.at(steps)
			if !ok || val != 0 {
				continue
			}

			// 回到起点
			if nextPoint == start {
				continue
			}

			curSteps, _ := cur.at(steps) // 获取到当前点的值
			steps[nextPoint.i][nextPoint.j] = curSteps + 1

			q = append(q, nextPoint)
		}

		// 当找到迷宫终点时，我们就没必要继续往下走了，跳出结束
		if steps[end.i][end.j] != 0 {
			break
		}
	}
	return steps
}

func tempFile() {
	var mazeMap [][]int = make([][]int, 6)
	for i  := range mazeMap {
        mazeMap[i] = make([]int, 5)
	}
	// 值为 1， 表示当前位置是一堵墙
	mazeMap[0][1] = 1
	mazeMap[1][3] = 1
	mazeMap[2][1] = 1
	mazeMap[2][3] = 1
	mazeMap[3][0] = 1
	mazeMap[3][1] = 1
	// mazeMap[3][2] = 1
	// mazeMap[3][3] = 1
	// mazeMap[3][4] = 1
	mazeMap[4][1] = 1
	mazeMap[4][4] = 1
	mazeMap[5][1] = 1
	// mazeMap[5][4] = 1
	for _, v := range mazeMap {
		for _, v1 := range v {
			fmt.Print(v1, " ")
		}
		fmt.Println()
	}

	data := walk(mazeMap, Point{0, 0}, Point{3, 2})
	// data := walk(mazeMap, Point{0, 0}, Point{len(mazeMap) - 1, len(mazeMap[0]) - 1})
	if data != nil && data[3][2] == 0 {
	// if data != nil && data[len(data)-1][len(data[len(mazeMap)-1])-1] == 0 {
		fmt.Println("不能抵达终点")
		return 
	}
	for _, row := range data {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

func main() {
	tempFile()
}