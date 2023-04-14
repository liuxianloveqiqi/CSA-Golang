package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	for {
		fmt.Println("Please input your guess")
		var guess int
		// 输入我们猜的数字
		_, err := fmt.Scanf("%d", &guess)
		// Go语言中处理错误的方法
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			return
		}
		fmt.Println("You guess is", guess)

		if guess < 0 || guess > maxNum {
			fmt.Println("范围错误，哥们这数字位于[0,100]")
			continue
		}

		if guess == secretNumber {
			fmt.Println("牛马运气，猜对了")
			break
		} else if guess < secretNumber {
			fmt.Println("猜小了")
		} else {
			fmt.Println("猜大了")
		}
	}
}
