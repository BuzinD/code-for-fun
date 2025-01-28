package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Field struct {
	M     int
	N     int
	Field [][]byte
}

func main() {
	var in *bufio.Reader = bufio.NewReader(os.Stdin)
	var out *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	casesQty := readCasesQty(in, out)

	for i := 0; i < casesQty; i++ {
		f, A, B := readInput(in, out)

		bMoreClosestToLeftUp := math.Sqrt(float64(A[0]*A[0]+A[1]*A[1])) - math.Sqrt(float64(B[0]*B[0]+B[1]*B[1]))

		if bMoreClosestToLeftUp > 0 {
			// fmt.Println("bMoreClosestToLeftUp")
			visited := make([][]bool, f.N)

			var path [][2]int
			for i := range visited {
				visited[i] = make([]bool, f.M)
			}

			visited[B[0]][B[1]] = true

			path = append(path, [2]int{B[0], B[1]})

			for i := range visited {
				visited[i] = make([]bool, f.M)
			}

			walk(B, f, &path, &visited, true)

			for _, p := range path {
				if f.Field[p[1]][p[0]] == 'B' {
					continue
				}
				f.Field[p[1]][p[0]] = 'b'
			}

			path = path[:0]

			for i := range visited {
				visited[i] = make([]bool, f.M)
			}

			visited[A[0]][A[1]] = true

			path = append(path, [2]int{A[0], A[1]})

			//отправляем гулять a
			walk(A, f, &path, &visited, false)

			for _, p := range path {
				if f.Field[p[1]][p[0]] == 'A' {
					continue
				}
				f.Field[p[1]][p[0]] = 'a'
			}

			for _, line := range f.Field {
				for _, ch := range line {
					fmt.Fprint(out, string(ch))
				}
				fmt.Fprintln(out, "")
			}

		} else {
			visited := make([][]bool, f.N)

			for i := range visited {
				visited[i] = make([]bool, f.M)
			}

			visited[A[0]][A[1]] = true
			var path [][2]int
			path = append(path, [2]int{A[0], A[1]})

			//отправляем гулять a
			walk(A, f, &path, &visited, true)

			for _, p := range path {
				if f.Field[p[1]][p[0]] == 'A' {
					continue
				}
				f.Field[p[1]][p[0]] = 'a'
			}

			path = path[:0]

			for i := range visited {
				visited[i] = make([]bool, f.M)
			}

			visited[B[0]][B[1]] = true
			path = append(path, [2]int{B[0], B[1]})

			//отправляем гулять B
			walk(B, f, &path, &visited, false)

			for _, p := range path {
				if f.Field[p[1]][p[0]] == 'B' {
					continue
				}
				f.Field[p[1]][p[0]] = 'b'
			}

			for _, line := range f.Field {
				for _, ch := range line {
					fmt.Fprint(out, string(ch))
				}
				fmt.Fprintln(out, "")
			}

		}
	}
}

func walk(coord [2]int, field Field, path *[][2]int, visited *[][]bool, left bool) bool {

	var ways [4][2]int

	if left {
		ways = [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	} else {
		ways = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	}

	if coord[0] == 0 && coord[1] == 0 || coord[0] == field.N-1 && coord[1] == field.M-1 {
		return true
	}

	for _, way := range ways {
		newX := coord[0] + way[0]
		newY := coord[1] + way[1]

		if newY >= field.M || newX >= field.N || newY < 0 || newX < 0 {
			continue
		}

		if field.Field[newY][newX] == '#' || field.Field[newY][newX] == 'a' || field.Field[newY][newX] == 'A' || field.Field[newY][newX] == 'B' || field.Field[newY][newX] == 'b' {
			continue
		}

		if (*visited)[newX][newY] { //проверить что там еще не были
			continue
		}

		point := [2]int{newX, newY}

		//добавить в путь
		*path = append(*path, point)

		(*visited)[newX][newY] = true

		if walk(point, field, path, visited, left) {
			return true
		}

		*path = (*path)[0 : len(*path)-1]
	}

	return false
}

func readCasesQty(in *bufio.Reader, out *bufio.Writer) int {

	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)

	casesQty, err := strconv.Atoi(line)
	if err != nil {
		fmt.Fprintln(out, "Ошибка чтения количества кейсов", err)
	}
	return casesQty
}

func readInput(in *bufio.Reader, out *bufio.Writer) (Field, [2]int, [2]int) {
	line, err := in.ReadString('\n')

	if err != nil {
		fmt.Fprintln(out, "Ошибка чтения количества столбцов и строк", err)
	}

	nMstr := strings.TrimSpace(line)

	nS := strings.Fields(nMstr)

	m, _ := strconv.Atoi(nS[0])
	n, _ := strconv.Atoi(nS[1])

	field := make([][]byte, m)

	// fmt.Println(len(field), cap(field))

	var a, b [2]int

	for i := 0; i < m; i++ {
		line, _ = in.ReadString('\n')
		line = strings.TrimSpace(line)

		lineByte := make([]byte, n)
		for j := 0; j < n; j++ {
			lineByte[j] = line[j]

			if lineByte[j] == 'A' {
				a[0] = j
				a[1] = i
			}
			if lineByte[j] == 'B' {
				b[0] = j
				b[1] = i
			}
		}
		field[i] = lineByte
	}

	return Field{m, n, field}, a, b
}
