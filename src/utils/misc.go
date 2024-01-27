package utils

import (
	"fmt"
	"os"
)

func CheckIfExistNodeModules() (bool) {
	fmt.Println("Checking if exist node_modules folder...")
	if _, err := os.Stat("node_modules"); os.IsNotExist(err) {
		return false
	}
	return true
}