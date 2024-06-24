package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

// Задание расплывчатая, т.к. не указана верхняя граница чисел, поэтому первая реализация работает для чисел
// a = 2^x, b = 2^y, где х+у < 62 без переполнения, т.к. они приводятся к типу int64
// также можно использовать пакет "math/big" - вторая реализация
func main() {
	a := 1 << 25
	b := 1 << 37

	fmt.Printf("a = %d, b = %d, a * b = %d\n", a, b, multiply(a, b))
	fmt.Printf("a = %d, b = %d, b / a = %d\n", a, b, divide(b, a))
	fmt.Printf("a = %d, b = %d, a + b = %d\n", a, b, sum(a, b))
	fmt.Printf("a = %d, b = %d, b - a = %d\n", a, b, subtract(b, a))
	fmt.Printf("a = %d, b = %d, a - b = %d\n", a, b, subtract(a, b))
	fmt.Println()
	fmt.Println("----=Используем math.big=----")
	fmt.Println()
	andromedaDistInLightYears := new(big.Int)
	andromedaDistInLightYears.SetString("2637000", 10)
	kmInLightYear := big.NewInt(299793 * 31556952)
	andromedaDistInKm := new(big.Int).Mul(andromedaDistInLightYears, kmInLightYear)

	fmt.Println("Расстояние до галактики Андромеды составляет", andromedaDistInLightYears, "световых лет.")
	fmt.Println("Что в км составляет", andromedaDistInKm)
	proximaCentauriDistInLightYears := big.NewInt(4)
	diff := new(big.Int).Sub(andromedaDistInKm, new(big.Int).Mul(proximaCentauriDistInLightYears, kmInLightYear))
	fmt.Println("Что на " + diff.String() + " км больше, чем до Проксима Центавры")
	fmt.Println("Чтобы вы понимали: за время полета до Андромеды, можно полететь до Проксимы и обратно " + new(big.Int).Div(andromedaDistInLightYears, new(big.Int).Mul(proximaCentauriDistInLightYears, big.NewInt(2))).String() + " раз")

}
func multiply(a int, b int) int64 {
	return int64(a) * int64(b)
}

func divide(dividend int, divisor int) int64 {
	return int64(dividend) / int64(divisor)
}

func sum(a int, b int) int64 {
	return int64(a) + int64(b)
}

func subtract(minuend int, subtrahend int) int64 {
	return int64(minuend) - int64(subtrahend)
}
