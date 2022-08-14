package main

import (
	"fmt"
	"strconv"
	"strings"
)

//1. Познакомьтесь с алгоритмом сортировки вставками. Напишите приложение, которое принимает на вход набор целых
//чисел и отдаёт на выходе его же в отсортированном виде. Сохраните код, он понадобится нам в дальнейших уроках.
func main() {
	var b string
	fmt.Println("Введите целые числа через запятую")
	fmt.Scan(&b)
	var slice []int
	for _, str := range strings.Split(b, ",") {
		intt, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("%s не число:%v\n", str, err)
			return
		}
		slice = append(slice, intt)
	}
	sortSlice(slice)
	fmt.Printf("Ваша последовательнотсть чисел после сортировки: %v\n", slice)
}

func sortSlice(array []int) {
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}
}
