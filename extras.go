package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clear() {
	if runtime.GOOS == "linux" {
		out, _ := exec.Command("clear").Output()
		fmt.Println(string(out))

	}
}
func home() string {
	return os.Getenv("HOME")
}

func configfl() string {
	return os.Getenv("HOME") + "/.config/termgame/gamesav.var"
}
