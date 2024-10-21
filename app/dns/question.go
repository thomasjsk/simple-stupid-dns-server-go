package dns

import (
	"encoding/binary"
)

type RawQuestion []byte

func (question RawQuestion) ToQuestion() Question {
	cursor := uint(0)

	domainLen := uint(question[cursor])
	cursor++

	domain := string(question[cursor : cursor+domainLen])
	cursor += domainLen

	extensionLen := uint(question[cursor])
	cursor++

	extension := string(question[cursor : cursor+extensionLen])
	cursor += extensionLen

	name := domain + "." + extension

	return Question{
		Name:  encodeDomain(name),
		Type:  uint16(1),
		Class: uint16(1),
	}
}

type Question struct {
	Name  string
	Type  uint16
	Class uint16
}

func (question Question) ToBytes() []byte {
	data := []byte(question.Name)
	data = binary.BigEndian.AppendUint16(data, question.Type)
	data = binary.BigEndian.AppendUint16(data, question.Class)

	return data
}
