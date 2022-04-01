package main

import (
	"monolang/emitter"
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
	code := m.ToByteCode()

	f.Write(code)
}
