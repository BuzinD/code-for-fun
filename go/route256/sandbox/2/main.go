package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	var out *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Считываем количество тестов
	var cases int
	fmt.Fscanln(in, &cases) // Используем Fscanln, чтобы убрать \n из буфера

	for i := 0; i < cases; i++ {
		var invalid bool = false
		// Считываем количество чисел
		var q int
		fmt.Fscanln(in, &q) // Считываем количество чисел

		// Считываем сами числа
		line, _ := in.ReadString('\n')
		numsStr := strings.Fields(strings.TrimSpace(line))
		nums := make([]int, q)

		for j, numStr := range numsStr {
			nums[j], _ = strconv.Atoi(numStr)
		}

		// Считываем ожидаемый результат
		expectedLine, _ := in.ReadString('\n')
		expectedLine = expectedLine[0 : len(expectedLine)-1] //убираем \n

		//количество считанных чисел из вывода
		qty := 0

		// Сортируем массив чисел
		sort.Ints(nums)

		var sb strings.Builder
		for j, c := range expectedLine {

			//если встречаем пробел вначале
			if j == 0 && c == ' ' {
				// fmt.Fprintln(out, "space in prefix")
				invalid = true
				break
			} else if c != ' ' && c != '-' && (c < '0' || c > '9') { //это недопустимый символ
				// fmt.Fprintln(out, "unexpected symbol "+string(c))
				invalid = true
				break
			} else if c == ' ' && expectedLine[j-1] == ' ' { // два подряд пробела
				// fmt.Fprintln(out, "double spaces j: "+strconv.Itoa(j))
				invalid = true
				break
			} else if c == ' ' && j == len(expectedLine)-1 { //последний символ в строке пробел
				// fmt.Fprintln(out, "last symbol is space ")
				invalid = true
				break
				//пора сравнивать накопленные
			} else if c == ' ' { //наткнулись на первый пробел
				qty++
				//слишком много чисел
				if len(nums) < qty || sb.String() != strconv.Itoa(nums[qty-1]) { //текущая подстрока не равна отсортированной
					// fmt.Fprintln(out, "too much numbers ")
					invalid = true
					break
				}
				//будем считывать следующее число
				sb.Reset()
			} else if j == len(expectedLine)-1 { //дошли до конца строки и символ не ' '
				qty++
				sb.WriteRune(c)
				//кол-во элементов меньше или последнее считанное число не равно последнему ожидаемому числу
				if len(nums) < qty || sb.String() != strconv.Itoa(nums[qty-1]) {
					// fmt.Fprintln(out, "last number isn't equal or too long out")
					invalid = true
					break
				}
			} else {
				sb.WriteRune(c)
			}
		}

		if invalid || len(nums) != qty {
			fmt.Fprintln(out, "no")
			continue
		}

		fmt.Fprintln(out, "yes")
	}
}
