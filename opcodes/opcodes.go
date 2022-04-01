package opcodes

const (
	NullChar     = 0x00
	Block        = 0x02
	Loop         = 0x03
	Br           = 0x0c
	BrIf         = 0x0d
	End          = 0x0b
	Call         = 0x10
	GetLocal     = 0x20
	SetLocal     = 0x21
	I32Store8    = 0x3a
	I32Const     = 0x41 // Push a 32 bit integer to the stack
	F32Const     = 0x43 // Push a 32 bit float to the stack
	I32Eqz       = 0x45
	I32Eq        = 0x46
	F32Eq        = 0x5b
	F32LT        = 0x5d
	F32GT        = 0x5e
	I32And       = 0x71
	F32Add       = 0x92
	F32Sub       = 0x93
	F32Mul       = 0x94
	F32Div       = 0x95
	I32TruncF32S = 0xa8
)
