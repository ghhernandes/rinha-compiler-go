package ast

type (
	Term     any
	BinaryOp string
)

const (
	Add BinaryOp = "Add"
	Sub          = "Sub"
	Mul          = "Mul"
	Div          = "Div"
	Rem          = "Rem"
	Eq           = "Eq"
	Neq          = "Neq"
	Lt           = "Lt"
	Gt           = "Gt"
	Lte          = "Lte"
	Gte          = "Gte"
	And          = "And"
	Or           = "Or"
)

const (
	INT      = "Int"
	STR      = "Str"
	BOOL     = "Bool"
	VAR      = "Var"
	FUNCTION = "Function"
	CALL     = "Call"
	LET      = "Let"
	IF       = "If"
	BINARY   = "Binary"
	TUPLE    = "Tuple"
	PRINT    = "Print"
)

type (
	File struct {
		Name       string   `json:"name"`
		Expression Term     `json:"expression"`
		Location   Location `json:"location"`
	}

	Location struct {
		Start    int    `json:"start"`
		End      int    `json:"end"`
		Filename string `json:"filename"`
	}

	Parameter struct {
		Text     string   `json:"text"`
		Location Location `json:"location"`
	}

	Var struct {
		Kind     string   `json:"kind"`
		Text     string   `json:"text"`
		Location Location `json:"location"`
	}

	Function struct {
		Kind       string      `json:"kind"`
		Parameters []Parameter `json:"parameters"`
		Value      Term        `json:"value"`
		Location   Location    `json:"location"`
	}

	Call struct {
		Kind      string   `json:"kind"`
		Callee    Term     `json:"callee"`
		Arguments []Term   `json:"arguments"`
		Location  Location `json:"location"`
	}

	Let struct {
		Kind     string    `json:"kind"`
		Name     Parameter `json:"name"`
		Value    Term      `json:"value"`
		Next     Term      `json:"next"`
		Location Location  `json:"location"`
	}

	Str struct {
		Kind     string   `json:"kind"`
		Value    string   `json:"value"`
		Location Location `json:"location"`
	}

	Int struct {
		Kind     string   `json:"kind"`
		Value    int32    `json:"value"`
		Location Location `json:"location"`
	}

	Bool struct {
		Kind     string   `json:"kind"`
		Value    bool     `json:"value"`
		Location Location `json:"location"`
	}

	If struct {
		Kind      string   `json:"kind"`
		Condition Term     `json:"condition"`
		Then      Term     `json:"then"`
		Otherwise Term     `json:"otherwise"`
		Location  Location `json:"location"`
	}

	Binary struct {
		Kind     string   `json:"kind"`
		Lhs      Term     `json:"lhs"`
		Op       BinaryOp `json:"op"`
		Rhs      Term     `json:"rhs"`
		Location Location `json:"location"`
	}

	Tuple struct {
		Kind     string   `json:"kind"`
		First    Term     `json:"first"`
		Second   Term     `json:"second"`
		Location Location `json:"location"`
	}

	Print struct {
		Kind     string   `json:"kind"`
		Value    Term     `json:"value"`
		Location Location `json:"location"`
	}
)
