package main

import (
	"fmt"
	_ "os"
	_ "os/exec"
)
func lhome() {
	clear()
	lvl := [10]string{
		"┌──────────┘ └────────────────┐",
		"│                             │",
		"│                             │",
		"│                             │",
		"│                             │",
		"│                             │",
		"│                             │",
		"│                             │",
		"│                             │",
		"└─────────────────────────────┘"}
	for i := 0; i < 10; i++ {
		fmt.Println(lvl[i])
	}
}

func lvl1() {

}
