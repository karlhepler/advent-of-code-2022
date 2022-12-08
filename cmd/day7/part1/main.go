package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fs"
	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var x = fs.NewExplorer()

	filepath := "cmd/day7/example"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		s := string(line)

		// no need to do anything if ls or dir
		if s == "$ ls" || strings.HasPrefix(s, "dir") {
			return nil
		}

		if strings.HasPrefix(s, "$ cd") {
			x.Chdir(s[5:])
			return nil
		}

		info := strings.Split(s, " ")
		x.CreateFile(info[1], it.Must2(strconv.Atoi(info[0])))

		return nil
	}))

	var total int
	for _, dirsize := range x.GetDirSizes() {
		if dirsize <= 100000 {
			total += dirsize
		}
	}

	fmt.Println(total)
}
