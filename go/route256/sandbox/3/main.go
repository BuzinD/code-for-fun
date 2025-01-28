package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Dir    `json:"folders"`
}

const BAD = ".hack"

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
		var invalid int = 0
		// Считываем количество чисел
		var strQty int
		line, _ := in.ReadString('\n')
		strQty, _ = strconv.Atoi(strings.TrimSpace(line))

		var sb strings.Builder

		for j := 0; j < strQty; j++ {
			line, _ := in.ReadString('\n')
			line = strings.TrimSpace(line)
			sb.WriteString(line)
		}

		var dat Dir
		if err := json.Unmarshal([]byte(sb.String()), &dat); err != nil {
			fmt.Println("error")
			fmt.Println(err)
		}
		sb.Reset()

		var visited map[string]bool = map[string]bool{dat.Dir: true}

		dfs(dat.Dir, &dat, visited, &invalid, false)

		fmt.Fprintln(out, invalid)
	}
}

func dfs(path string, dir *Dir, visited map[string]bool, counter *int, infected bool) {

	//если сюда пришли и еще не были заражены проверяем все файлы
	if infected == false {
		for _, file := range dir.Files {
			if strings.HasSuffix(file, BAD) {
				infected = true
				break
			}
		}
	}

	//стали инфицированными или пришли из инфицированного каталога
	if infected {
		//увеличиваем счетчик на количество файлов в каталоге
		*counter += len(dir.Files)
	}

	for _, subDirectory := range dir.Folders {
		var newPath strings.Builder
		newPath.WriteString(path)
		newPath.WriteRune('/')
		newPath.WriteString(subDirectory.Dir)

		if _, ok := visited[newPath.String()]; !ok {
			//добавляем в посещенные
			visited[newPath.String()] = true
			//запускаем рекурсию
			dfs(newPath.String(), &subDirectory, visited, counter, infected)
		}
	}
}
