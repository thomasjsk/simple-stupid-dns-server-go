package dns

import (
	"encoding/binary"
)

type RawHeaderFlags []byte

func (rawHeaderFlags RawHeaderFlags) toHeaderFlags() HeaderFlags {
	flags := binary.BigEndian.Uint16(rawHeaderFlags)

	opcode := (flags >> 11) & 0xF
	rcode := uint16(4)
	if opcode == 0 {
		rcode = uint16(0)
	}

	return HeaderFlags{
		QR: 1,
		//QR:     (flags >> 15) & 0x1,
		OPCODE: opcode,
		AA:     (flags >> 10) & 0x1,
		TC:     (flags >> 9) & 0x1,
		RD:     (flags >> 8) & 0x1,
		RA:     (flags >> 7) & 0x1,
		Z:      (flags >> 4) & 0x3,
		RCODE:  rcode,
	}
}

type HeaderFlags struct {
	QR     uint16
	OPCODE uint16
	AA     uint16
	TC     uint16
	RD     uint16
	RA     uint16
	Z      uint16
	RCODE  uint16
}

func (headerFlags HeaderFlags) toUint16() uint16 {
	return headerFlags.QR<<15 | headerFlags.OPCODE<<11 | headerFlags.AA<<10 | headerFlags.TC<<9 |
		headerFlags.RD<<8 | headerFlags.RA<<7 | headerFlags.Z<<3 | headerFlags.RCODE
}

type RawHeader []byte

func (rawHeader RawHeader) toHeader() Header {
	return Header{
		ID:      binary.BigEndian.Uint16(rawHeader[0:2]),
		Flags:   RawHeaderFlags(rawHeader[2:4]).toHeaderFlags(),
		QDCOUNT: 1,
		ANCOUNT: 1,
		NSCOUNT: 1,
		ARCOUNT: 1,
	}
}

type Header struct {
	ID      uint16
	Flags   HeaderFlags
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}

func (header Header) toBytes() []byte {
	data := make([]byte, 12)
	binary.BigEndian.PutUint16(data[0:2], header.ID)
	binary.BigEndian.PutUint16(data[2:4], header.Flags.toUint16())
	binary.BigEndian.PutUint16(data[4:6], header.QDCOUNT)
	binary.BigEndian.PutUint16(data[6:8], header.ANCOUNT)
	binary.BigEndian.PutUint16(data[8:10], header.NSCOUNT)
	binary.BigEndian.PutUint16(data[10:12], header.ARCOUNT)

	return data
}
