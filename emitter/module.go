package emitter

import "monolang/opcodes"

type Module struct {
	Magic    []byte // identifies the binary as a WASM binary
	Version  []byte // identifies the WASM version
	Sections []Section
}

var (
	magicHeader   []byte = []byte{opcodes.NullChar, 'a', 's', 'm'}                            // 0asm
	versionHeader []byte = []byte{0x01, opcodes.NullChar, opcodes.NullChar, opcodes.NullChar} // Version 1
)

func (m *Module) ToByteCode() ByteCode {
	var code []byte
	code = append(code, m.Magic...)
	code = append(code, m.Version...)

	for _, s := range m.Sections {
		code = append(code, s.ToByteCode()...)
	}

	return code
}

func (m *Module) AddSection(s Section) {
	m.Sections = append(m.Sections, s)
}

func NewModule() Module {
	return Module{
		Magic:   magicHeader,
		Version: versionHeader,
	}
}
