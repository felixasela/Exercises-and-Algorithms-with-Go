// ClickCount

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	var clickCount int

	for {
		clearScreen()
		fmt.Printf("Contador de clics: %d\n", clickCount)
		fmt.Println("Presiona Enter para hacer clic (q para salir):")

		input := getUserInput()
		if input == "q" {
			break
		} else if input == "" {
			clickCount++
		}
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func clearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
