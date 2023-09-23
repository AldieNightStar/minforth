package minforth

const (
	// Args: message
	OP_PRINT uint8 = iota

	// Args: cell_name
	OP_PRINT_FLUSH

	// Args: name value
	OP_SET

	// Args: result a b
	OP_ADD

	// Args: result a b
	OP_SUB

	// Args: result a b
	OP_MUL

	// Args: result a b
	OP_DIV

	// Args: result cell_name at
	OP_CELL_READ

	// Args: value cell_name at
	OP_CELL_WRITE

	// Args: seconds
	OP_WAIT

	// Args: pos
	OP_JUMP

	// Args: pos condition a b.
	//   Conditions:
	//   * lessThan
	//   * lessThanEq
	//   * equal
	//   * greaterThanEq
	//   * greaterThan
	//   * notEqual
	OP_JUMP_COND

	// Args: param_name block_name value
	OP_CONTROL

	// =====================
	// Special Operations
	// =====================

	// Args: value cell_name
	OP_SPEC_PUSH

	// Args: result cell_name
	OP_SPEC_POP

	// Args: cell_name
	OP_SPEC_DUPE

	// Args: cell_name
	OP_SPEC_DROP

	// Args: label_name
	OP_SPEC_DEF_LABEL

	// Args: label_name
	OP_SPEC_JUMP

	// Args: var_name cell_name
	OP_SPEC_SET_VAR

	// Args: var_name cell_name
	OP_SPEC_GET_VAR

	// Args: jump_pos cell_name
	OP_SPEC_LT

	// Args: jump_pos cell_name
	OP_SPEC_EQ

	// Args: jump_pos cell_name
	OP_SPEC_GT

	// Args: jump_pos cell_name
	OP_SPEC_NEQ

	// Args: jump_pos cell_name
	OP_SPEC_GTE

	// Args: jump_pos cell_name
	OP_SPEC_LTE

	// Args: 0
	OP_NONE
)
