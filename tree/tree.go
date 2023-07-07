package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	space = "  "
	line  = "│ "
	mid   = "├─"
	last  = "└─"
)

var dirCnt int
var fileCnt int

var levelFlag map[int]bool

func walkDir(dir string, depth int) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	levelFlag[depth] = true
	isLast := false
	for i, file := range files {
		absPath := filepath.Join(dir, file.Name())
		if i == len(files)-1 {
			isLast = true
		}
		s, err := os.Stat(absPath)
		if err != nil {
			fmt.Println(err)
		}
		writeLine(file.Name(), depth, isLast)
		if s.IsDir() {
			dirCnt++
			walkDir(absPath, depth+1)
		} else {
			fileCnt++
		}

	}

}

func showTree(dir string) {
	levelFlag = make(map[int]bool)
	fmt.Printf("%s\n", dir)
	walkDir(dir, 0)
	fmt.Printf("有%d目录%d个文件", dirCnt, fileCnt)
}

func writeLine(name string, depth int, isLast bool) {
	var builder strings.Builder
	for i := 0; i < depth; i++ {
		if levelFlag[i] {
			builder.WriteString(line)
		} else {
			builder.WriteString(space)
		}
	}
	if isLast {
		builder.WriteString(last)
		builder.WriteString(name)
		levelFlag[depth] = false
	} else {
		builder.WriteString(mid)
		builder.WriteString(name)
	}
	fmt.Println(builder.String())
}

func main() {
	dir := os.Args[1]
	showTree(dir)
}
