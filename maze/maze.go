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
	// 初始化稍后要走的路
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	
	// 开始探险
	q := []Point{start} // q = {0, 0}
	for len(q) > 0 {
		cur := q[0] // 获取起点的坐标
		q = q[1:] // 删除当前的位置

		fmt.Println("cur", cur)
		fmt.Println("q", q)

		for _, dir := range dirs {
			next := cur.add(dir) // 寻找下一个坐标的正确位置
			// 是否越界
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			// 是否走过
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			// 回到起点
			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			q = append(q, next)
		}
	}
	return steps
}

func main() {
	var mazeMap [][]int = make([][]int, 6)
	for i  := range mazeMap {
        mazeMap[i] = make([]int, 5)
    }
	mazeMap[0][1] = 1
	mazeMap[1][3] = 1
	mazeMap[2][1] = 1
	mazeMap[2][3] = 1
	mazeMap[3][0] = 1
	mazeMap[3][1] = 1
	mazeMap[3][2] = 1
	mazeMap[4][1] = 1
	mazeMap[4][4] = 1
	mazeMap[5][1] = 1
	for _, v := range mazeMap {
		for _, v1 := range v {
			fmt.Print(v1, " ")
		}
		fmt.Println()
	}

	data := walk(mazeMap, Point{0, 0}, Point{len(mazeMap) - 1, len(mazeMap) - 1})
	for _, row := range data {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}