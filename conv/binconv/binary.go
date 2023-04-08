package binconv

import (
	"encoding/binary"
	"io"
)

func WriteBytes(w io.Writer, order binary.ByteOrder, data []byte) {
	binary.Write(w, order, int16(len(data)))
	binary.Write(w, order, data)

}

func WriteUTF(w io.Writer, order binary.ByteOrder, data string) (err error) {
	err = binary.Write(w, order, int16(len(data)))
	if err != nil {
		return
	}
	err = binary.Write(w, order, []byte(data))
	if err != nil {
		return
	}
	return
}

func ReadBool(r io.Reader, order binary.ByteOrder) bool {
	var val bool
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadInt8(r io.Reader, order binary.ByteOrder) int8 {
	var val int8
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadInt16(r io.Reader, order binary.ByteOrder) int16 {
	var val int16
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadInt32(r io.Reader, order binary.ByteOrder) int32 {
	var val int32
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadInt64(r io.Reader, order binary.ByteOrder) int64 {
	var val int64
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadUint8(r io.Reader, order binary.ByteOrder) uint8 {
	var val uint8
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadUint16(r io.Reader, order binary.ByteOrder) uint16 {
	var val uint16
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadUint32(r io.Reader, order binary.ByteOrder) uint32 {
	var val uint32
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadUint64(r io.Reader, order binary.ByteOrder) uint64 {
	var val uint64
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadFloat32(r io.Reader, order binary.ByteOrder) float32 {
	var val float32
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadFloat64(r io.Reader, order binary.ByteOrder) float64 {
	var val float64
	binary.Read(r, binary.BigEndian, &val)
	return val
}

func ReadBytes(r io.Reader, order binary.ByteOrder) (data []byte) {
	var tlen int16
	binary.Read(r, order, &tlen)
	data = make([]byte, tlen)
	binary.Read(r, order, &data)
	return
}

func ReadUTF(r io.Reader, order binary.ByteOrder) (data string) {
	tlen := int16(0)
	binary.Read(r, order, &tlen)
	tbs := make([]byte, tlen)
	binary.Read(r, order, &tbs)
	data = string(tbs)
	return
}
