package main

import (
	"fmt"
	"os"
)

func main() {

}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
