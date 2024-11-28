package main

import (
	"bufio"
	"fmt"
	"github.com/mariot/scanner"
	"os"
	"strings"
)

func run(source string) {
	currentInterpreter := scanner.NewScanner(source)
	result := currentInterpreter.Expr()

	fmt.Println(result)
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			break
		} else if text == "" {
			continue
		} else {
			run(text)
		}
	}
}

func main() {
	runPrompt()
}
