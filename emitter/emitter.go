package emitter

import (
	"monolang/monoparser"
)

// /0asm aka the string nullasm
func getMagicHeader() []byte {
	return []byte{0x00, 'a', 's', 'm'}
}

// Version 1
func getVersionHeader() []byte {
	return []byte{0x01, 0x00, 0x00, 0x00}
}

func GenerateCode(ast monoparser.AST) []byte {
	code := make([]byte, 0)
	// append headers
	code = append(code, getMagicHeader()...)
	code = append(code, getVersionHeader()...)
	// type section -> 5 bytes long -> number of types of variables -> function type -> number of input params -> number of output params -> result type i32
	// function signatures
	code = append(code, []byte{0x01, 0x05, 0x01, 0x60, 0x00, 0x01, 0x7f}...)

	// function section -> section size 2 byte long -> number of functions -> index of function
	code = append(code, []byte{0x03, 0x02, 0x01, 0x00}...)

	// export section -> section size 8 bytes long -> number of exports 1 -> length of export name 4 -> export name main -> export kind 0 for func -> index of exported function
	code = append(code, []byte{0x07, 0x08, 0x01, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x00, 0x00}...)

	// code section -> section size 7 bytes -> number of functions 1 -> function body size five bytes -> number of local declerations
	code = append(code, []byte{0x0a, 0x07, 0x01, 0x05, 0x00}...)

	// push a numeric const on the stack instruction, int literal, return
	// code = append(code, []byte{0x41, byte(42), 0x0f}...)

	// // section size
	// code = append(code, 0x00)
	// code = append(code, 0x00)

	// emitExpression := func(e monoparser.Expression) {
	// 	switch e.Type {
	// 	case monoparser.NumberLiteral:
	// 		code = append(code, opcodes.F32Const) // push a 32 bit float to the stack instruction
	// 		f, _ := strconv.ParseFloat(e.Value, 32)
	// 		code = append(code, byte(f)) // pushing the actual float
	// 	}
	// }

	// for _, node := range ast {
	// 	switch node.Type {
	// 	case monoparser.PrintStatement:
	// 		emitExpression(node.Expression)   // push the data to be printed onto the stack
	// 		code = append(code, opcodes.Call) // push the call instruction
	// 		code = append(code, 0)            // call the function at index 0
	// 	}
	// }

	return code
}
