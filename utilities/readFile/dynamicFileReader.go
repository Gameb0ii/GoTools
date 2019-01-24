package main

import (
	"readtxt"
)

func main() {
	r := readtxt.New("README.md")
	r.ReadFile()
}