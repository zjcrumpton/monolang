package emitter

type SectionCode byte
type ByteCode []byte

// Section Types
const (
	CUSTOM_SECTION     SectionCode = 0x00
	TYPE_SECTION       SectionCode = 0x01 // contains function signatures
	IMPORT_SECTION     SectionCode = 0x02
	FUNC_SECTION       SectionCode = 0x03 // wires up the type signatures to the code
	TABLE_SECTION      SectionCode = 0x04
	LINEAR_MEM_SECTION SectionCode = 0x05
	GLOBAL_SECTION     SectionCode = 0x06
	EXPORT_SECTION     SectionCode = 0x07
	START_SECTION      SectionCode = 0x08
	ELEMENT_SECTION    SectionCode = 0x09
	CODE_SECTION       SectionCode = 0x0a // contains function bodies
	DATA_SECTION       SectionCode = 0x0b
)

// Signature Types for TYPE_SECTION
const (
	FUNC   = 0x60
	INT_32 = 0x7f
)

// External Encoding
const (
	EXT_FUNC = 0x00
)

type Section struct {
	ID      byte   // identifies the type of section
	Size    int32  // identifies how many bytes the section will contain
	Content []byte // the relevant section data (should be the same length as size)
}

type FuncSig struct {
	Params  []byte
	Results []byte
}

func (s *Section) ToByteCode() ByteCode {
	bc := []byte{s.ID, byte(s.Size)}
	bc = append(bc, s.Content...)
	return bc
}

func (s *Section) AddContent(c []byte) {
	s.Content = append(s.Content, c...)
}

func (s *Section) AddFuncSignatures(sigs []FuncSig) {
	f := []byte{byte(len(sigs))}

	for _, sig := range sigs {
		f = append(f, FUNC)

		numOfParams := len(sig.Params)
		numOfResults := len(sig.Results)
		f = append(f, byte(numOfParams))
		f = append(f, sig.Params...)
		f = append(f, byte(numOfResults))
		f = append(f, sig.Results...)
		f = append(f)
	}

	s.AddContent(f)
	s.Size = int32(len(s.Content))
}

func (s *Section) AddFuncIndex(i int32) {
	s.Size = 2
	s.AddContent([]byte{0x01, byte(i)})
}

func (s *Section) AddImports(moduleName, exportName string, sigIndex int32) {
	nameLen := len(moduleName)
	exportLen := len(exportName)

	s.AddContent([]byte{byte(nameLen)})
	s.AddContent([]byte(moduleName))
	s.AddContent([]byte{byte(exportLen)})
	s.AddContent([]byte(exportName))
	s.AddContent([]byte{EXT_FUNC, byte(sigIndex)})
	s.Size = int32(len(moduleName) + len(exportName))
}

func (s *Section) AddExports(moduleNames []string) {
	var c []byte
	numOfExports := len(moduleNames)

	c = append(c, byte(numOfExports))
	for i, name := range moduleNames {
		c = append(c, byte(len(name)))
		c = append(c, []byte(name)...)
		c = append(c, EXT_FUNC)
		c = append(c, byte(i))
	}

	s.AddContent(c)
	s.Size = int32(len(c))
}

func NewSection(id SectionCode) Section {
	return Section{
		ID: byte(id),
	}
}
