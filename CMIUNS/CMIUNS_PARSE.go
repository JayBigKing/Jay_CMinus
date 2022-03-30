package CMIUNS

import (
	"fmt"
	"strconv"
)

func addChild(childKind int, childLexeme string, mode int) {
	nextNode := new(Node)
	//var nextNode Node = Node{nil, nil,(*nowNode).layer+1 ,childKind, childLexeme}
	nextNode.child = nil
	nextNode.layer = (*nowNode).layer + 1
	nextNode.kind = childKind
	nextNode.childrenNum = 0
	nextNode.lexeme = childLexeme
	nextNode.parent = nowNode
	nextNode.NodeNo = NoCount
	NoCount++
	if nextNode.layer > layersOfTree {
		layersOfTree = nextNode.layer
	}
	(*nowNode).child = append((*nowNode).child, nextNode)
	(*nowNode).childrenNum++
	if (*nowNode).childrenNum > maxNumOfChild {
		maxNumOfChild = (*nowNode).childrenNum
	}
	if mode == 1 {
		//有些时候需要执行新的孩子，有些时候不需要,主要是叶子结点的时候不需要
		nowNode = (*nowNode).child[len((*nowNode).child)-1]
	}
}
func returnToParentNode() {
	nowNode = (*nowNode).parent
}
func (p C_MIUNS) needVarAndEQ() bool {
	if assignFlag == true {
		return false
	}
	if p.nowToken.kind != ID {
		return false
	} else {
		i := p.tokenPoint - 1
		for true {
			i++
			if p.tokenList[i].kind == ASSIGN {
				assignFlag = true
				return true
			} else if p.tokenList[i].kind == SEMI || p.tokenList[i].kind == NUM || p.tokenList[i].kind == LPAREN || p.tokenList[i].kind == RPAREN {
				return false
			}

			if p.nowToken.kind == ENDFILE {
				return false
			}

		}
	}
	return false
}

func (p *C_MIUNS) declaration_list() {
	addChild(DECLARATION, "declaration", 1) //为当前的结点添加孩子结点
	p.declaration()
	returnToParentNode() //执行完推导以后，就回到双亲结点
	for p.nowToken.kind == _INT_ || p.nowToken.kind == _VOID_ {
		addChild(DECLARATION, "declaration", 1) //为当前的结点添加孩子结点
		p.declaration()
		returnToParentNode() //执行完推导以后，就回到双亲结点
	}
}

func (p *C_MIUNS) declaration() {
	//	p.nextToken()
	if p.nowToken.kind == _INT_ || p.nowToken.kind == _VOID_ {
		p.nextToken()
		if p.nowToken.kind == ID {
			p.nextToken()
			if p.nowToken.kind == LPAREN {
				p.returnToLastToken()
				p.returnToLastToken()
				addChild(FUN_DECLARATION, "fun_declaration", 1) //为当前的结点添加孩子结点
				p.fun_declaration()
				returnToParentNode() //执行完推导以后，就回到双亲结点
			} else {
				p.returnToLastToken()
				p.returnToLastToken()
				addChild(VAR_DECLARATION, "var_declaration", 1) //为当前的结点添加孩子结点
				p.var_declaration()
				returnToParentNode() //执行完推导以后，就回到双亲结点
			}
		} else {
			p.cmiunsError(ID, p.nowToken.lineno)
			p.nextToken()
		}
	} else {
		p.cmiunsError(NOTYPE, p.nowToken.lineno)
		p.nextToken()
	}
}

func (p *C_MIUNS) var_declaration() {

	addChild(TYPE_SPECIFIER, "type_specifier", 1) //为当前的结点添加孩子结点
	p.type_specifier()
	returnToParentNode() //执行完推导以后，就回到双亲结点
	if p.nowToken.kind == ID {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加孩子结点,ID
	} else {
		p.returnToLastToken()
		p.cmiunsError(ID, p.nowToken.lineno)
		p.nextToken()
		return
	}

	p.nextToken()

	if p.nowToken.kind == LSQUARE {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加孩子结点
		p.nextToken()
		if p.nowToken.kind == NUM {
			addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加孩子结点
			p.nextToken()
			if p.nowToken.kind == RSQUARE {
				addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加孩子结点
				p.nextToken()
			} else {
				p.cmiunsError(RSQUARE, p.nowToken.lineno)
				p.nextToken()
			}
		} else {
			p.cmiunsError(NUM, p.nowToken.lineno)
			p.nextToken()
		}

	}
	//else {
	//	p.returnToLastToken()
	//}

	for p.nowToken.kind == SEMI {
		p.nextToken()
	}

}

func (p *C_MIUNS) type_specifier() {
	if p.nowToken.kind == _INT_ || p.nowToken.kind == _VOID_ {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点
		p.nextToken()
	} else {
		p.cmiunsError(NOTYPE, p.nowToken.lineno)
		p.nextToken()
	}

}

func (p *C_MIUNS) fun_declaration() {
	addChild(TYPE_SPECIFIER, "type_specifier", 1) //为当前的结点添加孩子结点
	p.type_specifier()
	returnToParentNode()
	addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加孩子结点,ID，因为是检测过才进来的，所以肯定是ID
	p.nextToken()
	addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加孩子结点,(，因为是检测过才进来的，所以肯定是(
	p.nextToken()
	addChild(PARAMS, "params", 1) //为当前的结点添加孩子结点
	p.params()
	returnToParentNode() //执行完推导以后，就回到双亲结点

	if p.nowToken.kind == RPAREN {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加孩子结点,(，因为是检测过才进来的，所以肯定是(
		p.nextToken()
	} else {
		p.cmiunsError(RPAREN, p.nowToken.lineno)
		p.nextToken()
	}

	addChild(COMPOUND_STMT, "compound_stmt", 1) //为当前的结点添加孩子结点
	p.compound_stmt()
	returnToParentNode() //执行完推导以后，就回到双亲结点
}

func (p *C_MIUNS) params() {
	if p.nowToken.kind == _VOID_ {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,VOID
		p.nextToken()
	} else {
		addChild(PARAM_LIST, "param_list", 1) //为当前的结点添加孩子结点
		p.param_list()
		returnToParentNode() //执行完推导以后，就回到双亲结点
	}
}
func (p *C_MIUNS) param_list() {
	addChild(PARAM, "param", 1) //为当前的结点添加孩子结点
	p.param()
	returnToParentNode() //执行完推导以后，就回到双亲结点
	for p.nowToken.kind == COMMA {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,','
		p.nextToken()
		addChild(PARAM, "param", 1) //为当前的结点添加孩子结点
		p.param()
		returnToParentNode() //执行完推导以后，就回到双亲结点
	}

}
func (p *C_MIUNS) param() {
	addChild(TYPE_SPECIFIER, "type_specifier", 1) //为当前的结点添加孩子结点
	p.type_specifier()
	returnToParentNode() //执行完推导以后，就回到双亲结点
	if p.nowToken.kind == ID {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,ID
		p.nextToken()
		if p.nowToken.kind == LSQUARE {
			addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,[
			p.nextToken()
			if p.nowToken.kind == RSQUARE {
				addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,]
				p.nextToken()
			} else {
				p.cmiunsError(RSQUARE, p.nowToken.lineno)
				p.nextToken()
			}
		} else {
			return
		}
	} else {
		p.cmiunsError(ID, p.nowToken.lineno)
		p.nextToken()
	}

}

func (p *C_MIUNS) compound_stmt() {
	if p.nowToken.kind == LCURLY {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,{
		p.nextToken()
		addChild(LOCAL_DECLARATIONS, "local_declarations", 1) //为当前的结点添加孩子结点
		p.local_declarations()
		returnToParentNode()

		addChild(STATEMENT_LIST, "statement_list", 1) //为当前的结点添加孩子结点
		p.statement_list()
		returnToParentNode()
		if p.nowToken.kind == RCURLY {
			addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,}
			p.nextToken()
		} else {
			p.cmiunsError(RCURLY, p.nowToken.lineno)
			p.nextToken()
		}
	} else {
		p.cmiunsError(LCURLY, p.nowToken.lineno)
		p.nextToken()
	}
}

func (p *C_MIUNS) local_declarations() {
	if p.nowToken.kind == RCURLY {
		//p.returnToLastToken()
		return
	} else {
		if p.nowToken.kind == _INT_ || p.nowToken.kind == _VOID_ {
			for p.nowToken.kind == _INT_ || p.nowToken.kind == _VOID_ {
				addChild(VAR_DECLARATION, "var_declaration", 1) //为当前的结点添加孩子结点
				p.var_declaration()
				returnToParentNode() //执行完推导以后，就回到双亲结点
			}
		} else {
			addChild(EMPTY, "empty", 0) //为当前的结点添加叶子结点,}
		}
	}
}
func (p *C_MIUNS) statement_list() {
	if p.nowToken.kind == RCURLY {
		//p.returnToLastToken()
		return
	} else {
		if p.nowToken.kind == LCURLY || p.nowToken.kind == ID || p.nowToken.kind == LPAREN || p.nowToken.kind == _IF_ || p.nowToken.kind == _WHILE_ || p.nowToken.kind == _RETURN_ || p.nowToken.kind == SEMI {
			for p.nowToken.kind == LCURLY || p.nowToken.kind == ID || p.nowToken.kind == LPAREN || p.nowToken.kind == _IF_ || p.nowToken.kind == _WHILE_ || p.nowToken.kind == _RETURN_ || p.nowToken.kind == SEMI {
				addChild(STATEMENT, "statement", 1) //为当前的结点添加孩子结点
				p.statement()
				returnToParentNode() //执行完推导以后，就回到双亲结点
			}
		} else {
			addChild(EMPTY, "empty", 0) //为当前的结点添加叶子结点,}
		}
	}
}

func (p *C_MIUNS) statement() {
	if p.nowToken.kind == ID || p.nowToken.kind == LPAREN || p.nowToken.kind == SEMI {
		addChild(EXPRESSION_STMT, "expression_stmt", 1) //为当前的结点添加孩子结点
		if p.nowToken.kind != SEMI {
			p.expression_stmt()
		} else {
			addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,;
			p.nextToken()
		}
		returnToParentNode() //执行完推导以后，就回到双亲结点
	} else if p.nowToken.kind == _IF_ {
		addChild(SELECTION_STMT, "selection_stmt", 1) //为当前的结点添加孩子结点
		p.selection_stmt()
		returnToParentNode() //执行完推导以后，就回到双亲结点
	} else if p.nowToken.kind == _WHILE_ {
		addChild(SELECTION_STMT, "iteration_stmt", 1) //为当前的结点添加孩子结点
		p.iteration_stmt()
		returnToParentNode() //执行完推导以后，就回到双亲结点
	} else if p.nowToken.kind == _RETURN_ {
		addChild(RETURN_STMT, "return_stmt", 1) //为当前的结点添加孩子结点
		p.return_stmt()
		returnToParentNode() //执行完推导以后，就回到双亲结点
	} else if p.nowToken.kind == LCURLY {
		addChild(COMPOUND_STMT, "compound_stmt", 1) //为当前的结点添加孩子结点
		p.compound_stmt()
		returnToParentNode() //执行完推导以后，就回到双亲结点
	} else {
		p.cmiunsError(STATEMENT_LIST_ERR, p.nowToken.lineno)
		p.nextToken()
	}
}

func (p *C_MIUNS) expression_stmt() {
	if p.nowToken.kind == ID || p.nowToken.kind == LPAREN || p.nowToken.kind == NUM {
		addChild(EXPRESSION, "expression", 1)
		p.expression()
		returnToParentNode() //执行完推导以后，就回到双亲结点
	}
	if p.nowToken.kind == SEMI {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,;
		p.nextToken()
	} else {
		p.cmiunsError(SEMI, p.nowToken.lineno)
		p.nextToken()
	}
}
func (p *C_MIUNS) selection_stmt() {
	if p.nowToken.kind == _IF_ {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,if
		p.nextToken()
		if p.nowToken.kind == LPAREN {
			addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,(
			p.nextToken()
			addChild(EXPRESSION, "expression", 1)
			p.expression()
			returnToParentNode() //执行完推导以后，就回到双亲结点

			if p.nowToken.kind == RPAREN {
				addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,)
				p.nextToken()
				addChild(STATEMENT, "statement", 1)
				p.statement()
				returnToParentNode() //执行完推导以后，就回到双亲结点

				if p.nowToken.kind == _ELSE_ {
					addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,)
					p.nextToken()
					addChild(STATEMENT, "statement", 1)
					p.statement()
					returnToParentNode() //执行完推导以后，就回到双亲结点
				} else {
					return
				}

			} else {
				p.cmiunsError(RPAREN, p.nowToken.lineno)
				p.nextToken()
			}

		} else {
			p.cmiunsError(LPAREN, p.nowToken.lineno)
			p.nextToken()
		}
	} else {
		p.cmiunsError(_IF_, p.nowToken.lineno)
		p.nextToken()
	}
}
func (p *C_MIUNS) iteration_stmt() {
	if p.nowToken.kind == _WHILE_ {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,while
		p.nextToken()
		if p.nowToken.kind == LPAREN {
			addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,(
			p.nextToken()
			addChild(EXPRESSION, "expression", 1)
			p.expression()
			returnToParentNode() //执行完推导以后，就回到双亲结点

			if p.nowToken.kind == RPAREN {
				addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,)
				p.nextToken()
				addChild(STATEMENT, "statement", 1)
				p.statement()
				returnToParentNode() //执行完推导以后，就回到双亲结点

			} else {
				p.cmiunsError(RPAREN, p.nowToken.lineno)
				p.nextToken()
			}

		} else {
			p.cmiunsError(LPAREN, p.nowToken.lineno)
			p.nextToken()
		}
	} else {
		p.cmiunsError(_WHILE_, p.nowToken.lineno)
		p.nextToken()
	}
}
func (p *C_MIUNS) return_stmt() {
	if p.nowToken.kind == _RETURN_ {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,return
		p.nextToken()
		if p.nowToken.kind == ID || p.nowToken.kind == LPAREN || p.nowToken.kind == NUM {
			addChild(EXPRESSION_STMT, "expression_stmt", 1) //为当前的结点添加孩子结点
			p.expression_stmt()
			returnToParentNode() //执行完推导以后，就回到双亲结点
		}
	} else {
		p.cmiunsError(_RETURN_, p.nowToken.lineno)
		p.nextToken()
	}

}

func (p *C_MIUNS) expression() {
	/*
		for p.nowToken.kind == ID {
			addChild(VAR, "var", 1)
			p.myVar()
			returnToParentNode() //执行完推导以后，就回到双亲结点
			if p.nowToken.kind == ASSIGN {
				addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,return
				p.nextToken()
			} else {
				p.cmiunsError(ASSIGN, p.nowToken.lineno)
			}
		}
	*/
	for p.needVarAndEQ() == true {
		addChild(VAR, "var", 1)
		p.myVar()
		returnToParentNode()                            //执行完推导以后，就回到双亲结点
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,return
		p.nextToken()
		assignFlag = false
	}

	addChild(SIMPLE_EXPRESSION, "simple_expression", 1)
	p.simple_expression()
	returnToParentNode() //执行完推导以后，就回到双亲结点

}

func (p *C_MIUNS) myVar() {
	if p.nowToken.kind == ID {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
		p.nextToken()
		if p.nowToken.kind == LSQUARE {
			addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,]
			p.nextToken()
			addChild(EXPRESSION, "expression", 1)
			p.expression()
			returnToParentNode() //执行完推导以后，就回到双亲结点
			if p.nowToken.kind == RSQUARE {
				addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,]
				p.nextToken()
			} else {
				p.cmiunsError(RSQUARE, p.nowToken.lineno)
			}
		} else {
			return
		}
	} else {
		p.cmiunsError(ID, p.nowToken.lineno)
		p.nextToken()
	}
}

func (p *C_MIUNS) simple_expression() {
	addChild(ADDITIVE_EXPRESSION, "additive_expression", 1)
	p.additive_expression()
	returnToParentNode() //执行完推导以后，就回到双亲结点
	if p.nowToken.kind == LE || p.nowToken.kind == LT || p.nowToken.kind == GT || p.nowToken.kind == GE || p.nowToken.kind == EQ || p.nowToken.kind == NOTEQ {
		addChild(RELOP, "relop", 1)
		p.relop()
		returnToParentNode() //执行完推导以后，就回到双亲结点

		addChild(ADDITIVE_EXPRESSION, "additive_expression", 1)
		p.additive_expression()
		returnToParentNode() //执行完推导以后，就回到双亲结点

	}
}

func (p *C_MIUNS) relop() {
	if p.nowToken.kind == LE || p.nowToken.kind == LT || p.nowToken.kind == GT || p.nowToken.kind == GE || p.nowToken.kind == EQ || p.nowToken.kind == NOTEQ {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
		p.nextToken()
	} else {
		p.cmiunsError(RELOP, p.nowToken.lineno)
		p.nextToken()
	}
}
func (p *C_MIUNS) additive_expression() {
	addChild(TERM, "term", 1)
	p.term()
	returnToParentNode() //执行完推导以后，就回到双亲结点
	if p.nowToken.kind == PLUS || p.nowToken.kind == MIUNS {
		for p.nowToken.kind == PLUS || p.nowToken.kind == MIUNS {
			addChild(ADDOP, "addop", 1)
			p.addop()
			returnToParentNode() //执行完推导以后，就回到双亲结点

			addChild(TERM, "term", 1)
			p.term()
			returnToParentNode() //执行完推导以后，就回到双亲结点
		}
	}

}
func (p *C_MIUNS) addop() {
	if p.nowToken.kind == PLUS || p.nowToken.kind == MIUNS {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
		p.nextToken()
	} else {
		p.cmiunsError(ADDOP, p.nowToken.lineno)
		p.nextToken()
	}
}
func (p *C_MIUNS) term() {
	addChild(FACTOR, "factor", 1)
	p.factor()
	returnToParentNode() //执行完推导以后，就回到双亲结点

	if p.nowToken.kind == TIMES || p.nowToken.kind == OVER {
		addChild(MULOP, "mulop", 1)
		p.mulop()
		returnToParentNode() //执行完推导以后，就回到双亲结点

		//addChild(FACTOR, "factor", 1)
		//p.factor()
		//returnToParentNode() //执行完推导以后，就回到双亲结点
		addChild(SIMPLE_EXPRESSION, "simple_expression", 1)
		p.simple_expression()
		returnToParentNode() //执行完推导以后，就回到双亲结点
	}

}
func (p *C_MIUNS) mulop() {
	if p.nowToken.kind == TIMES || p.nowToken.kind == OVER {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
		p.nextToken()
	} else {
		p.cmiunsError(MULOP, p.nowToken.lineno)
		p.nextToken()
	}
}
func (p *C_MIUNS) factor() {
	if p.nowToken.kind == LPAREN {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
		p.nextToken()

		addChild(EXPRESSION, "expression", 1) //为当前的结点添加孩子结点
		p.expression()
		returnToParentNode() //执行完推导以后，就回到双亲结点

		if p.nowToken.kind == RPAREN {
			addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
			p.nextToken()
		} else {
			p.cmiunsError(RPAREN, p.nowToken.lineno)
			p.nextToken()
		}
	} else if p.nowToken.kind == ID {
		p.nextToken()
		if p.nowToken.kind == LPAREN {
			p.returnToLastToken()
			addChild(CALL, "call", 1) //为当前的结点添加孩子结点
			p.myCall()
			returnToParentNode() //执行完推导以后，就回到双亲结点
		} else {
			p.returnToLastToken()
			addChild(VAR, "var", 1) //为当前的结点添加孩子结点
			p.myVar()
			returnToParentNode() //执行完推导以后，就回到双亲结点
		}
	} else if p.nowToken.kind == NUM {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
		p.nextToken()
	} else {
		p.cmiunsError(FACTOR_ERR, p.nowToken.lineno)
		p.nextToken()
	}
}

func (p *C_MIUNS) myCall() {
	if p.nowToken.kind == ID {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
		p.nextToken()

		if p.nowToken.kind == LPAREN {
			addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
			p.nextToken()

			addChild(ARGS, "args", 1) //为当前的结点添加孩子结点
			p.args()
			returnToParentNode() //执行完推导以后，就回到双亲结点

			if p.nowToken.kind == RPAREN {
				addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
				p.nextToken()
			} else {
				p.cmiunsError(RPAREN, p.nowToken.lineno)
				p.nextToken()
			}

		} else {
			p.cmiunsError(LPAREN, p.nowToken.lineno)
			p.nextToken()
		}

	} else {
		p.cmiunsError(ID, p.nowToken.lineno)
		p.nextToken()
	}
}
func (p *C_MIUNS) args() {
	if p.nowToken.kind == ID || p.nowToken.kind == LPAREN {
		addChild(ARG_LIST, "arg_list", 1) //为当前的结点添加孩子结点
		p.arg_list()
		returnToParentNode() //执行完推导以后，就回到双亲结点
	}
}
func (p *C_MIUNS) arg_list() {
	addChild(EXPRESSION, "expression", 1) //为当前的结点添加孩子结点
	p.expression()
	returnToParentNode() //执行完推导以后，就回到双亲结点

	for p.nowToken.kind == COMMA {
		addChild(p.nowToken.kind, p.nowToken.lexeme, 0) //为当前的结点添加叶子结点,id
		p.nextToken()
		addChild(EXPRESSION, "expression", 1) //为当前的结点添加孩子结点
		p.expression()
		returnToParentNode() //执行完推导以后，就回到双亲结点

	}

}

func (p *C_MIUNS) cmiunsError(errorType int, lineNo int) {
	//根据标识输出错误信息
	errorCount++
	switch errorType {
	case NOEND:
		fmt.Printf("line %d,There is no End of the code", lineNo)
		break
	case NOTYPE:
		fmt.Printf("line %d,There is no Type of the declaration", lineNo)
		break
	case ID:
		fmt.Printf("line %d,There is no ID of the declaration", lineNo)
		break
	case RSQUARE:
		fmt.Printf("line %d,Expect a ]", lineNo)
		break
	case LSQUARE:
		fmt.Printf("line %d,Expect a [", lineNo)
		break
	case LCURLY:
		fmt.Printf("line %d,Expect a {", lineNo)
		break
	case RCURLY:
		fmt.Printf("line %d,Expect a }", lineNo)
		break
	case LPAREN:
		fmt.Printf("line %d,Expect a (", lineNo)
		break
	case RPAREN:
		fmt.Printf("line %d,Expect a )", lineNo)
		break
	case NUM:
		fmt.Printf("line %d,Expect a NUM", lineNo)
		break
	case STATEMENT_LIST_ERR:
		fmt.Printf("line %d,Expect a ID or ( or if or while or return or {", lineNo)
		break
	case SEMI:
		fmt.Printf("line %d,Expect a ;", lineNo)
		break
	case EQ:
		fmt.Printf("line %d,Expect a =", lineNo)
		break
	case _IF_:
		fmt.Printf("line %d,Expect a if", lineNo)
		break
	case _WHILE_:
		fmt.Printf("line %d,Expect a while", lineNo)
		break
	case _RETURN_:
		fmt.Printf("line %d,Expect a return", lineNo)
		break
	case RELOP:
		fmt.Printf("line %d,Expect a comparison operator", lineNo)
		break
	case ADDOP:
		fmt.Printf("line %d,Expect a addition or subtraction operators", lineNo)
		break
	case MULOP:
		fmt.Printf("line %d,Expect a multiplication or division operators", lineNo)
		break
	case FACTOR_ERR:
		fmt.Printf("line %d,Expect a ( or ID or NUM", lineNo)
		break

	default:
		fmt.Printf("line %d,%d", lineNo, errorType)
		break

	}
	fmt.Println()
}

func (p *C_MIUNS) nextToken() {
	p.nowToken = p.tokenList[p.tokenPoint]
	if p.tokenPoint < len(p.tokenList)-1 {
		p.tokenPoint++
	}
	for p.nowToken.kind == ERROR {
		errorCount++
		fmt.Printf("line %d,unexpect lex", p.nowToken.lineno)
		fmt.Println()
		p.nowToken = p.tokenList[p.tokenPoint]
		if p.tokenPoint < len(p.tokenList)-1 {
			p.tokenPoint++
		} else {
			break
		}
	}

}
func (p *C_MIUNS) returnToLastToken() {
	p.tokenPoint -= 2
	p.nowToken = p.tokenList[p.tokenPoint]
	p.tokenPoint++
}

func (p *C_MIUNS) Parse() bool {
	/*初始化根结点*/
	ParseTree.parent = nil
	ParseTree.kind = PROGRAM
	ParseTree.lexeme = "program"
	ParseTree.layer = 1
	ParseTree.childrenNum = 0
	ParseTree.NodeNo = NoCount
	NoCount++
	layersOfTree = 1
	/*令当前结点指向根结点*/
	nowNode = &ParseTree
	errorCount = 0
	p.tokenPoint = 0 //从第一个token开始分析
	p.nextToken()
	addChild(DECLARATION_LIST, "declaration_list", 1) //为当前的结点添加孩子结点
	p.declaration_list()                              //递归下降执行推导
	returnToParentNode()                              //执行完推导以后，就回到双亲结点
	if p.nowToken.kind != ENDFILE {
		p.cmiunsError(NOEND, p.nowToken.lineno)
		fmt.Println(strconv.Itoa(errorCount) + "errors")
		return false
	} else {
		if errorCount == 0 {
			fmt.Println("DONE")
			return true
		} else {
			fmt.Println(strconv.Itoa(errorCount) + " errors")
			return false
		}
	}
}
