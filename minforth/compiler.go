package minforth

func addOperation(typ int, code *Code, stackCell string) {
	code.Add(newOperation(OP_SPEC_POP, "VAL1", stackCell))
	code.Add(newOperation(OP_SPEC_POP, "VAL2", stackCell))
	code.Add(newOperation(typ, "VAL1", "VAL1", "VAL2"))
	code.Add(newOperation(OP_SPEC_PUSH, "VAL1", stackCell))
}

func Compile(stackCell string, messageCell string, src string) (string, error) {
	tokens := lex(src)
	code := newCode(stackCell, messageCell)
	for _, tok := range tokens {
		if tok == "+" {
			addOperation(OP_ADD, code, stackCell)
		} else if tok == "-" {
			addOperation(OP_SUB, code, stackCell)
		} else if tok == "*" {
			addOperation(OP_MUL, code, stackCell)
		} else if tok == "/" {
			addOperation(OP_DIV, code, stackCell)
		} else if tok == "print" {
			code.Add(newOperation(OP_SPEC_POP, "VAL1", stackCell))
			code.Add(newOperation(OP_PRINT, "VAL1"))
			code.Add(newOperation(OP_PRINT_FLUSH, messageCell))
		} else if tok == "wait" {
			code.Add(newOperation(OP_SPEC_POP, "VAL1", stackCell))
			code.Add(newOperation(OP_WAIT, "VAL1"))
		} else if tok == "dup" {
			code.Add(newOperation(OP_SPEC_POP, "VAL1", stackCell))
			code.Add(newOperation(OP_SPEC_PUSH, "VAL1", stackCell))
			code.Add(newOperation(OP_SPEC_PUSH, "VAL1", stackCell))
		} else {
			_, isNum := lexNumber(tok)
			if isNum {
				code.Add(newOperation(OP_SPEC_PUSH, tok, stackCell))
			}
		}
	}
	optimize(code)
	return code.String()
}
