package main

import (
	"fmt"
	"math"
	"os"
)

//1. Доработать калькулятор из методички: больше операций и валидации данных (проверка деления на 0,
//возведение в степень, факториал) + форматирование строки при выводе дробного числа ( 2 знака после точки)

func main() {
	var a, b, res float64
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, !): ")
	fmt.Scanln(&op)
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Обнаружено деление на 0")
			os.Exit(1)
		}
		res = a / b
	case "^":
		res = math.Pow(a, b)
	case "!":
		fmt.Println("Факториал посчитан от первого числа!")
		res = factorial(a)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}
func factorial(n float64) float64 {
	var fact float64 = 1
	for i := 1.0; i <= n; i++ {
		fact = fact * i
	}
	return fact
}
