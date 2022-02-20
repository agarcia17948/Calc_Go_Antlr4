
package main

import (
    "fmt"
    "parser"  // copy parser folder to goroot
    "github.com/antlr/antlr4/runtime/Go/antlr"    
)

type calcListener struct{
    *parser.BaseCalcListener
}

func main() {
    fmt.Println("Parser: ")
    is:= antlr.NewInputStream("1+2*3")
    lexer:= parser.NewCalcLexer(is)
    stream:= antlr.NewCommonTokenStream(lexer,antlr.TokenDefaultChannel)
    p:= parser.NewCalcParser(stream)
    antlr.ParseTreeWalkerDefault.Walk(&calcListener{},p.L())
    fmt.Println("________termino todo correctamente________")
}



//    expression := "3*(5+4)"
	//input := antlr.NewInputStream(expression)
	//lexer := parser.NewCalcLexer(input)
	//stream := antlr.NewCommonTokenStream(lexer, 0)
	//p := parser.NewCalcParser(stream)
	//p.BuildParseTrees = true
	//tree := p.Expr()
	//var visitor = Visitor{}
	//var result = visitor.Visit(tree)
	//fmt.Println(expression, "=", result)