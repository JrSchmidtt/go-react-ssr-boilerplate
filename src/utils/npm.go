package utils

import (
	"fmt"
	"os/exec"
)

func NpmInstall() error {
	fmt.Println("Installing npm packages...")
	if err := exec.Command("npm", "install").Run(); err != nil {
		return err
	}
	return nil
}