package main

import (
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// Constants
const (
	OneMSec   = 1000 * 1000       // One milliseond
	OneSec    = 1000 * 1000 * 100 // One second
	TextSpeed = 50                // Text showing speed
)

// Return Ghost and his height
func getGhost() string {
	ghost := ps1str() + " \n" +
		ps1str() + Bold("    ___\n") +
		ps1str() + Bold("  _/ @@\\\n") +
		ps1str() + Bold(" ( \\  0/__\n") +
		ps1str() + Bold("  \\    \\__)\n") +
		ps1str() + Bold("  /     \\\n") +
		ps1str() + Bold(" /       \\\n") +
		ps1str() + Bold(" ^^^^^^^^^\n") +
		ps1str()
	return ghost
}

// Show text from black to white
func textShowSlow(inString string) {
	for i := 232; i <= 255; i++ {
		os.Stdout.Write([]byte("\033[38;5;" + strconv.Itoa(i) + "m" + Bold(inString) + "\033[0m\r"))
		time.Sleep(OneMSec * TextSpeed)
		os.Stdout.Sync()
	}
}

// Show text from white to black
func textHideSlow(inString string) {
	for i := 255; i >= 232; i-- {
		os.Stdout.Write([]byte("\033[38;5;" + strconv.Itoa(i) + "m" + Bold(inString) + "\033[0m\r"))
		time.Sleep(OneMSec * TextSpeed)
		os.Stdout.Sync()
	}

}

// Clearing screen
func clearScreen() {
	os.Stdout.Write([]byte("\033[H\033[2J"))
	os.Stdout.Sync()
}

// Bold text
func Bold(inString string) string {
	return "\033[1m" + inString + "\033[0m"
}

// Creating PS1
func ps1str() string {
	return "ghost@shell:~$ "
}

// Show creds on exit
func showCreds() {
	os.Stdout.Write([]byte("\n\n"))
	os.Stdout.Write([]byte(centrifyText("Created")))
	os.Stdout.Write([]byte("\n"))
	os.Stdout.Write([]byte(centrifyText(" by")))
	os.Stdout.Write([]byte("\n\n"))
	os.Stdout.Write([]byte(Bold(centrifyText("Andrey Esin"))))
	os.Stdout.Write([]byte("\n\n"))
	os.Stdout.Write([]byte(centrifyText("[ twitter.com/la_stik ] [ t.me/la_stik ] [ andrey@esin.name ]")))
	os.Stdout.Write([]byte("\n\n"))
	os.Stdout.Write([]byte(Bold(centrifyText("Sources"))))
	os.Stdout.Write([]byte("\n"))
	os.Stdout.Write([]byte(centrifyText("[ github.com/esin/intheshell ]")))
	os.Stdout.Write([]byte("\n"))
	os.Stdout.Write([]byte(centrifyText("[ bitbucket.org/andrey_esin/intheshell ]")))
	os.Stdout.Write([]byte("\n"))

	os.Stdout.Sync()

	time.Sleep(OneSec * 3)
}

// Get terminal count
func getTTYSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	//println(string(out))
	if err != nil {
		//log.Fatal(err)
		appExit()
	}
	outStr := strings.Replace(string(out), "\n", "", -1)
	cols, err := strconv.Atoi(strings.Split(outStr, " ")[1])
	if err != nil {
		//log.Fatal(err)
		appExit()
	}
	rows, err := strconv.Atoi(strings.Split(outStr, " ")[0])
	if err != nil {
		//log.Fatal(err)
		appExit()
	}
	return cols, rows
}

// Return string, which can be showed in horizontal center of terminal
func centrifyText(inText string) string {
	cols, _ := getTTYSize()
	resultString := ""
	spacesCount := (cols / 2) - (len(inText) / 2)
	for i := 0; i < spacesCount; i++ {
		resultString = resultString + " "
	}
	resultString = resultString + inText
	return resultString
}

// Right exiting from application
func appExit() {
	clearScreen()
	showCreds()
	os.Stdout.Write([]byte("\033[?25h"))
	os.Exit(0)
}

func centerVertical() string {
	_, rows := getTTYSize()

	resultString := "\n"
	for i := 0; i < rows/2; i++ {
		resultString += "\n"
	}

	return resultString
}

func main() {
	// Catch ctrl-c
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-c
		appExit()
	}()

	// Don't allow using scp, ssh with params and similar
	args := os.Args
	if len(args) > 1 {
		os.Stdout.Write([]byte("Hey, just try:" + Bold(" ssh ghost@theshell.xyz")))
		os.Stdout.Write([]byte("\n"))
		os.Stdout.Write([]byte("\033[?25h"))
		os.Exit(0)
	}

	clearScreen()

	// Hide cursor
	os.Stdout.Write([]byte("\033[?25l"))

	//1 scene
	//Andrey Esin
	os.Stdout.Write([]byte(centerVertical()))
	textShowSlow(centrifyText("Andrey Esin"))
	time.Sleep(OneSec * 5)
	textHideSlow(centrifyText("Andrey Esin"))
	time.Sleep(OneSec * 1)
	clearScreen()
	// 2 scene
	// PRESENTS
	os.Stdout.Write([]byte(centerVertical()))
	textShowSlow(centrifyText("PRESENTS"))
	time.Sleep(OneSec * 5)
	textHideSlow(centrifyText("PRESENTS"))
	time.Sleep(OneSec * 1)
	clearScreen()

	// 3 scene
	// GHOST IN THE SHELL (bash)
	os.Stdout.Write([]byte(centerVertical()))
	textShowSlow(centrifyText("GHOST IN THE SHELL (bash)"))
	time.Sleep(OneSec * 5)
	textHideSlow(centrifyText("GHOST IN THE SHELL (bash)"))
	time.Sleep(OneSec * 1)
	clearScreen()

	cols, _ := getTTYSize()
	ghost := getGhost()
	spaces := " "
	for i := 0; i < cols-15-len(ps1str()); i++ {
		clearScreen()
		spaces += " "

		result := strings.Replace(ghost, ps1str(), ps1str()+spaces, -1)
		//println(result)
		os.Stdout.Write([]byte(result))
		time.Sleep(OneMSec * 50)
		os.Stdout.Sync()
	}

	clearScreen()

	//GHOST IN THE SHELL (bash)
	os.Stdout.Write([]byte(centerVertical()))
	textShowSlow(centrifyText("THE END"))
	time.Sleep(OneSec * 5)
	textHideSlow(centrifyText("THE END"))
	time.Sleep(OneSec * 1)

	appExit()
}
