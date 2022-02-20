// nombre de la gramatica  igual al nombre del archivo
grammar Calc;
 

//tokens
MUL: '*';
ADD: '+';
LB: '(';
RB: ')';
DIGIT: [0-9]+;
WS: [\r\n\t]+ -> skip;

// los nombres de los no terminales
// deben ir en minusculas
//rules

l: e EOF;
e: e '+' t   # Sum
 | t         # PasaT
;         
t: t '*' f   # Mul
 | f         # PasaF
;
f: '(' e ')' # PasaE
 | DIGIT     # Digit
;


// Sum, PassT, etc   nombre para enter o exit de la interfaz
// se debe poner nombre a todas las producciones
