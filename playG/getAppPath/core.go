package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	a := flag.String("info", "", "")
	flag.Parse()
	fmt.Println(GetAppPath())
	fmt.Println(GetAppDirectory())
	fmt.Println(filepath.Dir("test"))
	fmt.Println(*a)
}

func GetAppDirectory() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)

	return path
}
