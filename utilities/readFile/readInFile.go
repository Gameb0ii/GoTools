package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

//this is how error handling in go is done, apparently.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

//pretty basic. Reads in a file. Mostly for learning purposes as of now
func main() {
	//Perhaps the most basic file reading task is slurping a file’s entire contents into memory.
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(filepath.Join(pwd,"README.md"))
	check(err)
	fmt.Print(string(dat))
}
