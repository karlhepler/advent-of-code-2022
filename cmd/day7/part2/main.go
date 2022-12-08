package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fs"
	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var x = fs.NewExplorer()

	filepath := "cmd/day7/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		s := string(line)

		// no need to do anything if ls
		if s == "$ ls" {
			return nil
		}

		if strings.HasPrefix(s, "$ cd") {
			x.Chdir(s[5:])
			return nil
		}

		if strings.HasPrefix(s, "dir") {
			x.Mkdir(s[4:])
			return nil
		}

		info := strings.Split(s, " ")
		x.CreateFile(info[1], it.Must2(strconv.Atoi(info[0])))

		return nil
	}))

	dirsizes := x.GetDirSizes()
	sort.Ints(dirsizes)

	var i = len(dirsizes) - 1
	for ; i >= 0; i-- {
		if dirsizes[i] < 8381165 {
			break
		}
	}

	fmt.Println(dirsizes[i+1]) // 7209796 is WRONG!
}
