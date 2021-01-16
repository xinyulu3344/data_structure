package ArrayList

import "fmt"

// 稀疏数组的值节点
type ValNode struct {
    row int
    col int
    val int
}

func main() {

    // 1. 先创建一个原始数组
    var chessMap [11][11]int
    chessMap[1][2] = 1 // 黑子
    chessMap[2][3] = 2 // 白子

    // 2. 输出原始数组
    fmt.Println("输出原始数组: ")
    for _, v := range chessMap {
        for _, v2 := range v {
            fmt.Printf("%d\t", v2)
        }
        fmt.Println()
    }

    var sparseArr []ValNode

    valNode := ValNode{
        row: 11,
        col: 11,
        val: 0,
    }
    sparseArr = append(sparseArr, valNode)
    for i, v := range chessMap {
        for j, v2 := range v {
            if v2 != 0 {
                // 创建一个ValNode
                valNode := ValNode{
                    row: i,
                    col: j,
                    val: v2,
                }
                sparseArr = append(sparseArr, valNode)
            }
        }
    }

    // 输出稀疏数组
    fmt.Println("输出稀疏数组: ")
    for i, valNode := range sparseArr {
        fmt.Printf("%d: %d\t%d\t%d\n", i, valNode.row, valNode.col, valNode.val)
    }
}
