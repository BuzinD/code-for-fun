package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Arrival struct {
	Idx  int
	Time int
}

func main() {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	var out *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Считываем количество тестов
	var casesQty int

	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)

	casesQty, err := strconv.Atoi(line)

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < casesQty; i++ {
		// Количество заказов
		var orderQty int
		line, _ := in.ReadString('\n')
		orderQty, _ = strconv.Atoi(strings.TrimSpace(line))

		//прибытие заказов
		var arrivalLine string
		arrivalLine, _ = in.ReadString('\n')
		arrivalLine = strings.TrimSpace(arrivalLine)

		arrivals := make([]Arrival, orderQty)

		for j, val := range strings.Fields(arrivalLine) {
			time, _ := strconv.Atoi(val)
			// время прибытия - индекс; значение - порядок
			arrivals[j] = Arrival{Idx: j, Time: time}
		}

		sort.Slice(arrivals, func(a, b int) bool {
			return arrivals[a].Time < arrivals[b].Time
		})

		//количество машин
		var carQty int
		line, _ = in.ReadString('\n')
		carQty, _ = strconv.Atoi(strings.TrimSpace(line))

		carSchedule := make([][4]int, carQty)

		for j := 0; j < carQty; j++ {
			line, _ = in.ReadString('\n')
			carSchedSlStr := strings.Fields(line)
			var carSched [4]int
			for k := 0; k < 3; k++ {
				carSched[k], _ = strconv.Atoi(carSchedSlStr[k])
			}
			carSched[3] = j
			carSchedule[j] = carSched
		}

		res := make([]int, orderQty)

		// Сортируем машины до обработки заказов
		sort.Slice(carSchedule, func(i, j int) bool {
			if carSchedule[i][0] == carSchedule[j][0] {
				return carSchedule[i][3] < carSchedule[j][3] // Если start одинаковые, сортируем по индексу
			}
			return carSchedule[i][0] < carSchedule[j][0]
		})

		//перебираем arrivals j-номер заказа
		lastCarIdx := 0
		for _, arrival := range arrivals {
			res[arrival.Idx] = -1
			for k := lastCarIdx; k < len(carSchedule); k++ {
				//если тачка заполнена
				if carSchedule[k][2] <= 0 {
					continue
				}
				//начало погрузки меньше arrivals окончание погрузки позже arrivals
				if arrival.Time >= carSchedule[k][0] && arrival.Time <= carSchedule[k][1] {
					carSchedule[k][2]--
					res[arrival.Idx] = carSchedule[k][3] + 1
					lastCarIdx = k
					break
				}

				//если дошли до машины у которой начальное время погрузки больше, то дальше нет смысла идти
				if carSchedule[k][0] > arrival.Time {
					break
				}
			}

		}

		for _, val := range res {
			fmt.Fprint(out, strconv.Itoa(val)+" ")
		}
		fmt.Fprintln(out, "")
	}
}
