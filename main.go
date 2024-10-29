package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/vit0rr/mumu/compiler"
	"github.com/vit0rr/mumu/evaluator"
	"github.com/vit0rr/mumu/lexer"
	"github.com/vit0rr/mumu/object"
	"github.com/vit0rr/mumu/parser"
	"github.com/vit0rr/mumu/repl"
	"github.com/vit0rr/mumu/vm"
)

func main() {
	var (
		useCompiler = flag.Bool("compiler", false, "Flag without argument to use compiler to run Monkey code.")
		filePath    = flag.String("file", "", "Flag to run Monkey code from file.")
		help        = flag.Bool("help", false, "List available commands and their descriptions.")
	)

	flag.Parse()

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if *help {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		flag.PrintDefaults()
		return
	}

	if *useCompiler && *filePath == "" {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Println("Compiler mode with REPL")

		repl.Start(os.Stdin, os.Stdout, *useCompiler)
		return
	} else if *useCompiler && *filePath != "" {
		if err := runFile(*filePath, *useCompiler); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			return
		}
	} else if !*useCompiler && *filePath == "" {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Println("Interpreter mode with REPL")

		repl.Start(os.Stdin, os.Stdout, *useCompiler)
	} else {
		if err := runFile(*filePath, *useCompiler); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			return
		}
	}
}

func runFile(filePath string, useCompiler bool) error {
	env := object.NewEnvironment()

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

	if useCompiler {
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			return fmt.Errorf("woops! Compilation failed: %s", err)
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			return fmt.Errorf("woops! Executing bytecode failed:\n %s", err)
		}

		stackTop := machine.StackTop()
		io.WriteString(os.Stdout, stackTop.Inspect())
		io.WriteString(os.Stdout, "\n")

		return nil
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	}

	return nil
}
