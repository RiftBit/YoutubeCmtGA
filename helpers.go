package main

import (
	"fmt"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
