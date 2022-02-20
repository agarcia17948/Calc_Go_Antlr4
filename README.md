
# USO DE ANTLR CON GO

-------------
## Instalacion de go en linux



### descargar el paquete de go para linux
 * https://go.dev/dl/

### ingresar a la carpeta de descargas con la terminal

*  $ cd Download

### eliminar la carpeta anterior si existiera

* $ sudo rm -rf /usr/local/go

### descomprimir el tar.gz en la carpeta /usr/local/go

* $ sudo tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz

### modificar el archivo bashrc

* $ nano ~/.bashrc

### agregar las variables GOPATH, PATH al archivo

* export GOROOT=/usr/local/go
* export GOPATH=$HOME/go
* export   PATH=$PATH:$GOPATH/bin:$GOROOT/bin



### guardar con  ctrl-x

### actualizar el shell ( actualizar el archivo bashrc )

* $ source ~/.bashrc

### comprobar la instalacion

* $ go -version

### output

* $ go version go1.12.1 linux/amd64

### crear la carpeta GO con subcarpetas BIN, SRC

* $ mkdir -p $HOME/go/{bin,src}


-------------
## instalacion de antlr 



### abrir una terminal y entrar a la carpeta:

* $ cd /usr/local/lib

### dentro de esta carpeta desgargar antlr.jar

* $ sudo curl -O https://www.antlr.org/download/antlr-4.9.2-complete.jar

### modificar el archibo bashrc y agregar las sig lines

* $ export CLASSPATH=".:/usr/local/lib/antlr-4.9.2-complete.jar:$CLASSPATH"
* $ alias antlr4='java -jar /usr/local/lib/antlr-4.9.2-complete.jar'
* $ alias grun='java org.antlr.v4.gui.TestRig'




### guardar con  ctrl-x

### actualizar el shell ( actualizar el archivo bashrc )

* $  source ~/.bashrc

### conprobar que antlr esta bien instalado

* $ antlr4

### output

* ANTLR Parser Generator  Version 4.9.2 . . .



-------------
### instalar el runtime package en GOPATH ( home/go )

### en una terminal ingresar a: cd home/go
* $ go get github.com/antlr/antlr4/runtime/Go/antlr


### crear una carpeta "calc" en cd home/go/src

* $ cd home/go/src 
* $ mkdir Calc

### crear un archivo Calc.g4 y main.go dentro de la carpeta Calc


* $ cd Calc 
* $ touch Calc.g4
* $ touch main.go


## cargar el archivo de la gramatica Calc.g4

```Go
// la gramatica
grammar Calc;
 
prog: expr;
// nombre de los no terminales con minusculas
// nombre de los terminales con mayusculas
expr: expr op=('*'|'/') expr  # OpBin
    | expr op=('+'|'-') expr  # OpBin
    | '(' expr ')'            # par
    | INT                     # num
    ;
 
INT : ('0'..'9')+ ;

```



### compilar el archivo gramatica ( generar la carpeta parser )

* $ antlr4 -Dlanguage=Go -o parser Calc.g4

### eliminar la carpeta parser si existe en Goroot

* $ sudo rm -rf /usr/local/go/src/parser

### copiar la carpeta generada(parser) a Goroot

* $ sudo mv parser /usr/local/go/src




## cargar el archivo de la main.go

```Go
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

```

### inicializar el proyecto  "carpeta_proyecto 'Calc' "

* $ go mod init Calc
* $ go mod tidy


### correr el main.go

* $ go run main.go







//
//
//

$ go env -w GO111MODULE=on

//
//
//

//https://github.com/go-graphics/go-gui-projects 

//
//
//

73 personas en la explicacion del proyecto
PABLO ANDRÃ‰S ROCA DOMINGUEZ21:11
3014066260101@ingenieria.usac.edu.gt
