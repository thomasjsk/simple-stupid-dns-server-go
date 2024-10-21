package dns

import "encoding/binary"

type AnswerDomain string

func (answerDomain AnswerDomain) ToAnswer() Answer {
	var _type uint16 = 1
	var class uint16 = 1
	var ttl uint32 = 60
	var rdlength uint16 = 4
	var rdata uint32 = 8888

	return Answer{
		name:     string(answerDomain),
		Type:     _type,
		class:    class,
		ttl:      ttl,
		rdlength: rdlength,
		rdata:    rdata,
	}
}

type Answer struct {
	name     string
	Type     uint16
	class    uint16
	ttl      uint32
	rdlength uint16
	rdata    uint32
}

func (answer Answer) ToBytes() []byte {
	data := []byte(answer.name)
	data = binary.BigEndian.AppendUint16(data, answer.Type)
	data = binary.BigEndian.AppendUint16(data, answer.class)
	data = binary.BigEndian.AppendUint32(data, answer.ttl)
	data = binary.BigEndian.AppendUint16(data, answer.rdlength)
	data = binary.BigEndian.AppendUint32(data, answer.rdata)

	return data
}
