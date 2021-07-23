package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Palindrome
	palA := "Sandal"
	palB := palindrome(palA)
	fmt.Println(palA, " = ", palB)

	nda := 7 % 5
	fmt.Println("Kabisat : ", nda)
	fmt.Println("--------------------------------------------------------------------")

	//Leap Year
	lpyear := leapyear(1900, 2020)
	fmt.Println("Leap Year : ", lpyear)
	fmt.Println("--------------------------------------------------------------------")

	//Reverse
	rev := Reverse("I am A Great human")
	fmt.Println("REVERSE WORD : ", rev)
	fmt.Println("--------------------------------------------------------------------")

	//Fabionacci
	arr := []int{15, 1, 3}
	arx := sum(arr)
	fa := GetFibonacci(arx)
	nilaiFA := fa[len(fa)-1] % arx
	fmt.Println("Nearest Fibonacci", nilaiFA)
	fmt.Println("--------------------------------------------------------------------")

	//FizzBuzz
	fizz := fizzbuzz(15)
	fmt.Println(fizz)
	fmt.Println("--------------------------------------------------------------------")

}

func palindrome(dataA string) bool {

	var ptrue = []string{"Radar", "Malam", "Kasur ini rusak", "Ibu Ratna antar ubi"}
	for _, ok := range ptrue {

		if dataA == ok {
			return true
		}

	}

	return false
}

func leapyear(yearA, yearB int) []int {

	// Cek kondisi jika tahun habis dibagi 400 maka tahun tersebut adalah tahun kabisat.
	// Cek kondisi jika tahun habis dibagi 100 maka tahun tersebut bukan tahun kabisat.
	// Cek kondisi jika tahun habisa dibagi 4 maka tahun tersebut adalah tahun kabisat.
	// Jika tahun tidak dapat habis dibagi 400, 100 dan 4 maka tahun tersebut bukan tahun kabisat.

	var data []int
	var oke = true //untuk tampungan boolean dibawah

	for i := yearA; i <= yearB; i++ {

		if i%400 == 0 {
			oke = true
		} else if i%100 == 0 {
			oke = false
		} else if i%4 == 0 {
			oke = true
		} else {
			oke = false
		}

		if oke == true {
			data = append(data, i)
		}

	}

	return data

}

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func GetFibonacci(n int) []int {

	//19
	var vB int
	var sliceA []int

	for i := 1; i <= n+n; i++ {
		vB += i
		sliceA = append(sliceA, vB)

		if vB > n {
			break
		}

	}

	return sliceA

}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func fizzbuzz(a int) []string {

	var number []string
	xx := ""

	for i := 1; i <= a; i++ {

		if i%3 == 0 && i%5 != 0 {
			xx = "Fizz"
		} else if i%5 == 0 && i%3 != 0 {
			xx = "Buzz"
		} else if (i%3 == 0) || (i%5 == 0) {
			xx = "FizzBuzz"
		} else {
			xx = ""
		}

		if xx != "" {
			number = append(number, xx)
		} else {
			number = append(number, strconv.Itoa(i))
		}

	}

	return number
}
