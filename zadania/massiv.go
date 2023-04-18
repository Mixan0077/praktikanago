package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Определяем размеры двумерного массива
	rows := 5
	cols := 10

	// Создаем двумерный массив
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// Заполняем массив уникальными случайными числами
	used := make(map[int]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// Генерируем случайное число
			var num int
			for {
				num = rand.Intn(rows * cols)
				if !used[num] {
					used[num] = true
					break
				}
			}
			matrix[i][j] = num
		}
	}

	// Находим строку с самым большим числом
	maxSum := 0
	maxRow := -1
	for i := 0; i < rows; i++ {
		rowSum := 0
		for j := 0; j < cols; j++ {
			rowSum += matrix[i][j]
		}
		if rowSum > maxSum {
			maxSum = rowSum
			maxRow = i
		}
	}

	// Выводим полученный двумерный массив и строку с самым большим числом
	fmt.Println("Массив:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Printf("Строка с максимальной суммой: %d\n", maxRow)
}