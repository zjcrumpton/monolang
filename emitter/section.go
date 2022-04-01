package emitter

type SectionCode byte
type ByteCode []byte

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

type Section struct {
	ID      byte   // identifies the type of section
	Size    uint32 // identifies how many bytes the section will contain
	Content []byte // the relevant section data (should be the same length as size)
}

func (s *Section) ToByteCode() ByteCode {
	bc := []byte{s.ID, byte(s.Size)}
	bc = append(bc, s.Content...)
	return bc
}

func (s *Section) AddContent(c []byte) {
	s.Content = append(s.Content, c...)
}

func NewSection(id SectionCode, s uint32) Section {
	return Section{
		ID:   byte(id),
		Size: s,
	}
}
