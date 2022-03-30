package CMIUNS

const baseIndex = 10
const TextEnd = -1

const (
	ID      = baseIndex + iota //letter标识符
	NUM                        //数字
	EQ                         //等于
	NOTEQ                      //不等于
	GE                         //大于或等于
	LE                         //小于或等于
	LT                         //小于
	GT                         //大于
	PLUS                       //加号
	MIUNS                      //减号
	TIMES                      //乘号
	OVER                       //除号
	ASSIGN                     //赋值
	LPAREN                     //"("
	RPAREN                     //")"
	LSQUARE                    //"["
	RSQUARE                    //"]"
	LCURLY                     //"{"
	RCURLY                     //"}"
	SEMI                       //分号
	COMMA                      //逗号
	_IF_
	_ELSE_
	_INT_
	_RETURN_
	_VOID_
	_WHILE_
)

const (
	START = _WHILE_ + 1 + iota
	DONE
	LCOMMENT
	MIDCOMMENT
	RCOMMENT
	INCOMMENT //处理注释
	INNUM
	INID
	INEQ    //处理等于或赋值
	INLE    //处理“小于”或者是“小于或等于”
	INGE    //处理“大于”或者是“大于或等于”
	INNOTEQ //处理不等于
	ERROR
	ENDFILE
)

const (
	NOEND = ENDFILE + 1 + iota
	NOTYPE
	NOID
	STATEMENT_LIST_ERR
	FACTOR_ERR
)
const (
	PROGRAM = FACTOR_ERR + 1 + iota
	DECLARATION_LIST
	DECLARATION
	VAR_DECLARATION
	TYPE_SPECIFIER
	FUN_DECLARATION
	PARAMS
	PARAM_LIST
	PARAM
	COMPOUND_STMT
	LOCAL_DECLARATIONS
	STATEMENT_LIST
	STATEMENT
	EXPRESSION_STMT
	SELECTION_STMT
	ITERATION_STMT
	RETURN_STMT
	EXPRESSION
	VAR
	SIMPLE_EXPRESSION
	RELOP
	ADDITIVE_EXPRESSION
	ADDOP
	TERM
	MULOP
	FACTOR
	CALL
	ARGS
	ARG_LIST
	EMPTY
)

type Node struct {
	child       []*Node
	parent      *Node
	layer       int
	childrenNum int
	kind        int
	lexeme      string
	NodeNo      int //编号
}

type Token struct {
	kind   int
	lexeme string
	lineno int //行号
}

var KeyWord = make(map[string]int)
var ParseTree Node
var nowNode *Node
var layersOfTree int
var maxNumOfChild int
var NoCount int = 0
var assignFlag = false
var errorCount = 0

type C_MIUNS struct {
	textBuffer     string
	tokenList      []Token
	textPoint      int
	tokenPoint     int
	nowToken       Token
	lineno         int //行号
	codingTextName string
}

func (c *C_MIUNS) Init(FileName string) {
	//"codingFile2.txt"
	KeyWord["if"] = _IF_
	KeyWord["else"] = _ELSE_
	KeyWord["int"] = _INT_
	KeyWord["return"] = _RETURN_
	KeyWord["void"] = _VOID_
	KeyWord["while"] = _WHILE_

	c.codingTextName = FileName
}
