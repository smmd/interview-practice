package main

import (
	"fmt"
	"github.com/smmd/interview-practice/jumping-numbers/solution"
)

func main()  {
	var n int
	var number int

	fmt.Print("Enter the positive integer N: ")

	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		panic(err)
	}

	fmt.Print("Enter the max number: ")

	_, err = fmt.Scanf("%d", &number)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The Jumping Numberes are: %#v", solution.JumpingNumbers(n, number))
}
