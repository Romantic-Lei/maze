package main
import (
	"fmt"
)

// 编写一个函数，完成找到路径
// mazeMap *[8][7]int 地图，保证是同一个地图，所以使用引用
// i, j 表示对地图的哪个点进行测试
func SetWay(mazeMap *[8][7]int, i int, j int) bool {
	// 分析出什么情况下找到出路
	// mazeMap[6][5] == 2
	if mazeMap[6][5] == 2 {
		return true
	} else {
		// 说明要继续找
		if mazeMap[i][j] == 0 {
			// 如果当前点没有走过，我们在进行尝试,寻找策略下右上左
			// 假设这个点是通的
			mazeMap[i][j] = 2
			if SetWay(mazeMap, i + 1, j) {
				// 向下探路
				return true
			} else if SetWay(mazeMap, i, j + 1) {
				// 向右探路
				return true
			} else if SetWay(mazeMap, i - 1, j) {
				// 向上探路
				return true
			} else if SetWay(mazeMap, i, j - 1) {
				// 向左探路
				return true
			} else {
				// 路全部不通，死路一条
				mazeMap[i][j] = 3
				return false
			}

		} else {
			// 说明这个点不能探测，是一堵墙
			return false
		}
	}
}

func main() {
	// 先创建一个二维数组模拟一个迷宫
	// 1. 如果元素的值为0， 则这里就是没有走过的点
	// 2. 如果元素的值为1， 则就是一堵墙
	// 3. 如果元素的值为2， 则这是一条通路
	// 4. 如果元素的值为3， 则是走过的点，但是走不通
	var mazeMap [8][7]int

	// 先把墙的边框设计出来
	for i := 0; i < len(mazeMap[0]); i++ {
		mazeMap[0][i] = 1
		mazeMap[7][i] = 1
	}
	for i := 0; i < len(mazeMap); i++ {
		mazeMap[i][0] = 1
		mazeMap[i][6] = 1
	}
	mazeMap[3][1] = 1
	mazeMap[3][2] = 1

	// 输出地图
	for _, v := range mazeMap {
		for _, v1 := range v {
			fmt.Print(v1)
		}
		fmt.Println()
	}

	SetWay(&mazeMap, 1, 1)
	fmt.Println("探测完毕，请验收")
	for _, v := range mazeMap {
		for _, v1 := range v {
			fmt.Print(v1, " ")
		}
		fmt.Println()
	}
	
}