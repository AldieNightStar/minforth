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

	tokens := lex(src)
	code := newCode(stackCell, messageCell)

	// Skipping feature
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
			code.Add(newOperation(OP_SPEC_GET_VAR, varGetName, stackCell))
		} else if varSetName != "" {
			code.Add(newOperation(OP_SPEC_SET_VAR, varSetName, stackCell))
		} else if jumpLabel != "" {
			code.Add(newOperation(OP_SPEC_JUMP, jumpLabel))
		} else if labelName != "" {
			code.Add(newOperation(OP_SPEC_DEF_LABEL, labelName))
		} else if tok == "+" {
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
			code.Add(newOperation(OP_SPEC_DUPE, stackCell))
		} else if tok == "drop" {
			code.Add(newOperation(OP_SPEC_DROP, stackCell))
		} else if tok == "<" {
			code.Add(newOperation(OP_SPEC_LT, "??", stackCell))
		} else if tok == "=" {
			code.Add(newOperation(OP_SPEC_EQ, "??", stackCell))
		} else if tok == ">" {
			code.Add(newOperation(OP_SPEC_GT, "??", stackCell))
		} else if tok == "<=" {
			code.Add(newOperation(OP_SPEC_LTE, "??", stackCell))
		} else if tok == ">=" {
			code.Add(newOperation(OP_SPEC_GTE, "??", stackCell))
		} else if tok == "!=" {
			code.Add(newOperation(OP_SPEC_NEQ, "??", stackCell))
		} else {
			_, isNum := lexNumber(tok)
			if isNum {
				code.Add(newOperation(OP_SPEC_PUSH, tok, stackCell))
			}
		}
	}

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

	// Expand all the code
	newops, err = expandAllWithReexpand(code.Operations)
	if err != nil {
		return "", err
	}
	code.Operations = newops

	return code.String()
}
