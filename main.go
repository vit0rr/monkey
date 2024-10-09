package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"

	"github.com/vit0rr/mumu/evaluator"
	"github.com/vit0rr/mumu/lexer"
	"github.com/vit0rr/mumu/object"
	"github.com/vit0rr/mumu/parser"
	"github.com/vit0rr/mumu/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		filePath := os.Args[1]
		fileExtension := strings.Split(filePath, ".")[1]

		if fileExtension != "monkey" {
			fmt.Fprintf(os.Stderr, "Invalid file extension: %s\n", fileExtension)
			return
		}

		if err := runFile(filePath); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			return
		}
	} else {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		repl.Start(os.Stdin, os.Stdout)
	}
}

func runFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	line := string(content)
	l := lexer.New(line)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		repl.PrintParserErrors(os.Stdout, p.Errors())
		return fmt.Errorf("parse errors in file")
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	}

	return nil
}
