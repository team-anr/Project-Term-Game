package main

import(
	"fmt"
	"os"
	"runtime"
	//"io/ioutil"
	"os/exec"
)


func main() {
	if _, err := os.Stat(os.Getenv("HOME") + "/.config/termgame/gamesav.var"); err != nil {
		fmt.Println("no game save found, creating save file...")
		_, err = exec.Command("mkdir", os.Getenv("HOME") + "/.config/termgame/", "-p").Output()
		fmt.Println(err)
		save, err := os.Create(os.Getenv("HOME") + "/.config/termgame/gamesav.var")
		save.WriteString("Platform=" + runtime.GOOS)
		if err != nil {
			fmt.Println("There was an error creating the save file at " + os.Getenv("HOME") + "/.config/termgame/gamesav.var")
			fmt.Println(err)
		} else {
			fmt.Println("Successfully created save file!")
		}
		fmt.Println("Starting game...")
	}
}