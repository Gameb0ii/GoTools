package readtxt

import (
"fmt"
"os"
"io/ioutil"
"path/filepath"
)

type readtxt struct{
	input string
}

func New(input string) readtxt {
	reader := readtxt{input}
	return reader
}

//this is how error handling in go is done, apparently.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (r readtxt) ReadFile(){
	//Perhaps the most basic file reading task is slurping a file’s entire contents into memory.
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(filepath.Join(pwd,r.input))
	check(err)
	fmt.Print(string(dat))
}