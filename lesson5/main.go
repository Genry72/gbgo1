package main

import (
	"fmt"
	"time"
)

//1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
//2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.** (Оно так же должно быть рекурсивным!)**

func main() {
	fmt.Println("Ведите число для расчета числа Фибоначчи")
	var n uint64
	fmt.Scanln(&n)
	t := time.Now()
	resMap := make(map[uint64]uint64, n)
	fmt.Printf("Результат: %s Время расчета оптимизированной функции: %d микросекунд\n", fmt.Sprint(fibonachi2(n, resMap)), time.Since(t).Microseconds())
	t2 := time.Now()
	fmt.Printf("Результат: %s Время расчета обычной функции: %d микросекунд\n", fmt.Sprint(fibonachi1(n)), time.Since(t2).Microseconds())
}

func fibonachi2(n uint64, resMap map[uint64]uint64) uint64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	} else {
		if val, ok := resMap[n]; ok {
			return val
		} else {
			resMap[n] = fibonachi2(n-1, resMap) + fibonachi2(n-2, resMap)
		}
		return resMap[n]
	}
}

func fibonachi1(n uint64) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonachi1(n-1) + fibonachi1(n-2)
	}
}
