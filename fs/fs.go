package fs

import (
	"path"
	"strings"
)

type Explorer struct {
	// map [filepath] filesize
	FS  map[string]int
	CWD string
}

func NewExplorer() *Explorer {
	return &Explorer{
		FS:  make(map[string]int),
		CWD: "/",
	}
}

func (x *Explorer) Chdir(dir string) {
	if strings.HasPrefix(dir, "/") {
		x.CWD = dir
		return
	}

	if dir == ".." {
		x.CWD = path.Dir(x.CWD)
		return
	}

	x.CWD = path.Join(x.CWD, dir)
}

func (x *Explorer) CreateFile(filename string, size int) {
	x.FS[path.Join(x.CWD, filename)] = size
}

func (x *Explorer) Mkdir(dirname string) {
	x.CreateFile(dirname, 0)
}

func (x *Explorer) GetDirSize(dir string) (dirsize int) {
	for filepath, filesize := range x.FS {
		if strings.HasPrefix(filepath, dir) {
			dirsize += filesize
		}
	}

	return dirsize
}

func (x *Explorer) GetDirSizes() []int {
	dirsizes := make([]int, 0)
	cache := make(map[string]int)

	for filepath := range x.FS {
		dir := path.Dir(filepath)
		if _, ok := cache[dir]; ok {
			continue
		}

		dirsize := x.GetDirSize(dir)
		dirsizes = append(dirsizes, dirsize)
		cache[dir] = dirsize
	}

	return dirsizes
}
