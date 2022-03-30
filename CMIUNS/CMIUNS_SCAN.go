package CMIUNS

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"unicode"
)

func (p *C_MIUNS) GetFileContain() bool {
	//获取文件的代码内容，放到textBuffer中
	data, err := ioutil.ReadFile(p.codingTextName)
	p.textBuffer = string(data)
	if err != nil {
		return false
	} else {
		return true
	}

}

func (c C_MIUNS) temp() {
	//用来观察一些输出的函数
	//    fmt.Printf("%s",c.textBuffer)
	fmt.Printf("%d", NUM)
}

func (p C_MIUNS) printCodingLine(lineNo int, startPos int) {
	//lineNo是之前记录的行号，startPos是要从那个位置开始输出
	//由于输出以一行为单位，所以先找到从startPos开始，第一个回车的位置，输出他们之间的字符串
	//	for ;byte(p.textBuffer[startPos])=='\r' || byte(p.textBuffer[startPos])=='\n';startPos++{
	//
	//	}
	fmt.Printf("%d:", lineNo)
	if byte(p.textBuffer[startPos]) == '\r' || byte(p.textBuffer[startPos]) == '\n' {
		fmt.Println()
		return
	}
	endPos := strings.Index(p.textBuffer[int(startPos):], "\r\n") + startPos
	if endPos > startPos {
		//由于startPos肯定不是回车换行了，那么endPos如果是回车换行的话，一定比startPos大
		fmt.Println(p.textBuffer[startPos:endPos])
	} else {
		//可能在字符串结尾的时候，没有回车换行
		fmt.Println(p.textBuffer[startPos:])
	}

}

func (c *C_MIUNS) inputCodeText() {
	temp := ""
	c.textBuffer = ""
	_, err := fmt.Scanf("%s\n", &temp)
	for {
		if err == io.EOF {
			break
		}
		c.textBuffer = c.textBuffer + (temp + "\n")
		temp = ""
		_, err = fmt.Scanf("%s\n", &temp)
	}

}

func (p *C_MIUNS) printNewToken(theToken Token) {
	//fmt.Print(string(p.lineno)+":f")
	fmt.Printf("     %d:", p.lineno)
	if theToken.kind >= _IF_ && theToken.kind <= _WHILE_ {
		fmt.Print("reserved word: " + theToken.lexeme)
	} else if theToken.kind == ID {
		fmt.Print("ID,NAME=" + theToken.lexeme)
	} else if theToken.kind == NUM {
		fmt.Print("NUM,VAL=" + theToken.lexeme)
	} else if theToken.kind == ERROR {
		fmt.Print("ERROR,VAL=" + theToken.lexeme)
	} else {
		fmt.Print(theToken.lexeme)
	}
	fmt.Println()
}

func (p *C_MIUNS) getNextChar() int {
	p.textPoint++
	if p.textPoint <= len(p.textBuffer) {
		return int(p.textBuffer[p.textPoint-1])
	} else {
		return -1
	}
}

func (p *C_MIUNS) ungetNextChar() {
	p.textPoint--
}

func (p *C_MIUNS) getToken() Token {
	tokenBeginIndex := p.textPoint
	tokenEndIndex := p.textPoint
	currentToken := 0
	var state int
	state = START
	const _EOF_ = 255
	var c byte
	var commentState int
	//	var save bool = false
	for state != DONE {
		c = byte(p.getNextChar()) //每次都在缓冲区中读一个数
		//		save = true
		switch state {
		case START:
			if unicode.IsDigit(rune(c)) == true {
				//扫描到数字，进入INNUM模式
				state = INNUM
			} else if unicode.IsLetter(rune(c)) == true && c != _EOF_ {
				//c-的标识符全是字母，保留字是特殊的标识符
				state = INID
			} else if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
				if c == '\n' {
					p.lineno++
					tokenBeginIndex++
					p.printCodingLine(p.lineno, p.textPoint)
					break
				}
				tokenBeginIndex++
			} else if c == '=' {
				//如果是等于，可能是赋值或者是求是否等于的符号（"=="）
				state = INEQ
			} else if c == '<' {
				//如果是小于，可能就是小于或者小于或等于
				state = INLE
			} else if c == '>' {
				//如果是大于，可能是大于或者大于或等于
				state = INGE
			} else if c == '!' {
				state = INNOTEQ
			} else if c == '/' {
				//这个时候可能是除法也可能是注释
				state = INCOMMENT
				commentState = LCOMMENT
				tokenBeginIndex++
			} else {
				state = DONE
				switch c {
				case '+':
					currentToken = PLUS
					break
				case '-':
					currentToken = MIUNS
					break
				case '*':
					currentToken = TIMES
					break
				//case '/':
				//	currentToken = OVER
				//	break
				case ';':
					currentToken = SEMI
					break
				case ',':
					currentToken = COMMA
					break
				case '(':
					currentToken = LPAREN
					break
				case ')':
					currentToken = RPAREN
					break
				case '[':
					currentToken = LSQUARE
					break
				case ']':
					currentToken = RSQUARE
					break
				case '{':
					currentToken = LCURLY
					break
				case '}':
					currentToken = RCURLY
					break
				default:
					currentToken = ERROR
					break
				}
				if int8(c) == TextEnd {
					currentToken = ENDFILE
				}

			}
			break

		case INCOMMENT:
			commentFlag := 1
			if int8(c) == TextEnd {
				state = DONE
				currentToken = ENDFILE
			}
			switch commentState {
			case LCOMMENT:
				if c == '*' {
					commentState = MIDCOMMENT
				} else {
					//因为注释里的东西是不需要的，所以每循环一次，起始的index要加1
					//但是如果是除号的话是需要的，为了反正起始的index加1，将标志对象置0
					tokenBeginIndex--
					currentToken = OVER
					p.ungetNextChar()
					commentState = OVER
					state = DONE
					commentFlag = 0
				}
				break
			case MIDCOMMENT:
				if c == '*' {
					commentState = RCOMMENT
				}
				break
			case RCOMMENT:
				if c == '/' {
					state = START
					break
				} else if c != '*' {
					commentState = MIDCOMMENT
				}
				break
			}
			//因为注释里的东西是不需要的，所以每循环一次，起始的index要加1
			if c == '\n' {
				p.lineno++
				p.printCodingLine(p.lineno, p.textPoint)
			}
			if commentFlag == 1 {
				tokenBeginIndex++
			}
			break

		case INNUM:
			if unicode.IsDigit(rune(c)) == false {
				p.ungetNextChar()
				state = DONE
				currentToken = NUM
			}
			break

		case INID:
			if unicode.IsDigit(rune(c)) == false && unicode.IsLetter(rune(c)) == false {
				p.ungetNextChar()
				state = DONE
				currentToken = ID
			}
			break

		case INEQ:
			state = DONE //之前是'='，接下来只有一次判断，即后面是不是跟着另一个'='
			if c == '=' {
				currentToken = EQ
			} else {
				p.ungetNextChar()
				currentToken = ASSIGN
			}
			break

		case INLE:
			state = DONE
			if c == '=' {
				currentToken = LE
			} else {
				p.ungetNextChar()
				currentToken = LT
			}
			break

		case INGE:
			state = DONE
			if c == '=' {
				currentToken = GE
			} else {
				p.ungetNextChar()
				currentToken = GT
			}
			break

		case INNOTEQ:
			state = DONE
			if c == '=' {
				currentToken = NOTEQ
			} else {
				p.ungetNextChar()
				currentToken = ERROR
			}
			break

		}

		tokenEndIndex = p.textPoint - 1
		if state == DONE && currentToken != ENDFILE {
			TokenLexeme := p.textBuffer[tokenBeginIndex : tokenEndIndex+1]
			if currentToken == ID && KeyWord[TokenLexeme] != 0 {
				p.tokenList = append(p.tokenList, Token{KeyWord[TokenLexeme], TokenLexeme, p.lineno})
			} else {
				p.tokenList = append(p.tokenList, Token{currentToken, TokenLexeme, p.lineno})
			}
		}

	}
	if currentToken != ENDFILE {
		p.printNewToken(p.tokenList[len(p.tokenList)-1])
		return p.tokenList[len(p.tokenList)-1]
	} else {
		p.tokenList = append(p.tokenList, Token{ENDFILE, "", p.lineno})
		return Token{ENDFILE, "", p.lineno}
	}

}

func (p *C_MIUNS) Scan() {
	if p.GetFileContain() == false {
		return
	}
	p.textPoint = 0
	p.lineno = 1
	p.printCodingLine(p.lineno, p.textPoint)
	var currentToken Token
	currentToken = p.getToken()
	for currentToken.kind != ENDFILE {
		currentToken = p.getToken()
	}
}
