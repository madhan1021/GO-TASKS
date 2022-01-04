package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dat, err := open("text.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dat)
	}
}

func open(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New("cant open the file")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.New("no contents exists")
	}
	fmt.Printf("%T\n", data)
	return string(data), nil
}
