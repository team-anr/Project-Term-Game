package main

import (
	"fmt"
	"os"

	"runtime"
	//"io/ioutil"
	"os/exec"
)

func main() {
	nosave := 0
	fmt.Println("Loading...")
	for i := 0; i < len(os.Args); i++ {
		//fmt.Println(os.Args[i])
		if os.Args[i] == "--nosave" || os.Args[i] == "-n" {
			nosave = 1
		}
		if os.Args[i] == "--delsave" || os.Args[i] == "-d" {
			_, err := exec.Command("rm", configfl()).Output()
			//fmt.Println(string(out), err)
			if err != nil {
				fmt.Println("Error: " + err.Error())
				fmt.Println("HELP: Try Searching \"rm command " + runtime.GOOS + " " + err.Error() + "\" online")
			} else {
				fmt.Println("Deleted Save")
			}

			os.Exit(0)
		}
	}
	if _, err := os.Stat(os.Getenv("HOME") + "/.config/termgame/gamesav.var"); err != nil && nosave == 0 {
		fmt.Println("no game save found, creating save file...")
		_, err = exec.Command("mkdir", os.Getenv("HOME")+"/.config/termgame/", "-p").Output()
		fmt.Println(err)
		_, err := os.Create(os.Getenv("HOME") + "/.config/termgame/gamesav.var")
		//save.WriteString("Platform=" + runtime.GOOS)
		if err != nil {
			fmt.Println("There was an error creating the save file at " + os.Getenv("HOME") + "/.config/termgame/gamesav.var")
			fmt.Println(err)
		} else {
			fmt.Println("Successfully created save file!")
		}
		fmt.Println("Starting game...")
	}
}
