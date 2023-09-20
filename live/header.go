package live

import "encoding/binary"

type PacketHeader [16]byte

func (ph PacketHeader) Length() uint32 {
	return binary.BigEndian.Uint32(ph[0:4])
}
func (ph PacketHeader) Version() MsgVersion {
	v := binary.BigEndian.Uint16(ph[6:8])
	return MsgVersion(v)
}
func (ph PacketHeader) Operation() Op {
	op := binary.BigEndian.Uint32(ph[8:12])
	return Op(op)
}

func NewPacketHeader(op Op, v MsgVersion, length uint32) (ph PacketHeader) {
	binary.BigEndian.PutUint32(ph[0:4], length)
	binary.BigEndian.PutUint16(ph[4:6], 16)
	binary.BigEndian.PutUint16(ph[6:8], uint16(v))
	binary.BigEndian.PutUint32(ph[8:12], uint32(op))
	return ph
}
