package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var slice []string
var operator string

func findOperator(text string) ([]string, string) {
	// Нахождение операнда и разбиение на элементы
	switch {
	case strings.Contains(text, "*"):
		operator = "*"
		slice = strings.Split(text, "*")
		return slice, operator
	case strings.Contains(text, "/"):
		operator = "/"
		slice = strings.Split(text, "/")
		return slice, operator
	case strings.Contains(text, "+"):
		operator = "+"
		slice = strings.Split(text, "+")
		return slice, operator
	case strings.Contains(text, "-"):
		operator = "-"
		slice = strings.Split(text, "-")
		return slice, operator
	default:
		fmt.Println("Ошибка! Это не математическая операция!")
		os.Exit(0)
	}
	return slice, operator
}

func checkSum(s []string) bool {
	if len(s) == 1 {
		fmt.Println("Ошибка! Строка не является математической операцией")
		os.Exit(0)
	} else if len(s) != 2 {
		fmt.Println("Ошибка формата! Принимаются только 2 операнда и 1 оператор")
		os.Exit(0)
	}
	return false
}

func checkConv(s []string) int {
	val1, err1 := strconv.Atoi(s[0])
	val2, err2 := strconv.Atoi(s[1])
	_, err3 := strconv.ParseFloat(s[0], 64)
	_, err4 := strconv.ParseFloat(s[1], 64)
	if err1 == nil && err2 == nil {
		if val1 > 10 || val1 < 1 || val2 > 10 || val2 < 1 {
			fmt.Println("Ошибка диапозона! Принимаются числа от 1 до 10 включительно")
			os.Exit(0)
			return 3
		}
		return 1
	} else if (err3 == nil && err1 != nil) || (err4 == nil && err2 != nil) {
		fmt.Println("Ошибка! Принимаются только целые числа")
		os.Exit(0)
	} else if err1 != nil && err2 != nil {
		return 2
	}
	fmt.Println("Ошибка формата записи!")
	os.Exit(0)
	return 3
}

func RimNum(s []string) []string {
	for ind := 0; ind < 2; ind++ {
		text1 := s[ind]
		result := 0
		for lent := len(text1); lent > 0; {
			if strings.Contains(text1, "IX") {
				result += 9
				lent -= 2
				text1 = strings.Replace(text1, "IX", "", 1)
			} else if strings.Contains(text1, "IV") {
				result += 4
				lent -= 2
				text1 = strings.Replace(text1, "IV", "", 1)
			} else if strings.Contains(text1, "X") {
				result += 10
				lent -= 1
				text1 = strings.Replace(text1, "X", "", 1)
			} else if strings.Contains(text1, "V") {
				result += 5
				lent -= 1
				text1 = strings.Replace(text1, "V", "", 1)
			} else if strings.Contains(text1, "I") {
				result += 1
				lent -= 1
				text1 = strings.Replace(text1, "I", "", 1)
			} else {
				fmt.Println("Неверный формат записи!")
				os.Exit(0)
			}
		}
		if text1 != "" {
			fmt.Println("Ошибка записи римских чисел!")
			os.Exit(0)
		} else {
			slice[ind] = strconv.Itoa(result)
		}
	}
	return slice
}

func ArabNum(s int) string {
	result := ""
	if s < 1 {
		fmt.Println("Ошибка! В римской системе счисления отсутствуют отрицательные числа и 0")
		os.Exit(0)
	}
	for s > 0 {
		if s >= 50 {
			if s == 100 {
				result += "C"
				s -= 100
			} else if s >= 90 {
				result += "XC"
				s -= 90
			} else if s >= 50 {
				result += "L"
				s -= 50
			}
		}
		if s >= 40 {
			result += "XL"
			s -= 40
		} else if s >= 10 {
			result += "X"
			s -= 10
		} else if s == 9 {
			result += "IX"
			s -= 9
		} else if s >= 5 {
			result += "V"
			s -= 5
		} else if s == 4 {
			result += "IV"
			s -= 4
		} else if s > 0 {
			result += "I"
			s -= 1
		}
	}
	return result
}

func operations(s []string, operator1 string) int {
	val1, _ := strconv.Atoi(s[0])
	val2, _ := strconv.Atoi(s[1])
	if operator1 == "+" {
		return val1 + val2
	} else if operator1 == "-" {
		return val1 - val2
	} else if operator1 == "/" {
		return val1 / val2
	} else {
		return val1 * val2
	}
}

func main() {
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите выражение: ")
		text, _ := input.ReadString('\n')
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, " ", "")
		findOperator(text) // нахождение оператора
		checkSum(slice)    // проверка на количество элементов
		if checkConv(slice) == 1 {
			fmt.Println(operations(slice, operator)) // операции над арабскими цифрами и вывод результата
		} else if checkConv(slice) == 2 {
			RimNum(slice) // перевод из араб. в рим. цифры
			if checkConv(slice) == 1 {
				fmt.Println(ArabNum(operations(slice, operator))) //операции, перевод из рим. в араб. цифры и вывод
			}
		}
	}
}
