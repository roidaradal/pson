package pson

import "github.com/zeroibot/syntax"

const jsonCfg string = `
tokens:

LIT_WS          =   \s+
LIT_LCURLY      =   \{
LIT_RCURLY      =   \}
LIT_LSQUARE     =   \[
LIT_RSQUARE     =   \]
LIT_COLON       =   :
LIT_COMMA       =   ,
KW_NULL         =   null
BOOLEAN         =   (true|false)
STRING          =   "[^"\\]*(?:\\.[^"\\]*)*"
NUMBER          =   -?(?:0|[1-9]\d*)(?:\.\d+)?(?:[eE][+-]?\d+)?

grammar:

<JSON>          =   KW_NULL | BOOLEAN | STRING | NUMBER | <LIST> | <OBJECT>
<LIST>          =   LIT_LSQUARE LIT_RSQUARE | LIT_LSQUARE <JSON> <ITEMS> LIT_RSQUARE
<ITEMS>         =   EPSILON | LIT_COMMA <JSON> <ITEMS>
<OBJECT>        =   LIT_LCURLY LIT_RCURLY | LIT_LCURLY STRING LIT_COLON <JSON> <ENTRIES> LIT_RCURLY
<ENTRIES>       =   EPSILON | LIT_COMMA STRING LIT_COLON <JSON> <ENTRIES>
`

func newJSONParser() (*syntax.Parser, error) {
	return syntax.NewParserFrom(jsonCfg)
}
