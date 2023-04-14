package main

import (
	"bufio"
	"fmt"
	"os"
)

// acwing给我报错了
func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// 去掉字符串最后的换行符
	str = str[:len(str)-1]

	for _, c := range str {
		fmt.Printf("%c ", c)
	}
}
