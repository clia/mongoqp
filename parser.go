package mongoqp

import (
	"fmt"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
)

// Exp the expression top element
type Exp struct {
	Properties []*Property `"{" @@* "}"`
}

// Property a property element
type Property struct {
	Key   string `@Ident ":"`
	Value *Value `@@`
}

// Value some type of value
type Value struct {
	String   *string   `  @String`
	Number   *float64  `| @Number`
	Property *Property `| "{" @@ "}"`
}

// Parser the parser type
type Parser struct{}

var (
	basicLexer = lexer.Must(ebnf.New(`
	Ident = (alpha | "$") { "_" | alpha | digit } .
	String = "\"" { "\u0000"…"\uffff"-"\""-"\\" | "\\" any } "\"" .
	Number = [ "-" | "+" ] ("." | digit) { "." | digit } .
	Punct = "!"…"/" | ":"…"@" | "["…` + "\"`\"" + ` | "{"…"~" .
	EOL = ( "\n" | "\r" ) { "\n" | "\r" }.
	Whitespace = ( " " | "\t" ) { " " | "\t" } .

	alpha = "a"…"z" | "A"…"Z" .
	digit = "0"…"9" .
	any = "\u0000"…"\uffff" .
`))

	parser = participle.MustBuild(&Exp{},
		participle.Lexer(basicLexer),
		participle.CaseInsensitive("Ident"),
		participle.Unquote("String"),
		participle.UseLookahead(2),
		participle.Elide("Whitespace"),
	)
)

// Parse parse the expression string
func (p *Parser) Parse(expression string) (*Exp, error) {

	exp := &Exp{}
	err := parser.ParseString(expression, exp)
	return exp, err
}

func test() {

	// parser, err := participle.Build(&Exp{})
	// if err != nil {
	// 	fmt.Printf("%s\n", err.Error())
	// }
	exp := &Exp{}
	err := parser.ParseString(`{ R_STAT: 10 }`, exp)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%#v\n", exp)
	fmt.Printf("%#v\n", exp.Properties[0])
	fmt.Printf("%#v\n", exp.Properties[0].Value)

	exp2 := &Exp{}
	err = parser.ParseString(`{ ERR_S: { $gte: 1 } }`, exp2)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%#v\n", exp2)
	fmt.Printf("%#v\n", exp2.Properties[0].Value)

}
