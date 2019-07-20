package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Welcome to the chan crawler!")
	// var x string
	// fmt.Scanf("%s", &x)
	var scanning [2]string
	scanning[0] = "8chan"
	scanning[1] = "/b/"
	for {
		makescreen(scanning)
		time.Sleep(time.Second)
	}

}

func makescreen(scanning [2]string) {

	clearscreen() // Clear the screen

	fmt.Println("########################Chan Crawler########################")
	fmt.Println("Currently scanning:")
	fmt.Printf("%s - %s", scanning[0], scanning[1])

}

func clearscreen() {

	var c *exec.Cmd
	var doClear = true

	switch runtime.GOOS {
	case "darwin":
	case "linux":
		c = exec.Command("clear")
	case "windows":
		c = exec.Command("cmd", "/c", "cls")
	default:
		doClear = false
	}
	if doClear {
		c.Stdout = os.Stdout
		c.Run()
	}
}
