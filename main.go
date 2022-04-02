package main

import (
	"fmt"
	"monolang/emitter"
	"monolang/opcodes"
	"os"
)

func main() {
	// content, err := ioutil.ReadFile("main.mono")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// program := string(content)

	// tokens := tokenizer.Tokenize(program)
	// ast := monoparser.Parse(tokens)
	// code := emitter.GenerateCode(ast)

	// fmt.Println(program)
	// fmt.Printf("%+v\n", tokens)
	// fmt.Printf("%+v\n", ast)
	// fmt.Printf("%+v\n", code)

	f, err := os.Create("./main.wasm")
	if err != nil {
		panic("file error")
	}
	defer f.Close()

	m := emitter.NewModule()

	// TYPE SECTION
	typeSection := emitter.NewSection(emitter.TYPE_SECTION)
	typeSection.AddFuncSignatures([]emitter.FuncSig{
		{
			Results: []byte{emitter.INT_32},
		},
	})
	m.AddSection(typeSection)

	// FUNC SECTION
	funcSection := emitter.NewSection(emitter.FUNC_SECTION)
	funcSection.AddFuncIndex(0)
	m.AddSection(funcSection)

	// EXPORT SECTION
	exportSection := emitter.NewSection(emitter.EXPORT_SECTION)
	exportSection.AddExports([]string{"main"})
	m.AddSection(exportSection)

	// CODE SECTION
	codeSection := emitter.NewSection(emitter.CODE_SECTION)
	codeSection.Size = 0x07
	codeSection.AddContent([]byte{0x01, 0x05, 0x00, opcodes.I32Const, byte(63), 0x0f, 0x0b})
	m.AddSection(codeSection)

	// // IMPORT SECTION
	// importSection := emitter.NewSection(emitter.IMPORT_SECTION)
	// importSection.AddImports("console", "log", 0)
	// m.AddSection(importSection)

	// imp := emitter.NewSection(emitter.IMPORT_SECTION, 10)
	// imp.AddContent([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e', ' ', 'l', 'o', 'g'})
	// m.AddSection(imp)

	code := m.ToByteCode()
	fmt.Println(code)

	f.Write(code)
}
