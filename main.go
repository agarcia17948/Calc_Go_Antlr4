/*
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
*/
package main

import (
	"fmt"
	"strconv"

	"parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type calcListener struct {
	*parser.BaseCalcListener
	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("empty stack")
	}
	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}


//----------------------

func (l *calcListener) ExitSum(c *parser.SumContext) {
	right, left := l.pop(), l.pop()
	l.push(left + right)
}

func (l *calcListener) ExitMul(c *parser.MulContext) {
	right, left := l.pop(), l.pop()
	l.push(left * right)
}

func (l *calcListener) ExitDigit(c *parser.DigitContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}
	l.push(i)
}
//--------------------------


func main() {
	fmt.Println("Parser:")

	// Setup the input
	is := antlr.NewInputStream("3*(5+4)")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.L())

	fmt.Println(listener.pop())
    fmt.Println("________termino todo correctamente________")
}







// go.mod 
// crear los paquetes go.mod



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