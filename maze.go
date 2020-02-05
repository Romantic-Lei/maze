package main
import (
	"fmt"
)

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

	for _, v := range mazeMap {
		for _, v1 := range v {
			fmt.Print(v1)
		}
		fmt.Println()
	}

	
}