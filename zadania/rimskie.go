package main

import (
	"bufio"
	"fmt"
	"os"
)

func romanToArabic(romanNum string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	prev := 0
	sum := 0
	for i := len(romanNum) - 1; i >= 0; i-- {
		curr := romanMap[romanNum[i]]
		if curr < prev {
			sum -= curr
		} else {
			sum += curr
		}
		prev = curr
	}
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите римское число: ")
	romanNum, _ := reader.ReadString('\n')
	arabicNum := romanToArabic(romanNum)
	fmt.Printf("%s = %d\n", romanNum, arabicNum)
}