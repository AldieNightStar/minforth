package minforth

func addOperation(typ uint8, code *Code, stackCell string) {
	code.Add(newOperation(OP_SPEC_POP, "VAL1", stackCell))
	code.Add(newOperation(OP_SPEC_POP, "VAL2", stackCell))
	code.Add(newOperation(typ, "VAL1", "VAL2", "VAL1"))
	code.Add(newOperation(OP_SPEC_PUSH, "VAL1", stackCell))
}

func Compile(stackCell string, messageCell string, src string) (string, error) {
	var err error
	var newops []*operation

	// Process tokens and enrich the *Code
	tokens := lex(src)
	code := newCode(stackCell, messageCell)
	processTokens(code, tokens)

	// Expand all the code
	newops, err = expandAllWithReexpand(code.Operations)
	if err != nil {
		return "", err
	}
	code.Operations = newops

	// Optimize the code at the end
	// We use three time optimization
	optimize(code, 3)

	// Process labels and jumps
	labels := takeLabels(code)
	newops, err = processJumps(code, labels)
	if err != nil {
		return "", err
	}
	code.Operations = newops

	// Process Logic labels: < > = <= >=
	newops, err = processLogicLabels(code, labels)
	if err != nil {
		return "", err
	}
	code.Operations = newops

	return code.String()
}

func processTokens(code *Code, tokens []string) {
	skips := 0
	for _, tok := range tokens {
		// If need to skip something
		if skips > 0 {
			skips -= 1
			continue
		}

		varGetName := lexVariableGetter(tok)
		varSetName := lexVariableSetter(tok)
		jumpLabel := lexJumpingToken(tok)
		labelName := lexLabel(tok)
		if varGetName != "" {
			code.Add(newOperation(OP_SPEC_GET_VAR, varGetName, code.StackCell))
		} else if varSetName != "" {
			code.Add(newOperation(OP_SPEC_SET_VAR, varSetName, code.StackCell))
		} else if jumpLabel != "" {
			code.Add(newOperation(OP_SPEC_JUMP, jumpLabel))
		} else if labelName != "" {
			code.Add(newOperation(OP_SPEC_DEF_LABEL, labelName))
		} else if tok == "+" {
			addOperation(OP_ADD, code, code.StackCell)
		} else if tok == "-" {
			addOperation(OP_SUB, code, code.StackCell)
		} else if tok == "*" {
			addOperation(OP_MUL, code, code.StackCell)
		} else if tok == "/" {
			addOperation(OP_DIV, code, code.StackCell)
		} else if tok == "print" {
			code.Add(newOperation(OP_SPEC_POP, "VAL1", code.StackCell))
			code.Add(newOperation(OP_PRINT, "VAL1"))
			code.Add(newOperation(OP_PRINT_FLUSH, code.MessageCell))
		} else if tok == "wait" {
			code.Add(newOperation(OP_SPEC_POP, "VAL1", code.StackCell))
			code.Add(newOperation(OP_WAIT, "VAL1"))
		} else if tok == "dup" {
			code.Add(newOperation(OP_SPEC_DUPE, code.StackCell))
		} else if tok == "drop" {
			code.Add(newOperation(OP_SPEC_DROP, code.StackCell))
		} else if tok == "<" {
			code.Add(newOperation(OP_SPEC_LT, "??", code.StackCell))
		} else if tok == "=" {
			code.Add(newOperation(OP_SPEC_EQ, "??", code.StackCell))
		} else if tok == ">" {
			code.Add(newOperation(OP_SPEC_GT, "??", code.StackCell))
		} else if tok == "<=" {
			code.Add(newOperation(OP_SPEC_LTE, "??", code.StackCell))
		} else if tok == ">=" {
			code.Add(newOperation(OP_SPEC_GTE, "??", code.StackCell))
		} else if tok == "!=" {
			code.Add(newOperation(OP_SPEC_NEQ, "??", code.StackCell))
		} else {
			_, isNum := lexNumber(tok)
			if isNum {
				code.Add(newOperation(OP_SPEC_PUSH, tok, code.StackCell))
			}
		}
	}
}
