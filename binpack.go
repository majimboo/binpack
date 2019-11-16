package binpack

import (
	"encoding/binary"
	"os"
)

// BinPack struct
type BinPack struct {
	fp         *os.File
	endianness binary.ByteOrder
}

// Open opens a file for reading
func Open(filepath string, isBigEndian bool) (*BinPack, error) {
	var b BinPack

	if isBigEndian {
		b.endianness = binary.BigEndian
	} else {
		b.endianness = binary.LittleEndian
	}

	fp, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	b.fp = fp

	return &b, nil
}

// ReadUInt8 reads a uint8 value and advances the offset
func (b *BinPack) ReadUInt8() (uint8, error) {
	var i uint8
	err := binary.Read(b.fp, b.endianness, &i)
	return i, err
}

// ReadUInt16 reads a uint16 value and advances the offset
func (b *BinPack) ReadUInt16() (uint16, error) {
	var i uint16
	err := binary.Read(b.fp, b.endianness, &i)
	return i, err
}

// ReadUInt32 reads a uint32 value and advances the offset
func (b *BinPack) ReadUInt32() (uint32, error) {
	var i uint32
	err := binary.Read(b.fp, b.endianness, &i)
	return i, err
}

// ReadUInt64 reads a int16 value and advances the offset
func (b *BinPack) ReadUInt64() (uint64, error) {
	var i uint64
	err := binary.Read(b.fp, b.endianness, &i)
	return i, err
}

// ReadInt8 reads a int8 value and advances the offset
func (b *BinPack) ReadInt8() (int8, error) {
	var i int8
	err := binary.Read(b.fp, b.endianness, &i)
	return i, err
}

// ReadInt16 reads a int16 value and advances the offset
func (b *BinPack) ReadInt16() (int16, error) {
	var i int16
	err := binary.Read(b.fp, b.endianness, &i)
	return i, err
}

// ReadInt32 reads a int32 value and advances the offset
func (b *BinPack) ReadInt32() (int32, error) {
	var i int32
	err := binary.Read(b.fp, b.endianness, &i)
	return i, err
}

// ReadInt64 reads a int16 value and advances the offset
func (b *BinPack) ReadInt64() (int64, error) {
	var i int64
	err := binary.Read(b.fp, b.endianness, &i)
	return i, err
}

// Seek forwards the offset
func (b *BinPack) Seek(offset int64, whence int) (int64, error) {
	return b.fp.Seek(offset, whence)
}
