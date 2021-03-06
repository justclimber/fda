package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/justclimber/fda/common/lang/fdalang"
)

func main() {
	sourceCode, _ := ioutil.ReadFile("example/example1")
	fmt.Printf("Running source code:\n%s\n", string(sourceCode))
	l := fdalang.NewLexer(string(sourceCode))
	p := fdalang.NewParser(l)

	astProgram, err := p.Parse()
	if err != nil {
		log.Fatalf("Parsing error: %s\n", err.Error())
	}
	env := fdalang.NewEnvironment()
	fmt.Println("Program output:")
	err = fdalang.NewExecAstVisitor().ExecAst(astProgram, env)
	if err != nil {
		log.Fatalf("Runtime error: %s\n", err.Error())
	}
	env.Print()
}
